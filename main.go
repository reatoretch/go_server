package main

import (
	"net"
    "go_server/game"
)

func main() {
	listener, err := net.Listen("tcp", ":1234")

	if err != nil {
        panic(err);
	}

    var room = game.NewRoom(0)
    ConnectionLoop(listener,0,room)

}

func ConnectionLoop(listener net.Listener, sequence int, room game.Room) {
    connection, err := listener.Accept()
    if err != nil {
        panic(err)
    }
    room.UserJoin(sequence, connection)
    return ConnectionLoop(listener,sequence + 1,room)
}
