package component

import (
	"time"

	"github.com/tomhjx/idea/metric"
	"github.com/tomhjx/idea/support"
)

type FastCGISpec struct {
	Version string //1.1
}

type FastCGIServer interface {
	FastCGISpec() *FastCGISpec
	Response() (time.Duration, *metric.Resources)
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

func (i *PHPFPMServer) fork() (duration time.Duration, resources *metric.Resources) {
	duration = 10 * time.Millisecond
	resources = &metric.Resources{
		CPUs:   1,
		Memory: 10,
	}
	return duration, resources
}

func (i *PHPFPMServer) accept() (duration time.Duration, resources *metric.Resources) {
	duration = 10 * time.Millisecond
	resources = &metric.Resources{
		CPUs:   1,
		Memory: 10,
	}
	return duration, resources
}

func (i *PHPFPMServer) Response() (duration time.Duration, resources *metric.Resources) {
	c := &support.Calculator{}
	c.Assemble(i.fork)
	c.Assemble(i.accept)
	c.Assemble(i.processor.Run)
	duration, resources = c.Serialize()
	return duration, resources
}
