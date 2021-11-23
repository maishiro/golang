package util

import (
	"encoding/base64"
	"fmt"
	"strings"
)

type Mail struct {
	Boudary    string
	From       string
	To         []string
	Cc         []string
	Bcc        []string
	Subject    string
	TextBody   string
	HtmlBody   string
	Attachment []byte
	AttachName string
}

func BuildMessage(mail Mail) string {

	msg := ""
	msg += fmt.Sprintf("From: %s\r\n", mail.From)

	if len(mail.To) > 0 {
		msg += fmt.Sprintf("To: %s\r\n", strings.Join(mail.To, ";"))
	}

	if len(mail.Cc) > 0 {
		msg += fmt.Sprintf("Cc: %s\r\n", strings.Join(mail.Cc, ";"))
	}

	msg += fmt.Sprintf("Subject: %s\r\n", mail.Subject)

	msg += "MIME-version: 1.0\r\n"
	if mail.TextBody != "" || mail.HtmlBody != "" || 0 < len(mail.Attachment) {
		msg += "Content-Type: multipart/alternative; boundary=\"" + mail.Boudary + "\"\r\n"
		msg += "\r\n"
	}

	// text/plain mail body
	if mail.TextBody != "" {
		msg += "--" + mail.Boudary + "\r\n"

		msg += "Content-Type: text/plain; charset=\"UTF-8\"\r\n"
		// msg += "Content-Transfer-Encoding: base64\r\n"
		msg += "\r\n"
		msg += mail.TextBody
		msg += "\r\n"
		msg += "\r\n"
	}

	// text/html mail body
	if mail.HtmlBody != "" {
		msg += "--" + mail.Boudary + "\r\n"

		msg += "Content-Type: text/html; charset=\"UTF-8\"\r\n"
		// msg += "Content-Transfer-Encoding: base64\r\n"
		msg += "\r\n"
		msg += mail.HtmlBody
		msg += "\r\n"
		msg += "\r\n"
	}

	// Attachment
	if 0 < len(mail.Attachment) {
		msg += "--" + mail.Boudary + "\r\n"

		msg += "Content-Type: text/plain; charset=\"UTF-8\"\r\n"
		msg += "Content-Transfer-Encoding: base64\r\n"
		msg += fmt.Sprintf("Content-Disposition: attachment; filename=%s\r\n", mail.AttachName)
		msg += "Content-ID: <words.txt>\r\n"

		msg += "\r\n"

		b := make([]byte, base64.StdEncoding.EncodedLen(len(mail.Attachment)))
		base64.StdEncoding.Encode(b, mail.Attachment)
		msg += string(b)

		msg += "\r\n"
		msg += "\r\n"
	}

	// mail body end
	if mail.TextBody != "" || mail.HtmlBody != "" || 0 < len(mail.Attachment) {
		msg += "--" + mail.Boudary + "\r\n"
	}

	return msg
}
