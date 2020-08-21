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
    color []string;
}


func (gameLogic *GameLogic) PlayerChange() bool{
	if gameLogic.CheckGameOver(){fmt.Println("gameOver");return true}
	for{
		gameLogic.TurnIdx=(gameLogic.TurnIdx+1)%4;
        for i:=0;i<4;i++{
            if gameLogic.PlayerRotation[i]-1==gameLogic.TurnIdx{
                gameLogic.TurnPlayer=i;
            }
        }
		if gameLogic.player[gameLogic.TurnPlayer].canPut(gameLogic){break}
        fmt.Println("pass")
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
	g:=GameLogic{field,[]Player{NewPlayer(0),NewPlayer(1),NewPlayer(2),NewPlayer(3)},0,[]int{1,2,3,4},0,[]string{},[]string{"red","blue","yellow","green"},}
	rand.Seed(time.Now().UnixNano())
    for i := range g.PlayerRotation {
        j := rand.Intn(i + 1)
        g.PlayerRotation[i], g.PlayerRotation[j] = g.PlayerRotation[j], g.PlayerRotation[i]
    }
    for i:=0;i<4;i++{
        if g.PlayerRotation[i]-1==g.TurnIdx{
            g.TurnPlayer=i;
        }
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
	fmt.Println("turnPlayer",gameLogic.TurnIdx,",client",gameLogic.PlayerRotation[playerId]-1,gameLogic.TurnPlayer)
	if !(gameLogic.TurnIdx==gameLogic.PlayerRotation[playerId]-1){
		fmt.Println("not turn player")
		fmt.Println(playerId,gameLogic.TurnIdx)
		fmt.Println(playerId,gameLogic.PlayerRotation)
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
	for i,v:=range gameLogic.player[gameLogic.TurnPlayer].blockIds{
		if v==blockId{
			idx=i
		}
	}
	if idx==-1{
		fmt.Println("duplicate block")
		fmt.Println(gameLogic.player[gameLogic.TurnPlayer].blockIds)
		return false
	}
    fmt.Println("delete ",gameLogic.TurnPlayer,gameLogic.player[gameLogic.TurnPlayer].blockIds[idx],gameLogic.TurnIdx)
	gameLogic.player[gameLogic.TurnPlayer].blockIds=append(gameLogic.player[gameLogic.TurnPlayer].blockIds[:idx],gameLogic.player[gameLogic.TurnPlayer].blockIds[idx+1:]...);
	gameLogic.field.easyDisp();
	if!gameLogic.field.putBlock(x,y,spin,blockId,gameLogic.PlayerRotation[playerId]){return false}
	gameLogic.history=append(gameLogic.history,strconv.Itoa(gameLogic.PlayerRotation[playerId]-1))
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


func (gameLogic *GameLogic) CreateInitMessage(userName []string, rate []int) ([]map[string]interface{}){
    messages:=[]map[string]interface{}{map[string]interface{}{},map[string]interface{}{},map[string]interface{}{},map[string]interface{}{}}
    for i := range gameLogic.color {
        j := rand.Intn(i + 1)
        gameLogic.color[i], gameLogic.color[j] = gameLogic.color[j], gameLogic.color[i]
    }
    for i:=0;i<4;i++{
        messages[i]["messageType"]="Init"
        m:=""
        for j:=0;j<4;j++{
            m+=gameLogic.color[j]
            if j!=3{m+=","}
            messages[i][gameLogic.color[gameLogic.PlayerRotation[j]-1]+"PlayerName"]=userName[j]
            messages[i][gameLogic.color[gameLogic.PlayerRotation[j]-1]+"Rate"]=rate[j]
        }
        messages[i]["PlayerRotation"]=m
        messages[i]["yourColor"]=gameLogic.color[gameLogic.PlayerRotation[i]-1]
    }
    return messages
}

func (gameLogic GameLogic) CreateUpdateMessage() []map[string]interface{}{
	message:=map[string]interface{}{}
	message["messageType"]="Update"
	message["Blocks"]=strings.Join(gameLogic.history[:], ",");
	message["Field"]="1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0";
	return []map[string]interface{}{message,message,message,message}
}

func (gameLogic GameLogic) CreateTerminateMessage(rates []int) ([]map[string]interface{}){
    message:=map[string]interface{}{}
    message["messageType"]="Terminate"
    rateChanges:=gameLogic.calcRates(rates,gameLogic.getRanking())
    for i:=0;i<4;i++{
        message[gameLogic.color[i]]=rateChanges[i]
    }
    return []map[string]interface{}{message,message,message,message}
}

//maybe Corresponds to SenderIndex
func (gameLogic GameLogic)getRanking() []int{
    numBlocks:=[]int{0,0,0,0}
    ranking:=[]int{0,0,0,0}
    for i:=0;i<4;i++{
        for _,v:=range gameLogic.player[i].blockIds{
            for _,j:=range kndBlock.array[v][0]{
                if(j){numBlocks[i]++}
            }
        }
    }
    for i:=0;i<4;i++{
        minv:=10000
        minId:=-1
        for j:=0;j<4;j++{
            fmt.Println(numBlocks[j])
            if numBlocks[j]<minv && ranking[j]==0{
                minId=j
                minv=numBlocks[j]
            }
        }
        ranking[minId]=i+1
    }
    return ranking
}

func (gameLogic GameLogic)calcRates(rates []int,ranking []int) []int{
    rateChanges:=[]int{0,0,0,0}
    sumRates:=0
    for _,i := range rates{sumRates+=i}
    for  i:=0; i<4; i++ {
        diffRates:=(sumRates-rates[i])/3-rates[i]
        if diffRates>400{diffRates=400}
        if diffRates < -400{diffRates=-400}

        if ranking[i]==1 || ranking[i]==2 {
            rateChanges[i]+=int(16.0+0.04*float64(diffRates))
        }else{
            diffRates*=-1
            rateChanges[i]-=int(16.0+0.04*float64(diffRates))
        }
        if ranking[i]==2 || ranking[i]==3{
            rateChanges[i]/=2
        }
    }
    return rateChanges
}
