package utils

import (
	"fmt"
	"runtime"
	"strconv"
	"testing"
	"time"
)

func Test_Job(t *testing.T) {
	queue := NewJobQueue(runtime.NumCPU())
	queue.Start()
	defer queue.Stop()

	for i := 0; i < 4*runtime.NumCPU(); i++ {
		queue.Submit(&TestJob{strconv.Itoa(i)})
	}
}

type TestJob struct {
	ID string
}

func (t *TestJob) Process() {
	fmt.Printf("Processing job '%s'\n", t.ID)
	time.Sleep(1 * time.Second)
}
