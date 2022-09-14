package task

import (
	"github.com/robfig/cron"
	"github.com/spf13/viper"
)

const modules = "task_block:"

type Task struct{}

func NewTask() *Task {
	return &Task{}
}

func (t *Task) Run() {
	c := cron.New()
	c.AddFunc(viper.GetString("task.block_times"), t.taskBlock)
	c.Start()
}

func (t *Task) taskBlock() {

	//统计ip

}
