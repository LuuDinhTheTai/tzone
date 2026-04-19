package email

import (
	"fmt"
	"net/smtp"
	"os"
)

func SendOTP(to string, otp string, purpose string) error {
	host := os.Getenv("SMTP_HOST")
	port := os.Getenv("SMTP_PORT")
	user := os.Getenv("SMTP_USER")
	pass := os.Getenv("SMTP_PASS")
	from := os.Getenv("SMTP_FROM")

	if host == "" || port == "" || user == "" || pass == "" {
		return fmt.Errorf("email service is not configured")
	}

	if from == "" {
		from = user
	}

	subject := "TZone verification code"
	body := fmt.Sprintf("Your verification code for %s is: %s\nThis code expires in 5 minutes.", purpose, otp)
	msg := []byte("To: " + to + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"MIME-Version: 1.0\r\n" +
		"Content-Type: text/plain; charset=UTF-8\r\n\r\n" +
		body + "\r\n")

	auth := smtp.PlainAuth("", user, pass, host)
	addr := host + ":" + port

	return smtp.SendMail(addr, auth, from, []string{to}, msg)
}
