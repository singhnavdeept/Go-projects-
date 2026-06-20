package scheduler

import (
	"github.com/robfig/cron/v3"
)

// JobRunner wraps the cron scheduler to run tasks periodically.
type JobRunner struct {
	cronManager *cron.Cron
}

// NewJobRunner creates a new JobRunner.
func NewJobRunner() *JobRunner {
	// Using standard 5-field cron spec (e.g. "0 */6 * * *")
	return &JobRunner{
		cronManager: cron.New(),
	}
}

// AddSchedule schedules a function to run at the spec interval.
func (j *JobRunner) AddSchedule(spec string, cmd func()) (cron.EntryID, error) {
	return j.cronManager.AddFunc(spec, cmd)
}

// Start starts the cron scheduler in its own goroutine.
func (j *JobRunner) Start() {
	j.cronManager.Start()
}

// Stop stops the cron scheduler, waiting for active jobs to finish.
func (j *JobRunner) Stop() {
	j.cronManager.Stop()
}
