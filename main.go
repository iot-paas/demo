package main

import (
	"net/http"
	"time"

	"github.com/brianvoe/gofakeit"
	"github.com/gin-gonic/gin"
)

func init() {
	gofakeit.Seed(0)
}

type DeviceInfo struct {
	ID                int    `json:"id"`
	Type              string `json:"type"`
	Sn                string `json:"sn"`
	OwnerID           int    `json:"owner_id"`
	OwnerName         string `json:"owner_name"`
	Description       string `json:"description"`
	DataFrequency     int    `json:"data_frequency"`
	DataFrequencyUnit string `json:"data_frequency_unit"`
}

type DeviceData struct {
	ID          int    `json:"id"`
	DataID      int    `json:"data_id"`
	Time        string `json:"time"`
	Temperature int    `json:"temperature"`
	Humidity    int    `json:"humidity"`
}

var deviceInfoList = map[string]DeviceInfo{
	"123": {
		ID:                123,
		Type:              "RAINBOW_IOT_T1",
		Sn:                "A10000101",
		OwnerID:           1,
		OwnerName:         "京东物流上海物流仓储中心",
		Description:       "温湿度检测设备",
		DataFrequency:     30,
		DataFrequencyUnit: "DAY",
	},
	"234": {
		ID:                234,
		Type:              "RAINBOW_IOT_T2",
		Sn:                "A10000102",
		OwnerID:           1,
		OwnerName:         "京东物流上海物流仓储中心",
		Description:       "温湿度检测设备",
		DataFrequency:     30,
		DataFrequencyUnit: "DAY",
	},
}

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("/device/:id", func(c *gin.Context) {
		id := c.Param("id")
		c.JSON(http.StatusOK, deviceInfoList[id])
	})
	r.GET("/device/:id/data", func(c *gin.Context) {
		id := c.Param("id")
		deviceInfo, exist := deviceInfoList[id]
		if !exist {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "device not found",
			})
			return
		}
		c.JSON(http.StatusOK, DeviceData{
			ID:          deviceInfo.ID,
			DataID:      gofakeit.Number(1, 100000),
			Time:        time.Now().Format("2006-01-02 15:04:05"),
			Temperature: gofakeit.Number(0, 40),
			Humidity:    gofakeit.Number(0, 100),
		})
	})

	r.POST("/api/iot/user/login", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"roles":        []string{"admin"},
			"introduction": "I am a super administrator",
			"avatar":       "https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif",
			"name":         "Super Admin",
			"token":        "xxxx00001",
			"code":         20000,
		})
	})

	r.GET("/api/iot/user/info", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"roles":        []string{"admin"},
			"introduction": "I am a super administrator",
			"avatar":       "https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif",
			"name":         "Super Admin",
			"token":        "xxxx00001",
			"code":         20000,
		})
	})
	r.POST("/api/iot/user/logout", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"data": "success",
			"code": 20000,
		})
	})

	r.Run(":8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
