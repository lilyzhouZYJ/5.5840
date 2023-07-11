package mr

//
// RPC definitions.
//
// remember to capitalize all names.
//

import (
	"os"
	"strconv"
)

// Add your RPC definitions here.

type AssignTaskArgs struct {
}

type AssignTaskReply struct {
	FileName string
	TaskType string
	TaskID   int
	// map specific fields
	NReduce int
	// reduce specific fields
}

type NotifyCompletionArgs struct {
	TaskType  string
	TaskID    int
	FileNames []string // filenames of intermediate files
}

type NotifyCompletionReply struct {
}

// Cook up a unique-ish UNIX-domain socket name
// in /var/tmp, for the coordinator.
// Can't use the current directory since
// Athena AFS doesn't support UNIX-domain sockets.
func coordinatorSock() string {
	s := "/var/tmp/5840-mr-"
	s += strconv.Itoa(os.Getuid())
	return s
}
