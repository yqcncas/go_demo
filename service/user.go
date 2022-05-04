package service

import (
	"bilibili_demo/define"
	"bilibili_demo/helper"
	"bilibili_demo/models"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// @Summary 获取用户详情
// @Description 描述
// @Tags 用户
// @Param identity query string false "请输入用户id"  // @Params 参数名 query 类型 是否必填 描述
// @Success 200 {string} json "{"code": 200, data: ""}" // 200成功 返回的 json {""}
// @Router /user-detail [get] // 路由
func GetUserDetail(c *gin.Context) {
	identity := c.Query("identity")
	if identity == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "缺少用户唯一标识",
		})
		return
	}
	data := new(models.UserBasic)

	err := models.DB.Omit("passwrod").Where("identity = ?", identity).Find(&data).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": data,
	})

}

// @Summary 接口概述
// @Description 描述
// @Tags 用户
// @Param username formData string true "用户名"  // @Params 参数名 query 类型 是否必填 描述
// @Param passwrod formData string true "密码"  // @Params 参数名 query 类型 是否必填 描述
// @Success 200 {string} json "{"code": 200, "msg":"", data: ""}" // 200成功 返回的 json {""}
// @Router /login [post] // 路由
func Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("passwrod")
	if password == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"data": "必填项为空",
		})
		return
	}
	password = helper.GetMD5(password)
	data := new(models.UserBasic)
	print(username, password)
	// print()
	err := models.DB.Where("name = ? AND passwrod = ?", username, password).First(&data).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "用户名或密码错误",
		})
		return
	}
	token, err := helper.MakeToken(data.Identity, data.Name)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": map[string]interface{}{
			"token": token,
		},
	})

}

// @Summary 接口概述
// @Description 描述
// @Tags 用户
// @Param email formData string true "邮箱"  // @Params 参数名 query 类型 是否必填 描述
// @Success 200 {string} json "{"code": 200, "msg":"", data: ""}" // 200成功 返回的 json {""}
// @Router /email [post] // 路由
func SendEmail(c *gin.Context) {
	email := c.PostForm("email")
	if email == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "参数不正确",
		})
		return
	}
	code := helper.GetRandom()
	models.RDB.Set(c, email, code, time.Second*300)
	err := helper.SendEmail(email, code)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "发送成功",
	})
}

// @Summary 接口概述
// @Description 描述
// @Tags 用户
// @Param mail formData string true "邮箱"  // @Params 参数名 query 类型 是否必填 描述
// @Param code formData string true "邮箱"  // @Params 参数名 query 类型 是否必填 描述
// @Param name formData string true "邮箱"  // @Params 参数名 query 类型 是否必填 描述
// @Param passwrod formData string true "邮箱"  // @Params 参数名 query 类型 是否必填 描述
// @Param phone formData string true "邮箱"  // @Params 参数名 query 类型 是否必填 描述
// @Success 200 {string} json "{"code": 200, "msg":"", data: ""}" // 200成功 返回的 json {""}
// @Router /register [post] // 路由
func Register(c *gin.Context) {

	mail := c.PostForm("mail")
	code := c.PostForm("code")
	name := c.PostForm("name")
	passwrod := c.PostForm("passwrod")
	phone := c.PostForm("phone")
	if mail == "" || code == "" || name == "" || passwrod == "" || phone == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "校验不通过",
		})
		return
	}

	sysCode, err := models.RDB.Get(c, mail).Result()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  err.Error(),
		})
		return
	}
	if sysCode != code {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "验证码不正确",
		})
		return
	}
	var cnt int64
	var user = &models.UserBasic{}

	err = models.DB.Model(user).Where("mail = ? ", mail).Count(&cnt).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  err,
		})
		return
	}
	if cnt > 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "用户已存在",
		})
		return
	}

	uuid := helper.GetUUID() // 为了获取每次不一样的值
	data := &models.UserBasic{
		Identity: uuid,
		Name:     name,
		Passwrod: helper.GetMD5(passwrod),
		Phone:    phone,
		Mail:     mail,
	}
	err = models.DB.Create(data).Error // 插入数据库
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  err.Error(),
		})
		return
	}
	// 生成token
	token, err := helper.MakeToken(uuid, name)

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": map[string]interface{}{
			"token": token,
		},
	})

}

// @Summary 接口概述
// @Description 描述
// @Tags 用户
// @Param page query int false "分页"  // @Params 参数名 query 类型 是否必填 描述
// @Param size query int false "分页"  // @Params 参数名 query 类型 是否必填 描述
// @Success 200 {string} json "{"code": 200, "msg":"", data: ""}" // 200成功 返回的 json {""}
// @Router /rank-list [get] // 路由
func GetRankList(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", define.DefaultPage))
	size, err := strconv.Atoi(c.DefaultQuery("size", define.DefaultSize))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "参数不正确",
		})
		return
	}
	var total int64

	page = (page - 1) * size

	list := make([]*models.UserBasic, 0)
	err = models.DB.Model(new(models.UserBasic)).
		Count(&total).
		Order("finish_problem_num DESC, submit_num ASC").
		Offset(page).Limit(size).Find(&list).Error

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg": map[string]interface{}{
			"list":  list,
			"total": total,
		},
	})
}
