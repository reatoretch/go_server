package gameLogic
import "fmt"

type KndBlock struct{
    height int;
    width int;
    array [][][]bool;
}

var kndBlock=KndBlock{
	5,5,
	[][][]bool{
            {
                    {
                            false, false, false, false, false,
                            false, false, false, false, false,
                            false, true , false, false, false,
                            false, false, false, false, false,
                            false, false, false, false, false,
                    },
                    {
                            false, false, false, false, false,
                            false, false, false, false, false,
                            false, true, false, false, false,
                            false, false, false, false, false,
                            false, false, false, false, false,
                    },
                    {
                            false, false, false, false, false,
                            false, false, false, false, false,
                            false, true, false, false, false,
                            false, false, false, false, false,
                            false, false, false, false, false,
                    },
                    {
                            false, false, false, false, false,
                            false, false, false, false, false,
                            false, true , false, false, false,
                            false, false, false, false, false,
                            false, false, false, false, false,
                    },
            },
            {
                    {
                            false, false, false, false, false,
                            false, false, false, false, false,
                            false, true , true , false, false,
                            false, false, false, false, false,
                            false, false, false, false, false,
                    },
                    {
                            false, false, false, false, false,
                            false, false, false, false, false,
                            false, false, true , false, false,
                            false, false, true , false, false,
                            false, false, false, false, false,
                    },
                    {
                            false, false, false, false, false,
                            false, false, false, false, false,
                            false, false, false, false, false,
                            false, true , true , false, false,
                            false, false, false, false, false,
                    },
                    {
                            false, false, false, false, false,
                            false, false, false, false, false,
                            false, true , false, false, false,
                            false, true , false, false, false,
                            false, false, false, false, false,
                    },
            },
            {
                    {
                            false, false, false, false, false,
                            false, false, false, false, false,
                            false, true,  true, false, false,
                            false, false, true, false, false,
                            false, false, false, false, false,
                    },
                    {
                            false, false, false, false, false,
                            false, false, false, false, false,
                            false, false, true , false, false,
                            false, true , true , false, false,
                            false, false, false, false, false,
                    },
                    {
                            false, false, false, false, false,
                            false, false, false, false, false,
                            false, true , false, false, false,
                            false, true , true , false, false,
                            false, false, false, false, false,
                    },
                    {
                            false, false, false, false, false,
                            false, false, false, false, false,
                            false, true , true , false, false,
                            false, true , false, false, false,
                            false, false, false, false, false,
                    },
            },
            {
                    {
                            false, false, false, false, false,
                            false, false, false, false, false,
                            false, true , true , true , false,
                            false, false, false, false, false,
                            false, false, false, false, false,
                    },
                    {
                            false, false, false, false, false,
                            false, false, true , false, false,
                            false, false, true , false, false,
                            false, false, true , false, false,
                            false, false, false, false, false,
                    },
                    {
                            false, false, false, false, false,
                            false, false, false, false, false,
                            false, true , true , true , false,
                            false, false, false, false, false,
                            false, false, false, false, false,
                    },
                    {
                            false, false, false, false, false,
                            false, false, true , false, false,
                            false, false, true , false, false,
                            false, false, true , false, false,
                            false, false, false, false, false,
                    },
            },
            {
                    {
                            false, false, false, false, false,
                            false, false, true , false, false,
                            false, true , true , true , false,
                            false, false, false, false, false,
                            false, false, false, false, false,
                    },
                    {
                            false, false, false, false, false,
                            false, false, true , false, false,
                            false, false, true , true , false,
                            false, false, true , false, false,
                            false, false, false, false, false,
                    },
                    {
                            false, false, false, false, false,
                            false, false, false, false, false,
                            false, true , true , true , false,
                            false, false, true , false, false,
                            false, false, false, false, false,
                    },
                    {
                            false, false, false, false, false,
                            false, false, true , false, false,
                            false, true , true , false, false,
                            false, false, true , false, false,
                            false, false, false, false, false,
                    },
            },
            {
                    {
                            false, false, false, false, false,
                            false, false, false, false, false,
                            false, true , true , true , true ,
                            false, false, false, false, false,
                            false, false, false, false, false,
                    },
                    {
                            false, false, false, false, false,
                            false, false, true , false, false,
                            false, false, true , false, false,
                            false, false, true , false, false,
                            false, false, true , false, false,
                    },
                    {
                            false, false, false, false, false,
                            false, false, false, false, false,
                            false, true , true , true , true ,
                            false, false, false, false, false,
                            false, false, false, false, false,
                    },
                    {
                            false, false, false, false, false,
                            false, false, true , false, false,
                            false, false, true , false, false,
                            false, false, true , false, false,
                            false, false, true , false, false,
                    },
            },
            {
                    {
                            false, false, false, false, false,
                            false, false, false, true , false,
                            false, true , true , true , false,
                            false, false, false, false, false,
                            false, false, false, false, false,
                    },
                    {
                            false, false, false, false, false,
                            false, false, true , false, false,
                            false, false, true , false, false,
                            false, false, true , true , false,
                            false, false, false, false, false,
                    },
                    {
                            false, false, false, false, false,
                            false, false, false, false, false,
                            false, true , true , true , false,
                            false, true , false, false, false,
                            false, false, false, false, false,
                    },
                    {
                            false, false, false, false, false,
                            false, true , true , false, false,
                            false, false, true , false, false,
                            false, false, true , false, false,
                            false, false, false, false, false,
                    },
            },
            {
                    {
                            false, false, false, false, false,
                            false, false, true , true , false,
                            false, true , true , false, false,
                            false, false, false, false, false,
                            false, false, false, false, false,
                    },
                    {
                            false, false, false, false, false,
                            false, true , false, false, false,
                            false, true , true , false, false,
                            false, false, true , false, false,
                            false, false, false, false, false,
                    },
                    {
                            false, false, false, false, false,
                            false, false, true , true , false,
                            false, true , true , false, false,
                            false, false, false, false, false,
                            false, false, false, false, false,
                    },
                    {
                            false, false, false, false, false,
                            false, true , false, false, false,
                            false, true , true , false, false,
                            false, false, true , false, false,
                            false, false, false, false, false,
                    },
            },
            {
                    {
                            false, false, false, false, false,
                            false, true , false, false, false,
                            false, true , true , true , true ,
                            false, false, false, false, false,
                            false, false, false, false, false,
                    },
                    {
                            false, false, false, false, false,
                            false, false, true , true , false,
                            false, false, true , false, false,
                            false, false, true , false, false,
                            false, false, true , false, false,
                    },
                    {
                            false, false, false, false, false,
                            false, false, false, false, false,
                            true , true , true , true , false,
                            false, false, false, true , false,
                            false, false, false, false, false,
                    },
                    {
                            false, false, true , false, false,
                            false, false, true , false, false,
                            false, false, true , false, false,
                            false, true , true , false, false,
                            false, false, false, false, false,
                    },
            },
            {
                    {
                            false, false, false, false, false,
                            false, false, true , false, false,
                            false, false, true , false, false,
                            false, true , true , true , false,
                            false, false, false, false, false,
                    },
                    {
                            false, false, false, false, false,
                            false, true , false, false, false,
                            false, true , true , true , false,
                            false, true , false, false, false,
                            false, false, false, false, false,
                    },
                    {
                            false, false, false, false, false,
                            false, true , true , true , false,
                            false, false, true , false, false,
                            false, false, true , false, false,
                            false, false, false, false, false,
                    },
                    {
                            false, false, false, false, false,
                            false, false, false, true , false,
                            false, true , true , true , false,
                            false, false, false, true , false,
                            false, false, false, false, false,
                    },
            },
            {
                    {
                            false, false, false, false, false,
                            false, true , false, false, false,
                            false, true , false, false, false,
                            false, true , true , true , false,
                            false, false, false, false, false,
                    },
                    {
                            false, false, false, false, false,
                            false, true , true , true , false,
                            false, true , false, false, false,
                            false, true , false, false, false,
                            false, false, false, false, false,
                    },
                    {
                            false, false, false, false, false,
                            false, true , true , true , false,
                            false, false, false, true , false,
                            false, false, false, true , false,
                            false, false, false, false, false,
                    },
                    {
                            false, false, false, false, false,
                            false, false, false, true , false,
                            false, false, false, true , false,
                            false, true , true , true , false,
                            false, false, false, false, false,
                    },

            },
            {
                    {
                            false, false, false, false, false,
                            false, false, false, false, false,
                            false, false, true , true , true ,
                            false, true , true , false, false,
                            false, false, false, false, false,
                    },
                    {
                            false, false, false, false, false,
                            false, true , false, false, false,
                            false, true , true , false, false,
                            false, false, true , false, false,
                            false, false, true , false, false,
                    },
                    {
                            false, false, false, false, false,
                            false, false, false, false, false,
                            false, false, true , true , false,
                            true , true , true , false, false,
                            false, false, false, false, false,
                    },
                    {
                            false, false, true , false, false,
                            false, false, true , false, false,
                            false, false, true , true , false,
                            false, false, false, true , false,
                            false, false, false, false, false,
                    },

            },
            {
                    {
                            false, false, false, false, false,
                            false, false, false, true , false,
                            false, true , true , true , false,
                            false, true , false, false, false,
                            false, false, false, false, false,
                    },
                    {
                            false, false, false, false, false,
                            false, true , true , false, false,
                            false, false, true , false, false,
                            false, false, true , true , false,
                            false, false, false, false, false,
                    },
                    {
                            false, false, false, false, false,
                            false, false, false, true , false,
                            false, true , true , true , false,
                            false, true , false, false, false,
                            false, false, false, false, false,
                    },
                    {
                            false, false, false, false, false,
                            false, true , true , false, false,
                            false, false, true , false, false,
                            false, false, true , true , false,
                            false, false, false, false, false,
                    },
            },
            {
                    {
                            false, false, false, false, false,
                            false, false, false, false, false,
                            true , true , true , true , true ,
                            false, false, false, false, false,
                            false, false, false, false, false,
                    },
                    {
                            false, false, true , false, false,
                            false, false, true , false, false,
                            false, false, true , false, false,
                            false, false, true , false, false,
                            false, false, true , false, false,
                    },
                    {
                            false, false, false, false, false,
                            false, false, false, false, false,
                            true , true , true , true , true ,
                            false, false, false, false, false,
                            false, false, false, false, false,
                    },
                    {
                            false, false, true , false, false,
                            false, false, true , false, false,
                            false, false, true , false, false,
                            false, false, true , false, false,
                            false, false, true , false, false,
                    },
            },
            {
                    {
                            false, false, false, false, false,
                            false, false, false, false, false,
                            false, true , true , true , false,
                            false, true , true , false, false,
                            false, false, false, false, false,
                    },
                    {
                            false, false, false, false, false,
                            false, true , true , false, false,
                            false, true , true , false, false,
                            false, false, true , false, false,
                            false, false, false, false, false,
                    },
                    {
                            false, false, false, false, false,
                            false, false, true , true , false,
                            false, true , true , true , false,
                            false, false, false, false, false,
                            false, false, false, false, false,
                    },
                    {
                            false, false, false, false, false,
                            false, false, true , false, false,
                            false, false, true , true , false,
                            false, false, true , true , false,
                            false, false, false, false, false,
                    },


            },
            {
                    {
                            false, false, false, false, false,
                            false, false, true , true , false,
                            false, true , true , false, false,
                            false, true , false, false, false,
                            false, false, false, false, false,
                    },
                    {
                            false, false, false, false, false,
                            false, true , true , false, false,
                            false, false, true , true , false,
                            false, false, false, true , false,
                            false, false, false, false, false,
                    },
                    {
                            false, false, false, false, false,
                            false, false, false, true , false,
                            false, false, true , true , false,
                            false, true , true , false, false,
                            false, false, false, false, false,
                    },

                    {
                            false, false, false, false, false,
                            false, true , false, false, false,
                            false, true , true , false, false,
                            false, false, true , true , false,
                            false, false, false, false, false,
                    },
            },
            {
                    {
                            false, false, false, false, false,
                            false, false, false, false, false,
                            false, true , true , true , false,
                            false, true , false, true, false,
                            false, false, false, false, false,
                    },
                    {
                            false, false, false, false, false,
                            false, true , true , false, false,
                            false, false, true , false, false,
                            false, true , true , false, false,
                            false, false, false, false, false,
                    },
                    {
                            false, false, false, false, false,
                            false, true , false, true , false,
                            false, true , true , true , false,
                            false, false, false, false, false,
                            false, false, false, false, false,
                    },
                    {
                            false, false, false, false, false,
                            false, false, true , true , false,
                            false, false, true , false, false,
                            false, false, true , true, false,
                            false, false, false, false, false,
                    },
            },
            {
                    {
                            false, false, false, false, false,
                            false, false, true , true , false,
                            false, true , true , false, false,
                            false, false, true , false, false,
                            false, false, false, false, false,
                    },
                    {
                            false, false, false, false, false,
                            false, false, true , false, false,
                            false, true , true , true , false,
                            false, false, false, true , false,
                            false, false, false, false, false,
                    },
                    {
                            false, false, false, false, false,
                            false, false, true , false, false,
                            false, false, true , true , false,
                            false, true , true , false, false,
                            false, false, false, false, false,
                    },
                    {
                            false, false, false, false, false,
                            false, true , false, false, false,
                            false, true , true , true , false,
                            false, false, true , false, false,
                            false, false, false, false, false,
                    },


            },
            {
                    {
                            false, false, false, false, false,
                            false, false, true , false, false,
                            false, true , true , true , false,
                            false, false, true , false, false,
                            false, false, false, false, false,
                    },
                    {
                            false, false, false, false, false,
                            false, false, true , false, false,
                            false, true , true , true , false,
                            false, false, true , false, false,
                            false, false, false, false, false,
                    },
                    {
                            false, false, false, false, false,
                            false, false, true , false, false,
                            false, true , true , true , false,
                            false, false, true , false, false,
                            false, false, false, false, false,
                    },
                    {
                            false, false, false, false, false,
                            false, false, true , false, false,
                            false, true , true , true , false,
                            false, false, true , false, false,
                            false, false, false, false, false,
                    },
            },
            {
                    {
                            false, false, false, false, false,
                            false, false, true , false, false,
                            false, true , true , true , true ,
                            false, false, false, false, false,
                            false, false, false, false, false,
                    },
                    {
                            false, false, false, false, false,
                            false, false, true , false, false,
                            false, false, true , true , false,
                            false, false, true , false, false,
                            false, false, true , false, false,
                    },
                    {
                            false, false, false, false, false,
                            false, false, false, false, false,
                            true , true , true , true , false,
                            false, false, true , false, false,
                            false, false, false, false, false,
                    },
                    {
                            false, false, true , false, false,
                            false, false, true , false, false,
                            false, true , true , false, false,
                            false, false, true , false, false,
                            false, false, false, false, false,
                    },
           },
    },
}


func main() {
  for i := 0; i < kndBlock.height; i++{
    for j := 0; j < kndBlock.width; j++{
      fmt.Print(kndBlock.array[0][0][i*kndBlock.width+j]);
      fmt.Print(" ");
    }
    fmt.Println("");
  }
}
