package service

import (
	"bilibili_demo/define"
	"bilibili_demo/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary 获取提交列表
// @Description 描述
// @Tags 标题组名
// @Param page query int false "请输入分页"  // @Params 参数名 query 类型 是否必填 描述
// @Param size query int false "请输入分页个数"  // @Params 参数名 query 类型 是否必填 描述
// @Param problem_identity query string false "问题唯一值"  // @Params 参数名 query 类型 是否必填 描述
// @Param user_identity query string false "用户唯一值"  // @Params 参数名 query 类型 是否必填 描述
// @Param status query int false "状态"
// @Success 200 {string} json "{"code": 200, "msg":"", data: ""}" // 200成功 返回的 json {""}
// @Router /submit-list [get] // 路由
func GetSubmitList(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", define.DefaultPage))
	size, _ := strconv.Atoi(c.DefaultQuery("size", define.DefaultSize))
	var total int64
	page = (page - 1) * size
	problemIdentity := c.Query("problem_identity")
	user_identity := c.Query("user_identity")
	status, _ := strconv.Atoi(c.Query("status"))

	tx := models.GetSubmitList(problemIdentity, user_identity, status)
	data := make([]models.SubmitBasic, 0)
	err := tx.Count(&total).Offset(page).Limit(size).Find(&data).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"data": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": map[string]interface{}{
			"list":  data,
			"total": total,
		},
	})

}
