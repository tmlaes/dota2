package main

import "time"

func freshTimes() {
	ticker := time.NewTicker(time.Second)
	for range ticker.C {
		times.Set(time.Now().Format(formmat))
	}
}
