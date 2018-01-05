package main

import (
	"./mysort"
	"fmt"
)
func main(){
	data := mysort.PreData(10*1000)

	mysort.Bubble(data)
	fmt.Println(data)
	data = mysort.PreData(10000)
	mysort.Quick(data)
	fmt.Println(data)
}