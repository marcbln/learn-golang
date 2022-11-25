package handlers

import (
	"github.com/julienschmidt/httprouter"
	"mymodule10-application-monitoring/metrics"
	"net/http"
)

// eine closure in golang: here we can inject stuff (eg database)
func GetRoot(msg string, mc *metrics.MetricsCollection) HttpHandler {
	const counterName = "/"
	mc.RegisterCounter(counterName)

	return func(rw http.ResponseWriter, req *http.Request, _ httprouter.Params) {
		defer mc.IncrementCounter(counterName)
		rw.WriteHeader(http.StatusOK)
		rw.Write([]byte(msg))
	}
}
