package model

type Account struct {
	Id       int     `json:"id"`
	Username string  `json:"username"`
	Money    float64 `json:"money"`
}

type AccountRecord struct {
	Id         int     `json:"id"`
	Username   string  `json:"username"`
	AddMoney   float64 `json:"add_money" gorm:"type:decimal"`
	OldMoney   float64 `json:"old_money" gorm:"type:decimal"`
	TotalMoney float64 `json:"total_money" gorm:"type:decimal"`
	AddTime    int     `json:"add_time" gorm:"type:int"`
}
