package game

import (
    "fmt"
    "net"
    "time"
)

type Room struct {
    RoomId int
    Channel chan Notification
    Observers Observer
}

func (room Room) Delete() {
    close(room.Channel)
}

func NewRoom(roomId int) *Room {
    room := new(Room)
    room.RoomId = roomId
    room.Channel = make(chan Notification)
    room.Observers = Observer{ Senders: make([]Sender, 0, 4), Subject: room.Channel }
    go func() { room.Observers.WaitNotice() }()
    return room
}

func (room Room)UserJoin(sequence int, connection net.Conn) {
    var receiver Receiver = Receiver{ Id: sequence, Connection: connection, Observer: room.Channel }
    go receiver.Start()

    //The game starts as soon as 4 members have gathered
    if sequence%4 == 3 {
	time.Sleep(time.Second * 1)
        room.Channel <- Notification{Type: Message}
        //For debugging
        fmt.Println(room.Observers.Game)
    }
}
