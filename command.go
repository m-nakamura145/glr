package main

import (
	"fmt"
	"github.com/robfig/cron"
)

type Sync struct{}

func (s *Sync) Help() string {
	return "glr sync Help"
}

func (s *Sync) Run(args []string) int {
	_, err := SyncRepository()

	if err != nil {
		return 1
	}

	return 0
}

func (s *Sync) Synopsis() string {
	return "Synchronize local git repository with remote at once"
}

// status have repository path and schedule cron string
type status struct {
	c            *cron.Cron
	repositories []string
	schedules    []string
}

var sharedStatus *status = newStatus()

func newStatus() *status {
	return &status{
		c: cron.New(),
	}
}

// singleton
func GetStatus() *status {
	return sharedStatus
}

type StatusStart struct{}

func (s *StatusStart) Help() string {
	return "glr status start Help"
}

func (s *StatusStart) Run(args []string) int {

	status := GetStatus()
	status.c.AddFunc("@hourly", func() {
		SyncRepository()
	})
	fmt.Println("set job schedule")
	status.c.Start()

	return 0
}

func (s *StatusStart) Synopsis() string {
	return "Synchronize local git repository with remote"
}

type StatusStop struct{}

func (s *StatusStop) Help() string {
	return "glr status stop Help"
}

func (s *StatusStop) Run(args []string) int {

	status := GetStatus()
	status.c.Stop()
	fmt.Println("stop job scheduler")

	return 0
}

func (s *StatusStop) Synopsis() string {
	return "Stop cron scheduler"
}
