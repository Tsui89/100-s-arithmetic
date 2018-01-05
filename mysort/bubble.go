package mysort

import (
	"math/rand"
	"time"
	"../timeCost"
	"fmt"
)

func PreData(len int)([]int){
	var data []int
	r:=rand.New(rand.NewSource(time.Now().UnixNano()))
	for i:=0;i<len;i++{
		data = append(data,r.Intn(len))
	}
	return data
}

func Bubble(data []int)[]int  {
	defer fmt.Println("Bubble Sort: Used seconds ",timeCost.Cost(time.Now()))
	for i:=0;i<len(data);i++{
		for j:=i+1;j<len(data);j++{
			if data[i] > data[j]{
				data[i],data[j] = data[j],data[i]
			}
		}
	}
	return data
}