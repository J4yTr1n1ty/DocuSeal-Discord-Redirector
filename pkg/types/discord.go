package types

type DiscordMessageWebhookPayload struct {
	Content   string         `json:"content"`
	Username  string         `json:"username"`
	AvatarURL string         `json:"avatar_url"`
	Embeds    []DiscordEmbed `json:"embeds"` // Maximum of 10 embeds
}

type DiscordEmbed struct {
	Title       string                `json:"title"`
	Type        string                `json:"type"`
	Description string                `json:"description"`
	URL         string                `json:"url"`
	Timestamp   string                `json:"timestamp"`
	Color       int                   `json:"color"`
	Footer      DiscordEmbedFooter    `json:"footer"`
	Image       DiscordEmbedImage     `json:"image"`
	Thumbnail   DiscordEmbedThumbnail `json:"thumbnail"`
	Video       DiscordEmbedVideo     `json:"video"`
	Provider    DiscordEmbedProvider  `json:"provider"`
	Author      DiscordEmbedAuthor    `json:"author"`
	Fields      []DiscordEmbedField   `json:"fields"`
}

type DiscordEmbedField struct {
	Name   string `json:"name"`
	Value  string `json:"value"`
	Inline bool   `json:"inline"`
}

type DiscordEmbedFooter struct {
	Text         string `json:"text"`
	IconURL      string `json:"icon_url"`
	ProxyIconURL string `json:"proxy_icon_url"`
}

type DiscordEmbedImage struct {
	URL      string `json:"url"`
	ProxyURL string `json:"proxy_url"`
	Height   int    `json:"height"`
	Width    int    `json:"width"`
}

type DiscordEmbedVideo struct {
	URL      string `json:"url"`
	ProxyURL string `json:"proxy_url"`
	Height   int    `json:"height"`
	Width    int    `json:"width"`
}

type DiscordEmbedThumbnail struct {
	URL      string `json:"url"`
	ProxyURL string `json:"proxy_url"`
	Height   int    `json:"height"`
	Width    int    `json:"width"`
}

type DiscordEmbedProvider struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type DiscordEmbedAuthor struct {
	Name         string `json:"name"`
	URL          string `json:"url"`
	IconURL      string `json:"icon_url"`
	ProxyIconURL string `json:"proxy_icon_url"`
}
