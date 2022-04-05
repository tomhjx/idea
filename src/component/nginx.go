package component

import (
	"time"

	"github.com/tomhjx/idea/metric"
)

type NginxConfig struct {
	WorkerProcesses   int
	WorkerConnections int
	KeepaliveTimeout  int
}

type Nginx struct {
	Version  string
	Config   NginxConfig
	requests int
}

func (i *Nginx) startup() *metric.RunTime {
	// create master
	// create listenfd
	// master fork n worker
	// master wait request
	r := metric.NewRunTime()
	r.Resources = &metric.Resources{
		CPUs:   1,
		Memory: 1024,
	}
	r.Resources.Times(i.Config.WorkerProcesses)
	return r
}

func (i *Nginx) acceptHttp() *metric.RunTime {
	r := metric.NewRunTime()

	// n worker grab a accept mutex
	r.Duration = 10 * time.Millisecond

	// 每个socket占用内存在15~20k之间
	r.Resources = &metric.Resources{
		CPUs:   1,
		Memory: 15 << 10,
	}
	r.Resources.Times(i.requests)
	return r
}

func (i *Nginx) UpstreamHttpPassFastCGI(requests int, backend FastCGIServer) *metric.RunTime {
	i.requests = requests
	r := metric.NewRunTime()
	r.Serialize(i.startup)
	r.Serialize(i.acceptHttp)
	r.Serialize(backend.Response)
	return r
}
