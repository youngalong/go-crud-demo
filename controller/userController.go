package controller

import (
	"github.com/gin-gonic/gin"
	"go-crud-demo/model"
	"go-crud-demo/utils"
	"net/http"
	"strconv"
)

// 增
func AddUser(context *gin.Context) {
	var user *model.User
	context.BindJSON(&user)
	err := utils.DB.Create(&user).Error
	if err != nil {
		context.JSON(http.StatusOK, gin.H{
			"code": 400,
			"data": err,
			"msg":  "增加失败",
		})
	} else {
		context.JSON(http.StatusOK, gin.H{
			"code": 200,
			"data": "",
			"msg":  "增加成功",
		})
	}
}

// 删
func DeleteUser(context *gin.Context) {
	id, ok := context.Params.Get("id")
	if ok {
		utils.DB.Where("id=?", id).Delete(&model.User{})
		context.JSON(http.StatusOK, gin.H{
			"code": 200,
			"data": "",
			"msg":  "删除成功",
		})
	} else {
		context.JSON(http.StatusOK, gin.H{
			"code": 400,
			"data": "",
			"msg":  "未接收到id请求",
		})
	}
}

// 改
func UpdateUser(context *gin.Context) {
	user := &model.User{}
	id := context.Param("id")
	utils.DB.Where("id=?", id).Find(&user)
	if user != nil {
		err := context.BindJSON(&user)
		if err != nil {
			context.JSON(http.StatusOK, gin.H{
				"code": 400,
				"data": err,
				"msg":  "参数错误",
			})
		} else {
			utils.DB.Save(&user)
			context.JSON(http.StatusOK, gin.H{
				"code": 200,
				"data": user,
				"msg":  "修改成功",
			})
		}
	} else {
		context.JSON(http.StatusOK, gin.H{
			"code": 400,
			"data": "",
			"msg":  "未接收到id请求",
		})
	}
}

// 查
func GetUserList(context *gin.Context) {
	var userList []*model.User
	var offSetVal, total int
	pageNo, _ := strconv.Atoi(context.Query("pageNo"))
	pageSize, _ := strconv.Atoi(context.Query("pageSize"))
	if pageNo == 0 && pageSize == 0 {
		offSetVal = -1
	} else {
		offSetVal = (pageNo - 1) * pageSize
	}
	err := utils.DB.Model(&userList).Count(&total).Limit(pageSize).Offset(offSetVal).Find(&userList).Error
	if err != nil {
		context.JSON(http.StatusOK, gin.H{
			"code": 400,
			"data": err,
			"msg":  "查询失败",
		})
	} else {
		context.JSON(http.StatusOK, gin.H{
			"code": 200,
			"data": userList,
			"msg":  "查询成功",
		})
	}
}
func GetUserByName(context *gin.Context) {
	user := &model.User{}
	name, exist := context.Params.Get("name")
	if exist {
		utils.DB.Where("name=?", name).Find(&user)
		if user.ID != 0 {
			context.JSON(http.StatusOK, gin.H{
				"code": 200,
				"data": user,
				"msg":  "查询成功",
			})
		} else {
			context.JSON(http.StatusOK, gin.H{
				"code": 400,
				"data": "",
				"msg":  "查询失败",
			})
		}
	} else {
		context.JSON(http.StatusOK, gin.H{
			"code": 400,
			"data": "",
			"msg":  "未接收到name请求",
		})
	}
}
