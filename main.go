package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/influxdata/influxdb-client-go/v2"
	"mime"
	"net/http"
	"qrcodeProject/common"
	"qrcodeProject/service"
)

func main() {
	// 初始化 Redis
	common.InitRedis()

	client := influxdb2.NewClient("http://192.168.2.10:8086", "aNoGae9Dh4t4T849gW4xJp0QBm6kfj7fpsI9K2OSzA9kJi1x_ecO9BVjoxTwKZ56umveOmr3jWXT1oXbyH2Atw==")
	// always close client at the end
	defer client.Close()

	writeAPI := client.WriteAPI("zzwl", "zzwl")

	// write line protocol
	writeAPI.WriteRecord(fmt.Sprintf("stat,unit=temperature avg=%f,max=%f", 23.5, 45.0))
	// Flush writes
	writeAPI.Flush()

	r := gin.Default()

	// 解决跨域
	r.Use(Cors())

	// 创建 下单 二维码
	r.GET("/qrScan", service.GetCreateQRCode)

	// 改变 二维码 状态
	r.GET("/qrChange", service.ChangeQrState)

	// 获取 二维码 识别状态
	r.GET("/qrStatus", service.ReadQrStatus)
	// 验证 二维码 并拉起支付
	r.GET("/qrVerify", service.VerifyQrCode)

	//group := r.Group("/api")
	//{
	//	// 创建 下单 二维码
	//	group.GET("/qrScan", service.GetCreateQRCode)
	//
	//	// 改变 二维码 状态
	//	group.GET("/qrChange", service.ChangeQrState)
	//
	//	// 获取 二维码 识别状态
	//	group.GET("/qrStatus", service.ReadQrStatus)
	//	// 验证 二维码 并拉起支付
	//	group.GET("/qrVerify", service.VerifyQrCode)
	//}

	_ = mime.AddExtensionType(".js", "application/javascript")
	//locExe, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	//r.StaticFS("/public", http.Dir(filepath.Join(locExe, "./public")))
	//dir := http.Dir(filepath.Join(locExe, "web"))
	r.StaticFS("/public", http.Dir("./public"))
	r.Static("./js", "./public/js")
	r.Static("./css", "./public/css")

	// 二维码静态 文件地址
	r.StaticFS("/QrCode", http.Dir("./QrCode"))

	gin.SetMode(gin.DebugMode)

	r.Run(":9091") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

//func Cors() gin.HandlerFunc {
//	return cors.New(cors.Config{
//		AllowMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
//		//AllowHeaders:     []string{"*"},
//		AllowHeaders:     []string{"Authorization", "ts", "Accept", "Origin", "DNT", "X-CustomHeader", "Keep-Alive", "User-Agent", "X-Requested-With", "If-Modified-Since", "Cache-Control", "Content-Type", "Content-Range", "Range"},
//		ExposeHeaders:    []string{"Content-Length", "Authorization", "Content-Type"},
//		AllowCredentials: false,
//		AllowAllOrigins:  true,
//		MaxAge:           12 * time.Hour,
//	},
//	)
//}

//func Cors() gin.HandlerFunc {
//	return func(c *gin.Context) {
//		method := c.Request.Method
//
//		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
//
//		// XSRF-TOKEN  X-XSRF-TOKEN
//
//		//c.Header("Access-Control-Allow-Methods", "POST, GET, PUT, PATCH, DELETE, HEAD, OPTIONS")
//		c.Header("Access-Control-Allow-Origin", "*")
//		c.Header("Access-Control-Allow-Headers", "*")
//		c.Header("Access-Control-Allow-Methods", "*")
//		c.Header("Access-Control-Allow-Credentials", "true")
//		c.Header("Access-Control-Expose-Headers", "*")
//		if method == "OPTIONS" {
//			c.AbortWithStatus(http.StatusNoContent)
//		}
//		c.Next()
//	}
//}

func Cors() gin.HandlerFunc {
	return func(context *gin.Context) {
		method := context.Request.Method
		context.Header("Access-Control-Allow-Origin", "*")
		context.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		context.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		context.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		context.Header("Access-Control-Allow-Credentials", "true")
		if method == "OPTIONS" {
			context.AbortWithStatus(http.StatusNoContent)
		}
		context.Next()
	}
}
