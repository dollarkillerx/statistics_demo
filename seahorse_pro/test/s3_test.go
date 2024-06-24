package test

import (
	"bytes"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"log"
	"os"
	"testing"
)

func TestS3(t *testing.T) {
	// S3配置
	bucket := "jp-kyoto"
	key := "a.jpg"
	region := "us-east-1"
	endpoint := "https://s3.dollarkiller.com"
	accessKey := "NysFT683Xd9hUDlHMnzd"
	secretKey := "TyyBygmnymxJyDvFuINJbYITbHXF1kpduLe7Mvc4"

	// 创建一个AWS Session
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(region),
		Endpoint:    aws.String(endpoint),
		Credentials: credentials.NewStaticCredentials(accessKey, secretKey, ""),
	})
	if err != nil {
		log.Fatalf("Failed to create session: %v", err)
	}

	// 创建一个S3客户端
	svc := s3.New(sess)

	// 打开文件
	file, err := os.Open("a.png")
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	defer file.Close()

	// 读取文件内容
	fileInfo, _ := file.Stat()
	var size int64 = fileInfo.Size()
	buffer := make([]byte, size)
	file.Read(buffer)

	// 上传文件
	_, err = svc.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
		Body:   bytes.NewReader(buffer),
		ACL:    aws.String("public-read"),
	})
	if err != nil {
		log.Fatalf("Failed to upload data to %s/%s, %v", bucket, key, err)
	}

	fmt.Printf("Successfully uploaded %s to %s\n", key, bucket)
}

func TestP2(t *testing.T) {
	bucket := "jp-kyoto"
	endpoint := "https://s3.dollarkiller.com"
	accessKey := "NysFT683Xd9hUDlHMnzd"
	secretKey := "TyyBygmnymxJyDvFuINJbYITbHXF1kpduLe7Mvc4"
	// 设置 AWS 凭证
	awsCredentials := credentials.NewStaticCredentialsFromCreds(
		credentials.Value{
			AccessKeyID:     accessKey,
			SecretAccessKey: secretKey,
		},
	)
	// 创建 AWS 会话
	sess, err := session.NewSession(&aws.Config{
		Region:           aws.String("us-west-2"), //"us-west-2"),
		Credentials:      awsCredentials,
		Endpoint:         aws.String(endpoint),
		S3ForcePathStyle: aws.Bool(true),
	})
	if err != nil {
		fmt.Println("Failed to create AWS session:", err)
		panic(err)
	}

	// 创建 S3 上传管理器
	uploader := s3manager.NewUploader(sess)

	// 打开要上传的文件
	file, err := os.Open("a.png")
	if err != nil {
		fmt.Println("Failed to open file:", err, ",filepath:", "a.png")
		panic(err)
	}
	defer file.Close()

	//key := file.Name()
	// 上传文件到 S3
	result, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucket),      //"your-bucket-name"),
		Key:    aws.String("/up/a.png"), //"path/to/upload/file.txt"),
		Body:   file,
		ACL:    aws.String(s3.BucketCannedACLPublicRead),
	})
	if err != nil {
		fmt.Println("Failed to upload file:", err)
		panic(err)
	}

	fmt.Println(result)
}
