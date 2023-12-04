package handlers

import "time"

func ConvertSeconds(unixSeconds int64) string {
	t := time.Unix(unixSeconds, 0)
	rfc3339Time := t.Format(time.RFC3339)
	return rfc3339Time
}

func ConvertTime(unixSeconds int64) time.Time {

	t := time.Unix(unixSeconds, 0)

	return t
}
