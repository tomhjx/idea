package metric

import (
	"time"
)

type RunTime struct {
	OperationName string
	Duration      time.Duration
	Resources     *Resources
	Childs        []*RunTime
}

func NewRunTime() *RunTime {
	return &RunTime{
		Duration: 0,
		Resources: &Resources{
			CPUs:   0,
			Memory: 0,
			Disk:   0,
		},
	}
}

// Summation all (e.g. Duration, Resources)
func (i *RunTime) Serialize(handler func() *RunTime) {
	r := handler()
	i.Duration += r.Duration
	i.Resources.Add(r.Resources)
	i.Childs = append(i.Childs, r)
}

// Use max Duration, only summation Resources
func (i *RunTime) Deserialize(handler func() *RunTime) {

	r := handler()
	if i.Duration < r.Duration {
		i.Duration = r.Duration
	}
	i.Resources.Add(r.Resources)
	i.Childs = append(i.Childs, r)
}
