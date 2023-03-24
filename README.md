# go-scheduler

A lightweight and easy-to-use scheduling library written in Go, designed for adding jobs and running them at a specified time. It provides a convenient way to schedule tasks to be executed in the future without the need for manual intervention.

With go-scheduler, you can schedule jobs to run at a specific time or on a recurring basis. It also provides the ability to cancel jobs that have been scheduled but should no longer be run.

The library is designed to be easy to integrate into existing Go applications, with a simple API and minimal dependencies. It uses a priority queue to efficiently manage scheduled tasks and ensure they run in the correct order.

To use go-scheduler, simply create a new scheduler instance and add jobs to it using the AddJob method. You can specify the time when the job should run and the function that should be executed when the job is run. If you need to cancel a job, you can use the CancelJob method.


## Usage
```go
package main

import (
	"fmt"
	"time"
        "scheduler/pkg/scheduler"
)

func main() {
	// Create a new scheduler instance
	scheduler := scheduler.NewScheduler()

	// Schedule a job to run in 5 seconds
        job1 := scheduler.AddJob(scheduler.Job{
		Function: func() {
			fmt.Println("Job 1 executed")
		},
		StartTime: time.Now().Add(5 * time.Second),
	})

	// Schedule a job to run in 10 seconds
	job2 := scheduler.AddJob(scheduler.Job{
		Function: func() {
			fmt.Println("Job 2 executed")
		},
		StartTime: time.Now().Add(10 * time.Second),
	})

	// Cancel the first job
	scheduler.CancelJob(job1)

	// Wait for the jobs to execute
	time.Sleep(15 * time.Second)
}
```

## Contributing
Contributions to this project are welcome. If you find any bugs or have any suggestions, please open an issue or submit a pull request.

