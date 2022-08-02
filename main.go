package main

import (
	"fmt"
	"os"
	"os/signal"
	"scheduler/pkg/scheduler"
	"syscall"
	"time"
)

func main() {
	s := scheduler.NewScheduler()
	id1 := s.AddJob(scheduler.Job{
		Function: func() {
			fmt.Println("Hey")
		},
		StartTime: time.Now().Add(6 * time.Second),
	})
	s.AddJob(scheduler.Job{
		Function: func() {
			fmt.Println("Hey2")
		},
		StartTime: time.Now().Add(6 * time.Second),
	})
	s.CancelJob(id1)

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)
	<-sig
}
