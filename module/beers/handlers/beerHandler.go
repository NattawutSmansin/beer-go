package handlers

import (
	"beer/database"
	"beer/enum"
	"beer/helper"
	"beer/module/beers/models"
	"beer/module/beers/usecases"
	"fmt"
	"os"
	"strconv"
	"strings"

	notifyHandlers "beer/module/notifies/handlers"
	notifyRepositories "beer/module/notifies/repositories"
	notifyUsecases "beer/module/notifies/usecases"

	categoryHandles "beer/module/catagories/handlers"
	categoryModel "beer/module/catagories/models"
	categoryRepositories "beer/module/catagories/repositories"
	categoryUsecases "beer/module/catagories/usecases"

	uploadHandles "beer/module/uploads/handlers"
	uploadModel "beer/module/uploads/models"
	uploadRepositories "beer/module/uploads/repositories"
	uploadUsecases "beer/module/uploads/usecases"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/labstack/gommon/log"
)

type BeerHandler interface {
	CreateBeer(c *gin.Context) error
	UpdateBeer(c *gin.Context) error
	DeleteBeer(c *gin.Context) error
	DataBeer(c *gin.Context) error
	ListBeer(c *gin.Context) error
}

type beerHttpHandler struct {
	beerUsecase usecases.BeerUsecase
}

func NewBeerdHttpHandler(beerUsecase usecases.BeerUsecase) BeerHandler {
	return &beerHttpHandler{
		beerUsecase: beerUsecase,
	}
}

// @Summary เพิ่มข้อมูล Beer
// @Description ตัวอย่างการใช้งาน API เบื้องต้น
// @Tags Beer
// @Accept json
// @Produce json
// @Param request body models.CreateBeer true "Beer Data"
// @Router /beer [POST]
func (h *beerHttpHandler) CreateBeer(c *gin.Context) error {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// บันทึกข้อมูลที่ได้รับจาก request body ลงใน struct
	reqBody := &models.CreateBeer{}

	// Bind JSON หรือ form-data กับ struct reqBody
	if err := c.ShouldBindJSON(reqBody); err != nil { // ใช้ ShouldBindJSON สำหรับ JSON payload
		log.Errorf("Error binding request body: %v", err)
		helper.ResponseJson(c, enum.Fail, gin.H{"message": "Bad request"})
		return err
	}

	dbMySql := database.ConMySQLDatabase() // ฟังก์ชันที่ใช้สร้างการเชื่อมต่อ MySql
	dataCategory, err := GetDataCategory(int(reqBody.CategoryId), dbMySql)
	if err != nil {
		helper.ResponseJson(c, enum.NotFound, gin.H{"message": "Not Found Category"})
		return err
	}

	// Process the create data beer
	err = h.beerUsecase.BeerCreateDataProcess(reqBody) // Call the BeerCreateDataProcess method
	if err != nil {
		// If there was an error creating the beer, respond with a JSON message indicating that an error occurred.
		helper.ResponseJson(c, enum.Error, gin.H{}, err.Error())
		return err // Return the error
	}

	dataSource := gin.H{
		"Title":  "เพิ่มข้อมูลเบียร์",
		"Detail": "ชื่อเบียร์: " + reqBody.Name + "\n ประเภทเบียร์ซ " + dataCategory.Name + "\n รายละเอียด: " + reqBody.Description,
	}

	// สร้าง instance ของ database.DatabaseMongo ที่ใช้เชื่อมต่อ MongoDB
	dbMongo := database.ConMongoDatabase() // ฟังก์ชันที่ใช้สร้างการเชื่อมต่อ MongoDB
	CreateNotifyGo(dataSource, dbMongo)

	// ส่งกลับข้อมูลที่บันทึกสำเร็จ
	helper.ResponseJson(c, enum.Created, gin.H{})
	return nil
}

// @Summary อัปเดตข้อมูล Beer
// @Description ตัวอย่างการใช้งาน API เบื้องต้น
// @Tags Beer
// @Accept json
// @Produce json
// @Param id path int true "Beer ID"
// @Param request body models.UpdateBeer true "Beer Data"
// @Router /beer/{id} [PUT]
func (h *beerHttpHandler) UpdateBeer(c *gin.Context) error {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// ดึง ID จาก URL parameters
	beerIdParam := c.Param("id")
	if beerIdParam == "" {
		log.Errorf("Error: beer ID is missing in URL")
		helper.ResponseJson(c, enum.Fail, gin.H{"message": "Beer ID is required"})
		return nil
	}

	beerId, err := strconv.ParseUint(beerIdParam, 10, 32) // แปลง ID ที่รับมาจาก URL (string) เป็น uint32
	if err != nil {
		log.Errorf("Error parsing beer ID: %v", err)
		helper.ResponseJson(c, enum.Fail, gin.H{"message": "Invalid beer ID"})
		return nil
	}

	// บันทึกข้อมูลที่ได้รับจาก request body ลงใน struct
	reqBody := &models.UpdateBeer{}

	// Bind JSON หรือ form-data กับ struct reqBody
	if err := c.ShouldBindJSON(reqBody); err != nil { // ใช้ ShouldBindJSON สำหรับ JSON payload
		log.Errorf("Error binding request body: %v", err)
		helper.ResponseJson(c, enum.Fail, gin.H{"message": "Bad request"})
		return err
	}

	dbMySql := database.ConMySQLDatabase() // ฟังก์ชันที่ใช้สร้างการเชื่อมต่อ MySql
	dataCategory, err := GetDataCategory(int(reqBody.CategoryId), dbMySql)
	if err != nil {
		helper.ResponseJson(c, enum.NotFound, gin.H{"message": "Not Found Category"})
		return err
	}

	// แปลงข้อมูลจาก reqBody เป็น Beer struct ที่พร้อมอัปเดต
	beerData := &models.UpdateBeer{
		Id:           uint32(beerId), // แปลง beerId ที่ได้รับจาก URL เป็น uint32
		Name:         reqBody.Name,
		Description:  reqBody.Description,
		CategoryId:   reqBody.CategoryId,
		ImageFileIds: reqBody.ImageFileIds,
	}

	// เรียกใช้งาน Usecase เพื่ออัปเดตข้อมูล
	err = h.beerUsecase.BeerUpdateDataProcess(beerData) // ใช้ beerUsecase ในการอัปเดต
	if err != nil {
		helper.ResponseJson(c, enum.Error, gin.H{}, err.Error())
		return err // ถ้ามีข้อผิดพลาดจะส่งกลับ error
	}

	// ข้อมูลที่ใช้สำหรับแจ้งเตือน
	dataSource := gin.H{
		"Title":  "อัปเดตข้อมูลเบียร์",
		"Detail": "ชื่อเบียร์: " + reqBody.Name + "\n ประเภทเบียร์ซ " + dataCategory.Name + "\n รายละเอียด: " + reqBody.Description,
	}

	// สร้าง instance ของ database.DatabaseMongo ที่ใช้เชื่อมต่อ MongoDB
	dbMongo := database.ConMongoDatabase() // ฟังก์ชันที่ใช้สร้างการเชื่อมต่อ MongoDB
	CreateNotifyGo(dataSource, dbMongo)

	// ส่งกลับข้อมูลที่บันทึกสำเร็จ
	helper.ResponseJson(c, enum.Accepted, gin.H{})
	return nil
}

// @Summary ลบข้อมูล Beer
// @Description ตัวอย่างการใช้งาน API เบื้องต้น
// @Tags Beer
// @Accept json
// @Produce json
// @Param id path int true "Beer ID"
// @Router /beer/{id} [DELETE]
func (h *beerHttpHandler) DeleteBeer(c *gin.Context) error {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// ดึง ID จาก URL parameters
	beerIdParam := c.Param("id")
	if beerIdParam == "" {
		log.Errorf("Error: beer ID is missing in URL")
		helper.ResponseJson(c, enum.Fail, gin.H{"message": "Beer ID is required"})
		return nil
	}

	beerId, err := strconv.ParseUint(beerIdParam, 10, 32) // แปลง ID ที่รับมาจาก URL (string) เป็น uint32
	if err != nil {
		log.Errorf("Error parsing beer ID: %v", err)
		helper.ResponseJson(c, enum.Fail, gin.H{"message": "Invalid beer ID"})
		return nil
	}

	beer, err := h.beerUsecase.BeerDataProcess(uint32(beerId))
	if err != nil || beer == nil {
		helper.ResponseJson(c, enum.NotFound, gin.H{}, err.Error())
		return err // ถ้ามีข้อผิดพลาดจะส่งกลับ error
	}

	dbMySql := database.ConMySQLDatabase() // ฟังก์ชันที่ใช้สร้างการเชื่อมต่อ MySql
	dataCategory, err := GetDataCategory(int(beer.CategoryId), dbMySql)
	if err != nil {
		helper.ResponseJson(c, enum.NotFound, gin.H{"message": "Not Found Category"})
		return err
	}

	// ข้อมูลที่ใช้สำหรับแจ้งเตือน
	dataSource := gin.H{
		"Title":  "ลบข้อมูลเบียร์",
		"Detail": "ชื่อเบียร์: " + beer.Name + "\n ประเภทเบียร์ซ " + dataCategory.Name + "\n รายละเอียด: " + beer.Description,
	}

	// เรียกใช้งาน Usecase เพื่ออัปเดตข้อมูล
	err = h.beerUsecase.BeerDataDelete(uint32(beerId)) // ใช้ beerUsecase ในการอัปเดต
	if err != nil {
		helper.ResponseJson(c, enum.Error, gin.H{}, err.Error())
		return err // ถ้ามีข้อผิดพลาดจะส่งกลับ error
	}

	// สร้าง instance ของ database.DatabaseMongo ที่ใช้เชื่อมต่อ MongoDB
	dbMongo := database.ConMongoDatabase() // ฟังก์ชันที่ใช้สร้างการเชื่อมต่อ MongoDB
	CreateNotifyGo(dataSource, dbMongo)

	// ส่งกลับข้อมูลที่บันทึกสำเร็จ
	helper.ResponseJson(c, enum.Deleted, gin.H{})
	return nil
}

// @Summary เรียกข้อมูล Beer ราย item
// @Description ตัวอย่างการใช้งาน API เบื้องต้น
// @Tags Beer
// @Accept json
// @Produce json
// @Param id path int true "Beer ID"
// @Router /beer/{id} [GET]
func (h *beerHttpHandler) DataBeer(c *gin.Context) error {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Errorf("Error loading .env file: %v", err)
		helper.ResponseJson(c, enum.Fail, gin.H{"message": "Internal server error"})
		return err
	}

	// Retrieve the beer ID from the URL parameters
	beerIdParam := c.Param("id")
	if beerIdParam == "" {
		log.Errorf("Error: beer ID is missing in URL")
		helper.ResponseJson(c, enum.Fail, gin.H{"message": "Beer ID is required"})
		return fmt.Errorf("beer ID is missing")
	}

	// Convert beer ID from string to uint32
	beerId, err := strconv.ParseUint(beerIdParam, 10, 32)
	if err != nil {
		log.Errorf("Error parsing beer ID: %v", err)
		helper.ResponseJson(c, enum.Fail, gin.H{"message": "Invalid beer ID"})
		return fmt.Errorf("invalid beer ID: %v", err)
	}

	// Call the usecase to process the beer data
	resultBeer, err := h.beerUsecase.BeerDataProcess(uint32(beerId))
	if err != nil {
		log.Errorf("Error processing beer data: %v", err)
		helper.ResponseJson(c, enum.Error, gin.H{"message": "Error processing beer data"})
		return err // Return error if there's an issue
	}

	// call usecase category
	dbMySql := database.ConMySQLDatabase() // ฟังก์ชันที่ใช้สร้างการเชื่อมต่อ MySql
	dataCategory, err := GetDataCategory(int(resultBeer.CategoryId), dbMySql)
	if err != nil {
		helper.ResponseJson(c, enum.NotFound, gin.H{"message": "Not Found Category"})
		return err
	}

	// convert to array
	imageFileIds := ExtractImageFileIds(resultBeer.ImageFileIds)

	data := gin.H{
		"id":          resultBeer.Id,
		"name":        resultBeer.Name,
		"description": resultBeer.Description,
		"catetgory": gin.H{
			"id":   dataCategory.Id,
			"name": dataCategory.Name,
		},
		"uploads": []gin.H{},
	}

	for index, imageFileId := range imageFileIds {
		fmt.Printf("Index: %d, ImageFileId: %d\n", index, imageFileId)

		getDataFile, err := GetDataFile(int(imageFileId), dbMySql)
		fmt.Println(getDataFile.FileName)
		if err != nil {
			log.Errorf("Error fetching data for imageFileId %d: %v", imageFileId, err)
			continue // Skip this iteration if there's an error fetching data
		}

		// Append the file details (as a gin.H map) to the uploads slice
		data["uploads"] = append(data["uploads"].([]gin.H), gin.H{
			"id":   getDataFile.Id,
			"name": getDataFile.FileName,
			"path": getDataFile.FilePath,
			"size": getDataFile.FileSize,
			"type": getDataFile.FileType,
			"url":  os.Getenv("SERVER_NAME") + getDataFile.FilePath + "/" + getDataFile.FileName,
		})
	}

	// Send back the processed beer data in the response
	helper.ResponseJson(c, enum.Success, data)
	return nil // Return nil on success
}

// @Summary เรียกข้อมูล Beer ทั้งหมด
// @Description ตัวอย่างการใช้งาน API เบื้องต้น
// @Tags Beer
// @Accept json
// @Produce json
// @Param name query string False "Beer Name"
// @Param page query int False "Page (optional, default: 1)"
// @Param limit query int False "Limit (optional, default: 10)"
// @Router /beer [GET]
func (h *beerHttpHandler) ListBeer(c *gin.Context) error {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Errorf("Error loading .env file: %v", err)
		helper.ResponseJson(c, enum.Fail, gin.H{"message": "Internal server error"})
		return err
	}

	// Retrieve the beer ID from the URL parameters
	beerNameParam := c.Query("name")
	pageParam := c.Query("page")
	limitParam := c.Query("limit")

	// Default values for pagination if not provided
	page, err := strconv.Atoi(pageParam)
	if err != nil || page <= 0 {
		page = 1
	}
	limit, err := strconv.Atoi(limitParam)
	if err != nil || limit <= 0 {
		limit = 10
	}

	// Call the usecase to get the data
	listBeer, total, err := h.beerUsecase.BeerListDataProcess(beerNameParam, page, limit)
	if err != nil {
		helper.ResponseJson(c, enum.Fail, gin.H{"message": err.Error()})
		return nil
	}

	dbMySql := database.ConMySQLDatabase() // ฟังก์ชันที่ใช้สร้างการเชื่อมต่อ MySql

	data := []gin.H{}

	// Iterate through the list of beers
	for _, beerItem := range listBeer {
		// ดึงข้อมูลหมวดหมู่ของเบียร์
		dataCategory, err := GetDataCategory(int(beerItem.CategoryId), dbMySql)
		if err != nil {
			helper.ResponseJson(c, enum.NotFound, gin.H{"message": "Not Found Category"})
			return err
		}

		// ดึงข้อมูล uploads (ImageFileIds)
		imageFileIds := ExtractImageFileIds(beerItem.ImageFileIds)

		uploads := []gin.H{}
		for _, imageFileId := range imageFileIds {
			// ดึงข้อมูลไฟล์แต่ละรายการ
			getDataFile, err := GetDataFile(imageFileId, dbMySql)
			if err != nil {
				log.Errorf("Error fetching data for imageFileId %d: %v", imageFileId, err)
				continue // ข้ามถ้ามีปัญหา
			}

			// เพิ่มไฟล์ลงใน uploads
			uploads = append(uploads, gin.H{
				"id":   getDataFile.Id,
				"name": getDataFile.FileName,
				"path": getDataFile.FilePath,
				"size": getDataFile.FileSize,
				"type": getDataFile.FileType,
				"url":  os.Getenv("SERVER_NAME") + "/uploads" + getDataFile.FilePath + "/" + getDataFile.FileName,
			})
		}

		// เพิ่มข้อมูลเบียร์ลงในลิสต์
		data = append(data, gin.H{
			"id":          beerItem.Id,
			"name":        beerItem.Name,
			"description": beerItem.Description,
			"category": gin.H{
				"id":   dataCategory.Id,
				"name": dataCategory.Name,
			},
			"uploads": uploads,
		})
	}

	// Response with the paginated data
	helper.ResponseJsonPaginate(c, enum.Success, data, page, limit, int(total))

	return nil
}

func GetDataFile(uploadId int, db database.Database) (*uploadModel.UploadData, error) {
	uploadRepository := uploadRepositories.NewUploadRepository(db)
	uploadUsecases := uploadUsecases.NewUploadUsecaseImpl(uploadRepository)
	uploadHandle := uploadHandles.NewUploadHttpHandler(uploadUsecases)

	uploadData, err := uploadHandle.GetUpload(uploadId)
	if err != nil || uploadData == nil {
		return nil, err
	}

	return uploadData, nil
}

func GetDataCategory(categoryId int, db database.Database) (*categoryModel.GetCategory, error) {
	categoryRepository := categoryRepositories.NewCategoryRepository(db)
	categoryUsecases := categoryUsecases.NewCategoryUsecaseImpl(categoryRepository)
	categoryHandle := categoryHandles.NewCategoryHttpHandler(categoryUsecases)

	categoryData, err := categoryHandle.GetDataCategory(categoryId)
	if err != nil || categoryData == nil {
		return nil, err
	}

	return categoryData, nil
}

func CreateNotifyGo(content gin.H, dbMongo database.DatabaseMongo) {
	// สร้าง NotifyRepository ก่อน
	notifyRepository := notifyRepositories.NewNotifyRepository(dbMongo)
	notifyUsecase := notifyUsecases.NewNotifyUsecaseImpl(notifyRepository)
	notifyHandler := notifyHandlers.NewNotifyHttpHandler(notifyUsecase)

	// Pass the content to the Store method Create Data Noti
	notifyHandler.Store(content)
}

func ExtractImageFileIds(imageFileIds string) []int {
	imageFileIds = strings.Trim(imageFileIds, "[]")
	parts := strings.Split(imageFileIds, ",")
	var imageFiles []int
	for _, part := range parts {
		// Convert each part to integer using strconv.Atoi (it returns an error if conversion fails)
		id, err := strconv.Atoi(part)
		if err != nil {
			log.Fatalf("Error converting string to int: %v", err)
		}
		imageFiles = append(imageFiles, id)
	}

	return imageFiles
}
