package models

type Nav struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Url     int    `json:"url"`
	State   int    `json:"state"`
	Sort    int    `json:"sort"`
	AddTime int64  `json:"add_time"`
}

func (u Nav) TableName() string {
	return "nav"
}
