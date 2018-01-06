package main

import (
	"./mysort"
	"log"
)
func main(){
	lenth := 10*1000
	data := mysort.PreData(lenth)

	d1:=make([]int,lenth)
	copy(d1,data)
	//fmt.Println(d1)
	mysort.Bubble(d1)

	d2:=make([]int,lenth)
	copy(d2,data)
	//fmt.Println(d2)
	mysort.Quick(d2)


	d3:=make([]int,lenth)
	copy(d3,data)
	//fmt.Println(d2)
	sr := mysort.Bucket(d3)
	//fmt.Println(d1)
	//fmt.Println(d2)
	for i:=0;i<len(data);i++{
		if d1[i]!=sr[i]{
			log.Fatalln("not")
		}
	}
}