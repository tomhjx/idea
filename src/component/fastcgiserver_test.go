package component

import (
	"testing"
	"time"

	"github.com/tomhjx/idea/metric"
)

func TestPHPFPMServer(t *testing.T) {
	p := NewPHPProcessor("7.4.28", func() *metric.RunTime {
		r := metric.NewRunTime()
		r.Duration = 100 * time.Millisecond
		r.Resources = &metric.Resources{
			CPUs:   15,
			Memory: 200,
		}
		return r
	})

	fpm := NewPHPFPMServer(&FastCGISpec{Version: "1.1"}, p)
	r := fpm.Response()
	t.Logf("duration: %v s", r.Duration.Seconds())
	t.Logf("cpu: %d m", r.Resources.CPUs)
	t.Logf("memory: %d Byte", r.Resources.Memory)
}
