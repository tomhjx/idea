package component

import (
	"log"
	"time"

	"github.com/tomhjx/idea/metric"
)

type ProgrammingSpec struct {
	Type    string //php, golang
	Version string //1.17
}

type Processor interface {
	ProgrammingSpec() *ProgrammingSpec
	Run() *metric.RunTime
}

type PHPProcessor struct {
	programmingSpec *ProgrammingSpec
	logic           func() *metric.RunTime
}

func NewPHPProcessor(version string, logic func() *metric.RunTime) *PHPProcessor {
	return &PHPProcessor{
		programmingSpec: &ProgrammingSpec{Type: "PHP", Version: version},
		logic:           logic,
	}
}

func (i *PHPProcessor) ProgrammingSpec() *ProgrammingSpec {
	return i.programmingSpec
}

func (i *PHPProcessor) zendCompile() *metric.RunTime {
	log.Println("PHP Processor Compile By Zend.")
	r := metric.NewRunTime()
	r.Duration = 100 * time.Millisecond
	r.Resources = &metric.Resources{
		CPUs:   10,
		Memory: 10,
	}
	return r
}

func (i *PHPProcessor) zendExecute() *metric.RunTime {
	log.Println("PHP Processor Execute By Zend.")
	r := metric.NewRunTime()
	r.Duration = 100 * time.Millisecond
	r.Resources = &metric.Resources{
		CPUs:   10,
		Memory: 10,
	}
	return r
}

func (i *PHPProcessor) Run() *metric.RunTime {
	r := metric.NewRunTime()
	r.Serialize(i.zendCompile)
	r.Serialize(i.zendExecute)
	r.Serialize(i.logic)
	return r
}
