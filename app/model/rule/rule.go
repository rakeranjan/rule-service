package rule

import "fmt"

func Try() {
	fmt.Println("from rule")
}

// RulesData contains all the rules
var rulesData map[string]*Rule

// AddRule add rule to in-memory
func AddRule(coupon string, rule *Rule) {
	if rulesData == nil {
		rulesData = make(map[string]*Rule)
	}
	rulesData[coupon] = rule
}

//GetRule(coupon string) to find the rul from rulesdata
func GetRule(coupon string) *Rule {
	return rulesData[coupon]
}

// RuleData is used to for rule data
type RuleData struct {
	Coupon string `json:"coupon"`
	Rule   *Rule  `json:"rule"`
}

//Rule mode for forming rule
type Rule struct {
	Type                  string                 `json:"type"`
	Refferal              *RefferalCountDiscount `json:"refferal"`
	UserCounts            *UserCounts            `json:"userCounts"`
	MonthlyDayDiscounts   *MonthlyDayDiscount    `json:"monthlyDayDiscounts"`
	MonthDayDiscounts     *MonthDayDiscount      `json:"monthDayDiscounts"`
	WeekDayDiscounts      *WeekDayDiscount       `json:"weekDayDiscounts"`
	BatteryLevelDiscounts *BatteryLevelDiscounts `json:"batteryLevelDiscounts"`
	OrderAmounts          *OrderAmounts          `json:"orderAmounts"`
	UserType              *UserType              `json:"userType"`
	Default               *Default               `json:"default"`
}

// RefferalCountDiscount array of range discounts for no of refferals like 0-10, 11-50, 50-100,
type RefferalCountDiscount struct {
	RefferalDiscounts []*RangeDiscount `json:"refferalDiscounts"`
}

// UserCounts will have Array of UserCount based on count for having various action
type UserCounts struct {
	UserCounts []*UserCount `json:"userCounts"`
}

// UserCount will have discount based on count used for 1st time 1 will used for others repeted will used like for 3 every 3rd will be used
type UserCount struct {
	Count    int       `json:"count"`
	Discount *Discount `json:"discount"`
}

// MonthlyDayDiscount will contains array of MonthlyDiscount for which MONTH & DAYS OF A WEEK discount is applicable
type MonthlyDayDiscount struct {
	MonthlyDayDiscounts []*MonthlyDiscount `json:"monthlyDayDiscounts"`
}

// MonthlyDiscount will contains array of DayDiscount for which MONTH OF YEAR & DAYS OF A MONTH discount is applicable like 1st month of JAN
type MonthlyDiscount struct {
	Month            int               `json:"month"`
	MonthDayDiscount *MonthDayDiscount `json:"monthDayDiscount"`
}

// MonthDayDiscount will contains array of DayDiscount for which DAYS OF A MONTH discount is applicable
type MonthDayDiscount struct {
	MonthDayDiscounts []*DayDiscount `json:"monthDayDiscounts"`
}

// WeekDayDiscount will contains array of DayDiscount for which DAYS OF A WEEK discount is applicable
type WeekDayDiscount struct {
	WeekDayDiscounts []*DayDiscount `json:"weekDayDiscounts"`
}

// DayDiscount has doay discount for day like 1,2 can be used for week or month
type DayDiscount struct {
	Day      int       `json:"day"`
	Discount *Discount `json:"discount"`
}

// BatteryLevelDiscounts array of range discounts for amount like 0-10, 11-50, 50-100
type BatteryLevelDiscounts struct {
	BatteryLevelDiscounts *RangeDiscount `json:"batteryLevelDiscounts"`
}

// OrderAmounts array of range discounts for amount like 0-100, 101-500, 500-1000
type OrderAmounts struct {
	OrderAmounts []*RangeDiscount `json:"orderAmounts"`
}

// RangeDiscount is used wherever range comes to picture for giving discounts
type RangeDiscount struct {
	From     int       `json:"from"`
	To       int       `json:"to"`
	Discount *Discount `json:"discount"`
}

// UserType covers discount specific to UserType
type UserType struct {
	NormalDiscount  *Discount `json:"normalDsicount"`
	PremiumDiscount *Discount `json:"premiumDiscount"`
}

type Default struct {
	Default *Discount `json:"default"`
}

// Discount for type, percent & fixed
type Discount struct {
	Type    string `json:"type"` // can be max, min, fixed percent
	Fixed   int    `json:"fixed"`
	Percent int    `json:"percent"`
}
