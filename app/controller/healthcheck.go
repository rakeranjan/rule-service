package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"rule-service/app/serializer"
	"rule-service/config"
)

//HealthCheck is a function for HealthCheck url
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	m := make(map[string]bool)
	m["database"] = checkDatabaseConnection()
	data := &serializer.JSONResponse{
		Data: m,
		Meta: responseMeta(m),
	}
	log.Println(data.Meta)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(responseStatus(m))
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}

func checkDatabaseConnection() bool {
	response := true
	if config.Getdb() == nil {
		response = false
	}
	return response
}

func responseMeta(m map[string]bool) string {
	var msg string
	if responseStatus(m) == 200 {
		msg = "successful"
	} else {
		msg = "HealthCheck failed"
	}
	return msg
}

func responseStatus(m map[string]bool) int {
	res := true
	var status_code int
	for _, v := range m {
		res = res && v
	}
	if res {
		status_code = 200
	} else {
		status_code = 500
	}
	return status_code
}
