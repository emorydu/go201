package v2

import (
	"fmt"
	"net/smtp"

	"github.com/jordan-wright/email"
)

type EmailSenderAdapter struct {
	e *email.Email
}

func (adapter *EmailSenderAdapter) Send(subject, from string,
	to []string, content string, mailserver string, a smtp.Auth) error {
	adapter.e.Subject = subject
	adapter.e.From = from
	adapter.e.To = to
	adapter.e.Text = []byte(content)
	return adapter.e.Send(mailserver, a)
}

func ExampleSendMailWithDisclaimer() {
	adapter := &EmailSenderAdapter{e: email.NewEmail()}
	err := SendMailWithDisclaimer(adapter, "emorydu email test", "emorydu@gmail.com",
		[]string{"DEST_MAILBOX"}, "Welcome Emory.Du", "smtp.163.com:25",
		smtp.PlainAuth("", "emorydu@gmail.com", "emorydu@gmail.com.password", "smtp.163.com"))
	if err != nil {
		fmt.Printf("SendMail error: %s\n", err)
		return
	}
	fmt.Println("SendMail OK!")
}
