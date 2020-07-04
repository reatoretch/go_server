package main

import (
    "go_server/game"
    "net"
    // For debugging
    //"reflect"
)

func main() {
    listener, err := net.Listen("tcp", ":1234")

    if err != nil {
        panic(err);
    }
    var room *game.Room
    ConnectionLoop(listener,0,room)

}

func ConnectionLoop(listener net.Listener, sequence int, room *game.Room) {
    connection, err := listener.Accept()
    if err != nil {
        panic(err)
    }
    if sequence%4 == 0 {
        room = game.NewRoom(sequence/4)
    }
    room.UserJoin(sequence, connection)
    ConnectionLoop(listener,sequence + 1,room)
}
