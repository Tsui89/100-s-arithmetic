package main

import (
	"github.com/Tsui89/100-s-arithmetic/mysort"
	"sync"
	"runtime"
)
func main(){
	runtime.GOMAXPROCS(4)

	var w sync.WaitGroup

	lenth := 10*1000
	data := mysort.PreData(lenth)

	d1:=make([]int,lenth)
	copy(d1,data)
	w.Add(1)
	go func() {
		mysort.Bubble(d1)
		w.Done()
	}()

	d2:=make([]int,lenth)
	copy(d2,data)
	w.Add(1)
	go func() {
		mysort.Quick(d2)
		w.Done()
	}()


	d3:=make([]int,lenth)
	copy(d3,data)
	w.Add(1)
	go func() {
		mysort.Bucket(d3)
		w.Done()
	}()
	w.Wait()
	//fmt.Println(d1)
	//fmt.Println(d2)
	//for i:=0;i<len(data);i++{
	//	if d1[i]!=sr[i]{
	//		log.Fatalln("not")
	//	}
	//}
}