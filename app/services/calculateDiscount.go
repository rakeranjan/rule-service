package services

import (
	"fmt"
	"rule-service/app/model/rule"
	"rule-service/app/model/user"
	"strings"
)

func Try() {
	fmt.Println("from service")
}

type Response struct {
	Price           int    `json:"price"`
	Discount        int    `json:"discount"`
	DiscountedPrice int    `json:"discountedPrice"`
	Msg             string `json:"msg"`
}

func CalculateDiscount(userData *user.User) *Response {
	resp := Response{}
	ruleData := rule.GetRule(strings.TrimSpace(userData.Coupon))
	if ruleData == nil {
		resp.Msg = "No Coupon code"
		return &resp
	}
	con := decideCondition(userData, ruleData)
	var discount int
	switch ruleData.Type {
	case "max":
		discount = maxDiscount(con)
	case "min":
		discount = maxDiscount(con)
	default:
		discount = maxDiscount(con)
	}
	resp.Price = userData.OrderAmount
	resp.Discount = discount
	resp.DiscountedPrice = resp.Price - resp.Discount
	resp.Msg = "Discount applied"
	return &resp
}

func decideCondition(userData *user.User, ruleData *rule.Rule) []int {
	res := make([]int, 0)
	res = append(res, calculatePrice(userData.OrderAmount, defaultCalculation(ruleData.Default)))
	fmt.Println("res :----", res)
	res = append(res, calculatePrice(userData.OrderAmount, orderAmountsCalculation(userData, ruleData.OrderAmounts)))
	fmt.Println("res :----", res)
	// referal := referalCalculation(ruleData.Refferal)
	// fmt.Println("referal :-", referal)
	// userCount := userCountCalculation(ruleData.UserCounts)
	// fmt.Println("userCount :- ", userCount)
	// monthlyDayDiscount := ruleData.MonthlyDayDiscount
	// monthDayDiscounts := ruleData.MonthDayDiscounts
	// weekDayDiscounts := ruleData.WeekDayDiscounts
	// batteryLevelDiscounts := ruleData.BatteryLevelDiscounts
	// orderAmounts := ruleData.OrderAmounts
	// userType := ruleData.UserType
	// defaultCal := defaultCalculation(ruleData.Default)
	return res
}

func defaultCalculation(defaultData *rule.Default) *rule.Discount {
	var disc *rule.Discount
	if defaultData != nil {
		disc = defaultData.Default
	}
	return disc
}

func orderAmountsCalculation(userData *user.User, order *rule.OrderAmounts) *rule.Discount {
	var disc *rule.Discount
	if order == nil || order.OrderAmounts == nil {
		return disc
	}
	fmt.Println("orders :- ", order, "2nd leve", order.OrderAmounts)
	for i := 0; i < len(order.OrderAmounts); i++ {
		rangedisc := order.OrderAmounts[i]
		if userData.OrderAmount >= rangedisc.From && userData.OrderAmount < rangedisc.To {
			fmt.Println(userData.OrderAmount, rangedisc.From, "-", rangedisc.To, "-", rangedisc.Discount)
			disc = rangedisc.Discount
		}
	}
	return disc
}

func calculatePrice(amount int, disc *rule.Discount) (discount int) {
	var value int
	if amount == 0 || disc == nil {
		return value
	}
	percentDsicount := (amount * disc.Percent) / 100
	fmt.Println("percent :- ", percentDsicount)
	fixedDsicount := disc.Fixed
	fmt.Println("fixed :- ", fixedDsicount)
	switch disc.Type {
	case "max":
		value = max(percentDsicount, fixedDsicount)
		fmt.Println("value :- ", value)
	case "min":
		value = min(percentDsicount, fixedDsicount)
	case "fixed":
		value = fixedDsicount
	case "percent":
		value = percentDsicount
	default:
		value = max(percentDsicount, fixedDsicount)
	}
	return value
}

func max(first, second int) int {
	value := first
	if second > first {
		value = second
	}
	return value
}

func min(first, second int) int {
	value := first
	if second < first {
		value = second
	}
	return value
}

func maxDiscount(discounts []int) int {
	var value int
	for i := 0; i < len(discounts); i++ {
		if discounts[i] > value {
			value = discounts[i]
		}
	}
	return value
}

func MinDiscount(discounts []int) int {
	var value int
	for i := 0; i < len(discounts); i++ {
		if discounts[i] < value {
			value = discounts[i]
		}
	}
	return value
}
