package main

import (
	"Week02/service"
	"fmt"
	"math/rand"
)

func main(){
	s:=service.NewService()
	var id=rand.Intn(1000)
	_,err:=s.GetOrderNameByOrderId(id)
	fmt.Printf("%v\n",err)
}
