package metric

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResources(t *testing.T) {

	cpu1 := 1
	cpu2 := 10

	r1 := &Resources{CPUs: cpu1}
	r2 := &Resources{CPUs: cpu2}
	r1.Add(r2)
	assert.Equal(t, cpu1+cpu2, r1.CPUs)

	t.Logf("r1.CPUs: %d", r1.CPUs)

}
