package helper

import (
	"crypto/tls"
	"math/rand"
	"net/smtp"
	"strconv"
	"time"

	"github.com/jordan-wright/email"
)

func SendEmail(toUserEmail, code string) error {
	e := email.NewEmail()
	e.From = "ych <yqcncas@163.com>"
	e.To = []string{toUserEmail}
	// e.Bcc = []string{"test_bcc@example.com"}
	// e.Cc = []string{"test_cc@example.com"}
	e.Subject = "我是主题测试" // 主题
	// e.Text = []byte("Text Body is, of course, supported!")
	e.HTML = []byte("您的验证码是: <h1>" + code + "</h1>")
	// err := e.Send("smtp.163.com:587", smtp.PlainAuth("", "yqcncas@163.com", "IRZICUAHGJXVEHLO", "smtp.163.com"))
	err := e.SendWithTLS("smtp.163.com:587", smtp.PlainAuth("", "yqcncas@163.com", "IRZICUAHGJXVEHLO", "smtp.163.com"), &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         "smtp.163.com",
	})
	return err
}

func GetRandom() string {
	rand.Seed(time.Now().Unix())
	// rand.Intn(len(ranDomString)
	s := ""
	for i := 0; i < 6; i++ {
		s += strconv.Itoa(rand.Intn(10))
	}
	return s
}
