package component

import (
	"time"

	"github.com/tomhjx/idea/metric"
)

type FastCGISpec struct {
	Version string //1.1
}

type FastCGIServer interface {
	FastCGISpec() *FastCGISpec
	Response() *metric.RunTime
}

type PHPFPMServer struct {
	fastCGISpec *FastCGISpec
	processor   Processor
}

func NewPHPFPMServer(fastCGISpec *FastCGISpec, processor Processor) *PHPFPMServer {
	return &PHPFPMServer{
		fastCGISpec: fastCGISpec,
		processor:   processor,
	}
}

func (i *PHPFPMServer) FastCGISpec() *FastCGISpec {
	return i.fastCGISpec
}

func (i *PHPFPMServer) fork() *metric.RunTime {
	r := metric.NewRunTime()
	r.Duration = 10 * time.Millisecond
	r.Resources = &metric.Resources{
		CPUs:   1,
		Memory: 10,
	}
	return r
}

func (i *PHPFPMServer) accept() *metric.RunTime {
	r := metric.NewRunTime()
	r.Duration = 10 * time.Millisecond
	r.Resources = &metric.Resources{
		CPUs:   1,
		Memory: 10,
	}
	return r
}

func (i *PHPFPMServer) process() *metric.RunTime {
	return i.processor.Run()
}

func (i *PHPFPMServer) Response() *metric.RunTime {

	r := metric.NewRunTime()
	r.Serialize(i.fork)
	r.Serialize(i.accept)
	r.Serialize(i.process)
	return r
}
