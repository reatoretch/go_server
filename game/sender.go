package game

import (
    "encoding/json"
    "net"
)

type Sender struct {
    Id int
    Connection net.Conn
}

func (sender Sender) SendMessage(message map[string]interface{}) {
    s, _ := json.Marshal(message)

    var buf = []byte(s);

    _, error := sender.Connection.Write(buf)
    if error != nil {
        panic(error);
    }
}
