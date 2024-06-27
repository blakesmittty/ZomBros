package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/google/uuid"

	"github.com/gorilla/handlers"
	"github.com/gorilla/websocket"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

type GameRoom struct {
	ID string
	Players map[*websocket.Conn]*Player
}

type Player struct {
	ID int 
	Username string 
}
var rooms = make(map[string]*GameRoom)


var upgrader = websocket.Upgrader{
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true // allowing all requests
	},
}

var db *sql.DB

func initDB() {
	var err error
	connStr := os.Getenv("DB_CONN_STR")
	if connStr == "" {
		log.Fatal("DB_CONN_STR env is not set")
	}

	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("successfully connected to database")
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func registerPlayer(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "invalid request method", http.StatusMethodNotAllowed)
		return
	}

	username := r.FormValue("username")
	email := r.FormValue("email")
	password := r.FormValue("password")

	log.Printf("received registration request: username=%s, email=%s", username, email)

	passwordHash, err := hashPassword(password)
	if err != nil {
		log.Printf("Error hashing password %v", err)
		http.Error(w, "Error hashing password", http.StatusInternalServerError)
		return
	}

	log.Printf("password hashed successfully for user %s", username)

	var playerId int
	query := `INSERT INTO players (username, email, password_hash) VALUES ($1, $2, $3) RETURNING id`
	err = db.QueryRow(query, username, email, passwordHash).Scan(&playerId)
	if err != nil {
		log.Printf("Error inserting new player into database %v", err)
		http.Error(w, "Error creating account", http.StatusInternalServerError)
		return
	}
	
	log.Printf("Account created successfully with Player ID: %d", playerId)
	fmt.Fprintf(w, "Account created successfully with Player ID: %d", playerId)
}

func loginPlayer(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "invalid request method", http.StatusMethodNotAllowed)
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")

	log.Printf("Received login request: username=%s", username)

	var storedHash string
	var playerId int
	query := `SELECT id, password_hash FROM players WHERE TRIM(username) = $1`
	err := db.QueryRow(query, username).Scan(&playerId, &storedHash)
	if err != nil {
		log.Printf("Error retreiving data: %v", err)
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	if !checkPasswordHash(password, storedHash) {
		log.Printf("Invalid password for user: username=%s", username)
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	log.Printf("User logged in successfully: username=%s, playerId=%d", username, playerId)

	response := map[string]interface{} {
		"id": playerId,
		"message": "login successful",
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func saveProgress(w http.ResponseWriter, r *http.Request) {
	playerID := r.FormValue("player_id")
	level := r.FormValue("level")
	score := r.FormValue("score")

	query := `INSERT INTO game_progress (player_id, level, score) VALUES ($1, $2, $3)
			ON CONFLICT (player_id) DO UPDATE SET level = EXCLUDED.level, score = EXCLUDED.score, updated_at = CURRENT_TIMESTAMP`
	_, err := db.Exec(query, playerID, level, score)
	if err != nil {
		http.Error(w, "error saving progress", http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "progress saved successfully")
}

func loadProgress(w http.ResponseWriter, r *http.Request) {
	playerID := r.FormValue("player_id")
	
	var level, score int
	query := `SELECT level, score FROM game_progress WHERE player_id = $1`
	err := db.QueryRow(query, playerID).Scan(&level, &score)
	if err != nil {
		http.Error(w, "Error loading progress", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Level: %d, Score: %d", level, score)
}

func unlockItem(w http.ResponseWriter, r *http.Request) {
	playerID := r.FormValue("player_id")
	itemID := r.FormValue("item_id")

	query := `INSERT INTO player_items (player_id, item_id) VALUES ($1, $2)`
	_, err := db.Exec(query, playerID, itemID)
	if err != nil {
		http.Error(w, "Error unlocking item", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "item successfully unlocked")
}

func getUnlockedItems(w http.ResponseWriter, r *http.Request) {
	playerID := r.FormValue("player_id")
	rows, err := db.Query(`SELECT i.id, i.name FROM items i INNER JOIN player_items pi ON i.id = pi.item_id WHERE pi.player_id = $1`, playerID)
	if err != nil {
		http.Error(w, "error fetching unlocked items", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var items []map[string]interface{}
	for rows.Next() {
		var id int
		var name string
		if err := rows.Scan(&id, &name); err != nil {
			log.Fatal(err)
		}
		items = append(items, map[string]interface{} {
			"id": id,
			"name": name,
		})
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	for _, item := range items {
		fmt.Printf("Item ID: %d, Item Name: %s\n", item["id"], item["name"])
	}
}


func createRoom() string {
	roomID := uuid.New().String()
	rooms[roomID] = &GameRoom{
		ID: roomID,
		Players: make(map[*websocket.Conn]*Player),
	}
	fmt.Printf("created game room with ID: %s", roomID)
	return roomID
}

func createRoomEndpoint(w http.ResponseWriter, r *http.Request) {
	roomID := createRoom()
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(roomID))
	//TODO return user information ie. id => query database
}

func joinRoom(roomID string, ws *websocket.Conn, playerID int) error {

	room, exists := rooms[roomID]

	if !exists {
		return fmt.Errorf("room does not exist")
	}

	if len(room.Players) >= 4 {
		return fmt.Errorf("sorry, this room's full")
	}
	var player Player
	err := db.QueryRow(`SELECT id, username FROM players WHERE id = $1`, playerID).Scan(&player.ID, &player.Username)
	if err != nil {
		return fmt.Errorf("could not retrieve player info: %v", err)
	}

	room.Players[ws] = &player
	notifyPlayers(room)
	return nil
}

func leaveRoom(roomID string, ws *websocket.Conn) {

	room, exists := rooms[roomID]

	if !exists {
		return
	}


	delete(room.Players, ws)
	if len(room.Players) == 0 {

		delete(rooms, roomID)

	} else {
		notifyPlayers(room)
	}
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Printf("could not upgrade to websocket: %v\n", err)
		return
	}
	defer func() {
		fmt.Println("closing websocket connection")
		ws.Close()
	}()

	playerIDstr := r.URL.Query().Get("playerId")
	roomID := r.URL.Query().Get("roomID")
	username := r.URL.Query().Get("username")
	if playerIDstr == "" || roomID == "" || username == "" {
		ws.WriteMessage(websocket.TextMessage, []byte("Error: missing playerID, roomID, or username"))
		return
	}

	playerID, err := strconv.Atoi(playerIDstr)
	if err != nil {
		ws.WriteMessage(websocket.TextMessage, []byte("Error: invalid playerID"))
		return
	}

	player := &Player{ID: playerID, Username: username}

	room, exists := rooms[roomID]
	if !exists {
		room = &GameRoom{ID: roomID, Players: make(map[*websocket.Conn]*Player)}
		rooms[roomID] = room
	}
	room.Players[ws] = player
	notifyPlayers(room)

	//read initial message
	_, msg, err := ws.ReadMessage()
	if err != nil {
		fmt.Printf("error reading initial message: %v\n", err)
		log.Printf("error reading initial message: %v\n", err)
		return
	}

	//expecting message to say create to make a room or join
	message := string(msg)
	if message == "create" {
		roomID = createRoom()
		fmt.Printf("game room created: %s\n", roomID)
		ws.WriteMessage(websocket.TextMessage, []byte(roomID))
	} else if strings.HasPrefix(message, "join:") {
		roomID = strings.TrimPrefix(message, "join:")
		err := joinRoom(roomID, ws, playerID)
		if err != nil {
			ws.WriteMessage(websocket.TextMessage, []byte("Error: "+err.Error()))
			return
		}
		fmt.Printf("playerID: %s joined room: %s\n",playerIDstr, roomID)
	}

	for {
		_, msg, err := ws.ReadMessage()
		if err != nil {
			fmt.Printf("error when reading message: %v\n", err)
			leaveRoom(roomID, ws)
			break
		}
		fmt.Printf("%s sent: %s\n", ws.RemoteAddr(), string(msg))
		broadcastMessage(roomID, msg)
	}
}

func notifyPlayers(room *GameRoom) {

	playerList := make([]Player, 0, len(room.Players))
	for _, player := range room.Players {
		playerList = append(playerList, *player)
	}
	data, err := json.Marshal(playerList)
	if err != nil {
		log.Printf("Error marshaling player list: %v", err)
		return
	}
	for conn := range room.Players {
		err := conn.WriteMessage(websocket.TextMessage, data)
		if err != nil {
			log.Printf("websocket write error: %v", err)
			fmt.Printf("error sending player list: %v\n", err)
			conn.Close()
			delete(room.Players, conn)
		}
	}
}

func broadcastMessage(roomID string, msg []byte) {
	room, exists := rooms[roomID]
	if !exists {
		return
	}
	for conn := range room.Players {
		err := conn.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			log.Printf("websocket write error: %v", err)
			conn.Close()
			delete(room.Players, conn)
		}
	}
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	initDB()

	mux := http.NewServeMux()
	mux.HandleFunc("/ws", handleConnections)

	corsHandler := handlers.CORS(
		handlers.AllowedOrigins([]string {"http://localhost:3000"}),
		handlers.AllowedMethods([]string{"GET", "POST"}),
		handlers.AllowedHeaders([]string{"Content-Type"}),
		
	)(mux)

	mux.HandleFunc("/register", registerPlayer)
	mux.HandleFunc("/login", loginPlayer)
	mux.HandleFunc("/saveProgress", saveProgress)
	mux.HandleFunc("/loadProgress", loadProgress)
	mux.HandleFunc("/unlockItem", unlockItem)
	mux.HandleFunc("/getUnlockedItems", getUnlockedItems)
	mux.HandleFunc("/createRoom", createRoomEndpoint)

	fmt.Println("server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", corsHandler))
}