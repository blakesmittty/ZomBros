// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.19.4
// source: gamestate.proto

package gamestate

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Vector2D struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	X float32 `protobuf:"fixed32,1,opt,name=x,proto3" json:"x,omitempty"`
	Y float32 `protobuf:"fixed32,2,opt,name=y,proto3" json:"y,omitempty"`
}

func (x *Vector2D) Reset() {
	*x = Vector2D{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gamestate_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Vector2D) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Vector2D) ProtoMessage() {}

func (x *Vector2D) ProtoReflect() protoreflect.Message {
	mi := &file_gamestate_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Vector2D.ProtoReflect.Descriptor instead.
func (*Vector2D) Descriptor() ([]byte, []int) {
	return file_gamestate_proto_rawDescGZIP(), []int{0}
}

func (x *Vector2D) GetX() float32 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *Vector2D) GetY() float32 {
	if x != nil {
		return x.Y
	}
	return 0
}

type Player struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int32     `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Username  string    `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Position  *Vector2D `protobuf:"bytes,3,opt,name=position,proto3" json:"position,omitempty"`
	Health    int32     `protobuf:"varint,4,opt,name=health,proto3" json:"health,omitempty"`
	Character string    `protobuf:"bytes,5,opt,name=character,proto3" json:"character,omitempty"`
	IsReady   bool      `protobuf:"varint,6,opt,name=isReady,proto3" json:"isReady,omitempty"`
	AimAngle  float32   `protobuf:"fixed32,7,opt,name=aimAngle,proto3" json:"aimAngle,omitempty"`
	Inventory []*Item   `protobuf:"bytes,8,rep,name=inventory,proto3" json:"inventory,omitempty"`
	MoveX     float32   `protobuf:"fixed32,9,opt,name=move_x,json=moveX,proto3" json:"move_x,omitempty"`
	MoveY     float32   `protobuf:"fixed32,10,opt,name=move_y,json=moveY,proto3" json:"move_y,omitempty"`
}

func (x *Player) Reset() {
	*x = Player{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gamestate_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Player) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Player) ProtoMessage() {}

func (x *Player) ProtoReflect() protoreflect.Message {
	mi := &file_gamestate_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Player.ProtoReflect.Descriptor instead.
func (*Player) Descriptor() ([]byte, []int) {
	return file_gamestate_proto_rawDescGZIP(), []int{1}
}

func (x *Player) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Player) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *Player) GetPosition() *Vector2D {
	if x != nil {
		return x.Position
	}
	return nil
}

func (x *Player) GetHealth() int32 {
	if x != nil {
		return x.Health
	}
	return 0
}

func (x *Player) GetCharacter() string {
	if x != nil {
		return x.Character
	}
	return ""
}

func (x *Player) GetIsReady() bool {
	if x != nil {
		return x.IsReady
	}
	return false
}

func (x *Player) GetAimAngle() float32 {
	if x != nil {
		return x.AimAngle
	}
	return 0
}

func (x *Player) GetInventory() []*Item {
	if x != nil {
		return x.Inventory
	}
	return nil
}

func (x *Player) GetMoveX() float32 {
	if x != nil {
		return x.MoveX
	}
	return 0
}

func (x *Player) GetMoveY() float32 {
	if x != nil {
		return x.MoveY
	}
	return 0
}

type ShootEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayerId  int32   `protobuf:"varint,1,opt,name=playerId,proto3" json:"playerId,omitempty"`
	Direction float32 `protobuf:"fixed32,2,opt,name=direction,proto3" json:"direction,omitempty"`
	Timestamp int64   `protobuf:"varint,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Weapon    string  `protobuf:"bytes,4,opt,name=weapon,proto3" json:"weapon,omitempty"`
}

func (x *ShootEvent) Reset() {
	*x = ShootEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gamestate_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShootEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShootEvent) ProtoMessage() {}

func (x *ShootEvent) ProtoReflect() protoreflect.Message {
	mi := &file_gamestate_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShootEvent.ProtoReflect.Descriptor instead.
func (*ShootEvent) Descriptor() ([]byte, []int) {
	return file_gamestate_proto_rawDescGZIP(), []int{2}
}

func (x *ShootEvent) GetPlayerId() int32 {
	if x != nil {
		return x.PlayerId
	}
	return 0
}

func (x *ShootEvent) GetDirection() float32 {
	if x != nil {
		return x.Direction
	}
	return 0
}

func (x *ShootEvent) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *ShootEvent) GetWeapon() string {
	if x != nil {
		return x.Weapon
	}
	return ""
}

type Bullet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string    `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Position   *Vector2D `protobuf:"bytes,2,opt,name=position,proto3" json:"position,omitempty"`
	Direction  float32   `protobuf:"fixed32,3,opt,name=direction,proto3" json:"direction,omitempty"`
	Speed      float32   `protobuf:"fixed32,4,opt,name=speed,proto3" json:"speed,omitempty"`
	Projectile string    `protobuf:"bytes,5,opt,name=projectile,proto3" json:"projectile,omitempty"`
}

func (x *Bullet) Reset() {
	*x = Bullet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gamestate_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Bullet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Bullet) ProtoMessage() {}

func (x *Bullet) ProtoReflect() protoreflect.Message {
	mi := &file_gamestate_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Bullet.ProtoReflect.Descriptor instead.
func (*Bullet) Descriptor() ([]byte, []int) {
	return file_gamestate_proto_rawDescGZIP(), []int{3}
}

func (x *Bullet) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Bullet) GetPosition() *Vector2D {
	if x != nil {
		return x.Position
	}
	return nil
}

func (x *Bullet) GetDirection() float32 {
	if x != nil {
		return x.Direction
	}
	return 0
}

func (x *Bullet) GetSpeed() float32 {
	if x != nil {
		return x.Speed
	}
	return 0
}

func (x *Bullet) GetProjectile() string {
	if x != nil {
		return x.Projectile
	}
	return ""
}

type Item struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name   string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Id     int32  `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Amount int32  `protobuf:"varint,3,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *Item) Reset() {
	*x = Item{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gamestate_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Item) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Item) ProtoMessage() {}

func (x *Item) ProtoReflect() protoreflect.Message {
	mi := &file_gamestate_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Item.ProtoReflect.Descriptor instead.
func (*Item) Descriptor() ([]byte, []int) {
	return file_gamestate_proto_rawDescGZIP(), []int{4}
}

func (x *Item) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Item) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Item) GetAmount() int32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

type PlayerInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MoveX      float32 `protobuf:"fixed32,1,opt,name=move_x,json=moveX,proto3" json:"move_x,omitempty"`
	MoveY      float32 `protobuf:"fixed32,2,opt,name=move_y,json=moveY,proto3" json:"move_y,omitempty"`
	IsShooting bool    `protobuf:"varint,3,opt,name=is_shooting,json=isShooting,proto3" json:"is_shooting,omitempty"`
	AimAngle   float32 `protobuf:"fixed32,4,opt,name=aim_angle,json=aimAngle,proto3" json:"aim_angle,omitempty"`
}

func (x *PlayerInput) Reset() {
	*x = PlayerInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gamestate_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayerInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayerInput) ProtoMessage() {}

func (x *PlayerInput) ProtoReflect() protoreflect.Message {
	mi := &file_gamestate_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayerInput.ProtoReflect.Descriptor instead.
func (*PlayerInput) Descriptor() ([]byte, []int) {
	return file_gamestate_proto_rawDescGZIP(), []int{5}
}

func (x *PlayerInput) GetMoveX() float32 {
	if x != nil {
		return x.MoveX
	}
	return 0
}

func (x *PlayerInput) GetMoveY() float32 {
	if x != nil {
		return x.MoveY
	}
	return 0
}

func (x *PlayerInput) GetIsShooting() bool {
	if x != nil {
		return x.IsShooting
	}
	return false
}

func (x *PlayerInput) GetAimAngle() float32 {
	if x != nil {
		return x.AimAngle
	}
	return 0
}

type Zombie struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int32     `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Position *Vector2D `protobuf:"bytes,2,opt,name=position,proto3" json:"position,omitempty"`
	Health   int32     `protobuf:"varint,3,opt,name=health,proto3" json:"health,omitempty"`
}

func (x *Zombie) Reset() {
	*x = Zombie{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gamestate_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Zombie) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Zombie) ProtoMessage() {}

func (x *Zombie) ProtoReflect() protoreflect.Message {
	mi := &file_gamestate_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Zombie.ProtoReflect.Descriptor instead.
func (*Zombie) Descriptor() ([]byte, []int) {
	return file_gamestate_proto_rawDescGZIP(), []int{6}
}

func (x *Zombie) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Zombie) GetPosition() *Vector2D {
	if x != nil {
		return x.Position
	}
	return nil
}

func (x *Zombie) GetHealth() int32 {
	if x != nil {
		return x.Health
	}
	return 0
}

type GameMap struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GameMap) Reset() {
	*x = GameMap{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gamestate_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GameMap) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GameMap) ProtoMessage() {}

func (x *GameMap) ProtoReflect() protoreflect.Message {
	mi := &file_gamestate_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GameMap.ProtoReflect.Descriptor instead.
func (*GameMap) Descriptor() ([]byte, []int) {
	return file_gamestate_proto_rawDescGZIP(), []int{7}
}

func (x *GameMap) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type GameState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Players []*Player `protobuf:"bytes,1,rep,name=players,proto3" json:"players,omitempty"`
	Zombies []*Zombie `protobuf:"bytes,2,rep,name=zombies,proto3" json:"zombies,omitempty"`
	Bullets []*Bullet `protobuf:"bytes,3,rep,name=bullets,proto3" json:"bullets,omitempty"`
	Map     *GameMap  `protobuf:"bytes,4,opt,name=map,proto3" json:"map,omitempty"`
}

func (x *GameState) Reset() {
	*x = GameState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gamestate_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GameState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GameState) ProtoMessage() {}

func (x *GameState) ProtoReflect() protoreflect.Message {
	mi := &file_gamestate_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GameState.ProtoReflect.Descriptor instead.
func (*GameState) Descriptor() ([]byte, []int) {
	return file_gamestate_proto_rawDescGZIP(), []int{8}
}

func (x *GameState) GetPlayers() []*Player {
	if x != nil {
		return x.Players
	}
	return nil
}

func (x *GameState) GetZombies() []*Zombie {
	if x != nil {
		return x.Zombies
	}
	return nil
}

func (x *GameState) GetBullets() []*Bullet {
	if x != nil {
		return x.Bullets
	}
	return nil
}

func (x *GameState) GetMap() *GameMap {
	if x != nil {
		return x.Map
	}
	return nil
}

var File_gamestate_proto protoreflect.FileDescriptor

var file_gamestate_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x67, 0x61, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x09, 0x67, 0x61, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x74, 0x65, 0x22, 0x26, 0x0a, 0x08,
	0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x32, 0x44, 0x12, 0x0c, 0x0a, 0x01, 0x78, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x02, 0x52, 0x01, 0x78, 0x12, 0x0c, 0x0a, 0x01, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x01, 0x79, 0x22, 0xae, 0x02, 0x0a, 0x06, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2f, 0x0a, 0x08, 0x70,
	0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e,
	0x67, 0x61, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72,
	0x32, 0x44, 0x52, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06,
	0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x68, 0x65,
	0x61, 0x6c, 0x74, 0x68, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65,
	0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74,
	0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x69, 0x73, 0x52, 0x65, 0x61, 0x64, 0x79, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x07, 0x69, 0x73, 0x52, 0x65, 0x61, 0x64, 0x79, 0x12, 0x1a, 0x0a, 0x08,
	0x61, 0x69, 0x6d, 0x41, 0x6e, 0x67, 0x6c, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x02, 0x52, 0x08,
	0x61, 0x69, 0x6d, 0x41, 0x6e, 0x67, 0x6c, 0x65, 0x12, 0x2d, 0x0a, 0x09, 0x69, 0x6e, 0x76, 0x65,
	0x6e, 0x74, 0x6f, 0x72, 0x79, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x67, 0x61,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x09, 0x69, 0x6e,
	0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x15, 0x0a, 0x06, 0x6d, 0x6f, 0x76, 0x65, 0x5f,
	0x78, 0x18, 0x09, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x6d, 0x6f, 0x76, 0x65, 0x58, 0x12, 0x15,
	0x0a, 0x06, 0x6d, 0x6f, 0x76, 0x65, 0x5f, 0x79, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05,
	0x6d, 0x6f, 0x76, 0x65, 0x59, 0x22, 0x7c, 0x0a, 0x0a, 0x53, 0x68, 0x6f, 0x6f, 0x74, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x1c, 0x0a, 0x09, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x09, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a,
	0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x16, 0x0a, 0x06, 0x77,
	0x65, 0x61, 0x70, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x77, 0x65, 0x61,
	0x70, 0x6f, 0x6e, 0x22, 0x9d, 0x01, 0x0a, 0x06, 0x42, 0x75, 0x6c, 0x6c, 0x65, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x2f,
	0x0a, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x13, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x56, 0x65, 0x63,
	0x74, 0x6f, 0x72, 0x32, 0x44, 0x52, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x1c, 0x0a, 0x09, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x09, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a,
	0x05, 0x73, 0x70, 0x65, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x73, 0x70,
	0x65, 0x65, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x69, 0x6c,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x69, 0x6c, 0x65, 0x22, 0x42, 0x0a, 0x04, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x79, 0x0a, 0x0b, 0x50, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x15, 0x0a, 0x06, 0x6d, 0x6f, 0x76, 0x65, 0x5f, 0x78,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x6d, 0x6f, 0x76, 0x65, 0x58, 0x12, 0x15, 0x0a,
	0x06, 0x6d, 0x6f, 0x76, 0x65, 0x5f, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x6d,
	0x6f, 0x76, 0x65, 0x59, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x73, 0x5f, 0x73, 0x68, 0x6f, 0x6f, 0x74,
	0x69, 0x6e, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x69, 0x73, 0x53, 0x68, 0x6f,
	0x6f, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x1b, 0x0a, 0x09, 0x61, 0x69, 0x6d, 0x5f, 0x61, 0x6e, 0x67,
	0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x08, 0x61, 0x69, 0x6d, 0x41, 0x6e, 0x67,
	0x6c, 0x65, 0x22, 0x61, 0x0a, 0x06, 0x5a, 0x6f, 0x6d, 0x62, 0x69, 0x65, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x2f, 0x0a, 0x08,
	0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13,
	0x2e, 0x67, 0x61, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x56, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x32, 0x44, 0x52, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a,
	0x06, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x68,
	0x65, 0x61, 0x6c, 0x74, 0x68, 0x22, 0x1d, 0x0a, 0x07, 0x47, 0x61, 0x6d, 0x65, 0x4d, 0x61, 0x70,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x22, 0xb8, 0x01, 0x0a, 0x09, 0x47, 0x61, 0x6d, 0x65, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x12, 0x2b, 0x0a, 0x07, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e,
	0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x07, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x12,
	0x2b, 0x0a, 0x07, 0x7a, 0x6f, 0x6d, 0x62, 0x69, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x11, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x5a, 0x6f, 0x6d,
	0x62, 0x69, 0x65, 0x52, 0x07, 0x7a, 0x6f, 0x6d, 0x62, 0x69, 0x65, 0x73, 0x12, 0x2b, 0x0a, 0x07,
	0x62, 0x75, 0x6c, 0x6c, 0x65, 0x74, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e,
	0x67, 0x61, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x42, 0x75, 0x6c, 0x6c, 0x65, 0x74,
	0x52, 0x07, 0x62, 0x75, 0x6c, 0x6c, 0x65, 0x74, 0x73, 0x12, 0x24, 0x0a, 0x03, 0x6d, 0x61, 0x70,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x2e, 0x47, 0x61, 0x6d, 0x65, 0x4d, 0x61, 0x70, 0x52, 0x03, 0x6d, 0x61, 0x70, 0x42,
	0x13, 0x5a, 0x11, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x67, 0x61, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_gamestate_proto_rawDescOnce sync.Once
	file_gamestate_proto_rawDescData = file_gamestate_proto_rawDesc
)

func file_gamestate_proto_rawDescGZIP() []byte {
	file_gamestate_proto_rawDescOnce.Do(func() {
		file_gamestate_proto_rawDescData = protoimpl.X.CompressGZIP(file_gamestate_proto_rawDescData)
	})
	return file_gamestate_proto_rawDescData
}

var file_gamestate_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_gamestate_proto_goTypes = []any{
	(*Vector2D)(nil),    // 0: gamestate.Vector2D
	(*Player)(nil),      // 1: gamestate.Player
	(*ShootEvent)(nil),  // 2: gamestate.ShootEvent
	(*Bullet)(nil),      // 3: gamestate.Bullet
	(*Item)(nil),        // 4: gamestate.Item
	(*PlayerInput)(nil), // 5: gamestate.PlayerInput
	(*Zombie)(nil),      // 6: gamestate.Zombie
	(*GameMap)(nil),     // 7: gamestate.GameMap
	(*GameState)(nil),   // 8: gamestate.GameState
}
var file_gamestate_proto_depIdxs = []int32{
	0, // 0: gamestate.Player.position:type_name -> gamestate.Vector2D
	4, // 1: gamestate.Player.inventory:type_name -> gamestate.Item
	0, // 2: gamestate.Bullet.position:type_name -> gamestate.Vector2D
	0, // 3: gamestate.Zombie.position:type_name -> gamestate.Vector2D
	1, // 4: gamestate.GameState.players:type_name -> gamestate.Player
	6, // 5: gamestate.GameState.zombies:type_name -> gamestate.Zombie
	3, // 6: gamestate.GameState.bullets:type_name -> gamestate.Bullet
	7, // 7: gamestate.GameState.map:type_name -> gamestate.GameMap
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_gamestate_proto_init() }
func file_gamestate_proto_init() {
	if File_gamestate_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_gamestate_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Vector2D); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_gamestate_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*Player); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_gamestate_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*ShootEvent); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_gamestate_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*Bullet); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_gamestate_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*Item); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_gamestate_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*PlayerInput); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_gamestate_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*Zombie); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_gamestate_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*GameMap); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_gamestate_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*GameState); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_gamestate_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_gamestate_proto_goTypes,
		DependencyIndexes: file_gamestate_proto_depIdxs,
		MessageInfos:      file_gamestate_proto_msgTypes,
	}.Build()
	File_gamestate_proto = out.File
	file_gamestate_proto_rawDesc = nil
	file_gamestate_proto_goTypes = nil
	file_gamestate_proto_depIdxs = nil
}
