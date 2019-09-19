package main

import (
	"fmt"
	"github.com/Liberatys/Sanctuary/communication"
)

func main(){
	get := communication.NewGetRequestOverUrl("http://google.de")
	fmt.Println(get.SendRequest())
}