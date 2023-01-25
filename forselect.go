package main

import "fmt"

func main()  {
     bufferedchannel := make(chan string, 3)
	 chars := []string{"a","b","c"}
	 for _, s := range chars {
		 select{
		 	case bufferedchannel <- s:
		 }
	 }
	 close(bufferedchannel)

	 for result := range bufferedchannel{
		 fmt.Println(result)
	 }
	
}