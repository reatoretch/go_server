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
    room.Observers = Observer{ UserNames: make([]string, 4, 4), UserRates: make([]int, 4, 4), Senders: make([]Sender, 0, 4), Subject: room.Channel, Status: Wait ,RoomId: room.RoomId}
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

func (room *Room)Reconnect(connection net.Conn,userId int){
    id:=-1
    for i:=0;i<4;i++{
        if room.Observers.Senders[i].Id==userId{
            id=i
        }
    }
    if id==-1{
        connection.Close()
        return
    }
    //make new receiver
    var receiver Receiver = Receiver{ Id: id, Connection: connection, Observer: room.Channel }
    go receiver.WaitMessage()
    //sender reconfigure
    room.Observers.Senders[id].Connection=connection
    room.Observers.Senders[id].DummyFlag=false
    messages:=room.Observers.Game.CreateInitMessage(room.Observers.UserNames,room.Observers.UserRates)
    messages[id]["RoomID"]=room.Observers.RoomId
    messages[id]["UserID"]=userId
    room.Observers.Senders[id].SendMessage(messages[id])
    messages=room.Observers.Game.CreateUpdateMessage()
    room.Observers.Senders[id].SendMessage(messages[id])
}
