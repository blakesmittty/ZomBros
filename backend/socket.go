package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)


var upgrader = websocket.Upgrader{
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true // allowing all requests
	},
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Printf("error upgrading to websocket: %v", err)
		return
	}
	defer ws.Close()

	roomID := r.URL.Query().Get("roomID")
	if roomID == "" {
		log.Println("Room ID not provided")
		return
	}

	roomsMutex.RLock()
	room, exists := rooms[roomID]
	roomsMutex.RUnlock()
	if !exists {
		log.Println("Room not found: ", roomID)
		return
	}

	playerIDstr := r.URL.Query().Get("playerID")
	playerID, err := strconv.Atoi(playerIDstr)
	if err != nil {
		log.Println("invalid username")
		return
	}
	log.Printf("Received connection request for roomID: %s, playerID: %s", roomID, playerIDstr)
	player := &Player{
		ID: playerID,
		Username: r.URL.Query().Get("username"),
		Position: Vector2D{X: 0, Y: 0},
		Health: 100,
	}
	room.mutex.Lock()
	room.Players[ws] = player
	room.mutex.Unlock()

	sendPlayerList(room)

	// message loop 
	for {
		_, msg, err := ws.ReadMessage()
		if err != nil {
			log.Println(err)
			room.mutex.Lock()
			delete(room.Players, ws)
			//room.removePlayer(ws)
			room.mutex.Unlock()
			sendPlayerList(room)
			break
		}

		var data map[string]interface{}
		err = json.Unmarshal(msg, &data)
		if err != nil {
		    protoErr := handleProto(msg, player, room)//readPlayerInput(msg) 
			if protoErr != nil {
				log.Println("Error reading proto:", protoErr)
				continue
			}
		} else {
			switch data["type"] {
			case "selectCharacter":
				log.Println("select character message:", data)
				character := data["character"].(string)
				isReady := data["isReady"].(bool)
				player.Character = character
				player.IsReady = isReady
				room.mutex.Lock()
				room.Players[ws] = player
				room.mutex.Unlock()
				sendPlayerList(room)
				checkAllPlayersReady(room)
			
			default: 
				log.Println("unknown message type:", data["type"])
			}
		}
	}
}

func checkAllPlayersReady(room *GameRoom) {
	//room.mutex.Lock()
	//defer room.mutex.Unlock()
	for _, player := range room.Players {
		if !player.IsReady {
			return
		}
	}
	go gameLoop(room)
}

func (gr *GameRoom)  getPlayerList() []*Player {
	playerList := make([]*Player, 0, len(gr.Players))
	for _, player := range gr.Players {
		if player.Username != "" {
			playerList = append(playerList, player)
		}
	}
	return playerList
}

func sendPlayerList(room *GameRoom) {
	room.mutex.RLock()
	playerList := room.getPlayerList()
	room.mutex.RUnlock()
	message := map[string]interface{}{
		"type": "playerList",
		"playerList": playerList,
	}
	messageJSON, _ := json.Marshal(message)
	room.mutex.RLock()
	for playerConn := range room.Players {
		if err := playerConn.WriteMessage(websocket.TextMessage, messageJSON); err != nil {
			log.Println(err)
			delete(room.Players, playerConn)
			playerConn.Close()
		}
	}
	room.mutex.RUnlock()
}

func createRoom(w http.ResponseWriter, r *http.Request) {
	roomID := uuid.New().String()
	newRoom := &GameRoom{
		ID: roomID,
		Players: make(map[*websocket.Conn]*Player),
		Zombies: make(map[int]*Zombie),
		Map: GameMap{Name: "default_map"},
		Round: 0,
		MaxZombiesPerRound: 30,
	}
	roomsMutex.Lock()
	rooms[roomID] = newRoom
	roomsMutex.Unlock()

	//go gameLoop(newRoom)
	log.Printf("Created new room: %s", roomID)

	json.NewEncoder(w).Encode(map[string]string{"roomID" : roomID})
}

func joinRoom(w http.ResponseWriter, r *http.Request) {
	roomID := r.URL.Query().Get("roomID")

	roomsMutex.Lock()
	room, exists := rooms[roomID]
	roomsMutex.Unlock()
	if exists {
		if len(room.Players) >= 4 {
			json.NewEncoder(w).Encode(map[string]interface{}{
				"success": false,
				"message": "Room is full",
			})
			return
		}
		json.NewEncoder(w).Encode(map[string]interface{}{
			"success": true,
			"roomID": roomID,
			"playerCount": len(room.Players),
		})
	} else {
		json.NewEncoder(w).Encode(map[string]interface{}{
			"success": false,
			"message": "Room not found",
		})
	}
}

func (room *GameRoom) removePlayer(ws *websocket.Conn) {
	room.mutex.Lock()
	defer room.mutex.Unlock()

	player, exists := room.Players[ws]
	if !exists {
		return
	}

	delete(room.Players, ws)

	room.broadcastPlayerDisconnect(player)

	ws.Close()
}

func (room *GameRoom) broadcastPlayerDisconnect(player *Player) {
	for ws := range room.Players {
		err := ws.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("Player %d disconnected", player.Username)))
		if err != nil {
			log.Printf("Error sending disconnect message: %v", err)
			ws.Close()
			delete(room.Players, ws)
		}
	}
}