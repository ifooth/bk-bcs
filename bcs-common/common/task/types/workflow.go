package types

// Workflow 支持FinalTask的任务流
type Workflow struct {
	Task      *Task
	FinalTask *Task
}
