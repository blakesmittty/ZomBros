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
	}
	room.mutex.Lock()
	room.Players[ws] = player
	room.mutex.Unlock()

	sendPlayerList(room)

	for {
		_, msg, err := ws.ReadMessage()
		if err != nil {
			log.Println(err)
			room.mutex.Lock()
			delete(room.Players, ws)
			room.mutex.Unlock()
			sendPlayerList(room)
			break
		}
		if len(msg) == 13 {
			input, err := readPlayerInput(msg)
			if err != nil {
            	log.Println("Error reading player input:", err)
            	continue
        	}
			handlePlayerInput(player, input)
		} else {
			var data map[string]interface{}
			err := json.Unmarshal(msg, &data)
			if err != nil {
				log.Println("Error parsing websocket message:", err)
				continue
			}

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
			default: 
				log.Println("unknown message type:", data["type"])
			}
		}
		//handleMessage(room, ws, messageType, p)
	}
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

// possibly not needed currently
func handleMessage(room *GameRoom, conn *websocket.Conn, messageType int, payload []byte) {
    // Implement your game logic here
    // This could include updating player positions, handling attacks, etc.
    // You'll need to broadcast updates to all players in the room

    for playerConn := range room.Players {
        if err := playerConn.WriteMessage(messageType, payload); err != nil {
            log.Println(err)
            delete(room.Players, playerConn)
            playerConn.Close()
        }
    }
}

func createRoom(w http.ResponseWriter, r *http.Request) {
	roomID := uuid.New().String()
	newRoom := &GameRoom{
		ID: roomID,
		Players: make(map[*websocket.Conn]*Player),
	}
	rooms[roomID] = newRoom
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

/*
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
	fmt.Printf("After joining a room: Rooms: %v, rooms[roomid]: %v, rooms[roomid].players: %v\n\n", rooms, rooms[roomID], rooms[roomID].Players)
	notifyPlayers(room)
	return nil
}
*/
/*
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
	*/