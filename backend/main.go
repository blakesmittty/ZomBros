package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/websocket"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

type GameRoom struct {
	ID string
	Players map[*websocket.Conn]*Player
	mutex sync.RWMutex
	Zombies map[int]*Zombie
	Bullets []*Bullet
	Map GameMap
	DifficultyMultiplier float64
	Round int
	MaxZombiesPerRound int
	MaxZombiesOnScreen int
	ZombieSpawnRate time.Duration
}

type Player struct {
	ID int `json:"id"`
	Username string `json:"username"`
	Position Vector2D
	Health int
	Inventory []Item
	IsReady bool `json:"isReady"`
	Character string `json:"character"`
	lastUpdate time.Time
	AimAngle float32
	mutex sync.Mutex
	MoveX float32
	MoveY float32
}
var (
	rooms = make(map[string]*GameRoom)
	roomsMutex sync.RWMutex
)

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
		"username": username,
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
	mux.HandleFunc("/createRoom", createRoom)
	mux.HandleFunc("/joinRoom", joinRoom)

	fmt.Println("server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", corsHandler))
}