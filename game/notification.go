package game

import (
    "net"
)

type NotificationType int

const  (
    InitGame NotificationType = iota
    Update
    Join
    Defect
)

type Notification struct {
    Type NotificationType
    ClientId int
    Connection net.Conn
    Message []map[string]interface{}
}
