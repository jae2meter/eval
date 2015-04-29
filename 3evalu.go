package main

import (
    "fmt"
    "log"
    "os/exec"
    "strconv"

)

func main() {

var nowTimeStr string
//var nowTimeInt int;

//  Get current time in seconds since epoch, string 

    nowTime, err := exec.Command("date", "+%s").Output()
    if err != nil {
        log.Fatal(err)
    }
    nowTimeStr = string (nowTime);

// Convert string to integer
//    nowTimeInt, err := strconv.Atoi(nowTimeStr)
    nowTimeInt, err := strconv.ParseUint( "1430231908\n",0,64)
    nowTimeIntsta, err := strconv.ParseUint(nowTimeStr,0,64)
    fmt.Println(nowTimeIntsta)
//    err := fmt.Sscan(nowTime, &nowTimeInt)
    if err != nil {
        log.Fatal(err)
    }
    
    
    fmt.Printf("Time in seconds is %d\n", nowTimeInt)
}
