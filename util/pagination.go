package util

import (
	"github.com/gin-gonic/gin"
	"github.com/k1/go-blog/conf"
	"github.com/unknwon/com"
)

func GetPage(c *gin.Context) int {
	result := 0
	page, _ := com.StrTo(c.Query("page")).Int()
	if page > 0 {
		result = (page - 1) * conf.Conf.PageSize
	}
	return result
}
