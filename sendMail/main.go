package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/smtp"
	"sendmail/util"
)

func main() {

	// Sender
	from := "from@gmail.com"
	password := "<Email Password>"

	// Receiver
	to := []string{
		"send1@example.com",
		"send2@example.com",
		"send3@example.com",
		"send4@example.com",
	}
	cc := []string{
		"cc1@example.com",
		"cc2@example.com",
	}
	//
	bcc := []string{
		"bcc1@example.com",
	}

	boundary := "my-boundary-365"

	// smtp configuration
	smtpHost := "localhost"
	smtpPort := "25"

	subject := "test mail"
	body_text := "email body message"
	body_html := `<p>email <b>body</b> message</p>`

	var attachment []byte = []byte{}
	attach_name := "sample.txt"

	// data := readFile("out.png")
	// attachment = data
	// attach_name = "out.png"

	request := util.Mail{
		Attachment: attachment,
		AttachName: attach_name,
		Boudary:    boundary,
		From:       from,
		To:         to,
		Cc:         cc,
		Bcc:        bcc,
		Subject:    subject,
		TextBody:   body_text,
		HtmlBody:   body_html,
	}

	msg := util.BuildMessage(request)

	// Authentication
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Send email
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, []byte(msg))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Email Sent Success")
}

func readFile(fileName string) []byte {

	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}

	return data
}
