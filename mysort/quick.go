package mysort

import (
	"fmt"
	"time"
	"github.com/Tsui89/100-s-arithmetic/timeCost"
)

func quickSort(data Sort,l,r int){

	if l>=r{
		return
	}

	//取中间位置的值
	//base := vd[(l+r)/2]

	//将中间值换到起点位置
	//data[l],data[(l+r)/2] = data[(l+r)/2],data[l]
	data.Swap(l,(l+r)/2)
	//base:=data[l]
	i,j:=l,r
	for i<j{
		//从右侧开始找小值
		for i<j&&data.Bigger(j,l){j--}
		//将小值放到起点
		//if i<j&&!data.Bigger(j,l){
		//	vd[i]=vd[j]
		//	i++
		//}
		//从左侧开始找大值
		for i<j&&data.Bigger(l,i){i++}
		//将大值放到j位置
		//if i<j&&!data.Bigger(l,i){
		//	vd[j]=vd[i]
		//}
		if i<j{
				data.Swap(i,j)
		}
	}
	data.Swap(i,l)
	//vd[i] = base
	//data=Sort{vd}
	quickSort(data,l,i-1)
	quickSort(data,i+1,r)
}

func Quick(data Sort)Sort{
	t :=time.Now()

	//defer fmt.Println("Quick Sort: Used seconds ",timeCost.Cost(t))
	quickSort(data,0,data.Len()-1)
	fmt.Println("Quick Sort: Used seconds ",timeCost.Cost(t))
	return data
}
