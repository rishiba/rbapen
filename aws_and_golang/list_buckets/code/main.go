// Function to list the buckets.
package main

import "fmt"

func main() {
	fmt.Println("vim-go")
	svc := s3.New(session.New())
	input := &s3.ListBucketsInput{}

	result, err := svc.ListBuckets(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

func getAccessKeyAndSecretKey() (string, string) {
	aws_access_key := "test"
	aws_secret_key := "test2"
	return aws_access_key, aws_secret_key
}

func getSdkConfig(aws_access_key, aws_secret_key, region) *Config {

	token := ""
	credentials := credentials.NewStaticCredentials(aws_access_key, aws_secret_key, token)
	config := aws.NewConfig().WithRegion(region).WithCredentials(credentials)
	return config
}

type FileDetails struct {
	filePath string
	fileData []byte
	fileType string
	fileSize int64
}

func getS3Service(config *Config) *S3 {
}

func getFilesInDirectory(dirName) []string {

}

func createPutObjectInput(bucketName, file *FileDetails) {
}
