package component

import (
	"log"
	"time"

	"github.com/tomhjx/idea/metric"
	"github.com/tomhjx/idea/support"
)

type ProgrammingSpec struct {
	Type    string //php, golang
	Version string //1.17
}

type Processor interface {
	ProgrammingSpec() *ProgrammingSpec
	Run() (time.Duration, *metric.Resources)
}

type PHPProcessor struct {
	programmingSpec *ProgrammingSpec
	logic           func() (time.Duration, *metric.Resources)
}

func NewPHPProcessor(version string, logic func() (time.Duration, *metric.Resources)) *PHPProcessor {
	return &PHPProcessor{
		programmingSpec: &ProgrammingSpec{Type: "PHP", Version: version},
		logic:           logic,
	}
}

func (i *PHPProcessor) ProgrammingSpec() *ProgrammingSpec {
	return i.programmingSpec
}

func (i *PHPProcessor) zendCompile() (duration time.Duration, resources *metric.Resources) {
	log.Println("PHP Processor Compile By Zend.")
	duration = 100 * time.Millisecond

	resources = &metric.Resources{
		CPUs:   10,
		Memory: 10,
	}
	return duration, resources
}

func (i *PHPProcessor) zendExecute() (duration time.Duration, resources *metric.Resources) {
	log.Println("PHP Processor Execute By Zend.")

	duration = 100 * time.Millisecond

	resources = &metric.Resources{
		CPUs:   10,
		Memory: 32 * 1024 * 1024,
	}
	return duration, resources
}

func (i *PHPProcessor) Run() (duration time.Duration, resources *metric.Resources) {
	c := &support.Calculator{}
	duration, resources = c.Assemble(i.zendCompile).Assemble(i.zendExecute).Assemble(i.logic).Serialize()
	return duration, resources
}
