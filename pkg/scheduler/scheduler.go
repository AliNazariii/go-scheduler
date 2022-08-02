package scheduler

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"
)

const TimeFormat = time.ANSIC

type Job struct {
	ID        string
	Function  func()
	StartTime time.Time
}

type Scheduler interface {
	AddJob(job Job) string
	CancelJob(id string)
}

type SchedulerImpl struct {
	l        sync.Mutex
	buckets  map[string]map[string]Job
	latestID int
}

func NewScheduler() Scheduler {
	buckets := make(map[string]map[string]Job)
	scheduler := &SchedulerImpl{
		buckets: buckets,
	}
	scheduler.start()
	return scheduler
}

func (s *SchedulerImpl) start() {
	ticker := time.NewTicker(1 * time.Second)

	go func() {
		for {
			select {
			case t := <-ticker.C:
				go func(t time.Time) {
					currentTime := t.Format(TimeFormat)
					fmt.Println(currentTime)
					for _, v := range s.buckets[currentTime] {
						go v.Function()
					}
				}(t)
			}
		}
	}()
}

func (s *SchedulerImpl) CancelJob(id string) {
	timeBucket := strings.Split(id, "$")[0]
	delete(s.buckets[timeBucket], id)
}

func (s *SchedulerImpl) AddJob(job Job) string {
	timeFormat := job.StartTime.Format(TimeFormat)
	s.l.Lock()
	s.latestID++
	id := timeFormat + "$" + strconv.Itoa(s.latestID)
	s.l.Unlock()
	if s.buckets[timeFormat] == nil {
		s.buckets[timeFormat] = map[string]Job{}
	}
	s.buckets[timeFormat][id] = job
	return id
}
