package helper

import (
	"fmt"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// getFileName generates a unique file name, optionally adding a file extension
func GetFileName(fileType string) string {
	return fmt.Sprintf("%d%s", time.Now().UnixNano(), fileType)
}

// getPathUpload generates the upload path based on year, month, and day
func GetPathUpload(basePath string) string {
	now := time.Now()
	return filepath.Join(basePath, fmt.Sprintf("%d", now.Year()), fmt.Sprintf("%02d", now.Month()), fmt.Sprintf("%02d", now.Day()), fmt.Sprintf("%d", time.Now().UnixNano()))
}

// uploadFile handles the file upload process
func UploadFile(file *multipart.FileHeader, savePath string, saveFilename string) (map[string]interface{}, error) {

	savePath = GetPathUpload("uploads/" + savePath)
	// Define upload path
	folderPath := filepath.Join("public", savePath)

	// Create directories if they don't exist
	if err := os.MkdirAll(folderPath, os.ModePerm); err != nil {
		return nil, fmt.Errorf("failed to create directory: %w", err)
	}

	// Generate filename
	if saveFilename == "" {
		saveFilename = strings.TrimSuffix(file.Filename, filepath.Ext(file.Filename))
	}
	extension := filepath.Ext(file.Filename)
	filename := saveFilename + extension
	filePath := filepath.Join(folderPath, filename)

	// Save the file
	outFile, err := os.Create(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to create file: %w", err)
	}
	defer outFile.Close()

	src, err := file.Open()
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer src.Close()

	if _, err := outFile.ReadFrom(src); err != nil {
		return nil, fmt.Errorf("failed to write file: %w", err)
	}

	// Gather file info
	fileInfo, err := outFile.Stat()
	if err != nil {
		return nil, fmt.Errorf("failed to get file info: %w", err)
	}

	// Prepare output data
	output := map[string]interface{}{
		"file_path":          savePath,
		"original_file_name": file.Filename,
		"file_name":          filename,
		"file_size":          fileInfo.Size(),
		"file_type":          file.Header.Get("Content-Type"),
	}

	return output, nil
}
