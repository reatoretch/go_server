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
}

func (observer *Observer) WaitNotice() {
    notice := <-observer.Subject

    switch notice.Type {
    case InitGame:
        observer.Game = gameLogic.NewGameLogic()
	    messages:=observer.Game.CreateInitMessage(observer.UserNames,observer.UserRates)
        observer.Status = Started
        for i := range observer.Senders {
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
		observer.Game.PlayerChange()
	}
	break

    case Join:
        observer.Senders = appendSender(notice.ClientId, notice.Connection, observer.Senders)
        observer.UserNames[notice.ClientId] = notice.UserName
        observer.UserRates[notice.ClientId] = notice.Rate
	    fmt.Printf("%d:%s:%d Join, now menber count is %d\n", notice.ClientId,observer.UserNames[notice.ClientId], observer.UserRates[notice.ClientId], len(observer.Senders))

        break

    case Defect:
        observer.Senders = removeSender(notice.ClientId, observer.Senders);
        fmt.Printf("Client %d defect, now menber count is %d\n", notice.ClientId, len(observer.Senders))
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
