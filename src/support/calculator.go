package support

import (
	"time"

	"github.com/tomhjx/idea/metric"
)

type Calculator struct {
	list []func() (time.Duration, *metric.Resources)
}

func (i *Calculator) Assemble(do func() (time.Duration, *metric.Resources)) *Calculator {
	i.list = append(i.list, do)
	return i
}

func (i *Calculator) Serialize() (duration time.Duration, resources *metric.Resources) {

	resources = &metric.Resources{}

	for _, f := range i.list {
		d, r := f()
		duration += d
		resources.Add(r)
	}

	return duration, resources
}

func (i *Calculator) Deserialize() (duration time.Duration, resources *metric.Resources) {

	resources = &metric.Resources{}

	for _, f := range i.list {
		d, r := f()
		if duration < d {
			duration = d
		}
		resources.Add(r)
	}

	return duration, resources
}
