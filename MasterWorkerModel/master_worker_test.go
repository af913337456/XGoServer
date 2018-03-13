package master_worker

import (
	"testing"
	"time"
	"strconv"
)

var index = 0

func TestMaster_StartWorkers(t *testing.T) {

	master := NewMaster(50,500)
	master.StartWorkers()

	go func() {
		for{
			time.Sleep(time.Second*1)
			index++
			master.AddJob(NewJob(strconv.Itoa(index)))
		}
	}()
	time.Sleep(time.Second*300)

}
