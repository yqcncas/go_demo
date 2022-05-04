package service

import (
	"bilibili_demo/define"
	"bilibili_demo/models"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// GetProblemList
// @Summary 问题列表
// @Description 描述
// @Tags 公共方法
// @Param page query int false "分页页数"
// @Param size query int false "分页个数"
// @Param keyword query string false "关键字"
// @Param category_identity query string false "category_identity"
// @Success 200 {string} json "{"code": 200, data: "", total: ""}"
// @Router /problem [get]
func GetProblemList(c *gin.Context) {
	page, err := strconv.Atoi(c.DefaultQuery("page", define.DefaultPage))
	size, _ := strconv.Atoi(c.DefaultQuery("size", define.DefaultSize))
	keyword := c.Query("keyword")
	category_identity := c.Query("category_identity")
	if err != nil {
		return
	}
	page = (page - 1) * size

	data := make([]*models.ProblemBasic, 0)
	var total int64

	tx := models.GetProblemList(keyword, category_identity)
	err = tx.Count(&total).Offset(page).Limit(size).Find(&data).Error
	if err != nil {
		log.Println("GetProblemList List error", err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": map[string]interface{}{
			"list":  data,
			"total": total,
		},
		"total": total,
	})
	// c.String(http.StatusOK, "成功")
}

// GetProblemListDetail
// @Summary 问题详情
// @Description 描述
// @Tags 公共方法
// @Param identity query string true "详情"
// @Success 200 {string} json "{"code": 200, data: "", total: ""}"
// @Router /problem-detail [get]
func GetProblemListDetail(c *gin.Context) {
	identity := c.Query("identity")
	if identity == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "问题唯一标识不能为空",
		})
		return
	}
	data := models.ProblemBasic{}
	err := models.DB.Preload("ProblemCategories").Preload("ProblemCategories.CategoryBasic").
		Where("identity = ? ", identity).First(&data).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  "查询内容不存在",
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  err.Error(),
			})
		}

		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": data,
	})
}
