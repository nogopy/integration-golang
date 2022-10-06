package model

type Phone struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Brand string `json:"brand"`
}

func (a *Phone) TableName() string {
	return "phone"
}
