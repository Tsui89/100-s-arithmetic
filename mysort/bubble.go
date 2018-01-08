package mysort

import (
	"time"
	"github.com/Tsui89/100-s-arithmetic/timeCost"
	"fmt"
)


func Bubble(data Sort)Sort {
	t:=time.Now()
	//defer fmt.Println("Bubble Sort: Used seconds ",timeCost.Cost(time.Now()))
	for i:=0;i<data.Len();i++{
		for j:=i+1;j<data.Len();j++{
			if b,err:=data.Bigger(i,j);err==nil&&b==true{
				data.Swap(i,j)
			}
		}
	}
	fmt.Println("Bubble Sort: Used seconds ",timeCost.Cost(t))

	return data
}