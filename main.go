package main

import (
	"github.com/Tsui89/100-s-arithmetic/mysort"
	"sync"
	"runtime"
	"log"
	"os"
)

type sortData []int


func (sd sortData)Len()int{
	return len(sd)
}

func (sd sortData)Bigger(i,j int)(bool){

	return sd[i]>=sd[j]
}

func (sd sortData)Swap(i,j int){

	sd[i],sd[j] = sd[j],sd[i]
}

func main(){
	runtime.GOMAXPROCS(4)
	logger := log.New(os.Stdout,"sort ",log.Ltime)
	var w sync.WaitGroup
	var sd1 sortData
	var sd2 sortData
	//var sd3 sortData

	length := 10*1
	data := mysort.PreData(length)

	sd1 = append(sd1,data...)
	sd2 = append(sd2,data...)
	//sd3 = append(sd1,data...)

	w.Add(1)
	go func() {
		logger.Println(sd1)
		mysort.Bubble(sd1)
		logger.Println(sd1)
		w.Done()
	}()
	w.Add(1)
	go func() {
		logger.Println(sd2)
		mysort.Quick(sd2)
		logger.Println(sd2)
		w.Done()
	}()

	w.Wait()
	//w.Add(1)
	//go func() {
	//	mysort.Bucket(sd3)
	//	w.Done()
	//}()
	//w.Wait()
	//fmt.Println(d1)
	//fmt.Println(d2)
	//for i:=0;i<len(data);i++{
	//	if d1[i]!=sr[i]{
	//		log.Fatalln("not")
	//	}
	//}
}