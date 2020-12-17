package main

import (
	"fmt"
	"time"
)



func main() {

	arr := []interface{}{}

	for i:=0;i<653;i++{
		arr = append(arr, i)
	}

	var mappw = func(in interface{}){
		intt := in.(int)
		fmt.Println(intt,":peek")
	}

	var filterF = func(in interface{}) bool{
		intt := in.(int)
		fmt.Println(intt,":filter")
		return intt%2==0
	}

	//var flatF = func(in interface{}) []interface{}{
	//	intt := in.(int)
	//	fmt.Println(intt,":flat")
	//	return []interface{}{intt+2,intt-3}
	//}

	start:= time.Now()
	output := GetStream(arr).Peek(mappw).Filter(filterF).Execute()
	end := time.Now()

	fmt.Println("done!")
	for _, o := range output {
		it := o.(int)
		fmt.Println(it)
	}

	fmt.Println(end.Sub(start).Seconds())

}
