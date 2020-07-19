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
    modeSelector:=game.NewModeSelector()
    ConnectionLoop(listener,0,modeSelector)

}

func ConnectionLoop(listener net.Listener, sequence int, modeSelector *game.ModeSelector) {
    connection, err := listener.Accept()
    if err != nil {
        panic(err)
    }
    go modeSelector.Start(connection,sequence)
    ConnectionLoop(listener,sequence + 1,modeSelector)
}
