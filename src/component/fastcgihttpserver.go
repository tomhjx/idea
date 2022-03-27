package component

import (
	"time"

	"github.com/tomhjx/idea/metric"
	"github.com/tomhjx/idea/support"
)

type FastCGIHttpServer struct {
	httpSpec *HttpSpec
	backend  FastCGIServer
	doAccept func(int, FastCGIServer) (time.Duration, *metric.Resources)
	Requests int
}

func NewFastCGIHttpServer(httpSpec *HttpSpec, backend FastCGIServer, doAccept func(int, FastCGIServer) (time.Duration, *metric.Resources)) *FastCGIHttpServer {
	return &FastCGIHttpServer{
		httpSpec: httpSpec,
		backend:  backend,
		doAccept: doAccept,
	}
}

func (i *FastCGIHttpServer) HttpSpec() *HttpSpec {
	return i.httpSpec
}

func (i *FastCGIHttpServer) accept() (duration time.Duration, resources *metric.Resources) {
	duration, resources = i.doAccept(i.Requests, i.backend)
	return duration, resources
}

func (i *FastCGIHttpServer) Response() (duration time.Duration, resources *metric.Resources) {

	c := &support.Calculator{}
	c.Assemble(i.accept)
	duration, resources = c.Serialize()

	return duration, resources
}
