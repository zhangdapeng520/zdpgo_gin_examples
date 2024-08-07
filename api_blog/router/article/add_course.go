package article

import (
	gin "github.com/zhangdapeng520/zdpgo_gin"
	"github.com/zhangdapeng520/zdpgo_gin_examples/api_blog/g"
	"github.com/zhangdapeng520/zdpgo_gin_examples/api_blog/model"
	"time"
)

func addCourse(c *gin.Context) {
	// 解析请求参数
	var req requestArticle
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// 新增
	now := int(time.Now().Unix())
	err := g.GDB.Create(&model.CourseArticle{
		Title:       req.Title,
		Category:    req.Category,
		Description: req.Description,
		Content:     req.Content,
		Price:       req.Price,
		SaleNum:     0,
		GoodNum:     0,
		MoneyNum:    0,
		ViewNum:     0,
		AddTime:     now,
		UpdateTime:  now,
	}).Error

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	// 返回
	c.JSON(200, nil)
}
