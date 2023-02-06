package controllers

import (
	"context"
	"fmt"
	"gofiber-gorm/src/app/service"
	"gofiber-gorm/src/database/schema"
	"gofiber-gorm/src/pkg/config"
	"gofiber-gorm/src/pkg/helpers"
	"gofiber-gorm/src/pkg/modules/response"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
)

// GetUploads 	func gets all exists Uploads.
// @Description Get all exists Uploads.
// @Summary 		get all exists Uploads
// @Tags 				Upload
// @Accept 			json
// @Produce 		json
// @Success 		200 {string} status "Ok"
// @Security 		ApiKeyAuth
// @Router 			/v1/upload [get]
func FindAllUpload(c *fiber.Ctx) error {
	db := config.GetDB()

	uploadService := service.NewUploadService(db)
	data, total, err := uploadService.FindAll(c)

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"code":    http.StatusBadRequest,
			"message": "error to get uploads",
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"code":    http.StatusOK,
		"message": "data has been received",
		"data":    data,
		"total":   total,
	})
}

// GetUpload 		func gets Upload by given ID or 404 error.
// @Description Get Upload by given ID.
// @Summary 		get Upload by given ID
// @Tags 				Upload
// @Accept 			json
// @Produce 		json
// @Param 			id path string true "Upload ID"
// @Success 		200 {string} status "Ok"
// @Security 		ApiKeyAuth
// @Router 			/v1/upload/{id} [get]
func FindUploadById(c *fiber.Ctx) error {
	db := config.GetDB()
	id, err := uuid.Parse(c.Params("id"))

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"code":    http.StatusBadRequest,
			"message": "failed to get role",
		})
	}

	uploadService := service.NewUploadService(db)
	data, err := uploadService.FindById(id)

	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"code":    http.StatusNotFound,
			"message": "data not found or has been deleted",
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"code":    http.StatusOK,
		"message": "data has been received",
		"data":    data,
	})
}

// GetUploadByKeyFile 	func gets Upload by given ID or 404 error.
// @Description				 	Get Upload by given ID.
// @Summary				 			get Upload by given ID
// @Tags				 				Upload
// @Accept				 			json
// @Produce				 			json
// @Param				 				keyFile path string true "keyFile Upload"
// @Success				 			200 {string} status "Ok"
// @Security				 		ApiKeyAuth
// @Router				 			/v1/upload/presign/{keyFile} [get]
func PresignedUploadURL(c *fiber.Ctx) error {
	ctx := context.Background()
	MINIO_BUCKET_NAME := config.Env("MINIO_BUCKET_NAME", "gofiber")
	keyFile := c.Params("keyFile")

	expiresInDays := os.Getenv("MINIO_S3_EXPIRED")
	expiresIn, _ := strconv.Atoi(expiresInDays)                 // expires in days
	expiresPresign := time.Hour * 24 * time.Duration(expiresIn) // expires in 7 days

	// Set request parameters
	reqParams := make(url.Values)
	// reqParams.Set("response-content-disposition", fmt.Sprintf("attachment; filename=%s", keyFile)) // for download file

	minioClient, _ := config.MinIOConnection()
	presignedURL, err := minioClient.PresignedGetObject(ctx, MINIO_BUCKET_NAME, keyFile, expiresPresign, reqParams)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	signedURL := fmt.Sprintf("%s:%s%s?%s", presignedURL.Scheme, presignedURL.Host, presignedURL.Path, presignedURL.RawQuery)

	return c.JSON(fiber.Map{
		"message":    "Successfully uploaded file",
		"info":       presignedURL,
		"signed_url": signedURL,
	})
}

// CreateUpload func for creates a new Upload.
// @Description Create a new Upload.
// @Summary 		create a new Upload
// @Tags 				Upload
// @Accept 			mpfd
// @Produce 		json
// @Param 			fileUpload formData file true "File Upload"
// @Success 		200 {string} status "Ok"
// @Security 		ApiKeyAuth
// @Router 			/v1/upload [post]
func CreateUpload(c *fiber.Ctx) error {
	// db := config.GetDB()
	ctx := context.Background()
	// uploadSchema := new(schema.UploadSchema)

	MINIO_BUCKET_NAME := config.Env("MINIO_BUCKET_NAME", "gofiber")
	file, err := c.FormFile("fileUpload")

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		})
	}

	filePath := fmt.Sprintf("./public/tmp/uploads/%s", file.Filename)
	c.SaveFile(file, filePath)

	// Create minio connection.
	minioClient, err := config.MinIOConnection()
	if err != nil {
		// Return status 500 and minio connection error.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	filename := file.Filename
	size := file.Size
	mimetype := file.Header["Content-Type"][0]

	// Upload file
	info, err := minioClient.FPutObject(ctx, MINIO_BUCKET_NAME, filename, filePath, minio.PutObjectOptions{
		ContentType:        mimetype,
		ContentDisposition: fmt.Sprintf("inline; filename=%s", filename),
	})

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	log.Printf("Successfully uploaded %s of size %d\n", filename, size)

	return c.JSON(fiber.Map{
		"message": "Successfully uploaded file",
		"info":    info,
	})

	// // role service
	// uploadService := service.NewUploadService(db)
	// data, err := uploadService.Create(*uploadSchema)

	// if err != nil {
	// 	return c.Status(http.StatusBadRequest).JSON(fiber.Map{
	// 		"code":    http.StatusBadRequest,
	// 		"message": "failed to create role",
	// 	})
	// }

	// return c.Status(http.StatusOK).JSON(fiber.Map{
	// 	"code":    http.StatusOK,
	// 	"message": "data has been added",
	// 	"data":    data,
	// })
}

// UpdateUpload func for updates Upload by given ID.
// @Description Update Upload.
// @Summary 		update Upload
// @Tags 				Upload
// @Accept 			x-www-form-urlencoded
// @Produce 		json
// @Param 			id path string true "Upload ID"
// @Param 			name formData string true "Name"
// @Success 		200 {string} status "Ok"
// @Security 		ApiKeyAuth
// @Router 			/v1/upload/{id} [put]
func UpdateUpload(c *fiber.Ctx) error {
	db := config.GetDB()
	uploadSchema := new(schema.UploadSchema)

	id, err := uuid.Parse(c.Params("id"))

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"code":    http.StatusBadRequest,
			"message": "failed to get role",
		})
	}

	if err := helpers.ParseFormDataAndValidate(c, uploadSchema); err != nil {
		return c.Status(http.StatusBadRequest).JSON(response.HttpErrorResponse(err))
	}

	// role service
	uploadService := service.NewUploadService(db)
	data, err := uploadService.Update(id, *uploadSchema)

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"code":    http.StatusBadRequest,
			"message": "failed to update role",
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"code":    http.StatusOK,
		"message": "data has been updated",
		"data":    data,
	})
}

// RestoreUpload 	func for Restores Upload by given ID.
// @Description 	Restore Upload by given ID.
// @Summary 			Restore Upload by given ID
// @Tags 					Upload
// @Accept 				json
// @Produce 			json
// @Param 				id path string true "Upload ID"
// @Success 			200 {string} status "Ok"
// @Security 			ApiKeyAuth
// @Router 				/v1/upload/restore/{id} [put]
func RestoreUploadById(c *fiber.Ctx) error {
	db := config.GetDB()
	id, err := uuid.Parse(c.Params("id"))

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"code":    http.StatusBadRequest,
			"message": "failed to get role",
		})
	}

	uploadService := service.NewUploadService(db)
	err = uploadService.Restore(id)

	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"code":    http.StatusNotFound,
			"message": "data not found or has been deleted",
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"code":    http.StatusOK,
		"message": "data has been updated",
	})
}

// SoftDeleteUpload 	func for Soft Deletes Upload by given ID.
// @Description 			Soft Delete Upload by given ID.
// @Summary 					Soft Delete Upload by given ID
// @Tags 							Upload
// @Accept 						json
// @Produce 					json
// @Param 						id path string true "Upload ID"
// @Success 					200 {string} status "Ok"
// @Security 					ApiKeyAuth
// @Router 						/v1/upload/soft-delete/{id} [delete]
func SoftDeleteUploadById(c *fiber.Ctx) error {
	db := config.GetDB()
	id, err := uuid.Parse(c.Params("id"))

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"code":    http.StatusBadRequest,
			"message": "failed to get role",
		})
	}

	uploadService := service.NewUploadService(db)
	err = uploadService.SoftDelete(id)

	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"code":    http.StatusNotFound,
			"message": "data not found or has been deleted",
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"code":    http.StatusOK,
		"message": "data has been deleted",
	})
}

// ForceDeleteUpload 	func for Force Deletes Upload by given ID.
// @Description 			Force Delete Upload by given ID.
// @Summary 					Force Delete Upload by given ID
// @Tags 							Upload
// @Accept 						json
// @Produce 					json
// @Param 						id path string true "Upload ID"
// @Success 					200 {string} status "Ok"
// @Security 					ApiKeyAuth
// @Router 						/v1/upload/force-delete/{id} [delete]
func ForceDeleteUploadById(c *fiber.Ctx) error {
	db := config.GetDB()
	id, err := uuid.Parse(c.Params("id"))

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"code":    http.StatusBadRequest,
			"message": "failed to get role",
		})
	}

	uploadService := service.NewUploadService(db)
	err = uploadService.ForceDelete(id)

	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"code":    http.StatusNotFound,
			"message": "data not found or has been deleted",
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"code":    http.StatusOK,
		"message": "data has been deleted",
	})
}
