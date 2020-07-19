package game

import (
    "encoding/json"
    "net"
)

type Receiver struct {
    Id int
    Connection net.Conn
    Observer chan<- Notification
}

func (receiver *Receiver) Start(userName string ,rate int) {
	receiver.Observer <- Notification{ Type: Join, ClientId: receiver.Id, Connection: receiver.Connection ,UserName: userName,Rate:rate}
	receiver.WaitMessage();
}

func (receiver *Receiver) WaitMessage() {
    var buf = make([]byte, 1024);

    n, error := receiver.Connection.Read(buf);
    if error != nil {
        receiver.Observer <- Notification{ Type: Defect, ClientId: receiver.Id }
        return;
    }

    var jsonText map[string]interface{}
    //Ignore if the parse failed.
    if err := json.Unmarshal(buf[:n], &jsonText); err == nil {
        receiver.Observer <- Notification{ Type: Update, ClientId: receiver.Id, Message: []map[string]interface{}{jsonText}}
    }

    receiver.WaitMessage();
}
