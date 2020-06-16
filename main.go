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

    var channel = make(chan game.Notification);
    var observer game.Observer = game.Observer{ Senders: make([]game.Sender, 0, 4), Subject: channel };
    go observer.WaitNotice();
    waitClient(listener, 0, observer, channel);

}

func waitClient(listener net.Listener, sequence int, observer game.Observer, channel chan game.Notification) {
    connection, err := listener.Accept();

    if err != nil {
        panic(err);
    }

    var receiver game.Receiver = game.Receiver{ Id: sequence, Connection: connection, Observer: channel };
    go receiver.Start();

    waitClient(listener, sequence + 1, observer, channel);
}
