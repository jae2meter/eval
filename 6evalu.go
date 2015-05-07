package main

import (
	"fmt"
	"time"
	"log"
	"os"
	"strconv"
)

func main() {
	
	var fileName string;
	fileName = os.Args[1];
	//var maxDiffTime int;
	maxDiffTimeS := os.Args[2];
	maxDiffTime, err := strconv.ParseInt (maxDiffTimeS, 10, 64);
	var nowTime, fileTime int64;

	// fmt.Println (fileName);
	//  Get current time in seconds since epoch 
	now := time.Now();
	nowTime = now.Unix();
	
	// fmt.Printf("Time in seconds is %d\n", nowTime);

	//  Get file modificaton time in seconds since epoch 
	info, err := os.Stat(fileName)
	if err != nil {
		log.Fatal(err)
	}

	fileTime = info.ModTime().Unix();
	// fmt.Println(info.ModTime().Unix());
	//fmt.Println(nowTime)
	//fmt.Println (fileTime - fileTime);
	if (nowTime-fileTime) > maxDiffTime {
		fmt.Printf("Diff is big %d\n", (nowTime-fileTime));
	}
	os.Exit (1)
}
