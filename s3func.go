package main

import ( //"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/aws/session"

	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"os"
	"github.com/gorilla/mux"
	"net/http"
	"log"
)

type file struct {
	Filename , Bucketname ,Objectname string
}

/*func Configuration()  {
	accesid := "AKIAIOHZVBINWEK35PZA"                       //os.Getenv("s3accessid") ////
	accesskey := "I2ia8U3QQsQO8eF6Rce/P259ovcdl9JNBV8lIXLv" //os.Getenv("secretacceskey") ////
	token := ""
	cred := credentials.NewStaticCredentials(accesid, accesskey, token)
	conf := aws.NewConfig()
	conf.Credentials = cred
	conf.Region = aws.String("us-east-1")
	return conf
}*/
func Uploadobject(w http.ResponseWriter, req *http.Request) {
	var file1 file
	params := req.URL.Query()
	file1.Filename=params["filename"][0]
	file1.Bucketname=params["bucketname"][0]
	file1.Objectname=params["objectname"][0]
	fmt.Fprintln(w,file1)
	accesid := "AKIAIOHZVBINWEK35PZA"                       //os.Getenv("s3accessid") ////
	accesskey := "I2ia8U3QQsQO8eF6Rce/P259ovcdl9JNBV8lIXLv" //os.Getenv("secretacceskey") ////
	token := ""
	cred := credentials.NewStaticCredentials(accesid, accesskey, token)
	conf := aws.NewConfig()
	conf.Credentials = cred
	conf.Region = aws.String("us-east-1")
	sess := session.Must(session.NewSession(conf))
	upload := s3manager.NewUploader(sess)
	file := file1.Filename
	f, err := os.Open(file)
	if err != nil {
		fmt.Print("open:", err)
	}
	buckname := file1.Bucketname
	objname := file1.Objectname
	upoadinput := s3manager.UploadInput{Bucket: aws.String(buckname), Body: f, Key: aws.String(objname)}
	out,err := upload.Upload(&upoadinput)
	if err != nil {
		fmt.Print("upload", err)
	}
	fmt.Println("uploadoutput:",out)

}

func Listobject(w http.ResponseWriter, req *http.Request) {
	var file1 file
	params := req.URL.Query()
	file1.Bucketname=params["bucketname"][0]
	accesid := "AKIAIOHZVBINWEK35PZA"                       //os.Getenv("s3accessid") ////
	accesskey := "I2ia8U3QQsQO8eF6Rce/P259ovcdl9JNBV8lIXLv" //os.Getenv("secretacceskey") ////
	token := ""
	cred := credentials.NewStaticCredentials(accesid, accesskey, token)
	conf := aws.NewConfig()
	conf.Region = aws.String("us-east-1")
	conf.Credentials = cred
	svc := s3.New(session.New(conf))
	buckname := file1.Bucketname
	listobj, err := svc.ListObjects(&s3.ListObjectsInput{Bucket: aws.String(buckname)})
	if err != nil {
		fmt.Println(err)

	}
	fmt.Fprintln(w,listobj)

}

func Deleteobj(w http.ResponseWriter, req *http.Request) {
	var file1 file
	params := req.URL.Query()
	file1.Bucketname=params["bucketname"][0]
	file1.Objectname=params["objectname"][0]
	fmt.Fprintln(w,file1)
	accesid := "AKIAIOHZVBINWEK35PZA"                       //os.Getenv("s3accessid") ////
	accesskey := "I2ia8U3QQsQO8eF6Rce/P259ovcdl9JNBV8lIXLv" //os.Getenv("secretacceskey") ////
	token := ""
	cred := credentials.NewStaticCredentials(accesid, accesskey, token)
	conf := aws.NewConfig()
	conf.Region = aws.String("us-east-1")
	conf.Credentials = cred
	svc := s3.New(session.New(conf))
	buckname := file1.Bucketname
	objname := file1.Objectname
	deleteout, err := svc.DeleteObject(&s3.DeleteObjectInput{
		Bucket: aws.String(buckname),
		Key:    aws.String(objname),
	})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Fprintln(w,deleteout)

}

func Downloadobj(w http.ResponseWriter, req *http.Request) {
	var file1 file
	params := req.URL.Query()
	file1.Bucketname=params["bucketname"][0]
	file1.Objectname=params["objectname"][0]
	fmt.Fprintln(w,file1)
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
	buckname := file1.Bucketname
	objname := file1.Objectname
	n, err := downloader.Download(f, &s3.GetObjectInput{
		Bucket: aws.String(buckname),
		Key:    aws.String(objname),
	})
	if err != nil {
		fmt.Println("cannot download file", err)
	}
	fmt.Fprintln(w,"new file downloaded:", n)

}
