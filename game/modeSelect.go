package game

import (
    "encoding/json"
    "net"
    "fmt"
)

type ModeSelector struct {
    Rooms []*Room
    Index int
}

func (modeSelector *ModeSelector)AppendRoom(){
	modeSelector.Index+=1
	modeSelector.Rooms=append(modeSelector.Rooms,NewRoom(modeSelector.Index))
}

func NewModeSelector() *ModeSelector{
	modeSelector:=new(ModeSelector)
	modeSelector.Rooms=make([]*Room,0,100)
	modeSelector.Index=-1
	modeSelector.AppendRoom();
	return modeSelector
}

func (modeSelector *ModeSelector) Start(Connection net.Conn,UUID int) {
    var buf = make([]byte, 1024);
    n, error := Connection.Read(buf);
    if error != nil {
	modeSelector.Close(Connection)
        return;
    }
    var jsonText map[string]interface{}
    //Ignore if the parse failed.
    if err := json.Unmarshal(buf[:n], &jsonText); err == nil {
	    fmt.Println(jsonText)
	    fmt.Println(UUID)
	    User:=new(User)
	    User.Name="Default user"
	    User.Rate=1500

	    if modeSelector.Rooms[modeSelector.Index].GetStatus()!=Wait{
		    modeSelector.AppendRoom()
		    fmt.Println("appendRoom")
	    }
	    modeSelector.Rooms[modeSelector.Index].UserJoin(UUID,Connection,User.Name,User.Rate)
	    return
    }
    modeSelector.Close(Connection)
    return;

}

func (modeSelector *ModeSelector) Close(Connection net.Conn){
	Connection.Close()
}
