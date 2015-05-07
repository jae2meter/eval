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
	var err error;
	var nowTime, maxDiffTime, fileTime int64;

	// Get command line arguments
	// fmt.Println(len(os.Args), os.Args)

	if (os.Args[0] == 1) {
		log.Fatal(err)
	}
	fileName = os.Args[1];
	maxDiffTime, err = strconv.ParseInt (os.Args[2], 10, 64);
	if err != nil {
		log.Fatal(err)
	}

	//  Get current time in seconds since epoch 
	now := time.Now();
	nowTime = now.Unix();// fmt.Printf("Time in seconds is %d\n", nowTime);

	//  Get file modificaton time in seconds since epoch 
	info, err := os.Stat(fileName)
	if err != nil {
		log.Fatal(err)
	}

	// If time diff is larger than command line argument then return 1 
	fileTime = info.ModTime().Unix(); // fmt.Println(info.ModTime().Unix());
	if (nowTime-fileTime) > maxDiffTime {
		fmt.Printf("Diff is big %d\n", (nowTime-fileTime));
		os.Exit (1)
	}
	os.Exit (0)
}
