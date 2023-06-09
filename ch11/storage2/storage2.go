package storage2

import (
	"fmt"
	"log"
	"net/smtp"
)

func BytesInUse(username string)int64{return 980000000}
const template = "warning , you are using %d bytes of storage, %d%% of your quota"

const sender = "abc@example.com"
const password = "123"
const hostname = "smtp.example.com"

var notifyUser = func(username,msg string){
	auth := smtp.PlainAuth("",sender,password,hostname)
	err := smtp.SendMail(hostname+":587",auth,sender,[]string{username},[]byte(msg))
	if err !=nil{
		log.Printf("smtp.SendMail(%s)failed:%s",username,err)
	}
}

func CheckQuota(username string) {
	used := BytesInUse(username)
	const quota = 1000000000 //1GB
	percent := 100 * used / quota
	if percent < 90 {
		return //OK
	}
	msg := fmt.Sprintf(template,used,percent)
	notifyUser(username,msg)
}

