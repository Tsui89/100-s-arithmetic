package main

import (
	"./mysort"
	"fmt"
)
func main(){
	len := 10*1000
	data := mysort.PreData(len)

	d1:=make([]int,len)
	copy(d1,data)
	mysort.Bubble(d1)
	fmt.Println(d1)

	d2:=make([]int,len)
	copy(d2,data)
	mysort.Quick(d2)
	fmt.Println(d2)
}