package component

import (
	"time"

	"github.com/tomhjx/idea/metric"
)

type HttpServer interface {
	HttpSpec() *HttpSpec
	Response() (time.Duration, *metric.Resources)
}

type HttpSpec struct {
	TLS     string //1.1, 1.2, 1.3
	Version string //1.1
}
