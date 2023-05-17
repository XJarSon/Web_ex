package api

import (
	"back/models"
	setting "back/pkg"
	"back/pkg/crypto"
	"back/pkg/e"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ID struct {
	ID uint `json:"id"`
}

type ViewFaculty struct {
	CODE string `json:"code"`
	Name string `json:"name"`
}

func AddFaculty(c *gin.Context) {
	var faculty ViewFaculty
	code := e.SUCCESS
	if err := c.ShouldBindJSON(&faculty); err != nil {
		code = e.INVALID_PARAMS
		c.JSON(200, gin.H{
			"code": code,
			"msg":  e.GetMsg(code),
		})
		return
	}
	fid := models.AddFaculty(models.Faculty{
		FCODE: faculty.CODE,
		Name:  faculty.Name,
	})
	if fid == 0 {
		code = e.ERROR
	}
	c.JSON(200, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"fid":  fid,
	})
}

func DeleteFaculty(c *gin.Context) {
	var id ID
	code := e.SUCCESS
	if err := c.ShouldBindJSON(&id); err != nil {
		code = e.INVALID_PARAMS
		c.JSON(200, gin.H{
			"code": code,
			"msg":  e.GetMsg(code),
		})
		return
	}
	if err := models.DeleteFaculty(id.ID); err != nil {
		code = e.ERROR
	}
	c.JSON(200, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
	})
}

func AddCourse(c *gin.Context) {

	var course models.Course
	code := e.SUCCESS
	if err := c.ShouldBindJSON(&course); err != nil {
		code = e.INVALID_PARAMS
		c.JSON(200, gin.H{
			"code": code,
			"msg":  e.GetMsg(code),
		})
		return
	}

	if err := models.AddCourse(&course); err != nil {
		code = e.ERROR
	}

	c.JSON(200, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"cid":  course.CID,
	})
}

func Upload(c *gin.Context) {
	code := e.ERROR
	url := c.PostForm("url")
	id := c.PostForm("id")
	picture := ""
	var fileName string
	if id != "" {
		if url != "" {
			fileName = beego.Substr(url, 12, len(url)-16)
		} else {
			fileHeader, _ := c.FormFile("img")
			// 如果获取图片成功
			// 读取图片file并生成对应的md5
			file, _ := fileHeader.Open()
			fileName = crypto.EncryptFile(file)

			dst := "./material/static/img/" + fileName + ".jpg"
			err := c.SaveUploadedFile(fileHeader, dst)

			if err != nil {
				return
			}
		}
		picture = setting.HOST + "/static/img/" + fileName + ".jpg"
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"url":  picture,
	})
}

func DeleteCourse(c *gin.Context) {
	var id ID
	code := e.SUCCESS
	if err := c.ShouldBindJSON(&id); err != nil {
		code = e.INVALID_PARAMS
		c.JSON(200, gin.H{
			"code": code,
			"msg":  e.GetMsg(code),
		})
		return
	}
	if err := models.DeleteCourse(id.ID); err != nil {
		code = e.ERROR
	}
	c.JSON(200, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
	})
}

func AllCourses(c *gin.Context) {
	courses := models.GetAllCourses()
	c.JSON(200, gin.H{
		"code": e.SUCCESS,
		"msg":  e.GetMsg(e.SUCCESS),
		"data": courses,
	})
}

func AddFC(c *gin.Context) {
	var fc models.FC
	code := e.SUCCESS
	if err := c.ShouldBindJSON(&fc); err != nil {
		code = e.INVALID_PARAMS
		c.JSON(200, gin.H{
			"code": code,
			"msg":  e.GetMsg(code),
		})
		return
	}
	if err := models.AddFC(&fc); err != nil {
		code = e.ERROR
	}
	c.JSON(200, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
	})
}

func DeleteFC(c *gin.Context) {
	var fc models.FC
	code := e.SUCCESS
	if err := c.ShouldBindJSON(&fc); err != nil {
		code = e.INVALID_PARAMS
		c.JSON(200, gin.H{
			"code": code,
			"msg":  e.GetMsg(code),
		})
		return
	}

	fmt.Println(fc)
	if err := models.DeleteFC(&fc); err != nil {
		code = e.ERROR
	}

	c.JSON(200, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
	})
}
