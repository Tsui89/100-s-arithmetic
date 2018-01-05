package mysort

import (
	"fmt"
	"time"
	"../timeCost"
)

func quickSort(data []int,l,r int){
	if l>=r{
		return
	}
	base := data[l]
	i,j:=l,r
	for i<j{
		for ;i<j&&data[j]>=base;j--{}
		for ;i<j&&data[i]<=base;i++{}
		if i<j {
			data[i],data[j] = data[j],data[i]
		}
	}
	data[l],data[i] = data[i],data[l]
	quickSort(data,l,i-1)
	quickSort(data,i+1,r)
}
func Quick(data []int)[]int{
	defer fmt.Println("Quick Sort: Used seconds ",timeCost.Cost(time.Now()))

	quickSort(data,0,len(data)-1)
	return data
}
