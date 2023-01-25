package main

import "fmt"

func main(){
	mychannel := make(chan  string)
	anotherchannel := make(chan  string)

	go func(){
		mychannel <- "data"
	}()
	go func(){
		anotherchannel <- "cow"
	}()

	select{
		case msgFromMyChannel := <- mychannel:
			fmt.Println(msgFromMyChannel)
		case msgFromAnotherChannel := <- anotherchannel:
			fmt.Println(msgFromAnotherChannel)
	}
	
	
}