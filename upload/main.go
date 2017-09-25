package main

import ( //"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/aws/session"

	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"os"
)

func main() {
	accesid := "AKIAIOHZVBINWEK35PZA"//os.Getenv("s3accessid") ////
	accesskey := "I2ia8U3QQsQO8eF6Rce/P259ovcdl9JNBV8lIXLv"//os.Getenv("secretacceskey") ////
	token := ""
	cred := credentials.NewStaticCredentials(accesid, accesskey, token)
	conf := aws.NewConfig()
	conf.Credentials = cred
	conf.Region = aws.String("us-east-1")
	sess := session.Must(session.NewSession(conf))
	upload := s3manager.NewUploader(sess)
	f, err := os.Open("file")
	if err != nil {
		fmt.Print("open:",err)
	}
		upoadinput := s3manager.UploadInput{Bucket: aws.String("jenkins19"), Body: f,Key:aws.String("a")}
	_, err = upload.Upload(&upoadinput)
	if err != nil {
		fmt.Print("upload",err)
	}

}
