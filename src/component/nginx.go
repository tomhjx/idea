package component

import (
	"time"

	"github.com/tomhjx/idea/metric"
	"github.com/tomhjx/idea/support"
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

func (i *Nginx) startup() (duration time.Duration, resources *metric.Resources) {
	// create master
	// create listenfd
	// master fork n worker
	// master wait request
	resources = &metric.Resources{
		CPUs:   1,
		Memory: 1024,
	}
	resources.Times(i.Config.WorkerProcesses)

	return duration, resources
}

func (i *Nginx) acceptHttp() (duration time.Duration, resources *metric.Resources) {

	// n worker grab a accept mutex
	duration = 10 * time.Millisecond

	// 每个socket占用内存在15~20k之间
	resources = &metric.Resources{
		CPUs:   1,
		Memory: 15 << 10,
	}

	resources.Times(i.requests)

	return duration, resources
}

func (i *Nginx) UpstreamHttpPassFastCGI(requests int, backend FastCGIServer) (duration time.Duration, resources *metric.Resources) {
	i.requests = requests
	c := &support.Calculator{}
	c.Assemble(i.startup)
	c.Assemble(i.acceptHttp)
	c.Assemble(backend.Response)
	duration, resources = c.Serialize()
	return duration, resources
}
