package master_worker

import (
	"fmt"
	"time"
)

type Job struct {
	Tag string
}

func NewJob(tag string) Job {
	return Job{tag}
}

func (j Job) doJob()  {
	go func() {
		time.Sleep(time.Second*5)
		fmt.Println(j.Tag+" --- doing job...")
	}()
}