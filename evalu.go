package main

import (
    "fmt"
    "log"
    "os/exec"
)

func main() {
    nowTime, err := exec.Command("date", "+%s").Output()
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("The date is %s\n", nowTime)
}
