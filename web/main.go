package main

import (
	"axinger.xyz/ax-go-demo/web/dto"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"time"
)

func main() {
	// 默认使用了Logger,Recovery 2个中间件
	r := gin.Default()
	// 注册一个全局中间件
	r.Use(StatCost(), TokenHandler())
	r.LoadHTMLGlob("./resources/templates/*")

	//router.Static第一个参数为路由匹配地址，第二个参数为静态资源相对路径
	r.Static("../static", "resources/static")


	r.GET("/", func(c *gin.Context) {
		//c.JSON(http.StatusOK, gin.H{
		//	"message": "首页",
		//})

		h := gin.H{
			"msg": "首页",
			"title": "我是go web",
		}

		c.HTML(http.StatusOK,"login.html",h)

	})

	//r.Any("/test", func(c *gin.Context) {
	//
	//
	//})

	r.GET("/test1", func(c *gin.Context) {

		params := c.Params
		fmt.Println("params", params)
		//c.JSON(http.StatusOK, gin.H{
		//	"message": "首页",
		//})

		c.String(http.StatusOK, "%v", "bb")
	})

	// restful 风格
	r.GET("/test2/:name/*action", func(c *gin.Context) {
		params := c.Params

		fmt.Println("params", params)
		fmt.Println("params-action", c.Param("action"))

		c.JSON(http.StatusOK, gin.H{
			"message": "test2",
		})
	})
	// url 风格
	r.GET("/test3", func(c *gin.Context) {
		// DefaultQuery 参数不存在,返回默认值
		// Query 返回空字串
		name := c.DefaultQuery("name", "")
		age := c.Query("age")

		fmt.Println(strconv.Atoi(age))

		fmt.Printf("name:%v,age:%v", name, age)

		c.JSON(http.StatusOK, gin.H{
			"message": "test3",
		})
	})

	// 参数绑定
	r.GET("/user", func(c *gin.Context) {

		var login dto.Login
		// get 用 ShouldBind
		// post 用 ShouldBindJSON
		err := c.ShouldBind(&login)

		if err != nil {
			fmt.Printf("err = %v\n", err)

			c.JSON(http.StatusBadRequest, gin.H{
				"message": "参数错误",
			})
			return
		}

		fmt.Printf("login = %v\n", login)

		c.JSON(http.StatusOK, gin.H{
			"message": "login",
		})

	})

	// HTTP重定向
	r.GET("/redirect", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "https://www.baidu.com/")
	})

	// 路由重定向
	r.GET("/redirect1", func(c *gin.Context) {
		// 指定重定向的URL
		c.Request.URL.Path = "/redirect2"
		r.HandleContext(c)
	})
	r.GET("/redirect2", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"hello": "world"})
	})

	// 路由组

	shopGroup := r.Group("/shop")
	{
		// http://localhost:8080/shop/index  要带前部分
		shopGroup.GET("/index", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "shop_index",
			})

		})
		shopGroup.GET("/cart", func(c *gin.Context) {

			c.JSON(http.StatusOK, gin.H{
				"message": "cart",
			})

		})
		shopGroup.POST("/checkout", func(c *gin.Context) {

			c.JSON(http.StatusOK, gin.H{
				"message": "checkout",
			})
		})
	}

	// 前面要有 冒号 :
	r.Run(":8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

// StatCost 是一个统计耗时请求耗时的中间件
func StatCost() gin.HandlerFunc {
	return func(c *gin.Context) {

		params := c.Params
		log.Println("StatCost params", params)
		log.Println("StatCost params-action", c.Param("action"))

		start := time.Now()
		c.Set("name", "小王子") // 可以通过c.Set在请求上下文中设置值，后续的处理函数能够取到该值
		// 调用该请求的剩余处理程序
		c.Next()
		// 不调用该请求的剩余处理程序
		// c.Abort()
		// 计算耗时
		cost := time.Since(start)
		log.Println("cost:", cost)
	}
}

func TokenHandler() gin.HandlerFunc {
	return func(c *gin.Context) {

		token := c.GetHeader("token")
		log.Println("token:", token)
		log.Println("FullPath:", c.FullPath())
		// FullPath: /shop/cart
		log.Println("FullPath:", c.FullPath())

		log.Println("FullPath:", (c.FullPath() != "/" || c.FullPath() != "/index") && token == "")

		//if !(c.FullPath() == "/" || c.FullPath() == "/index") && token == "" {
		//	c.Abort()
		//	c.JSON(http.StatusOK, gin.H{
		//		"message": "未授权",
		//	})
		//	return
		//}

		params := c.Params
		log.Println("StatCost params", params)
		log.Println("StatCost params-action", c.Param("action"))

		start := time.Now()
		c.Set("name", "小王子") // 可以通过c.Set在请求上下文中设置值，后续的处理函数能够取到该值
		// 调用该请求的剩余处理程序
		c.Next()
		// 不调用该请求的剩余处理程序

		// 计算耗时
		cost := time.Since(start)
		log.Println("cost:", cost)
	}
}
