package handlers

import (
	"beer/database"
	"beer/enum"
	"beer/helper"
	"beer/module/uploads/models"
	"beer/module/uploads/usecases"
	"os"

	notifyHandlers "beer/module/notifies/handlers"
	notifyRepositories "beer/module/notifies/repositories"
	notifyUsecases "beer/module/notifies/usecases"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/labstack/gommon/log"
)

type UploadHandler interface {
	DetectUpload(c *gin.Context) error
	GetUpload(uploadId int) (*models.UploadData ,error)
}

type uploadHttpHandler struct {
	uploadUsecase usecases.UploadUsecase
}

func NewUploadHttpHandler(uploadUsecase usecases.UploadUsecase) UploadHandler {
	return &uploadHttpHandler{
		uploadUsecase: uploadUsecase,
	}
}

func (h *uploadHttpHandler) DetectUpload(c *gin.Context) error {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// รับไฟล์จาก request
	file, err := c.FormFile("file") // "file" คือ key ที่ใช้ใน form-data
	if err != nil {
		log.Errorf("Error retrieving the file: %v", err)
		helper.ResponseJson(c, enum.Fail, gin.H{}, "File not found.")
		return err
	}

	result, err := helper.UploadFile(file, "files", "") // เรียกใช้ method อัปโหลดไฟล์
	if err != nil {
		log.Errorf("Error uploading the file: %v", err)
		helper.ResponseJson(c, enum.Error, gin.H{}, err.Error())
		return err
	}

	// บันทึกข้อมูลการอัปโหลดลงในฐานข้อมูล
	reqBody := &models.CreateUploadData{
		OriginalName: result["original_file_name"].(string),
		FileName:     result["file_name"].(string),
		FilePath:     result["file_path"].(string),
		FileSize:     float32(result["file_size"].(int64) / 1048576), // แปลงเป็นเมกะไบต์
		FileType:     result["file_type"].(string),
	}

	// Attempt to bind the request body to the struct
	if err := c.Bind(reqBody); err != nil {
		log.Errorf("Error binding request body: %v", err)
		helper.ResponseJson(c, enum.Fail, gin.H{"message": "Bad request"})
		return err
	}

	// Process the upload data
	id, err := h.uploadUsecase.UploadDataProcessing(reqBody) // เรียกใช้ UploadDataProcessing แค่ครั้งเดียว
	if err != nil {
		helper.ResponseJson(c, enum.Error, gin.H{}, err.Error())
		return err // Return the error
	}

	dataSource := gin.H{
		"Title":  "อัปโหลดไฟล์ข้อมูล",
		"Detail": "urls: " + os.Getenv("SERVER_NAME") + reqBody.FilePath + "/" + reqBody.FileName,
	}

	// สร้าง instance ของ database.DatabaseMongo ที่ใช้เชื่อมต่อ MongoDB
	dbMongo := database.ConMongoDatabase() // ฟังก์ชันที่ใช้สร้างการเชื่อมต่อ MongoDB
	CreateNotifyGo(dataSource, dbMongo)

	// ส่งกลับ ID ของการอัปโหลดที่สร้างขึ้น
	helper.ResponseJson(c, enum.Created, gin.H{"id": id})
	return nil // Return nil if everything is fine
}

func (h *uploadHttpHandler) GetUpload(uploadId int) (*models.UploadData ,error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dataUpload, err := h.uploadUsecase.GetUploadData(uploadId)
	if err != nil || dataUpload == nil {
		return nil, err
	}

	return dataUpload, nil
}

func CreateNotifyGo(content gin.H, dbMongo database.DatabaseMongo) {
	// สร้าง NotifyRepository ก่อน
	notifyRepository := notifyRepositories.NewNotifyRepository(dbMongo)
	notifyUsecase := notifyUsecases.NewNotifyUsecaseImpl(notifyRepository)
	notifyHandler := notifyHandlers.NewNotifyHttpHandler(notifyUsecase)

	// Pass the content to the Store method
	notifyHandler.Store(content)
}
