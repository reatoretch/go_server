package game

import (
    "net"
    "fmt"
    "go_server/game/gameLogic"
)


type Observer struct {
    Senders []Sender
    Subject <-chan Notification
    Game gameLogic.GameLogic
}

func (observer Observer) WaitNotice() {
    notice := <-observer.Subject

    switch notice.Type {
    case Message:
        for i := range observer.Senders {
            observer.Senders[i].SendMessage(notice.Message)
        }
        break

    case Secret:
        //define any sendData (usage: SendMessage(sendData))
        observer.Senders[searchSender(notice.ClientId,observer.Senders)].SendMessage()
        break

    case Join:
        observer.Senders = appendSender(notice.ClientId, notice.Connection, observer.Senders)
        fmt.Printf("Client %d Join, now menber count is %d\n", notice.ClientId, len(observer.Senders))
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
