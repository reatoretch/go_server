package main

import (
    "fmt"
    "net"
    "strings"
)

func main() {
    color_ind := map[string]uint32 {
        "red"   : 0,
        "blue"  : 1,
        "yellow": 2,
        "green" : 3,
    }
    color := map[uint32]string {
        0 : "red"   ,
        1 : "blue"  ,
        2 : "yellow",
        3 : "green" ,
    }
    listener, err := net.Listen("tcp", "0.0.0.0:1234")
    if err != nil {
        fmt.Printf("Listen error: %s\n", err)
        return
    }
    defer listener.Close()
    
    conn, err := listener.Accept()
    if err != nil {
        fmt.Printf("Accept error: %s\n", err)
        return
    }
    
    buf := make([]byte,1024)
    for {
        n, err := conn.Read(buf)
        if n == 0 {
            break
        }
        if err != nil {
            fmt.Printf("Read error: %s\n", err)
            return
        }
        now_color, ok := color_ind[strings.Replace(string(buf[:n]),"\n",
        "", 1)]
        if !ok {
            fmt.Printf("Array error: %s\n", err)
            return
        }
        //fmt.Printf(string(buf[:n]))
        //fmt.Println(color_ind[strings.Replace(string(buf[:n]), "\n", "", 1)])
        message := color[(now_color+1)%4]
        newmessage := strings.ToLower(message)
        conn.Write([]byte(newmessage + "\n"))
    }
    //defer conn.Close()
}
