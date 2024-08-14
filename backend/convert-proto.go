package main

import (
	pb "backend/gamestate"
	"errors"
	"log"
	"time"

	"google.golang.org/protobuf/proto"
)

func handleProto (data []byte, player *Player, room *GameRoom) (error) {
	var playerInput pb.PlayerInput
	now := time.Now()
	deltaTime := now.Sub(player.lastUpdate).Seconds()
	player.lastUpdate = now
	if err := proto.Unmarshal(data, &playerInput); err == nil {
		log.Println("received PlayerInput: ", playerInput)
		handlePlayerInput(player, *processPlayerInput(&playerInput), room, deltaTime)
		return nil
	}

	var shootEvent pb.ShootEvent
	if err := proto.Unmarshal(data, &shootEvent); err == nil {
		log.Println("received shoot event: ", shootEvent)
		handleShootEvent(player, room, &shootEvent)
		return nil
	}

	var gameState pb.GameState
	if err := proto.Unmarshal(data, &gameState); err == nil {
		log.Println("received gamestate:", gameState)
		convGameState := processGameState(&gameState)
		log.Println("converted gamestate:", convGameState)
		return nil
	}

	return errors.New("failed to unmarshal proto message")
}

func convertVector2D(protoVec *pb.Vector2D) Vector2D {
	return Vector2D{
		X: float64(protoVec.GetX()),
		Y: float64(protoVec.GetY()),
	}
}

func processPlayer(state *pb.Player) *Player {
	return &Player {
		Position: convertVector2D(state.Position),
		Health: int(state.GetHealth()),
		AimAngle: state.GetAimAngle(),
		MoveX: state.GetMoveX(),
		MoveY: state.GetMoveY(),
	}
}

func processPlayerInput(state *pb.PlayerInput) *PlayerInput {
	return &PlayerInput {
		MoveX: state.GetMoveX(),
		MoveY: state.GetMoveY(),
		IsShooting: state.GetIsShooting(),
		AimAngle: state.GetAimAngle(),
	}
}

func processZombie(state *pb.Zombie) *Zombie {
	return &Zombie{
		ID: int(state.Id),
		Position: convertVector2D(state.GetPosition()),
		Health: int(state.Health),
	}
}

func processMap(state *pb.GameMap) *GameMap {
	return &GameMap{
		Name: state.GetName(),
	}
}

func processGameState(state *pb.GameState) *GameState {
	gameState := &GameState{
		Players: make(map[string]*Player),
		Zombies: []*Zombie{},
		Map: processMap(state.Map),
	}

	for _, player := range state.GetPlayers() {
		gameState.Players[player.Username] = processPlayer(player)
	}

	for _, zombie := range state.GetZombies() {
		gameState.Zombies = append(gameState.Zombies, processZombie(zombie))
	}

	return gameState
}