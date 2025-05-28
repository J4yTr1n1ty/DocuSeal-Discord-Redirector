package types

type DocuSealEvent struct {
	EventType string `json:"event_type"` // form.viewed, form.started, form.completed, form.declined
	Timestamp string `json:"timestamp"`
	Data      struct {
		ID            int    `json:"id"`
		SubmissionID  int    `json:"submission_id"`
		Email         string `json:"email"`
		Phone         string `json:"phone"`
		Name          string `json:"name"`
		UserAgent     string `json:"ua"`
		IP            string `json:"ip"`
		SentAt        string `json:"sent_at"`
		OpenedAt      string `json:"opened_at"`
		CompletedAt   string `json:"completed_at"`
		DeclinedAt    string `json:"declined_at"`
		CreatedAt     string `json:"created_at"`
		UpdatedAt     string `json:"updated_at"`
		Status        string `json:"status"`
		DeclineReason string `json:"decline_reason"`
		Role          string `json:"role"`
		Values        []struct {
			Field string `json:"field"`
			Value string `json:"value"`
		}
		Documents []struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		} `json:"documents"`
		AuditLogUrl   string `json:"audit_log_url"`
		SubmissionUrl string `json:"submission_url"`
	} `json:"data"`
}

const (
	DocuSealEventTypeFormViewed    = "form.viewed"
	DocuSealEventTypeFormStarted   = "form.started"
	DocuSealEventTypeFormCompleted = "form.completed"
	DocuSealEventTypeFormDeclined  = "form.declined"
)
