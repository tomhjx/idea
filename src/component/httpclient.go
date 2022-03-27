package component

import (
	"time"

	"github.com/tomhjx/idea/metric"
)

type HttpBytes struct {
	Method int
	Host   int
	Uri    int
	Header int
	Body   int
}

type HttpClient struct {
	server HttpServer
	socket *Socket
	req    *HttpBytes
	resp   *HttpBytes
}

func NewHttpClient(network *NetWork, server HttpServer, req *HttpBytes, resp *HttpBytes) *HttpClient {
	return &HttpClient{
		server: server,
		socket: NewSocket(network),
		req:    req,
		resp:   resp,
	}
}

func (i *HttpClient) resolveAddr() *metric.RunTime {
	// dns
	return &metric.RunTime{
		Duration: 10 * time.Millisecond,
	}
}
func (i *HttpClient) connect() *metric.RunTime {
	i.socket.Connect()
	// tls session
	// socket tcp
	return &metric.RunTime{
		Duration: 10 * time.Millisecond,
	}
}
func (i *HttpClient) sendRequest() *metric.RunTime {

	// Host: developer.mozilla.org
	// [Host: ] = 6 bytes
	// [developer.mozilla.org] = `hostBytes`
	// headers
	// body
	bytes := i.req.Method + i.req.Host + i.req.Uri + i.req.Header + i.req.Body
	i.socket.Send(bytes)
	return &metric.RunTime{
		Duration: 10 * time.Millisecond,
	}
}
func (i *HttpClient) recvResponse() *metric.RunTime {
	// headers
	// body
	bytes := i.resp.Header + i.req.Body
	i.socket.Recv(bytes)

	return &metric.RunTime{
		Duration: 10 * time.Millisecond,
	}

}
func (i *HttpClient) close() *metric.RunTime {
	// socket tcp
	i.socket.Close()
	return &metric.RunTime{
		Duration: 10 * time.Millisecond,
	}

}

func (i *HttpClient) Request() *metric.RunTime {

	r := i.resolveAddr()
	r.Serialize(i.connect)
	r.Serialize(i.sendRequest)
	r.Serialize(i.recvResponse)
	r.Serialize(i.close)

	return r
}
