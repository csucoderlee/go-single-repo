package main

import "net/http"

type EChartInterface interface {
	Plot(w http.ResponseWriter, r *http.Request)
	Save(string) bool
	Name() string
	Type() string
}
