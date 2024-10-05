package models

type Article struct {
	ID            int    `json:"id"`
	Title         string `json:"title"`
	ArticleCateId int    `json:"article_cate_id"`
	State         string `json:"state"`
	AddTime       int64  `json:"add_time"`
}

func (u Article) TableName() string {
	return "article"
}
