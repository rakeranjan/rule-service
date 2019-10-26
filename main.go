package main

import (
	"net/http"
	"rule-service/app/controller"
	"rule-service/app/controller/v1"
	// "rule-service/config"
	// "app/controller/v1"

	"github.com/gorilla/mux"
)

func main() {
	// configuration := config.Config()
	// fmt.Println(configuration.ProgramName, " started")
	// rule.Try()
	// db := config.Getdb().Db.Collection("rule")
	// fmt.Printf("%T\n", db)

	route := mux.NewRouter()
	// route.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// 	w.WriteHeader(http.StatusNotFound)
	// })
	route.HandleFunc("/healthcheck", controller.HealthCheck).Methods("GET")
	route.HandleFunc("/rules", v1.CreateRule).Methods("POST")
	route.HandleFunc("/rules/get_price", v1.GetPrice).Methods("POST")
	route.HandleFunc("/rules/{coupon}", v1.GetRule).Methods("GET")
	// // route.HandleFunc("/healthcheck", timed(controller.HealthCheck)).Methods("GET")
	// route.HandleFunc("/api/v1/payment", v1.Create).Methods("POST")
	http.Handle("/", route)
	http.ListenAndServe(":3000", nil)
}

// func logRequestResponse(f func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		start := time.Now()
// 		f(w, r)
// 		end := time.Now()
// 		fmt.Println("the request took", end.Sub(start))
// 	}
// }
