package main

import (
	"github.com/aws/aws-sdk-go/aws/session"

	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"

	"os"
)

func main() {
	accesid := "AKIAIOHZVBINWEK35PZA"                       //os.Getenv("s3accessid") ////
	accesskey := "I2ia8U3QQsQO8eF6Rce/P259ovcdl9JNBV8lIXLv" //os.Getenv("secretacceskey") ////
	token := ""
	cred := credentials.NewStaticCredentials(accesid, accesskey, token)
	conf := aws.NewConfig()
	conf.Region = aws.String("us-east-1")
	conf.Credentials = cred
	sess := session.Must(session.NewSession(conf))
	downloader := s3manager.NewDownloader(sess)
	f, err := os.Create("file")
	if err != nil {
		fmt.Println("cannot create file", err)
	}
	n, err := downloader.Download(f, &s3.GetObjectInput{
		Bucket: aws.String("jenkins19"),
		Key:    aws.String("a"),
	})
	if err != nil {
		fmt.Println("cannot download file", err)
	}
	fmt.Println("new file downloaded:", n)

}
