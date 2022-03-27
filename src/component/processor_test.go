package component

import (
	"testing"
	"time"

	"github.com/tomhjx/idea/metric"
)

func TestPHPProcessor(t *testing.T) {
	p := NewPHPProcessor("7.4.28", func() (time.Duration, *metric.Resources) {
		return 100 * time.Millisecond, &metric.Resources{
			CPUs:   15,
			Memory: 200,
		}
	})

	d, r := p.Run()
	t.Logf("duration: %v s", d.Seconds())
	t.Logf("cpu: %d m", r.CPUs)
	t.Logf("memory: %d Byte", r.Memory)
}
