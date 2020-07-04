package game

import (
    "encoding/json"
    "net"
    "log"
)

type Receiver struct {
    Id int
    Connection net.Conn
    Observer chan<- Notification
}

func (receiver Receiver) Start() {
    receiver.Observer <- Notification{ Type: Join, ClientId: receiver.Id, Connection: receiver.Connection }
    receiver.WaitMessage();
}

func (receiver Receiver) WaitMessage() {
    var buf = make([]byte, 1024);

    n, error := receiver.Connection.Read(buf);
    if error != nil {
        receiver.Observer <- Notification{ Type: Defect, ClientId: receiver.Id }
        return;
    }

    var jsonText map[string]interface{}
    if err := json.Unmarshal(buf[:n], &jsonText); err != nil {
        log.Fatal(err)
    }

    receiver.Observer <- Notification{ Type: Update, ClientId: receiver.Id, Message: []map[string]interface{}{jsonText}}

    receiver.WaitMessage();
}
