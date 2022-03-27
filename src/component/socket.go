package component

import (
	"time"

	"github.com/tomhjx/idea/metric"
)

// tcp packet head: 22 bytes
// tcp packet payload: 1400 bytes

type Socket struct {
	network *NetWork
}

func NewSocket(network *NetWork) *Socket {
	return &Socket{network: network}
}

func (i *Socket) Bind() (duration time.Duration, resources *metric.Resources) {
	return duration, resources
}

func (i *Socket) Listen() (duration time.Duration, resources *metric.Resources) {
	return duration, resources
}

func (i *Socket) Connect() (duration time.Duration, resources *metric.Resources) {
	return duration, resources
}

func (i *Socket) Accept() (duration time.Duration, resources *metric.Resources) {
	return duration, resources
}

func (i *Socket) Send(bytes int) (duration time.Duration, resources *metric.Resources) {
	kbytes := float64(bytes / 1024)
	d := kbytes / i.network.UpRate
	duration = time.Duration(d) * time.Second
	return duration, resources
}

func (i *Socket) Recv(bytes int) (duration time.Duration, resources *metric.Resources) {
	kbytes := float64(bytes / 1024)
	d := kbytes / i.network.DownRate
	duration = time.Duration(d) * time.Second
	return duration, resources
}

func (i *Socket) Close() (duration time.Duration, resources *metric.Resources) {
	return duration, resources
}
