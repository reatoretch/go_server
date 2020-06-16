package game

import (
    "net"
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

    receiver.Observer <- Notification{ Type: Message, ClientId: receiver.Id, Message: string(buf[:n]) }

    receiver.WaitMessage();
}
