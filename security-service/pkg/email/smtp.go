package email

import (
	"crypto/rand"
	"encoding/hex"
	"io/ioutil"
	"net/smtp"
	"strings"
)

func SendEmail(to string, subject string, smtpUser string, smtpPass string, smtpHost string, code string) error {
	from := smtpUser

	htmlData, err := ioutil.ReadFile("/app/template/email_template.html")
	if err != nil {
		return err
	}

	htmlBody := strings.Replace(string(htmlData), "{code}", code, -1)

	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: " + subject + "\n" +
		"MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n" +
		htmlBody

	err = smtp.SendMail(smtpHost+":587",
		smtp.PlainAuth("", smtpUser, smtpPass, smtpHost),
		from, []string{to}, []byte(msg))

	if err != nil {
		return err
	}
	return nil
}

func GenerateCode() (string, error) {
	length := 6
	bytes := make([]byte, length)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}
