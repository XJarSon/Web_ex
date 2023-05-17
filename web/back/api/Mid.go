package api

import (
	"back/models"
	"back/pkg/e"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

func FindFaculties(c *gin.Context) {
	faculties, err := models.GetAllFaculties()
	code := e.SUCCESS
	if err != nil {
		code = e.ERROR
	}
	c.JSON(code, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": faculties,
	})
}

func FindCourses(c *gin.Context) {
	departmentId := c.Query("fid")

	id, _ := strconv.Atoi(departmentId)
	fmt.Println(id)
	courses, err := models.GetCoursesByFaculties(uint(id))
	fmt.Println(courses)
	code := e.SUCCESS
	if err != nil {
		code = e.ERROR
	}
	c.JSON(code, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": courses,
	})
}
