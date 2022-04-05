package component

import (
	"testing"
	"time"

	"github.com/tomhjx/idea/metric"
)

func TestFastCGIHttpServerOnNginxFPM(t *testing.T) {

	requests := 100000
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

	ngx := &Nginx{Version: "1.21.6", Config: NginxConfig{WorkerProcesses: 1}}
	server := NewFastCGIHttpServer(&HttpSpec{Version: "1.1"}, fpm, ngx.UpstreamHttpPassFastCGI)
	server.Requests = requests
	r := server.Response()
	t.Logf("duration: %v s", r.Duration.Seconds())
	t.Logf("cpu: %d m", r.Resources.CPUs)
	t.Logf("memory: %d Byte", r.Resources.Memory)

}
