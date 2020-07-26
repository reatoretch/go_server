package game

import (
    "encoding/json"
    "net"
)

type Sender struct {
    Id int
    Connection net.Conn
    DummyFlag bool
}

func (sender Sender) SendMessage(message map[string]interface{}) {
    s, _ := json.Marshal(message)

    var buf = append(s, []byte("\n")...);

    if sender.DummyFlag{
	return
    }

    _, error := sender.Connection.Write(buf)
    if error != nil {
        panic(error);
    }
}
