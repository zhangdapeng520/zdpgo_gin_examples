package article

import (
	gin "github.com/zhangdapeng520/zdpgo_gin"
	"github.com/zhangdapeng520/zdpgo_gin_examples/api_blog/g"
	"github.com/zhangdapeng520/zdpgo_gin_examples/api_blog/model"
	"strconv"
	"time"
)

func updateCourse(c *gin.Context) {
	// 解析请求参数
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var req requestArticle
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// 查找
	var article model.CourseArticle
	err = g.GDB.Find(&article, "id=?", id).First(&article).Error
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	// 修改
	article.Title = req.Title
	article.Category = req.Category
	article.Description = req.Description
	article.Content = req.Content
	article.Price = req.Price
	article.UpdateTime = int(time.Now().Unix())
	err = g.GDB.Save(&article).Error
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	// 返回
	c.JSON(200, nil)
}
