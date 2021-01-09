package date

import "time"

const (
	layout = "2006-01-02 15:04:05"
)

func GetNow() time.Time {
	return time.Now().UTC()
}

func GetNowDateString() string {
	return GetNow().Format(layout)
}
