package mail

import (
	"fmt"

	campaing "github.com/henrique998/email-N/internal/domain/campaign"
	"github.com/resend/resend-go/v2"
)

func SendMail(campaign *campaing.Campaing) error {
	fmt.Println("Sending email...")

	apiKey := "re_NWNnxN7F_EygmAofMXaDGbk77CUJrE6zx"

	client := resend.NewClient(apiKey)

	var emails []string

	for _, contact := range campaign.Contacts {
		emails = append(emails, contact.Email)
	}

	params := &resend.SendEmailRequest{
		From:    "onboarding@resend.dev",
		To:      emails,
		Subject: campaign.Name,
		Html:    campaign.Content,
		Cc:      nil,
		Bcc:     nil,
		ReplyTo: "henriquemonteiro037@gmail.com",
	}

	_, err := client.Emails.Send(params)
	if err != nil {
		return err
	}

	return nil
}
