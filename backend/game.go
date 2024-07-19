package main

import (
	pb "backend/gamestate"
	"log"
	"math"
	"time"

	"github.com/gorilla/websocket"
	"google.golang.org/protobuf/proto"
)

type GameState struct {
	Players map[string]*Player 
	Zombies []*Zombie
	Map *GameMap
	//Rules *GameRules
}

type GameMap struct {
	Name string
}

type PlayerInput struct {
	MoveX, MoveY float32
	IsShooting bool
	AimAngle float32
}

type Zombie struct {
	ID int
	Position Vector2D
	Health int
	Speed float64
}

type Vector2D struct {
	X, Y float64
}

type Item struct {
	ID string `json:"id"`
	Name string `json:"name"`
}

const PLAYER_SPEED = 100.0
const DEFAULT_ZOMBIE_SPEED = 100.0

func calculateDistance(a, b Vector2D) float64 {
	distX := a.X - b.X
	distY := a.Y - b.Y
	return math.Sqrt((distX * distX) + (distY * distY))
}

func gameLoop(room *GameRoom) {
	ticker := time.NewTicker(time.Millisecond * 50)
	defer ticker.Stop()

	var lastUpdate time.Time
	
	for {
		select {
		case <-ticker.C:
			now := time.Now()
			if !lastUpdate.IsZero() {
				deltaTime := now.Sub(lastUpdate).Seconds()
				//room.mutex.Lock()
				updateGameState(room, deltaTime)
				broadcastGameState(room)
				//room.mutex.Unlock()
			}
			lastUpdate = now
		}
	}
}

func handleZombies(room *GameRoom) {
	for _, player := range room.Players {
		if len(room.Zombies) < len(room.Players) {
			newZombie := &Zombie{
				ID: len(room.Zombies),
				Position: Vector2D{X: player.Position.X + 50, Y: player.Position.Y + 50},
				Health: 100,
				Speed: 1.0,
			}
			room.Zombies[newZombie.ID] = newZombie
		}
	}
}

func updateGameState(room *GameRoom, deltaTime float64) {
	// update player positions
	// move zombies
	// check for collisions
	// apply game rules

	room.mutex.Lock()
	defer room.mutex.Unlock()
	
	handleZombies(room)

	for _, zombie := range room.Zombies {
		if len(room.Players) > 0 {
			var closestPlayer *Player
			minDistance := math.MaxFloat64
			for _, player := range room.Players {
				distance := calculateDistance(zombie.Position, player.Position)
				if distance < minDistance {
					minDistance = distance
					closestPlayer = player
				}
			}
			if closestPlayer != nil {
				moveTowards(zombie, closestPlayer.Position, deltaTime)
				log.Printf("Zombie %d moved towards player at position (%.2f, %.2f)", zombie.ID, closestPlayer.Position.X, closestPlayer.Position.Y)
			}
		}
	}
}

func moveTowards(zombie *Zombie, target Vector2D, deltaTime float64) {
	distX := target.X - zombie.Position.X
	distY := target.Y - zombie.Position.Y
	distance := calculateDistance(zombie.Position, target)
	if distance > 0 {
		zombie.Position.X += (distX / distance) * DEFAULT_ZOMBIE_SPEED * deltaTime
		zombie.Position.Y += (distY / distance) * DEFAULT_ZOMBIE_SPEED * deltaTime
	}
}

func broadcastGameState(room *GameRoom) {
	// send updated game state to all players in the room

	room.mutex.RLock()
	defer room.mutex.RUnlock()

	log.Println("current players: ", room.Players)
	//log.Println("current zombies: ", room.Zombies)
	log.Println("current map: ", room.Map)

	gameState := &pb.GameState{
		Players: []*pb.Player{},
		Zombies: []*pb.Zombie{},
		Map: &pb.GameMap{Name: room.Map.Name},
	}

	for _, player := range room.Players {
		gameState.Players = append(gameState.Players, &pb.Player{
			Id: int32(player.ID),
			Username: player.Username,
			Position: &pb.Vector2D{X: float32(player.Position.X), Y: float32(player.Position.Y)},
			Health: int32(player.Health),
			Character: player.Character,
			IsReady: player.IsReady,
			AimAngle: player.AimAngle,
		})
	}

	
	for _, zombie := range room.Zombies {
		gameState.Zombies = append(gameState.Zombies, &pb.Zombie{
			Id: int32(zombie.ID),
			Position: &pb.Vector2D{X: float32(zombie.Position.X), Y: float32(zombie.Position.Y)},
			Health: int32(zombie.Health),
		})
	}
	
	log.Println("populated game state: ", gameState)

	data, err := proto.Marshal(gameState)
	if err != nil {
		log.Println("Error marshaling game state: ", err)
		return
	}

	for ws := range room.Players {
		if err := ws.WriteMessage(websocket.BinaryMessage, data); err != nil {
			log.Println("Error sending game state: ", err)
		}
		log.Println("sent gamestate: (proto) ", data)
		log.Println("sent gamestate: (unmarshal)", gameState)
	}
}

func handlePlayerInput(player *Player, input PlayerInput, deltaTime float64) {
	// updaet player state based on input
	// for example, move the player

	player.mutex.Lock()
	defer player.mutex.Unlock()

	// update player state based on input
	player.Position.X += float64(input.MoveX) * PLAYER_SPEED * deltaTime
	player.Position.Y += float64(input.MoveY) * PLAYER_SPEED * deltaTime

	/*
	if input.IsShooting {
		// handle shooting logic
	}
		*/
	player.AimAngle = input.AimAngle

	log.Printf("player %d input: %+v", player.ID, input)
}

/*
func updateZombies(zombies []*Zombie, players map[string]*Player) {
	for _, zombie := range zombies {
		nearestPlayer := findNearestPlayer(zombie, players)
		moveTowards(zombie, nearestPlayer)
	}
}

func checkCollisions(players map[string]*Player, zombies []*Zombie) {
	for _, player := range players {
		for _, zombie := range zombies {
			if isColliding(player, zombie) {
				handlePlayerZombieCollision(player, zombie)
			}
		}
	}
}
*/
func handleGameEvents(room *GameRoom) {
	// check for player damage
	// handle item pickups
	// spawn new zombies
	// check win/loss conditions?
}
