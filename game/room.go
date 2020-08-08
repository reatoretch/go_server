package game

import (
    "fmt"
    "net"
)

type Room struct {
    RoomId int
    Channel chan Notification
    Observers Observer
}

func (room Room) GetStatus() GameStatus{
    return room.Observers.Status
}

func (room Room) Delete() {
    close(room.Channel)
}

func NewRoom(roomId int) *Room {
    room := new(Room)
    room.RoomId = roomId
    room.Channel = make(chan Notification)
    room.Observers = Observer{ UserNames: make([]string, 4, 4), UserRates: make([]int, 4, 4), Senders: make([]Sender, 0, 4), Subject: room.Channel, Status: Wait }
    go func() { room.Observers.WaitNotice() }()
    return room
}

//I will be passing the UUID to the sequence variable in the future.
// variable name "sequence" is change at that ime.
func (room *Room)UserJoin(sequence int, connection net.Conn,userName string,rate int) {
    var receiver Receiver = Receiver{ Id: len(room.Observers.Senders), Connection: connection, Observer: room.Channel }
    room.Observers.Join(receiver.Id, connection, userName, rate)
    go receiver.WaitMessage()

    //Wait for the add sender.
	//time.Sleep(time.Second * 1)
    //The game starts as soon as 4 members have gathered
    if 4 <= len(room.Observers.Senders) {
        room.Channel <- Notification{Type: InitGame}
        //For debugging
        fmt.Println(room.Observers.Game)
    }
}
