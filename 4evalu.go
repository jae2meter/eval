package main

import (
	"fmt"
	"time"
)

func main() {

	var nowTimeInt int64;

	//  Get current time in seconds since epoch 
	now := time.Now();
	nowTimeInt = now.Unix();
	
	fmt.Printf("Time in seconds is %d\n", nowTimeInt);


}
