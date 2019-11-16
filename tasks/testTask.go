package tasks

import "time"
import "fmt"

func TestTask() {
	startTime := time.Now().Unix()
	fmt.Println("start time is: ", startTime)
	n := 0
	for ;; {
		fmt.Println("hello world.")
		n ++
		if n > 10 {
			break
		}
	}
	endTime := time.Now().Unix()
	fmt.Println("end time is: ", endTime)
}