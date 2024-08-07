package article

import (
	gin "github.com/zhangdapeng520/zdpgo_gin"
	"github.com/zhangdapeng520/zdpgo_gin_examples/api_blog/g"
	"github.com/zhangdapeng520/zdpgo_gin_examples/api_blog/model"
	"strconv"
)

func deleteCourse(c *gin.Context) {
	// 解析请求参数
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// 删除
	err = g.GDB.Delete(&model.CourseArticle{}, "id=?", id).Error
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	// 返回
	c.JSON(200, nil)
}
