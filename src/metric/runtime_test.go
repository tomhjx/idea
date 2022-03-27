package metric

import (
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type RunTimeExample struct {
	duration int
	cpus     int
	memory   int
}

func (i *RunTimeExample) Step1() *RunTime {

	return &RunTime{
		Duration: time.Duration(i.duration) * time.Second,
		Resources: &Resources{
			CPUs:   i.cpus,
			Memory: i.memory,
		},
	}
}
func (i *RunTimeExample) Step2() *RunTime {

	return &RunTime{
		Duration: time.Duration(i.duration) * time.Microsecond,
		Resources: &Resources{
			CPUs:   i.cpus * 2,
			Memory: i.memory * 2,
		},
	}

}

func TestRunTimeForAnonymous(t *testing.T) {

	duration1 := 10 * time.Second
	cpus1 := 10
	memory1 := 1024
	duration2 := 15 * time.Millisecond
	cpus2 := 12
	memory2 := 2 * 1024
	duration3 := 16 * time.Millisecond
	cpus3 := 13
	memory3 := 3 * 1024

	r1 := &RunTime{
		Duration: duration1,
		Resources: &Resources{
			CPUs:   cpus1,
			Memory: memory1,
		},
	}

	r1.Serialize(func() *RunTime {
		return &RunTime{
			Duration: duration2,
			Resources: &Resources{
				CPUs:   cpus2,
				Memory: memory2,
			},
		}
	})
	r1.Serialize(func() *RunTime {
		return &RunTime{
			Duration: duration3,
			Resources: &Resources{
				CPUs:   cpus3,
				Memory: memory3,
			},
		}
	})

	assert.Equal(t, duration1+duration2+duration3, r1.Duration)
	assert.Equal(t, cpus1+cpus2+cpus3, r1.Resources.CPUs)
	assert.Equal(t, memory1+memory2+memory3, r1.Resources.Memory)

	t.Logf("duration: %v s", r1.Duration.Seconds())
	t.Logf("cpu: %d m", r1.Resources.CPUs)
	t.Logf("memory: %d Byte", r1.Resources.Memory)

	r2 := &RunTime{
		Duration: duration1,
		Resources: &Resources{
			CPUs:   cpus1,
			Memory: memory1,
		},
	}

	r2.Deserialize(func() *RunTime {
		return &RunTime{
			Duration: duration2,
			Resources: &Resources{
				CPUs:   cpus2,
				Memory: memory2,
			},
		}
	})
	r2.Deserialize(func() *RunTime {
		return &RunTime{
			Duration: duration3,
			Resources: &Resources{
				CPUs:   cpus3,
				Memory: memory3,
			},
		}
	})

	assert.Equal(t, duration1, r2.Duration)
	assert.Equal(t, cpus1+cpus2+cpus3, r2.Resources.CPUs)
	assert.Equal(t, memory1+memory2+memory3, r2.Resources.Memory)

}

func TestRunTimeForExampleMethod(t *testing.T) {

	example := &RunTimeExample{
		duration: 10,
		cpus:     10,
		memory:   1024,
	}
	r1 := NewRunTime()

	runtime1 := example.Step1()
	runtime2 := example.Step2()

	r1.Serialize(example.Step1)
	r1.Serialize(example.Step2)

	t.Logf("runtime1.Duration:%v, runtime2.Duration:%v", runtime1.Duration, runtime2.Duration)

	assert.Equal(t, runtime1.Duration+runtime2.Duration, r1.Duration)
	assert.Equal(t, runtime1.Resources.CPUs+runtime2.Resources.CPUs, r1.Resources.CPUs)
	assert.Equal(t, runtime1.Resources.Memory+runtime2.Resources.Memory, r1.Resources.Memory)

	t.Logf("duration: %v s", r1.Duration.Seconds())
	t.Logf("cpu: %d m", r1.Resources.CPUs)
	t.Logf("memory: %d Byte", r1.Resources.Memory)

	r2 := NewRunTime()

	r2.Deserialize(example.Step1)
	r2.Deserialize(example.Step2)

	assert.Equal(t, runtime1.Duration, r2.Duration)
	assert.Equal(t, runtime1.Resources.CPUs+runtime2.Resources.CPUs, r2.Resources.CPUs)
	assert.Equal(t, runtime1.Resources.Memory+runtime2.Resources.Memory, r2.Resources.Memory)

	t.Logf("duration: %v s", r2.Duration.Seconds())
	t.Logf("cpu: %d m", r2.Resources.CPUs)
	t.Logf("memory: %d Byte", r2.Resources.Memory)

	// t.Logf("%d", 1<<10) // 1 KB
	// t.Logf("%d", 256<<20) // 256 MiB
	// t.Logf("%d", 8<<30) // 8 GiB

}

var RunTimes map[string]*RunTime

func putRunTimes(key string, val *RunTime) {
	if RunTimes == nil {
		RunTimes = make(map[string]*RunTime)
	}
	RunTimes[key] = val
}

func LocalCall() {
	putRunTimes("local", &RunTime{
		Duration: 10 * time.Second,
	})
}

func RemoteCall() {

	putRunTimes("remote", &RunTime{
		Duration: 10 * time.Second,
	})

}

type Base struct {
	id int
}

func (i *Base) ID() int {
	if i.id == 0 {
		i.id = 1111
	}
	return i.id
}

func (i *Base) String() string {
	return fmt.Sprintf("%s,%d", reflect.TypeOf(i), i.ID())
}

// caller: customer(调用方),supplier(服务方)
type Caller struct {
	Base
	Runtimes map[string]*RunTime
}

func (i *Caller) putRunTimes(key string, val *RunTime) {
	if i.Runtimes == nil {
		i.Runtimes = make(map[string]*RunTime)
	}
	i.Runtimes[key] = val
}

func (i *Caller) local() {
	i.putRunTimes("local", &RunTime{
		Duration: 10 * time.Second,
	})
}

func (i *Caller) remote() {
	i.putRunTimes("remote", &RunTime{
		Duration: 20 * time.Millisecond,
	})
}

func TestMyExample(t *testing.T) {
	// LocalCall()
	// RemoteCall()
	// t.Logf("%v", RunTimes)
	c := &Caller{}
	c.local()
	c.remote()
	t.Logf("%v", c.Runtimes)
	t.Logf("%d", c.ID())
	t.Logf("%d", c.ID())
	t.Logf("%d", c.ID())
	t.Logf("%s", c)
}
