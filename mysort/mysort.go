package mysort

import (
	"time"
	"math/rand"
)

func PreData(len int)([]int){
	var data []int
	r:=rand.New(rand.NewSource(time.Now().UnixNano()))
	for i:=0;i<len;i++{
		data = append(data,r.Intn(len))
	}
	return data
}

type Sort interface {
	Len()int
	Bigger(i,j int)(bool)
	Swap(i,j int)
}