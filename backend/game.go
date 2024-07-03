package main

import (
	"encoding/binary"
	"errors"
	"math"
	"time"
)

type GameState struct {
	Players map[string]*Player 
	Zombies []*Zombie
	//Map *GameMap
	//Rules *GameRules
}

type PlayerInput struct {
	MoveX, MoveY float32
	IsShooting bool
	AimAngle float32
}

type Zombie struct {
	ID string
	Position Vector2D
	Health int
}

type Vector2D struct {
	X, Y float64
}

type Item struct {
	ID string `json:"id"`
	Name string `json:"name"`
}

func gameLoop(room *GameRoom) {
	ticker := time.NewTicker(time.Millisecond * 50)
	defer ticker.Stop()

	for {
		select {
		case <- ticker.C:
			updateGameState(room)
			broadcastGameState(room)
		}
	}
}

func updateGameState(room *GameRoom) {
	// update player positions
	// move zombies
	// check for collisions
	// apply game rules
}

func broadcastGameState(room *GameRoom) {
	// send updated game state to all players in the room
}

func handlePlayerInput(player *Player, input PlayerInput) {
	// updaet player state based on input
	// for example, move the player
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

func readPlayerInput(data []byte) (PlayerInput, error) {
	if len(data) != 13 {
		return PlayerInput{}, errors.New("invalid input data length")
	}

	return PlayerInput {
		MoveX: math.Float32frombits(binary.LittleEndian.Uint32(data[0:4])),
		MoveY: math.Float32frombits(binary.LittleEndian.Uint32(data[4:8])),
		IsShooting: data[8] != 0,
		AimAngle: math.Float32frombits(binary.LittleEndian.Uint32(data[9:13])),
	}, nil
}