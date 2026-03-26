package basics

import (
	"fmt"
	"time"
)

// Map funciton are like key value pairs

func studentMap() {
	student := map[string]int{
		"Maths":   80,
		"Science": 79,
	}
	student["english"] = 88

	for subject, marks := range student {
		fmt.Println(subject, marks)
	}
	time.Sleep(time.Second * 15)
}
