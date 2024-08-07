package model

import gorm "github.com/zhangdapeng520/zdpgo_gorm"

type PythonArticle struct {
	gorm.Model
	Title       string `json:"title"`
	Description string `json:"description"`
	Content     string `json:"content" gorm:"type:longtext"`
	Author      string `json:"string"`
	GoodNum     int    `json:"good_num"`
	BadNum      int    `json:"bad_num"`
	ViewNum     int    `json:"view_num"`
}

type GolangArticle struct {
	gorm.Model
	Title       string `json:"title"`
	Description string `json:"description"`
	Content     string `json:"content" gorm:"type:longtext"`
	Author      string `json:"string"`
	GoodNum     int    `json:"good_num"`
	BadNum      int    `json:"bad_num"`
	ViewNum     int    `json:"view_num"`
}

type WebArticle struct {
	gorm.Model
	Title       string `json:"title"`
	Description string `json:"description"`
	Content     string `json:"content" gorm:"type:longtext"`
	Author      string `json:"string"`
	GoodNum     int    `json:"good_num"`
	BadNum      int    `json:"bad_num"`
	ViewNum     int    `json:"view_num"`
}

type DatabaseArticle struct {
	gorm.Model
	Title       string `json:"title"`
	Description string `json:"description"`
	Content     string `json:"content" gorm:"type:longtext"`
	Author      string `json:"string"`
	GoodNum     int    `json:"good_num"`
	BadNum      int    `json:"bad_num"`
	ViewNum     int    `json:"view_num"`
}

type CloudNativeArticle struct {
	gorm.Model
	Title       string `json:"title"`
	Description string `json:"description"`
	Content     string `json:"content" gorm:"type:longtext"`
	Author      string `json:"string"`
	GoodNum     int    `json:"good_num"`
	BadNum      int    `json:"bad_num"`
	ViewNum     int    `json:"view_num"`
}

type CourseArticle struct {
	gorm.Model
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Content     string  `json:"content" gorm:"type:longtext"`
	Author      string  `json:"string"`
	Price       float64 `json:"price" gorm:"type:decimal"`
	SaleNum     int     `json:"sale_num"`
	GoodNum     int     `json:"good_num"`
	BadNum      int     `json:"bad_num"`
	ViewNum     int     `json:"view_num"`
}

type OtherArticle struct {
	gorm.Model
	Title       string `json:"title"`
	Description string `json:"description"`
	Content     string `json:"content" gorm:"type:longtext"`
	Author      string `json:"string"`
	GoodNum     int    `json:"good_num"`
	BadNum      int    `json:"bad_num"`
	ViewNum     int    `json:"view_num"`
}
