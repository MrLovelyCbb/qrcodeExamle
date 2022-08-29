package service

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"qrcodeProject/common"
	"qrcodeProject/utils"
	"time"
)

type QRStruct struct {
	QrUrl   string `json:"qrUrl"`
	QrId    string `json:"qrId"`
	QrState int    `json:"qrState"`
}

// GetCreateQRCode 创建并获取二维码
func GetCreateQRCode(c *gin.Context) {
	sessionId := utils.GetSessionId(12)
	imgUrl := "http://192.168.2.20:9091/qrChange?qrId=" + sessionId
	// 03: 生成二维码
	qrCodeFileName := utils.CreateCompanyQR(imgUrl, sessionId)
	//qrCodeFileName := utils.CreateQrCode(imgUrl, sessionId)

	qrcodeUrl := "http://192.168.2.20:9091" + qrCodeFileName[1:]

	qrInfoData := QRStruct{
		QrId:    sessionId,
		QrUrl:   qrcodeUrl,
		QrState: 1,
	}
	setQrInfoById(&qrInfoData)

	c.JSON(200, gin.H{
		"code":      "200",
		"qrId":      sessionId,
		"qrCodeUrl": qrcodeUrl,
	})
}

// ChangeQrState 改变二维码状态
func ChangeQrState(c *gin.Context) {
	qrId, _ := c.GetQuery("qrId")

	qrStruct := getQrInfoById(qrId)
	if qrStruct != nil {
		qrStruct.QrState = 2
		setQrInfoById(qrStruct)
		fmt.Println("qrId=? do =>? ", qrId, qrStruct.QrState)

		fmt.Println("此二维码ID：" + qrId + "已被扫")
	} else {
		c.JSON(200, gin.H{
			"code":  200,
			"data":  "",
			"state": 3,
		})
	}
}

func ReadQrStatus(c *gin.Context) {
	qrId, _ := c.GetQuery("qrId")

	qrStruct := getQrInfoById(qrId)
	if qrStruct == nil {
		c.JSON(200, gin.H{
			"code":  200,
			"data":  "",
			"state": 3,
		})
	} else {
		c.JSON(200, gin.H{
			"code":  200,
			"data":  "",
			"state": qrStruct.QrState,
		})
	}
}

// VerifyQrCode 确认二维码被扫
func VerifyQrCode(c *gin.Context) {

	query, _ := c.GetQuery("qrId")

	fmt.Println("二维码已被确认 =============== > ", query)

	//c.Redirect(http.StatusTemporaryRedirect, "http://192.168.2.20:9091/public/")
	//c.Header("Access-Control-Allow-Origin ", "*")
	//c.Header("Origin", "*")
	//c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	//c.Redirect(301)
	//c.Header("Location", "http://www.baidu.com/")
	c.Redirect(http.StatusMovedPermanently, "https://www.baidu.com/")

	//c.JSON(200, gin.H{
	//	"message": "pong",
	//})
}

func getQrInfoById(qrId string) *QRStruct {
	val, _ := common.Rdb.Get(common.Ctx, qrId).Result()
	if val == "" {
		return nil
	}
	qrStruct := &QRStruct{}
	err := json.Unmarshal([]byte(val), qrStruct)
	if err != nil {
		fmt.Println(err)
	}
	return qrStruct
}

func setQrInfoById(qrInfoData *QRStruct) *QRStruct {
	qrInfo, _ := json.Marshal(qrInfoData)
	err := common.Rdb.Set(common.Ctx, qrInfoData.QrId, string(qrInfo), 10*time.Second).Err()
	if err != nil {
		panic(err)
	}
	return qrInfoData
}

func qrCreateCode() {

}
