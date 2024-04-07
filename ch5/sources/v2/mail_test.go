package v2

import (
	"net/smtp"
	"testing"
)

type FakeMailSender struct {
	subject string
	from    string
	to      []string
	content string
}

func (s *FakeMailSender) Send(subject, from string,
	to []string, content string, mailserver string, a smtp.Auth) error {
	s.subject = subject
	s.from = from
	s.to = to
	s.content = content
	return nil
}

func TestSendMailWithDisclaimer(t *testing.T) {
	s := &FakeMailSender{}
	err := SendMailWithDisclaimer(s, "gopher", "Emory163@163.com", []string{"Emory163@163.com"}, "Welcome, Emory.Du",
		"smtp.163.com:25", smtp.PlainAuth("", "Emory163@163.com", "Djf@2441cn", "smtp.163.com"))
	if err != nil {
		t.Fatalf("want: nil, actual: %s\n", err)
		return
	}
	want := "Welcome, Emory.Du" + "\n\n" + DISCLAIMER
	if s.content != want {
		t.Fatalf("want: %s, actual: %s\n", want, s.content)
	}
}
