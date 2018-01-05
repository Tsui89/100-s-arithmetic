package main

import (
	"./mysort"
)
func main(){
	data := mysort.PreData(10*10000)

	mysort.Bubble(data)
	data = mysort.PreData(10000)
	mysort.Quick(data)
}