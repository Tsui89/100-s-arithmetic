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
	//取中间位置的值
	base := data[(l+r)/2]

	//将中间值换到起点位置
	data[l],data[(l+r)/2] = data[(l+r)/2],data[l]

	i,j:=l,r
	for i<j{
		//从右侧开始找小值
		for i<j&&data[j]>=base{j--}
		//将小值放到起点
		if i<j&&data[j]<base{
			data[i]=data[j]
			i++
		}
		//从左侧开始找大值
		for i<j&&data[i]<=base{i++}
		//将大值放到j位置
		if i<j&&data[i]>base{
			data[j]=data[i]
		}
	}
	data[i] = base
	quickSort(data,l,i-1)
	quickSort(data,i+1,r)

}
func Quick(data []int)[]int{
	defer fmt.Println("Quick Sort: Used seconds ",timeCost.Cost(time.Now()))
	quickSort(data,0,len(data)-1)
	return data
}
