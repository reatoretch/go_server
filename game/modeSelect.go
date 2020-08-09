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
    fmt.Println(buf[:n])
    //Ignore if the parse failed.
    if err := json.Unmarshal(buf[:n], &jsonText); err == nil {
        fmt.Println("marshal")
        name, ok := jsonText["UserName"].(string);
        if !ok{return}
        rate, ok := jsonText["Rate"].(float64);
        if !ok{return}
        action, ok := jsonText["Action"].(string);
        if !ok{return}
        fmt.Println(action)
        fmt.Println(ok)


        User:=new(User)
        User.Name=name
        User.Rate=int(rate)

        if action=="Join" && modeSelector.Rooms[modeSelector.Index].GetStatus()!=Wait{
            modeSelector.AppendRoom()
            fmt.Println("appendRoom")
        }
        if action=="Join"{
            modeSelector.Rooms[modeSelector.Index].UserJoin(UUID,Connection,User.Name,User.Rate)
        }
        if action=="Reconnect"{
            //find room and dicition client
            roomId, ok := jsonText["RoomId"].(float64);
            if(!ok){
                modeSelector.Close(Connection)
                return
            }
            userId, ok := jsonText["UserId"].(float64);
            if(!ok){
                modeSelector.Close(Connection)
                return
            }
            if int(roomId)<0 || len(modeSelector.Rooms)<=int(roomId){
                modeSelector.Close(Connection)
                return
            }
            room:=modeSelector.Rooms[int(roomId)]
            if room.GetStatus()!=Started{
                modeSelector.Close(Connection)
                return
            }
            room.Reconnect(Connection,int(userId))
        }
        return
    }
    modeSelector.Close(Connection)
    return;

}

func (modeSelector *ModeSelector) Close(Connection net.Conn){
	Connection.Close()
}
