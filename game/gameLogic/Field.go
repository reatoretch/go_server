package gameLogic

import "fmt"

const (
    BLUE   = iota
    RED
    YELLOW
    GREEN
)
type Field struct{
  FieldWidth int;
  FieldHeight int;
  Board [20][20]int;
}

var field = Field{20,20,[20][20]int{{},}};


func (field *Field) CheckGameOver() bool{
	return true;
}

func (field Field) isIn(x int,y int) bool{
	if(x<0 || field.FieldWidth<=x || y<0 || field.FieldHeight<=y){return false;}
        return true;
}



func (field *Field) canPut(x int,y int,spin int,blockId int,color int) bool{
	var edge=false;
	var side=[][]int{{0,1},{0,-1},{1,0},{-1,0},};
	var cross=[][]int{{1,1},{-1,-1},{1,-1},{-1,1},};
	var targetBlock = kndBlock.array[blockId][spin]
	for i:=0;i<kndBlock.width*kndBlock.height;i++{
		if(!targetBlock[i]){continue;}
		if(!field.isIn(x+i%kndBlock.width,y+i/kndBlock.width)){return false}
		if(field.Board[y+i/kndBlock.width][x+i%kndBlock.width]!=0){return false}
		if((x+i%kndBlock.width==0 || x+i%kndBlock.width==field.FieldWidth-1) && (y+i/kndBlock.width==0 || y+i/kndBlock.width==field.FieldHeight-1) ){edge=true}

		for j:=0;j<4;j++{
			if(!field.isIn(x+i%kndBlock.width+side[j][0],y + i/ kndBlock.width+side[j][1])){continue}
			if(field.Board[y + i /kndBlock.width+side[j][1]][x+i%kndBlock.width+side[j][0]]==color){return false}
		}
		for j:=0;j<4;j++{
			if(!field.isIn(x+i%kndBlock.width+cross[j][0],y + i/ kndBlock.width+cross[j][1])){continue}
			if(field.Board[y + i /kndBlock.width+cross[j][1]][x+i%kndBlock.width+cross[j][0]]==color){edge=true;}
		}
	}
	return edge

}

func (field *Field) putBlock(x int,y int ,spin int,blockId int,color int) bool{
	if(!field.canPut(x,y,spin,blockId,color)){return false;}
	var targetBlock = kndBlock.array[blockId][spin]
	for i:=0;i<kndBlock.width*kndBlock.height;i++{
            if(!targetBlock[i]){continue;}
            field.Board[y + i / kndBlock.width][x + i % kndBlock.width] = color;
        }

	return true;
}

func (field *Field) easyDisp(){
	for i:=0;i<field.FieldWidth*field.FieldHeight;i++{
		if i%field.FieldWidth==0{fmt.Println()}
		fmt.Print(field.Board[i/field.FieldWidth][i%field.FieldHeight])
	}
}

func Test(){
  fmt.Println(field)
  fmt.Println(kndBlock.array[0][0])
  fmt.Println(field.putBlock(-1,-2,0,0,1),"true?")
  for i:=0;i<25;i++{
    fmt.Println(field.putBlock(-3+i/5,-3+i%5,0,0,1),i)
  }
  field.easyDisp()
}

