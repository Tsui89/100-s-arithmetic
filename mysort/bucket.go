package mysort

import (
	"fmt"
	"time"
	"../timeCost"
)

func max(data []int)int{
	max:=data[0]
	for _,v:=range data{
		if v >max{
			max =v
		}
	}
	return max
}
type bs struct {
	count int
}
func Bucket(data []int)[]int{
	defer fmt.Println("Bucket Sort: Used seconds ",timeCost.Cost(time.Now()))

	b :=make([]bs,max(data)+1)

	for _,v:=range data{
		b[v].count++
	}

	var r []int
	for i,v:=range b{
		for j:=0; v.count-j>0;j++{
			r= append(r,i)
		}
	}
	return r
}
