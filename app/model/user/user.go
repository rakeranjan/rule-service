package user

// User model
type User struct {
	Coupon         string `json:"coupon"`
	Type           string `json:"type"`
	OrderAmount    int    `json:"orderAmount"`
	LevelOfBattery int    `json:"levelOfBattery"`
	EntryCount     int    `json:"entrtCount"`
	RefferalCount  int    `json:"refferalCount"`
}
