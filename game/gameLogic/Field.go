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

func (field *Field)getCanPutList(color int,haveBlocks []int) [][]int {
	sides:=[][]int{{0,1},{1,0},{-1,0},{0,-1}}
	crosses:=[][]int{{1,1},{1,-1},{-1,1},{-1,-1}}
	edge_map:=map[int][][]int{0:{{19,19}},1:{{19,0}},2:{{0,19}},3:{{0,0}}}
	ans:=[][]int{}

	for x:=0;x<20;x++{
		for y:=0;y<20;y++{
			for i,cross:=range crosses{
				if field.isIn(x+cross[0],y+cross[1]) && field.Board[y+cross[1]][x+cross[0]]==color && field.Board[y][x]==0{
					flag:=true
					for _,side:=range sides{
						if field.isIn(x+side[0],y+side[1]) && field.Board[y+side[1]][x+side[0]]==color{
							flag=false
						}
					}
					if flag{
						edge_map[i]=append(edge_map[i],[]int{x,y})
					}
				}
			}
		}
	}
	for _,num :=range haveBlocks{
		for spin:=0;spin<4;spin++{
			for y:=-1;y<7;y++{
				for x:=-1;x<7;x++{
					for j,cross:=range crosses{
						if 0<=x-cross[0] && x-cross[0]<5 &&0<=y-cross[1] && y-cross[1]<5 &&  kndBlock.array[num][spin][(y-cross[1])*5+x-cross[0]]{
							flag:=true
							for _,side:=range sides{
								if 0<=x+side[0] && x+side[0]<5 && 0<=y+side[1] && y+side[1]<5 && kndBlock.array[num][spin][(y+side[1])*5+x+side[0]]{flag=false}
							}
							if flag{
								for _,place :=range edge_map[j]{
									if field.canPut(place[0]-x+cross[0],place[1]-y+cross[1],spin,num,color){
										ans=append(ans,[]int{num,spin,place[0]-x+cross[0],place[1]-y+cross[1]})
									}
								}
							}
						}
					}
				}
			}
		}
	}
	return ans
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

