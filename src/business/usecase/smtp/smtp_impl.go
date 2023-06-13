package smtp

import (
	"bytes"
	"crypto/tls"
	"path"

	"github.com/alitdarmaputra/fims-be/src/business/entity"
	"github.com/alitdarmaputra/fims-be/src/business/model"
	"github.com/alitdarmaputra/fims-be/src/config"
	"github.com/alitdarmaputra/fims-be/utils"
	"github.com/k3a/html2text"
	"gopkg.in/gomail.v2"
)

type SMTPUsecaseImpl struct {
	cfg config.SMTP
}

func InitSMTPUsecase(cfg config.SMTP) SmtpUsecase {
	return &SMTPUsecaseImpl{
		cfg: cfg,
	}
}

func (usecase *SMTPUsecaseImpl) SendMail(user *model.User, data *entity.EmailData) error {
	from := usecase.cfg.EmailFrom
	smtpPass := usecase.cfg.Password
	smtpUser := usecase.cfg.Username
	to := user.Email
	smtpHost := usecase.cfg.Host
	smtpPort := usecase.cfg.Port

	var body bytes.Buffer
	template, err := utils.ParseTemplateDir(path.Join("src", "docs", "template"))
	if err != nil {
		return err
	}

	template.ExecuteTemplate(&body, "verification_code", data)

	m := gomail.NewMessage()
	m.SetHeader("From", from)
	m.SetHeader("To", to)
	m.SetHeader("Subject", data.Subject)
	m.SetBody("text/html", body.String())
	m.AddAlternative("text/plain", html2text.HTML2Text(body.String()))

	d := gomail.NewDialer(smtpHost, smtpPort, smtpUser, smtpPass)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	err = d.DialAndSend(m)

	if err != nil {
		return err
	}

	return nil
}

func (usecase *SMTPUsecaseImpl) SendResetToken(user *model.User, data *entity.EmailData) error {
	from := usecase.cfg.EmailFrom
	smtpPass := usecase.cfg.Password
	smtpUser := usecase.cfg.Username
	to := user.Email
	smtpHost := usecase.cfg.Host
	smtpPort := usecase.cfg.Port

	var body bytes.Buffer
	template, err := utils.ParseTemplateDir(path.Join("src", "docs", "template"))
	if err != nil {
		return err
	}

	template.ExecuteTemplate(&body, "reset_code", data)

	m := gomail.NewMessage()
	m.SetHeader("From", from)
	m.SetHeader("To", to)
	m.SetHeader("Subject", data.Subject)
	m.SetBody("text/html", body.String())
	m.AddAlternative("text/plain", html2text.HTML2Text(body.String()))

	d := gomail.NewDialer(smtpHost, smtpPort, smtpUser, smtpPass)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	err = d.DialAndSend(m)

	if err != nil {
		return err
	}
	return nil
}

func (usecase *SMTPUsecaseImpl) SendUpdate(user *model.User, data *entity.EmailData) error {
	from := usecase.cfg.EmailFrom
	smtpPass := usecase.cfg.Password
	smtpUser := usecase.cfg.Username
	to := data.Email
	smtpHost := usecase.cfg.Host
	smtpPort := usecase.cfg.Port

	var body bytes.Buffer
	template, err := utils.ParseTemplateDir(path.Join("src", "docs", "template"))
	if err != nil {
		return err
	}

	template.ExecuteTemplate(&body, "update", data)

	m := gomail.NewMessage()
	m.SetHeader("From", from)
	m.SetHeader("To", to)
	m.SetHeader("Subject", data.Subject)
	m.SetBody("text/html", body.String())
	m.AddAlternative("text/plain", html2text.HTML2Text(body.String()))

	d := gomail.NewDialer(smtpHost, smtpPort, smtpUser, smtpPass)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	err = d.DialAndSend(m)

	if err != nil {
		return err
	}
	return nil
}
