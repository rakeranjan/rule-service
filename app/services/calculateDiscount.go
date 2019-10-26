package services

import (
	"fmt"
	"log"
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
		log.Println("[INFO]", "INVALID COUPON CODE")
		resp.Msg = "INVALID COUPON CODE"
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
	log.Println("[INFO]", "Discount for Default :-", res[0])
	res = append(res, calculatePrice(userData.OrderAmount, orderAmountsCalculation(userData, ruleData.OrderAmounts)))
	log.Println("[INFO]", "Discount for Order Amount :-", res[1])
	res = append(res, calculatePrice(userData.OrderAmount, userTypeCalculation(userData, ruleData.UserType)))
	log.Println("[INFO]", "Discount for User Type :-", res[2])
	res = append(res, calculatePrice(userData.OrderAmount, referalCalculation(userData, ruleData.Refferal)))
	log.Println("[INFO]", "Discount for Refferal :-", res[3])

	//-----Same way with other parameters discont can be calculated
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

func referalCalculation(userData *user.User, referal *rule.RefferalCountDiscount) *rule.Discount {
	var disc *rule.Discount
	if referal == nil || referal.RefferalDiscounts == nil {
		return disc
	}
	fmt.Println("orders :- ", referal, "2nd leve", referal.RefferalDiscounts)
	for i := 0; i < len(referal.RefferalDiscounts); i++ {
		rangedisc := referal.RefferalDiscounts[i]
		if userData.OrderAmount >= rangedisc.From && userData.OrderAmount < rangedisc.To {
			disc = rangedisc.Discount
		}
	}
	return disc
}

func userTypeCalculation(userData *user.User, userType *rule.UserType) *rule.Discount {
	var disc *rule.Discount

	if userType == nil || (userType.NormalDiscount == nil && userType.PremiumDiscount == nil) {
		fmt.Println("no usertype", userType)
		return disc
	}
	if userData.Type == "normal" {
		disc = userType.NormalDiscount
	} else if userData.Type == "premium" {
		disc = userType.PremiumDiscount
	}
	return disc
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
	fixedDsicount := disc.Fixed
	switch disc.Type {
	case "max":
		value = max(percentDsicount, fixedDsicount)
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
