package scheduler

import (
	"sync"

	"github.com/robfig/cron/v3"
)

type Scheduler struct {
	cron *cron.Cron
}

var instance *Scheduler
var once sync.Once

func GetScheduler() *Scheduler {
	once.Do(func() {
		instance = &Scheduler{
			cron: cron.New(),
		}
	})
	return instance
}

func (s *Scheduler) GetCron() *cron.Cron {
	return s.cron
}
