package gameLogic

import (
	"math/rand"
	"time"
	"fmt"
	"strconv"
	"strings"
)

type Player struct {
    blockIds []int
    Color int
}

func NewPlayer(color int) Player {
	blocks:=[]int{}
	for i:=0;i<20;i++{
		blocks=append(blocks,i)
	}
	return Player{blocks,color}
}

func (player Player) canPut(game *GameLogic) bool{
	fmt.Println(player.Color)
	for _,v:= range player.blockIds{
		for j:=0;j<900;j++{
			for spin:=0;spin<4;spin++{
				if game.field.canPut(-3+j/30,-3+j%30,spin,v,game.PlayerRotation[player.Color]){return true}
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
	fmt.Println(message,"input message")
	blockIdf, ok := message["BlockId"].(float64);
	if !ok{
		fmt.Println("blockid not valid")
		return false
	}
	spinf, ok := message["spin"].(float64);
	if !ok{
		fmt.Println("spin not valid")
		return false
	}
	xf, ok := message["x"].(float64);
	if !ok{
		fmt.Println("x not valid")
		return false
	}
	yf, ok := message["y"].(float64);
	if !ok{
		return false
	}
	blockId:=int(blockIdf)
	spin:=int(spinf)
	x:=int(xf)
	y:=int(yf)
	playerId=playerId%4


	fmt.Println("turnPlayer",gameLogic.TurnIdx,",client",playerId)
	if !(playerId==gameLogic.TurnIdx){
		fmt.Println("not turn player")
		fmt.Println(playerId,gameLogic.TurnIdx)

		return false
	}
	fmt.Println("playerId success");
	fmt.Println(x,y,spin,blockId)
	if !gameLogic.field.canPut(x,y,spin,blockId,gameLogic.PlayerRotation[playerId]){
		fmt.Println("cannot put");
		return false
	}
	fmt.Println("the hand ok");

	idx:=-1
	for i,v:=range gameLogic.player[playerId].blockIds{
		if v==blockId{
			idx=i
		}
	}
	if idx==-1{
		fmt.Println("duplicate block")
		return false
	}
	gameLogic.player[playerId].blockIds=append(gameLogic.player[playerId].blockIds[:idx],gameLogic.player[playerId].blockIds[idx+1:]...);
	gameLogic.field.easyDisp();

	if!gameLogic.field.putBlock(x,y,spin,blockId,gameLogic.PlayerRotation[playerId]){return false}

	gameLogic.history=append(gameLogic.history,strconv.Itoa(playerId))
	gameLogic.history=append(gameLogic.history,strconv.Itoa(x))
	gameLogic.history=append(gameLogic.history,strconv.Itoa(y))
	gameLogic.history=append(gameLogic.history,strconv.Itoa(blockId))
	gameLogic.history=append(gameLogic.history,strconv.Itoa(spin))
	fmt.Println(gameLogic.CreateRandomPutMessage(gameLogic.PlayerRotation[playerId],playerId))

	return true;
}

func (gameLogic GameLogic)CreateRandomPutMessage(color int,playerId int) (map[string]interface{}){
	candidate:=gameLogic.field.getCanPutList(color,gameLogic.player[playerId].blockIds)
	if len(candidate)==0{return nil}
	hand:=candidate[rand.Intn(len(candidate))]

	return map[string]interface{}{"BlockId":float64(hand[0]),"spin":float64(hand[1]),"x":float64(hand[2]),"y":float64(hand[3])}
}


func (gameLogic GameLogic) CreateInitMessage(userName []string, rate []int) ([]map[string]interface{}){
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
	message1["yourColor"]="red"
	message2["yourColor"]="blue"
	message3["yourColor"]="yellow"
	message4["yourColor"]="green"
    message1["redPlayerName"]=userName[0]
    message2["redPlayerName"]=userName[0]
    message3["redPlayerName"]=userName[0]
    message4["redPlayerName"]=userName[0]
    message1["bluePlayerName"]=userName[1]
    message2["bluePlayerName"]=userName[1]
    message3["bluePlayerName"]=userName[1]
    message4["bluePlayerName"]=userName[1]
    message1["yellowPlayerName"]=userName[2]
    message2["yellowPlayerName"]=userName[2]
    message3["yellowPlayerName"]=userName[2]
    message4["yellowPlayerName"]=userName[2]
    message1["greenPlayerName"]=userName[3]
    message2["greenPlayerName"]=userName[3]
    message3["greenPlayerName"]=userName[3]
    message4["greenPlayerName"]=userName[3]
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
