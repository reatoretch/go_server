package game

import (
    "net"
    "fmt"
    "go_server/game/gameLogic"
)

type GameStatus int

const (
    Wait GameStatus = iota
    Started
    Finished
)

type Observer struct {
    UserNames []string
    UserRates []int
    Senders []Sender
    Subject <-chan Notification
    Game gameLogic.GameLogic
    Status GameStatus
    RoomId int
}

func (observer *Observer) Close(){
	observer.Status=Finished
	for i:=0;i<4;i++{
		//connection close if the connection alive
		if !observer.Senders[i].DummyFlag{
			observer.Senders[i].Connection.Close()
		}
	}
}

func (observer *Observer) Join(ClientId int, Connection net.Conn, UserName string, Rate int) {
    observer.Senders = appendSender(ClientId, Connection, observer.Senders)
    observer.UserNames[ClientId] = UserName
    observer.UserRates[ClientId] = Rate
	fmt.Printf("%d:%s:%d Join, now menber count is %d\n", ClientId,observer.UserNames[ClientId], observer.UserRates[ClientId], len(observer.Senders))
    message:=map[string]interface{}{}
    message["messageType"]="Wait"
    message["nowWaitingPlayer"]=len(observer.Senders)
	for i := range observer.Senders {
		observer.Senders[i].SendMessage(message)
	}

}

func (observer *Observer) WaitNotice() {
    defer func() {
        //fmt.Println("GameEnd")
        if err := recover();err != nil {
            fmt.Println("Crash!:", err)
            observer.Status=Finished
            message:=map[string]interface{}{}
            message["messageType"]="Error"
            for i := range observer.Senders {
                observer.Senders[i].SendMessage(message)
                observer.Senders[i].Connection.Close()
            }
        }
    }()

    notice := <-observer.Subject

    switch notice.Type {
    case InitGame:
        observer.Game = gameLogic.NewGameLogic()
	    messages:=observer.Game.CreateInitMessage(observer.UserNames,observer.UserRates)
        observer.Status = Started
        for i := range observer.Senders {
            messages[i]["RoomID"]=observer.RoomId
            messages[i]["UserID"]=observer.Senders[i].Id
            observer.Senders[i].SendMessage(messages[i])
        }
        break

    case Update:
	if observer.Game.Update(notice.ClientId,notice.Message[0]){
		fmt.Println("field update!")
		messages:=observer.Game.CreateUpdateMessage()
		for i := range observer.Senders {
			observer.Senders[i].SendMessage(messages[i])
		}
		gameOver:=observer.Game.PlayerChange()
		for observer.Senders[observer.Game.TurnPlayer].DummyFlag && !gameOver{
			fmt.Println("dummy_loop")
			if observer.Game.Update(observer.Game.TurnPlayer,observer.Game.CreateRandomPutMessage(observer.Game.PlayerRotation[observer.Game.TurnPlayer],observer.Game.TurnPlayer)){
				fmt.Println("field_update!")
				messages:=observer.Game.CreateUpdateMessage()
				for i := range observer.Senders {
					observer.Senders[i].SendMessage(messages[i])
				}
			    gameOver=observer.Game.PlayerChange()
			}
		}
		if gameOver{
			messages:=observer.Game.CreateTerminateMessage(observer.UserRates)
			for i := range observer.Senders {
				observer.Senders[i].SendMessage(messages[i])
			}
			observer.Close()
			return
		}
	}

	break

    case Defect:
	    if observer.Status==Wait{
		    observer.Senders = removeSender(notice.ClientId, observer.Senders);
		    fmt.Printf("Client %d defect, now menber count is %d\n", notice.ClientId, len(observer.Senders))
		    message:=map[string]interface{}{}
		    message["messageType"]="Wait"
		    message["nowWaitingPlayer"]=len(observer.Senders)
		    for i := range observer.Senders {
			    observer.Senders[i].SendMessage(message)
		    }
	    }else if observer.Status==Started{
		    observer.Senders[notice.ClientId].DummyFlag=true
		    fmt.Printf("Client %d defect, change to dummy\n", notice.ClientId)
		    gameOver:=false
		    for observer.Senders[observer.Game.TurnPlayer].DummyFlag && !gameOver{
			    fmt.Println("dummy_loop")
			    if observer.Game.Update(observer.Game.TurnPlayer,observer.Game.CreateRandomPutMessage(observer.Game.PlayerRotation[observer.Game.TurnPlayer],observer.Game.TurnPlayer)){
				    fmt.Println("field_update!")
				    messages:=observer.Game.CreateUpdateMessage()
				    for i := range observer.Senders {
					    observer.Senders[i].SendMessage(messages[i])
				    }
			        gameOver=observer.Game.PlayerChange()
			    }
		    }
		    if gameOver{
			    messages:=observer.Game.CreateTerminateMessage(observer.UserRates)
			    for i := range observer.Senders {
				    observer.Senders[i].SendMessage(messages[i])
			    }
			    observer.Close()
			    return
		    }
	    }
        break

    default:

    }

    observer.WaitNotice();
}

func appendSender(senderId int, connection net.Conn, senders []Sender) []Sender {
    return append(senders, Sender{ Id: senderId, Connection: connection})
}

func removeSender(senderId int, senders []Sender) []Sender {
    var find = searchSender(senderId, senders)

    if (find == -1) {
        return senders;
    }

    return append(senders[:find], senders[find+1:]...);
}

func searchSender(senderId int, senders []Sender) int {
    var find = -1;

    for i:= range senders {
        if ( senders[i].Id == senderId ) {
            find = i;
            break;
        }
    }
    return find;
}
