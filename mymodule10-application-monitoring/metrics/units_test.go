package metrics_test

import (
	"github.com/stretchr/testify/assert"
	"mymodule10-application-monitoring/metrics"
	"testing"
)

func TestUnits(t *testing.T) {
	assert.Equal(t, "milliseconds", metrics.Milliseconds)
}
