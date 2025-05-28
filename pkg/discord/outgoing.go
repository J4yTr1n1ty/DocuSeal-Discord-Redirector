package discord

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/J4yTr1n1ty/DocuSeal-Discord-Redirector/pkg/config"
	"github.com/J4yTr1n1ty/DocuSeal-Discord-Redirector/pkg/types"
)

func AssembleMessage(event types.DocuSealEvent) error {
	payload := types.DiscordMessageWebhookPayload{
		Username: "DocuSeal",
	}

	switch event.EventType {
	case types.DocuSealEventTypeFormViewed:
		addViewEmbed(event, &payload)
		log.Println("Handling form.started event from " + event.Data.IP)
	case types.DocuSealEventTypeFormStarted:
		log.Println("Handling form.started event from " + event.Data.IP)
		addStartEmbed(event, &payload)
	case types.DocuSealEventTypeFormCompleted:
		log.Println("Handling form.completed event from " + event.Data.IP)
		addCompleteEmbed(event, &payload)
	case types.DocuSealEventTypeFormDeclined:
		log.Println("Handling form.declined event from " + event.Data.IP)
		addDeclineEmbed(event, &payload)
	default:
		return fmt.Errorf("unknown event type: %s", event.EventType)
	}

	return SendWebhook(payload)
}

func addViewEmbed(event types.DocuSealEvent, payload *types.DiscordMessageWebhookPayload) {
	openedAt, err := ConvertToDiscordTimestampWithFormat(event.Data.OpenedAt, "f")
	if err != nil {
		log.Println(err)
		openedAt = event.Data.OpenedAt
	}

	embed := types.DiscordEmbed{
		Title:       "Someone viewed a form",
		Type:        "rich",
		Description: "Click the title to view the submission.",
		URL:         event.Data.SubmissionUrl,
		Color:       0x4287f5,
		Fields: []types.DiscordEmbedField{
			{
				Name:   "Role",
				Value:  event.Data.Role,
				Inline: true,
			},
			{
				Name:   "Email",
				Value:  event.Data.Email,
				Inline: true,
			},
			{
				Name:   "Opened At",
				Value:  openedAt,
				Inline: false,
			},
			{
				Name:   "Status",
				Value:  event.Data.Status,
				Inline: true,
			},
		},
	}

	payload.Embeds = append(payload.Embeds, embed)
}

func addStartEmbed(event types.DocuSealEvent, payload *types.DiscordMessageWebhookPayload) {
	embed := types.DiscordEmbed{
		Title:       "Someone started a form",
		URL:         event.Data.SubmissionUrl,
		Description: "Click the title to view the submission.",
		Type:        "rich",
		Color:       0xE6D629,
		Fields: []types.DiscordEmbedField{
			{
				Name:   "Role",
				Value:  event.Data.Role,
				Inline: true,
			},
			{
				Name:   "Email",
				Value:  event.Data.Email,
				Inline: true,
			},
		},
	}

	payload.Embeds = append(payload.Embeds, embed)
}

func addCompleteEmbed(event types.DocuSealEvent, payload *types.DiscordMessageWebhookPayload) {
	var signatureImageURL string
	for _, value := range event.Data.Values {
		// Check if the value contains "signature" or if the signature image URL is already set
		if strings.Contains(strings.ToLower(value.Field), "signature") || signatureImageURL != "" {
			signatureImageURL = value.Value
		}
	}

	completedAt, err := ConvertToDiscordTimestampWithFormat(event.Data.CompletedAt, "f")
	if err != nil {
		log.Println(err)
		completedAt = event.Data.CompletedAt
	}

	embed := types.DiscordEmbed{
		Title:       "Someone completed a form",
		Type:        "rich",
		URL:         event.Data.SubmissionUrl,
		Description: "Click the title to view the submission.",
		Color:       0x38EB56,
		Fields: []types.DiscordEmbedField{
			{
				Name:   "Role",
				Value:  event.Data.Role,
				Inline: true,
			},
			{
				Name:   "Email",
				Value:  event.Data.Email,
				Inline: true,
			},
			{
				Name:  "Completed At",
				Value: completedAt,
			},
			{
				Name:  "Audit Log URL",
				Value: "[Click here](" + event.Data.AuditLogUrl + ")",
			},
			{
				Name: "Signature",
			},
		},
		Image: types.DiscordEmbedImage{
			URL: signatureImageURL,
		},
	}

	payload.Embeds = append(payload.Embeds, embed)
}

func addDeclineEmbed(event types.DocuSealEvent, payload *types.DiscordMessageWebhookPayload) {
	declinedAt, err := ConvertToDiscordTimestampWithFormat(event.Data.DeclinedAt, "f")
	if err != nil {
		log.Println(err)
		declinedAt = event.Data.DeclinedAt
	}

	embed := types.DiscordEmbed{
		Title:       "Someone declined a form",
		Type:        "rich",
		Description: "Click the title to view the submission.",
		URL:         event.Data.SubmissionUrl,
		Color:       0xFF0000,
		Fields: []types.DiscordEmbedField{
			{
				Name:   "Role",
				Value:  event.Data.Role,
				Inline: true,
			},
			{
				Name:   "Email",
				Value:  event.Data.Email,
				Inline: true,
			},
			{
				Name:  "Declined At",
				Value: declinedAt,
			},
			{
				Name:  "Audit Log URL",
				Value: "[Click here](" + event.Data.AuditLogUrl + ")",
			},
			{
				Name:  "Decline Reason",
				Value: event.Data.DeclineReason,
			},
		},
	}

	payload.Embeds = append(payload.Embeds, embed)
}

func SendWebhook(payload types.DiscordMessageWebhookPayload) error {
	webhookURL := config.Config.DiscordWebhookURL
	if webhookURL == "" {
		return fmt.Errorf("webhook URL not set")
	}

	body, err := json.Marshal(payload)
	if err != nil {
		log.Println("Error marshalling webhook payload: " + err.Error())
		return err
	}

	r, err := http.NewRequest("POST", webhookURL, bytes.NewBuffer(body))
	if err != nil {
		log.Println("Error creating webhook request: " + err.Error())
		return err
	}

	r.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(r)
	if err != nil {
		log.Println("Error sending webhook: " + err.Error())
		return err
	}
	if resp.StatusCode != http.StatusOK {
		defer resp.Body.Close()
		body, _ := io.ReadAll(resp.Body)
		log.Println("Error sending webhook: " + string(body) + " (" + string(resp.StatusCode) + ")")
		return fmt.Errorf("webhook returned status code %d", resp.StatusCode)
	}
	defer resp.Body.Close()

	return nil
}
