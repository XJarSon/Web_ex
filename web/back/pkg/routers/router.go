package routers

import (
	"back/api"
	"back/pkg"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetupRouter() *gin.Engine {

	gin.SetMode(setting.RunMode)
	r := gin.Default()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(Cors())
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Token")
	})

	// 静态文件处理
	{
		r.StaticFile("/favicon.ico", "./material/favicon.ico")
		r.StaticFile("/22_open.png", "./material/22_open.png")
		r.StaticFile("/22_close.png", "./material/22_close.png")
		r.StaticFile("/33_open.png", "./material/33_open.png")
		r.StaticFile("/33_close.png", "./material/33_close.png")
		r.Static("/static", "./material/static")
		r.LoadHTMLFiles("material/index.html")
	}

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", "./index.html")
	})

	// 建立基本页面
	{
		r.GET("/Login", func(c *gin.Context) {
			c.String(http.StatusOK, "/Login")
		})

		r.GET("/MainCourse", func(c *gin.Context) {
			c.String(http.StatusOK, "/MainCourse")
		})

	}
	// 用户鉴定
	{
		r.POST("/Check", api.CheckUser)
		r.POST("/Login", api.GetUser)
		r.POST("/Register", api.RegisterUser)

	}

	mid := r.Group("/Mid")
	{
		// 查询学院
		mid.GET("/Faculties", api.FindFaculties)

		// 查询应修课程
		mid.GET("/Courses", api.FindCourses)

	}

	education := r.Group("/MainCourse")
	{
		education.POST("/Faculty", api.AddFaculty)
		education.DELETE("/Faculty", api.DeleteFaculty)

		education.POST("/Course", api.AddCourse)
		education.DELETE("/Course", api.DeleteCourse)
		education.GET("/Course", api.AllCourses)

		education.POST("/FC", api.AddFC)
		education.DELETE("/FC", api.DeleteFC)

		education.POST("/Upload", api.Upload)
	}

	return r

}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin")
		if origin != "" {
			c.Header("Access-Control-Allow-Origin", "*") // 可将将 * 替换为指定的域名
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
			c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization, Token")
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
			c.Header("Access-Control-Allow-Credentials", "true")
		}
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
}
