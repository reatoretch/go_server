package main

import (
    "encoding/json"
    "fmt"
    "go_server/game"
    "io/ioutil"
    "log"
    "net"
)

type MoreInfo struct {
    Id int `json:"userId"`
    Friends []string `json:"userFriends"`
}

type Config struct {
    Info MoreInfo `json:"userInfo"`
}

func main() {
    listener, err := net.Listen("tcp", ":1234")

    if err != nil {
        panic(err);
    }

    bytes, err := ioutil.ReadFile("env_num.json")
    if err != nil {
        panic(err);
    }

    var config Config
    if err := json.Unmarshal(bytes, &config); err != nil {
        log.Fatal(err)
    }

    fmt.Println(config.Info.Friends[0])

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
