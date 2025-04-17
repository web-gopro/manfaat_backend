package mail

import (
	"fmt"
	"log"
	"net/smtp"
)

func SendAdminMail(otp string) error {

	var (
		fromGmail = "xaydarovjasur6999@gmail.com"
		password  = "pzlt skhd nvgq fxks"
	)
//nodirxaydarov911@gmail.com
	var to []string = []string{"xaydarovjasur6999@gmail.com"}
	smsHost := "smtp.gmail.com"
	port := "587"

	message := []byte("user data \n" + otp)

	byteM := []byte(message)

	auth := smtp.PlainAuth("", fromGmail, password, smsHost)

	err := smtp.SendMail(smsHost+":"+port, auth, fromGmail, to, byteM)

	if err != nil {
		log.Println("error", err)
		return err
	}

	fmt.Println("sms successfully")

	return nil
}


func SendMail(to []string, otp string) error {

	var (
	  fromGmail = "xaydarovjasur6999@gmail.com"
	  password  = "pzlt skhd nvgq fxks"
	)
  
	smsHost := "smtp.gmail.com"
	port := "587"
  
	message := []byte("Your otp code is\n" + otp)
  
	byteM := []byte(message)
  
	auth := smtp.PlainAuth("", fromGmail, password, smsHost)
  
	err := smtp.SendMail(smsHost+":"+port, auth, fromGmail, to, byteM)
  
	if err != nil {
	  log.Println("error", err)
	  return err
	}
  
	fmt.Println("sms successfully")
  
	return nil
  }