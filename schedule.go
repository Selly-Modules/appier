package appier

import (
	"fmt"
	"log"
	"time"

	"github.com/robfig/cron/v3"
)

// Job ...
type Job struct {
	Spec string
	Name string
	Cmd  func()
}

// Scheduler ...
type Scheduler struct {
	cron *cron.Cron
	jobs []*Job
}

// New ...
func newSchedule(jobs ...*Job) *Scheduler {
	l, _ := time.LoadLocation("Asia/Ho_Chi_Minh")
	c := cron.New(cron.WithSeconds(), cron.WithLocation(l))

	return &Scheduler{
		cron: c,
		jobs: jobs,
	}
}

// startSchedule ...
func (s *Scheduler) startSchedule() {
	for _, job := range s.jobs {
		if _, err := s.cron.AddFunc(job.Spec, job.Cmd); err != nil {
			log.Fatalf("Add job err: %v", err)
		}
		fmt.Printf("Job %s is started: %s\n", job.Name, job.Spec)
	}
	s.cron.Start()
}

// initSchedule ..
func initSchedule() {
	jobs := []*Job{
		{
			Spec: "*/ * * * *",
			Name: "Start job sync data to service appier",
			Cmd:  SyncToService{}.syncData,
		},
	}

	s := newSchedule(jobs...)
	s.startSchedule()
}
