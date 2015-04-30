package main

import (
	"fmt"
	"time"
	"log"
	"os"
)

func main() {

	args := os.Args
	fileName := 
	var nowTime, fileTime int64;

	//  Get current time in seconds since epoch 
	now := time.Now();
	nowTime = now.Unix();
	
	fmt.Printf("Time in seconds is %d\n", nowTime);

	//  Get file modificaton time in seconds since epoch 
	info, err := os.Stat("evalu.go")
	if err != nil {
		log.Fatal(err)
	}

	fileTime = info.ModTime().Unix();
//	fmt.Println(info.ModTime().Unix());
	fmt.Println(fileTime)
	fmt.Println (nowTime - fileTime);
	os.Exit (1)
}
