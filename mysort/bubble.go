package mysort

import (
	"time"
	"github.com/Tsui89/100-s-arithmetic/timeCost"
	"fmt"
)


func Bubble(data []int)[]int  {
	t:=time.Now()
	//defer fmt.Println("Bubble Sort: Used seconds ",timeCost.Cost(time.Now()))
	for i:=0;i<len(data);i++{
		for j:=i+1;j<len(data);j++{
			if data[i] > data[j]{
				data[i],data[j] = data[j],data[i]
			}
		}
	}
	fmt.Println("Bubble Sort: Used seconds ",timeCost.Cost(t))

	return data
}