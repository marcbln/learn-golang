package handlers

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"mymodule10-application-monitoring/metrics"
	"net/http"
)

type healthResponse struct {
	Uptime metrics.ValueWithUnit `json:"uptime"`
	Counts map[string]uint64     `json:"counts"`
}

// eine closure in golang: here we can inject stuff (eg database)
func GetHealth(mc *metrics.MetricsCollection) HttpHandler {

	const counterName = "/health"

	mc.RegisterCounter(counterName)

	return func(rw http.ResponseWriter, req *http.Request, _ httprouter.Params) {
		defer mc.IncrementCounter(counterName)
		healthResponse := healthResponse{
			Uptime: mc.Uptime.GetValueWithUnit(),
			Counts: mc.GetCountersMap(),
		}
		jsonResponse, err := json.Marshal(healthResponse)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}

		rw.Header().Set("Content-Type", "application/json")
		rw.WriteHeader(http.StatusOK)
		rw.Write(jsonResponse)
	}
}
