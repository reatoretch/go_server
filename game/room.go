package game

import (
    "net"
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
    go room.Observers.WaitNotice()
    return room
}

func (room Room)UserJoin(sequence int, connection net.Conn) {
    var receiver Receiver = Receiver{ Id: sequence, Connection: connection, Observer: room.Channel }
    go receiver.Start()
}
