package game

import (
    "net"
)

type User struct {
    Name string
    Rate int
    UUID int
    Connection net.Conn
}
