package router

import (
	"fmt"
	"net/http"
	"time"

	"github.com/UfiairENE/bantu_solution/internal/model"
	"github.com/UfiairENE/bantu_solution/pkg/middleware"
	"github.com/UfiairENE/bantu_solution/pkg/router/connection"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func Setup(validate *validator.Validate) *gin.Engine {
	r := gin.Default()

	// Middlewares
	r.Use(gin.Logger())
	r.Use(middleware.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.CORS())
	r.Use(gzip.Gzip(gzip.DefaultCompression))

	// ApiVersion := "v1"
	// FootprintUrl(r, validate, ApiVersion)

	type form struct {
		IpAddress   string `json:"ip_address" validate:"required"`
		DeviceInfo  string `json:"device_info" validate:"required"`
		BrowserType string `json:"browser_type" validate:"required"`
		Longitude   string   `json:"longitude" validate:"required"`
		Latitude    string   `json:"latitude" validate:"required"`
		City        string `json:"city" validate:"required"`
		Country     string `json:"country" validate:"required"`
	}

	r.POST("/add", func(c *gin.Context) {
		db := connection.Connection()
		req := form{}
		err := c.ShouldBind(&req)
		if err != nil {
			c.JSON(http.StatusBadRequest, struct{ message string }{message: err.Error()})
			return
		}

		err = validate.Struct(&req)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		
		fmt.Println(1)
		print := model.Footprint{
			CurrentTime: time.Now(),
			IPAddress:   req.IpAddress,
			DeviceInfo:  req.DeviceInfo,
			BrowserType: req.BrowserType,
			Longitude:   req.Longitude,
			Latitude:    req.Latitude,
			City:        req.City,
			Country:     req.Country,
		}
		
		tx := db.Create(&print)
		if tx.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": tx.Error.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "successfuly added to database", "data": print})
	})

	r.GET("/viewfootprint", func (c *gin.Context) {
		db := connection.Connection()
		var footprints []model.Footprint
		result :=db.Find(&footprints)
		if result.Error !=nil{
			c.JSON(http.StatusInternalServerError,gin.H{"message":result.Error.Error()})
			return 
		}
		c.JSON(http.StatusOK, gin.H{"message": "view footprint", "data": footprints})

	})
		

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"name":    "Not Found",
			"message": "Page not found.",
			"code":    400,
			"status":  http.StatusNotFound,
		})
	})

	return r
}
