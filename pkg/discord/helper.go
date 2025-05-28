package discord

import (
	"fmt"
	"log"
	"time"
)

// ConvertToDiscordTimestampWithFormat converts an ISO 8601 timestamp to Discord timestamp with specific format
// format options: t=short time, T=long time, d=short date, D=long date, f=short datetime, F=long datetime, R=relative
func ConvertToDiscordTimestampWithFormat(isoTimestamp, format string) (string, error) {
	parsedTime, err := time.Parse(time.RFC3339, isoTimestamp)
	if err != nil {
		log.Println("Error parsing timestamp: " + err.Error())
		return "", err
	}

	unixTimestamp := parsedTime.Unix()
	return fmt.Sprintf("<t:%d:%s>", unixTimestamp, format), nil
}
