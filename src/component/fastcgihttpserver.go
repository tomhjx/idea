package component

import (
	"time"

	"github.com/tomhjx/idea/metric"
)

type FastCGIHttpServer struct {
	httpSpec *HttpSpec
	backend  FastCGIServer
	doAccept func(int, FastCGIServer) *metric.RunTime
	Requests int
}

func NewFastCGIHttpServer(httpSpec *HttpSpec, backend FastCGIServer, doAccept func(int, FastCGIServer) *metric.RunTime) *FastCGIHttpServer {
	return &FastCGIHttpServer{
		httpSpec: httpSpec,
		backend:  backend,
		doAccept: doAccept,
	}
}

func (i *FastCGIHttpServer) HttpSpec() *HttpSpec {
	return i.httpSpec
}

func (i *FastCGIHttpServer) accept() *metric.RunTime {
	r := metric.NewRunTime()
	r.Duration = 10 * time.Millisecond
	r.Resources = &metric.Resources{
		CPUs:   1,
		Memory: 10,
	}
	return r
}

func (i *FastCGIHttpServer) Response() *metric.RunTime {

	r := metric.NewRunTime()
	r.Serialize(i.accept)
	return r
}
