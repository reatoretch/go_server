package gameLogic

import (
	"math/rand"
	"time"
	"fmt"
	"strconv"
	"strings"
)

type Player struct {
    blockIds map[int] bool
    Color int
}

func (player Player) canPut(game *GameLogic) bool{
	fmt.Println(player.Color)
	for i,v:= range player.blockIds{
		if !v{continue}
		for j:=0;j<900;j++{
			for spin:=0;spin<4;spin++{
				//fmt.Println(-3+j/30,-3+j%30,spin,i,game.PlayerRotation[player.Color])
				if game.field.canPut(-3+j/30,-3+j%30,spin,i,game.PlayerRotation[player.Color]){return true}
			}
		}
	}
	return false;
}


type GameLogic struct {
    field Field
    player []Player
    TurnIdx int;
    PlayerRotation []int;
    TurnPlayer int;
    history []string;
}

func NewPlayer(color int) Player {
	blocks:=map[int] bool{};
	for i:=0;i<20;i++{
		blocks[i]=true;
	}
	return Player{blocks,color}
}

func (gameLogic *GameLogic) PlayerChange() bool{
	if gameLogic.CheckGameOver(){fmt.Println("gameOver");return true}
	for{
		gameLogic.TurnIdx=(gameLogic.TurnIdx+1)%4;
		gameLogic.TurnPlayer=gameLogic.PlayerRotation[gameLogic.TurnIdx];
		if gameLogic.player[gameLogic.TurnIdx].canPut(gameLogic){break}
	}


	return false
}

func (gameLogic *GameLogic) CheckGameOver() bool{
	canPut:=false;
	for i:=0;i<4;i++{
		fmt.Println(gameLogic.player[i].canPut(gameLogic))
		canPut=canPut||gameLogic.player[i].canPut(gameLogic)
	}
	return !canPut;
}

func NewGameLogic() GameLogic{
	g:=GameLogic{field,[]Player{NewPlayer(0),NewPlayer(1),NewPlayer(2),NewPlayer(3)},0,[]int{1,2,3,4},0,[]string{}}
	rand.Seed(time.Now().UnixNano())
	for i := range g.PlayerRotation {
        j := rand.Intn(i + 1)
        g.PlayerRotation[i], g.PlayerRotation[j] = g.PlayerRotation[j], g.PlayerRotation[i]
	}
	return g
}


func (gameLogic *GameLogic) Update(playerId int,message map[string]interface{}) bool {
	fmt.Println(gameLogic)
	fmt.Println(message)
	fmt.Println(playerId)
	blockIdf, ok := message["BlockId"].(float64);
	if !ok{return false}
	spinf, ok := message["spin"].(float64);
	if !ok{return false}
	xf, ok := message["x"].(float64);
	if !ok{return false}
	yf, ok := message["y"].(float64);
	if !ok{return false}
	blockId:=int(blockIdf)
	spin:=int(spinf)
	x:=int(xf)
	y:=int(yf)
	playerId=playerId%4


	if !(playerId==gameLogic.TurnIdx){return false}
	fmt.Println("playerId success");
	fmt.Println(x,y,spin,blockId)
	if !gameLogic.field.canPut(x,y,spin,blockId,gameLogic.PlayerRotation[playerId]){return false}
	fmt.Println("the hand ok");

	if!gameLogic.field.putBlock(x,y,spin,blockId,gameLogic.PlayerRotation[playerId]){return false}

	gameLogic.history=append(gameLogic.history,strconv.Itoa(playerId));
	gameLogic.history=append(gameLogic.history,strconv.Itoa(x))
	gameLogic.history=append(gameLogic.history,strconv.Itoa(y));
	gameLogic.history=append(gameLogic.history,strconv.Itoa(blockId));
	gameLogic.history=append(gameLogic.history,strconv.Itoa(spin));
	gameLogic.player[playerId].blockIds[blockId]=false;
	gameLogic.field.easyDisp();

	return true;
}


func (gameLogic GameLogic) CreateInitMessage() ([]map[string]interface{}){
	message1:=map[string]interface{}{}
	message2:=map[string]interface{}{}
	message3:=map[string]interface{}{}
	message4:=map[string]interface{}{}
	message1["messageType"]="Init"
	message2["messageType"]="Init"
	message3["messageType"]="Init"
	message4["messageType"]="Init"
	message1["PlayerRotation"]="red,blue,yellow,green";
	message2["PlayerRotation"]="red,blue,yellow,green";
	message3["PlayerRotation"]="red,blue,yellow,green";
	message4["PlayerRotation"]="red,blue,yellow,green";
	message2["messageType"]="Init"
	message3["messageType"]="Init"
	message4["messageType"]="Init"
	message1["yourColor"]="red"
	message2["yourColor"]="blue"
	message3["yourColor"]="yellow"
	message4["yourColor"]="green"
	return []map[string]interface{}{message1,message2,message3,message4}
}

func (gameLogic GameLogic) CreateUpdateMessage() []map[string]interface{}{
	message:=map[string]interface{}{}
	message["messageType"]="Update"
	message["Blocks"]=strings.Join(gameLogic.history[:], ",");
	message["Field"]="1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0";
	return []map[string]interface{}{message,message,message,message}
}

func (gameLogic GameLogic) CreateTerminateMessage() (map[string]interface{},map[string]interface{},map[string]interface{},map[string]interface{}){
	message:=map[string]interface{}{}
	message["messageType"]="Terminate"
	return message,message,message,message
}
