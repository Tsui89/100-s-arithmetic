package main

import (
	"github.com/Tsui89/100-s-arithmetic/mysort"
	"sync"
	"runtime"
	"errors"
)

type sortData struct {
	data []int
}

func (sd sortData)Len()int{
	return len(sd.data)
}

func (sd sortData)Bigger(i,j int)(bool,error){
	if i>sd.Len() || j >sd.Len(){
		return false,errors.New("out of range")
	}
	return sd.data[i]>sd.data[j],nil
}

func (sd sortData)Swap(i,j int)error{
	if i>sd.Len() || j >sd.Len(){
		return errors.New("out of range")
	}
	sd.data[i],sd.data[j] = sd.data[j],sd.data[i]
	return nil
}

func main(){
	runtime.GOMAXPROCS(4)

	var w sync.WaitGroup
	var sd1 sortData
	var sd2 sortData
	var sd3 sortData

	length := 10*1000
	data := mysort.PreData(length)
	sd1.data=data
	sd2.data=data
	sd3.data=data


	w.Add(1)
	go func() {
		mysort.Bubble(sd1)
		w.Done()
	}()

	//w.Add(1)
	//go func() {
	//	mysort.Quick(sd2)
	//	w.Done()
	//}()
	//
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