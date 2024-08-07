package article

import (
	gin "github.com/zhangdapeng520/zdpgo_gin"
	"github.com/zhangdapeng520/zdpgo_gin_examples/api_blog/g"
	"github.com/zhangdapeng520/zdpgo_gin_examples/api_blog/model"
	"strconv"
)

func getAllCourse(c *gin.Context) {
	pageStr := c.DefaultQuery("page", "1")
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	sizeStr := c.DefaultQuery("size", "20")
	size, err := strconv.Atoi(sizeStr)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var articles []model.CourseArticle
	err = g.GDB.
		Limit(size).
		Offset((page - 1) * size).
		Order("update_time desc").
		Find(&articles).Error
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, articles)
}

func getCourse(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var article model.CourseArticle
	err = g.GDB.Find(&model.CourseArticle{}, "id=?", id).First(&article).Error
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, article)
}
