package main

import (
    "fmt"
    "log"
    "os/exec"
    "strconv"

)

func main() {

var nowTimeStr string
//var nowTimeInt int

//  Get current time in seconds since epoch, string 

    s, err := exec.Command("date", "+%s").Output()
    if err != nil {
        log.Fatal(err)
    }
    nowTimeStr = string(s)

// Convert string to integer
    nowTimeInt, err := strconv.Atoi(nowTimeStr); 
    if err != nil {
        log.Fatal(err)
    }
    
    
    fmt.Printf("Time in seconds is %d\n", nowTimeInt)
}
