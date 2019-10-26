package v1

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"rule-service/app/model/rule"
	"rule-service/app/model/user"
	"rule-service/app/serializer"
	"rule-service/app/services"
	"strings"

	"github.com/gorilla/mux"
)

// CreateRule controller for rules
func CreateRule(w http.ResponseWriter, r *http.Request) {

	var ruleData rule.RuleData
	decoder := json.NewDecoder(r.Body) //.Decode(&rule)
	err1 := decoder.Decode(&ruleData)
	// fmt.Println("couponCode :-", ruleData.Coupon)
	// fmt.Println("user_type :-", ruleData.Rule.UserType, "normal", ruleData.Rule.UserType.NormalDiscount.Type, "-", ruleData.Rule.UserType.NormalDiscount.Fixed, "-", ruleData.Rule.UserType.NormalDiscount.Percent)
	// fmt.Println("RefferalCountDiscounts :-", ruleData.Rule.Refferal)
	// fmt.Println("UserCounts :-", ruleData.Rule.UserCounts)
	// fmt.Println("MonthlyDayDiscounts :-", ruleData.Rule.MonthlyDayDiscounts)
	// fmt.Println("MonthDayDiscounts :-", ruleData.Rule.MonthDayDiscounts)
	// fmt.Println("WeekDayDiscounts :-", ruleData.Rule.WeekDayDiscounts)
	// fmt.Println("BatteryLevelDiscounts :-", ruleData.Rule.BatteryLevelDiscounts)
	// fmt.Println("OrderAmounts :-", ruleData.Rule.OrderAmounts, "-", ruleData.Rule.OrderAmounts.OrderAmounts, "-", ruleData.Rule.OrderAmounts.OrderAmounts[0].From, "-", ruleData.Rule.OrderAmounts.OrderAmounts[0].To, "=", ruleData.Rule.OrderAmounts.OrderAmounts[0].Discount.Type, "-", ruleData.Rule.OrderAmounts.OrderAmounts[0].Discount.Fixed, "-", ruleData.Rule.OrderAmounts.OrderAmounts[0].Discount.Percent)
	rule.AddRule(ruleData.Coupon, ruleData.Rule)
	a := rule.GetRule(ruleData.Coupon)
	if err1 != nil {
		log.Println("[ERROR]", err1)
	}
	data := &serializer.JSONResponse{
		Data: a,
		Meta: "rule created",
	}
	log.Println("[INFO]--[RULE CONTROLLER]", data.Meta)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(200)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}

func GetRule(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	log.Println("[INFO]", "fetching data for coupon :-", vars["coupon"])
	a := rule.GetRule(strings.TrimSpace(vars["coupon"]))
	data := &serializer.JSONResponse{
		Data: a,
		Meta: "rule created",
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(200)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}

func GetPrice(w http.ResponseWriter, r *http.Request) {
	var userData user.User
	decoder := json.NewDecoder(r.Body)
	err1 := decoder.Decode(&userData)
	fmt.Println(userData)
	if err1 != nil {
		log.Println("[ERROR]", err1)
	}
	res := services.CalculateDiscount(&userData)
	data := &serializer.JSONResponse{
		Data: res,
		Meta: "rule created",
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(200)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}
