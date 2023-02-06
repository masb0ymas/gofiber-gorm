package config

import (
	"context"
	"log"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

// MinIO Client
func MinIOConnection() (*minio.Client, error) {
	var err error
	ctx := context.Background()

	MINIO_HOST := Env("MINIO_HOST", "127.0.0.1")
	MINIO_PORT := Env("MINIO_PORT", "9000")
	MINIO_ACCESS_KEY := Env("MINIO_ACCESS_KEY", "")
	MINIO_SECRET_KEY := Env("MINIO_SECRET_KEY", "")
	MINIO_BUCKET_NAME := Env("MINIO_BUCKET_NAME", "gofiber")
	MINIO_REGION := Env("MINIO_REGION", "ap-southeast-1")

	endpoint := MINIO_HOST + ":" + MINIO_PORT
	useSSL := false

	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(MINIO_ACCESS_KEY, MINIO_SECRET_KEY, ""),
		Secure: useSSL,
	})

	if err != nil {
		log.Fatalln(err.Error())
	}

	err = minioClient.MakeBucket(ctx, MINIO_BUCKET_NAME, minio.MakeBucketOptions{
		Region: MINIO_REGION,
	})

	if err != nil {
		exists, err := minioClient.BucketExists(ctx, MINIO_BUCKET_NAME)

		if err == nil && exists {
			log.Printf("We already own %s\n", MINIO_BUCKET_NAME)
		} else {
			log.Fatalln(err.Error())
		}
	} else {
		log.Printf("Successfully created %s\n", MINIO_BUCKET_NAME)
	}

	return minioClient, err
}
