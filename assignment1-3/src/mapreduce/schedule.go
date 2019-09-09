package mapreduce

import (
		"sync"
)

// schedule starts and waits for all tasks in the given phase (Map or Reduce).
func (mr *Master) schedule(phase jobPhase) {
	var ntasks int
	var nios int // number of inputs (for reduce) or outputs (for map)
	switch phase {
	case mapPhase:
		ntasks = len(mr.files)
		nios = mr.nReduce
	case reducePhase:
		ntasks = mr.nReduce
		nios = len(mr.files)
	}

	debug("Schedule: %v %v tasks (%d I/Os)\n", ntasks, phase, nios)

	// All ntasks tasks have to be scheduled on workers, and only once all of
	// them have been completed successfully should the function return.
	// Remember that workers may fail, and that any given worker may finish
	// multiple tasks.
	//
	// TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO
	//

	var waitGroup sync.WaitGroup
	for i:= 0; i < ntasks; i++ {
		waitGroup.Add(1)
		var taskArgs DoTaskArgs
		taskArgs.JobName = mr.jobName
		taskArgs.Phase = phase
		taskArgs.TaskNumber = i
		taskArgs.NumOtherPhase = nios
		if phase == mapPhase {
			taskArgs.File = mr.files[i]
		}

		go func() {

			for {
				worker := <- mr.registerChannel
				if call(worker, "Worker.DoTask", &taskArgs, nil) {
					waitGroup.Done()
					mr.registerChannel <- worker
					break
				}	
			}
		} ()
	
	}
	waitGroup.Wait()

	debug("Schedule: %v phase done\n", phase)
}
