package handlers

import (
	"fmt"
	"time"
	"unicode"
)

func ConvertSeconds(unixSeconds int64) string {
	t := time.Unix(unixSeconds, 0)
	rfc3339Time := t.Format(time.RFC3339)
	return rfc3339Time
}

func ConvertTime(unixSeconds int64) time.Time {
	t := time.Unix(unixSeconds, 0)
	return t
}

func CapitalizeWords(input string) string {
	r := []rune(input)
	r[0] = unicode.ToUpper(r[0])
	return string(r)
}

func AdjustTimeTime(Timestamp string, TimeZone int) string {

	var remoteDate string

	// Calculate new timestime removing the timezone
	t, _ := time.Parse(time.RFC3339, Timestamp)
	timezoneDuration := time.Duration(TimeZone) * time.Second
	newTimestamp := t.Add(timezoneDuration)

	remoteDate = newTimestamp.Format(time.RFC3339)

	return remoteDate
}

func formatDate(input string) string {
	t, err := time.Parse("2006-01-02", input)
	if err != nil {
		fmt.Println(err)
	}
	outputDate := t.Format("02-01-2006")
	return outputDate
}
