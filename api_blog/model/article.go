package model

type CourseArticle struct {
	Title       string  `json:"title" gorm:"unique"`          // 标题
	Category    string  `json:"category"`                     // 分类
	Description string  `json:"description"`                  // 描述
	Content     string  `json:"content" gorm:"type:longtext"` // 内容
	Price       float64 `json:"price" gorm:"type:decimal"`    // 价格
	SaleNum     int     `json:"sale_num"`                     // 销量
	GoodNum     int     `json:"good_num"`                     // 点赞数量
	MoneyNum    int     `json:"money_num"`                    // 打赏数量
	ViewNum     int     `json:"view_num"`                     // 查看数量
	AddTime     int     `json:"add_time"`                     // 添加时间
	UpdateTime  int     `json:"update_time"`                  // 修改时间
}
