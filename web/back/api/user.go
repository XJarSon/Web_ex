package api

import (
	"back/models"
	"back/pkg/crypto"
	"back/pkg/detection"
	"back/pkg/e"
	"back/pkg/util"
	"fmt"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"regexp"
)

type user struct {
	Username string `valid:"Required; MaxSize(50)" json:"username"`
	Password string `valid:"Required; MaxSize(50)" json:"password"`
}

func CheckUser(c *gin.Context) {
	token := c.GetHeader("token")
	code := e.SUCCESS
	if !util.CheckToken(token) {
		code = e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
	})
}

func GetUser(c *gin.Context) {
	valid := validation.Validation{}
	var a user
	err := c.BindJSON(&a)
	if err != nil {
		return
	}
	// md5加密
	a.Password = crypto.Encrypt(a.Password)
	ok, _ := valid.Valid(&a)

	code := e.INVALID_PARAMS
	var userId uint
	if ok {
		userId = models.CheckUser(a.Username, a.Password)
		if userId > 0 {
			token, err := util.GenerateToken(a.Username, a.Username)
			if err != nil {
				code = e.ERROR_AUTH_TOKEN
			} else {
				c.Writer.Header().Add("token", token)
				code = e.SUCCESS
			}

		} else {
			code = e.ERROR_AUTH
		}
	} else {
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
		}
	}

	data := make(map[string]interface{})
	if userId > 0 {
		user, err := models.GetUserInfo(userId)
		// 获取ID不存在则返回参数错误
		if err != nil {
			code = e.ERROR_EXIST_USER
			c.JSON(http.StatusOK, gin.H{
				"code": code,
				"msg":  e.GetMsg(code),
			})
			return
		}
		data["id"] = user.ID
		data["username"] = user.Username
		err = models.UpdateVisitsNumber()
		if err != nil {
			log.Println(err)
		}
		system, _ := models.GetSystemMessage()
		fmt.Println(system)
		data["createTime"] = system.CreateTime.Format("2022-01-02 15:04:05")
		data["visitsNumber"] = system.VisitsNumber
		data["version"] = system.Version
		data["population"] = models.GetPopulationMessage()
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

func CheckString(str string) bool {
	//正则表达式
	re := regexp.MustCompile("^[a-zA-Z0-9_]+$")
	return re.MatchString(str)
}

// RegisterUser 注册用户
func RegisterUser(c *gin.Context) {
	var user models.Message
	err := c.BindJSON(&user)
	if err != nil {
		return
	}
	fmt.Println(user)
	code := e.SUCCESS
	// 检查输入信息是否合法
	if CheckString(user.Password) && detection.CheckSensitiveWord(user.Username) && detection.CheckSensitiveWord(user.Name) {
		user.Password = crypto.Encrypt(user.Password)
		err := models.AddUser(&user)
		if err != nil {
			code = e.ERROR_EXIST_USER
		}
	} else {
		code = e.CONTAINING_SENSITIVE_WORD
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
	})
}
