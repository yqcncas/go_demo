package gormtest

import (
	"crypto/tls"
	"net/smtp"
	"testing"

	"github.com/jordan-wright/email"
)

func TestSendEmail(t *testing.T) {
	e := email.NewEmail()
	e.From = "ych <yqcncas@163.com>"
	e.To = []string{"434664858@qq.com"}
	// e.Bcc = []string{"test_bcc@example.com"}
	// e.Cc = []string{"test_cc@example.com"}
	e.Subject = "我是主题测试" // 主题
	// e.Text = []byte("Text Body is, of course, supported!")
	e.HTML = []byte("您的验证码是: <h1>1234</h1>")
	// err := e.Send("smtp.163.com:587", smtp.PlainAuth("", "yqcncas@163.com", "IRZICUAHGJXVEHLO", "smtp.163.com"))
	err := e.SendWithTLS("smtp.163.com:587", smtp.PlainAuth("", "yqcncas@163.com", "IRZICUAHGJXVEHLO", "smtp.163.com"), &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         "smtp.163.com",
	})

	if err != nil {
		t.Fatal(err) // 当出现报错EOF时  需关闭SSL重试
	}
}

// IRZICUAHGJXVEHLO
