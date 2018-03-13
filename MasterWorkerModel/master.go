package master_worker

import (
	"fmt"
)

type Master struct {
	JobQueue chan Job
	Workers []Worker
	isStart bool
}

func NewMasterAndStartWorker(workerSize,jobQueueSize int) *Master {
	master := &Master{
		Workers:[]Worker{},
		isStart:false,
	}
	if master.JobQueue == nil {
		master.JobQueue = make(chan Job,jobQueueSize)
	}
	for i:=0;i<workerSize;i++ {
		worker := Worker{}
		master.Workers = append(master.Workers,worker)
		worker.startWork(master)
	}
	if workerSize >0 {
		master.isStart = true
	}
	return master
}

func NewMaster(workerSize,jobQueueSize int) *Master {
	master := &Master{
		Workers:[]Worker{},
		isStart:false,
	}
	if master.JobQueue == nil {
		master.JobQueue = make(chan Job,jobQueueSize)
	}
	for i:=0;i<workerSize;i++ {
		worker := Worker{}
		master.Workers = append(master.Workers,worker)
	}
	return master
}

func (m Master) StartWorkers()  {
	if m.isStart {
		fmt.Println("不重复开启 workers")
		return
	}
	size := len(m.Workers)
	if size >0 {
		m.isStart = true
	}
	for i:=0;i<size;i++ {
		(m.Workers)[i].startWork(&m)
	}
}

func (m Master) AddJob(job Job)  {
	fmt.Println("add a job "+job.Tag)
	m.JobQueue <- job
}






































