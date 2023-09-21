package controller

import (
	// "log"

	"gin/db"
	"gin/orm"
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func TestFunc(c *gin.Context) {

	c.JSON(200, gin.H{"msg": "hello world"})

}

func GetUser(c *gin.Context) {

	rows, err := db.DB.Query("select id,name,phone from users limit 10")

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	defer rows.Close()

	// 创建一个切片用于存储查询结果
	var users []gin.H

	// 迭代查询结果，将数据添加到切片中
	for rows.Next() {
		var id int

		var name, phone string
		err := rows.Scan(&id, &name, &phone)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error scanning database rows"})
		}
		users = append(users, gin.H{"id": id, "name": name, "phone": phone})
	}

	// 返回 JSON 格式的查询结果
	c.JSON(http.StatusOK, users)

	return
}

func OrmGetUser(c *gin.Context) {

	var id string

	id = c.Query("id")

	if id == "" {
		c.JSON(http.StatusInternalServerError, "请上传用户id")
		return
	}

	var user User
	result := orm.DB.Find(&user, "id = ?", id)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	if user == (User{}) {
		c.JSON(http.StatusInternalServerError, "数据信息未找到")
		return
	}

	c.JSON(http.StatusOK, user)
	return
}

func OrmGetUsers(c *gin.Context) {

	var user []User
	result := orm.DB.Find(&user)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	if user == nil {
		c.JSON(http.StatusInternalServerError, "数据信息未找到")
		return
	}

	c.JSON(http.StatusOK, user)

	return
}

func AscTest(c *gin.Context) {

	c.JSON(http.StatusOK, 1)
	return

}
