package gameLogic

import (
	"math/rand"
	"time"
	"fmt"
)

type Player struct {
    blockIds map[int] bool
    Color int
}

type GameLogic struct {
    field Field
    player []Player
    TurnIdx int;
    PlayerRotation []int;
    TurnPlayer int;
}

func NewPlayer(color int) Player {
	blocks:=map[int] bool{};
	for i:=0;i<20;i++{
		blocks[i]=true;
	}
	return Player{blocks,color}
}

func (gameLogic *GameLogic) PlayerChange() bool{
  gameLogic.TurnIdx=(gameLogic.TurnIdx+1)%4;
  gameLogic.TurnPlayer=gameLogic.PlayerRotation[gameLogic.TurnIdx];
  return gameLogic.CheckGameOver();
}

func (gameLogic *GameLogic) CheckGameOver() bool{
	return false;
}

func NewGameLogic() GameLogic{
	g:=GameLogic{field,[]Player{NewPlayer(BLUE),NewPlayer(RED),NewPlayer(GREEN),NewPlayer(YELLOW)},0,[]int{1,2,3,4},0}
	rand.Seed(time.Now().UnixNano())
	for i := range g.PlayerRotation {
        j := rand.Intn(i + 1)
        g.PlayerRotation[i], g.PlayerRotation[j] = g.PlayerRotation[j], g.PlayerRotation[i]
	}

	return g
}


func (gameLogic *GameLogic) Update(playerId int,message map[string]interface{}) bool {
	blockId, ok := message["BlockId"].(int);
	if !ok{return false}
	spin, ok := message["spin"].(int);
	if !ok{return false}
	fmt.Println("OK")
	x, ok := message["x"].(int);
	if !ok{return false}
	fmt.Println("OK")
	y, ok := message["y"].(int);
	if !ok{return false}
	fmt.Println("OK")

	if !(playerId==gameLogic.TurnIdx){return false}
	if !gameLogic.field.canPut(x,y,spin,blockId,gameLogic.PlayerRotation[playerId]){return false}

	return gameLogic.field.putBlock(x,y,spin,blockId,gameLogic.PlayerRotation[playerId])
}


func (gameLogic GameLogic) CreateInitMessage() map[string]interface{}{
	message:=map[string]interface{}{}
	message["messageType"]="Terminate"
	return message
}

func (gameLogic GameLogic) CreateUpdateMessage() map[string]interface{}{
	message:=map[string]interface{}{}
	message["messageType"]="Terminate"
	return message
}

func (gameLogic GameLogic) CreateTerminateMessage() map[string]interface{}{
	message:=map[string]interface{}{}
	message["messageType"]="Terminate"
	return message
}
