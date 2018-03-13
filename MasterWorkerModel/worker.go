package master_worker




type Worker struct {

}

func (w Worker) startWork(master *Master)  {
	go func() {
		for {
			select {
			case job := <-master.JobQueue:
				job.doJob()
			}
		}
	}()
}








