package component

import (
	"testing"
	"time"

	"github.com/tomhjx/idea/metric"
)

func TestPHPFPMServer(t *testing.T) {
	p := NewPHPProcessor("7.4.28", func() (time.Duration, *metric.Resources) {
		return 100 * time.Millisecond, &metric.Resources{
			CPUs:   15,
			Memory: 200,
		}
	})

	fpm := NewPHPFPMServer(&FastCGISpec{Version: "1.1"}, p)

	d, r := fpm.Response()
	t.Logf("duration: %v s", d.Seconds())
	t.Logf("cpu: %d m", r.CPUs)
	t.Logf("memory: %d Byte", r.Memory)
}
