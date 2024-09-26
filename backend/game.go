package main

import (
	pb "backend/gamestate"
	"fmt"
	"log"
	"math"
	"math/rand"
	"time"

	"github.com/gorilla/websocket"
	"google.golang.org/protobuf/proto"
)

type GameState struct {
	Players map[string]*Player 
	Zombies []*Zombie
	Map *GameMap
	Bullets []*Bullet
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

type Bullet struct {
	ID int
	Position Vector2D
	Direction float32
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
const DEFAULT_ZOMBIE_SPEED = 90.0
const DEFAULT_BULLET_SPEED = 300.0

const (
	MAP_MIN_X = 0
	MAP_MAX_X = 3000
	MAP_MIN_Y = 0
	MAP_MAX_Y = 3000
)

var zombieSpawns = []Vector2D{
	{X: 600, Y: 600},
	{X: 600, Y: 1200},
	{X: 600, Y: 1800},
	{X: 600, Y: 2400},
	{X: 1200, Y: 600},
	{X: 1200, Y: 1200},
	{X: 1200, Y: 1800},
	{X: 1200, Y: 2400},
	{X: 1800, Y: 600},
	{X: 1800, Y: 1200},
	{X: 1800, Y: 1800},
	{X: 1800, Y: 2400},
	{X: 2400, Y: 600},
	{X: 2400, Y: 1200},
	{X: 2400, Y: 1800},
	{X: 2400, Y: 2400},
}



func calculateDistance(a, b Vector2D) float64 {
	distX := a.X - b.X
	distY := a.Y - b.Y
	return math.Sqrt((distX * distX) + (distY * distY))
}

func generateBulletID() int {
	return int(time.Now().UnixNano())
}

func generateZombieID() int {
	return rand.Intn(10000000)
}
//broken
func (room *GameRoom) updateDifficulty() {
	room.DifficultyMultiplier = 1 + (float64(room.Round) * 0.1)
	room.ZombieSpawnRate = time.Duration(math.Max(1000, float64(5000 - (room.Round * 200)))) * time.Millisecond
	room.MaxZombiesPerRound = 30 + (room.Round * 4)
}
//broken
func startRound(room *GameRoom) {
	//room.mutex.Lock()
	//defer room.mutex.Unlock()
	//room.mutex.Lock()
	room.updateDifficulty()
	//room.mutex.Unlock()
	for i := 0; i < room.MaxZombiesPerRound; i++ {
		//spawnDelay := time.Duration(float64(room.ZombieSpawnRate) / room.DifficultyMultiplier)
		//time.Sleep(3 * time.Second)
		//room.mutex.Lock()
		spawnZombie(room)
		//room.mutex.Unlock()
		log.Printf("zombie spawned: %d\n", i+1)
	}
	log.Println("round started with difficulty multiplier: ", room.DifficultyMultiplier)
}

func getRandomZombieSpawnOutsideMap() Vector2D {
	edge := rand.Intn(4)
	var pos Vector2D
	switch edge {
	case 0:
		pos = Vector2D{X: float64(rand.Intn(MAP_MAX_X)), Y: float64(MAP_MIN_Y - 100)}
	case 1: 
		pos = Vector2D{X: float64(MAP_MAX_X + 100), Y: float64(rand.Intn(MAP_MAX_Y))}
	case 2: 
		pos = Vector2D{X: float64(rand.Intn(MAP_MAX_X)), Y: float64(MAP_MAX_Y + 100)}
	case 3: 
		pos = Vector2D{X: float64(MAP_MIN_X - 100), Y: float64(rand.Intn(MAP_MAX_Y))}
	}
	return pos
}


func getClosestZombieSpawn(room *GameRoom) Vector2D {
	// make better
	//room.mutex.Lock()
	//defer room.mutex.Unlock()
	minDistance := math.MaxFloat64
	spawnPoint := 0
	for _, player := range room.Players {
		for i, point := range zombieSpawns {
			dist := calculateDistance(player.Position, point)
			if dist < minDistance {
				spawnPoint = i
				minDistance = dist
			}
		}
	}
	
	return zombieSpawns[spawnPoint]
}
//broken
func spawnZombie(room *GameRoom) {
	newZombie := &Zombie{
		ID: generateZombieID(),
		//Position: getClosestZombieSpawn(room),
		Position: getRandomZombieSpawnOutsideMap(),
		Health: 100,
		Speed: DEFAULT_ZOMBIE_SPEED,
	}

	//room.mutex.Lock()
	room.Zombies[newZombie.ID] = newZombie
	//room.mutex.Unlock()
	log.Printf("Spawned zombie with id %d at position: (%f, %f)\n", newZombie.ID, newZombie.Position.X, newZombie.Position.Y)
}

func handleShootEvent(player *Player, room *GameRoom, shootEvent *pb.ShootEvent) {
	// determine weapon type here

	room.mutex.Lock()
	defer room.mutex.Unlock()

	//log.Printf("shoot event for player %d at pos (%.2f, %.2f) direction: %.2f", player.ID, player.Position.X, player.Position.Y, shootEvent.Direction)

	if shootEvent == nil {
		log.Println("shoot event nil")
		return
	}

	bullet := &Bullet{
		ID: generateBulletID(),
		Position: player.Position,
		Direction: shootEvent.Direction,
		Speed: DEFAULT_BULLET_SPEED,
	}


	//log.Printf("generated bullet id %d at pos(%.2f, %.2f) direction: %.2f", bullet.ID, bullet.Position.X, bullet.Position.Y, bullet.Direction)

	room.Bullets = append(room.Bullets, bullet)
	//log.Printf("bullet added to room. total bullets: %d", len(room.Bullets))
}

func updateBulletPosition(bullet *Bullet, deltaTime float64) {
	bullet.Position.X += math.Cos(float64(bullet.Direction)) * bullet.Speed * deltaTime
	bullet.Position.Y += math.Sin(float64(bullet.Direction)) * bullet.Speed * deltaTime
}

func filterActiveBullets(bullets []*Bullet) []*Bullet {
	var activeBullets []*Bullet
	for _, bullet := range bullets {
		if !bulletOutOfBounds(bullet) {
			activeBullets = append(activeBullets, bullet)
		} else {
			log.Printf("bullet out of bounds at position: (%f, %f) \n", bullet.Position.X, bullet.Position.Y)
		}
	}
	return activeBullets
}

func bulletOutOfBounds(bullet *Bullet) bool {
	return bullet.Position.X < 0 || bullet.Position.X > 3000 || bullet.Position.Y < 0 || bullet.Position.Y > 3000
}

func detectCollisions(room *GameRoom) {
	log.Printf("in detect collisions")
	room.mutex.Lock()
	defer room.mutex.Unlock()
	var remainingBullets []*Bullet 
	for _, bullet := range room.Bullets {
		bulletCollided := false
		for _, zombie := range room.Zombies {
			if calculateDistance(bullet.Position, zombie.Position) < 25.0 {
				zombie.Health -= 50
				if zombie.Health <= 0 {
					log.Printf("zombie id: %d is dead and will be deleted", zombie.ID)
					delete(room.Zombies, zombie.ID)
				}
				bulletCollided = true
				break
			}
		}
		if !bulletCollided {
			remainingBullets = append(remainingBullets, bullet)
		}
	}
	room.Bullets = remainingBullets
}

func gameLoop(room *GameRoom) {
	ticker := time.NewTicker(time.Millisecond * 50)
	defer ticker.Stop()

	var lastUpdate time.Time
	//room.Round = 1

	for {
		select {
		case <-ticker.C:
			now := time.Now()
			if !lastUpdate.IsZero() {
				deltaTime := now.Sub(lastUpdate).Seconds()
		
				updateGameState(room, deltaTime)
				detectCollisions(room)
				broadcastGameState(room)
		
			}
			lastUpdate = now
		}
	}
}

func updateGameState(room *GameRoom, deltaTime float64) {
	room.mutex.Lock()
	defer room.mutex.Unlock()
	//log.Println("current zombies: ", len(room.Zombies))
	//log.Println("zombies array: ", room.Zombies)
	//handleZombies(room, 50)
	
	if len(room.Zombies) == 0 {
		log.Printf("room round", room.Round)
		room.Round++
		startRound(room)
	}
	
	for _, bullet := range room.Bullets {
		updateBulletPosition(bullet, deltaTime)
	}

	room.Bullets = filterActiveBullets(room.Bullets)

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
		Bullets: []*pb.Bullet{},
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
			MoveX: player.MoveX,
			MoveY: player.MoveY,
		})
	}

	
	for _, zombie := range room.Zombies {
		gameState.Zombies = append(gameState.Zombies, &pb.Zombie{
			Id: int32(zombie.ID),
			Position: &pb.Vector2D{X: float32(zombie.Position.X), Y: float32(zombie.Position.Y)},
			Health: int32(zombie.Health),
		})
	}

	for _, bullet := range room.Bullets {
		gameState.Bullets = append(gameState.Bullets, &pb.Bullet{
			Id: fmt.Sprintf("%d", bullet.ID),
			Position: &pb.Vector2D{X: float32(bullet.Position.X), Y: float32(bullet.Position.Y)},
			Direction: bullet.Direction,
			Speed: float32(bullet.Speed),
		})
	}
	
	//log.Println("populated game state: ", gameState)

	data, err := proto.Marshal(gameState)
	if err != nil {
		log.Println("Error marshaling game state: ", err)
		return
	}

	for ws := range room.Players {
		if err := ws.WriteMessage(websocket.BinaryMessage, data); err != nil {
			log.Println("Error sending game state: ", err)
		}
		//log.Println("sent gamestate: (proto) ", data)
		//log.Println("sent gamestate: (unmarshal)", gameState)
	}
}

func handlePlayerInput(player *Player, input PlayerInput, room *GameRoom, deltaTime float64) {
	// updaet player state based on input
	// for example, move the player

	player.mutex.Lock()
	defer player.mutex.Unlock()

	// update player state based on input
	log.Println("player input: ", input.MoveX, input.MoveY)
	player.Position.X += float64(input.MoveX) * PLAYER_SPEED * deltaTime
	player.Position.Y += float64(input.MoveY) * PLAYER_SPEED * deltaTime

	player.AimAngle = input.AimAngle
	player.MoveX = input.MoveX
	player.MoveY = input.MoveY

	log.Printf("player %d input: %+v", player.ID, input)
	
	if input.IsShooting {
		handleShootEvent(player, room, &pb.ShootEvent{
			PlayerId: int32(player.ID),
			Direction: input.AimAngle,
		})
	}
			
}

func handleGameEvents(room *GameRoom) {
	// check for player damage
	// handle item pickups
	// spawn new zombies
	// check win/loss conditions?
}
