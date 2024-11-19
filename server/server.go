package server

import (
	"beer/database"
	"beer/enum"
	"beer/helper"
	"os"

	uploadHandlers "beer/module/uploads/handlers"
	uploadRepositories "beer/module/uploads/repositories"
	uploadUsecases "beer/module/uploads/usecases"

	beerHandlers "beer/module/beers/handlers"
	beerRepositories "beer/module/beers/repositories"
	beerUsecases "beer/module/beers/usecases"

	"github.com/gin-gonic/gin"
	"github.com/labstack/gommon/log"

	docs "beer/docs"

	swaggerFile "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Server interface {
	Start() // ฟังก์ชันเริ่มเซิร์ฟเวอร์
}

type ginServer struct {
	app     *gin.Engine // เปลี่ยนจาก *gin.H เป็น *gin.Engine
	db      database.Database
	dbMongo database.DatabaseMongo
}

func NewGinServer(db database.Database, dbMongo database.DatabaseMongo) Server {
	ginApp := gin.New()        // ใช้ gin.New() เพื่อสร้าง Gin engine ใหม่
	ginApp.Use(gin.Logger())   // เพิ่ม middleware logger
	ginApp.Use(gin.Recovery()) // เพิ่ม middleware recovery
	ginApp.Static("/uploads", "./public/uploads")

	return &ginServer{
		app:     ginApp,
		db:      db,
		dbMongo: dbMongo,
	}
}

// @Summary แสดงข้อความต้อนรับ
// @Description ตัวอย่างการใช้งาน API เบื้องต้น
// @Tags Welcome to My Beer
// @Router / [GET]
func (s *ginServer) Start() {
	host := os.Getenv("SERVER_HOST")
	port := os.Getenv("SERVER_PORT") // หรือสามารถตั้งค่าได้จาก ENV

	s.initializeSwaggerHttpHandler()
	s.app.GET("api/v1", func(c *gin.Context) {
		helper.ResponseJson(c, enum.Success, gin.H{"message": "Welcome to the Beer API!"})
	})

	s.initializeUploadHttpHandler() // upload file
	s.initializeBeerHttpHandler()

	if err := s.app.Run(host + ":" + port); err != nil {
		log.Fatal(err)
	}
}

func (s *ginServer) initializeSwaggerHttpHandler() {
	docs.SwaggerInfo.BasePath = "/api/v1"
	s.app.GET("api/v1/docs/*any", ginSwagger.WrapHandler(swaggerFile.Handler, ginSwagger.DefaultModelsExpandDepth(-1)))
}

// @Summary อัปโหลดไฟล์
// @Description ตัวอย่างการใช้งาน API เบื้องต้น
// @Tags Upload
// @Param file formData file false "ไฟล์ที่ต้องการอัปโหลด"
// @Router /upload [POST]
func (s *ginServer) initializeUploadHttpHandler() {
	uploadRepositories := uploadRepositories.NewUploadRepository(s.db)
	uploadUsecases := uploadUsecases.NewUploadUsecaseImpl(
		uploadRepositories,
	)

	uploadHandler := uploadHandlers.NewUploadHttpHandler(uploadUsecases)
	uploadRouters := s.app.Group("api/v1/upload")

	uploadRouters.POST("", func(c *gin.Context) {
		uploadHandler.DetectUpload(c)
	})
}

func (s *ginServer) initializeBeerHttpHandler() {
	beerRepository := beerRepositories.NewBeerRepository(s.db)
	beerUsecases := beerUsecases.NewBeerUsecaseImpl(
		beerRepository,
	)

	beerHandlers := beerHandlers.NewBeerdHttpHandler(beerUsecases)
	beerRouters := s.app.Group("api/v1/beer")

	beerRouters.GET("", func(c *gin.Context) {
		beerHandlers.ListBeer(c)
	})

	beerRouters.GET("/:id", func(c *gin.Context) {
		beerHandlers.DataBeer(c)
	})

	beerRouters.POST("", func(c *gin.Context) {
		beerHandlers.CreateBeer(c)
	})

	beerRouters.PUT("/:id", func(c *gin.Context) {
		beerHandlers.UpdateBeer(c)
	})

	beerRouters.DELETE("/:id", func(c *gin.Context) {
		beerHandlers.DeleteBeer(c)
	})
}
