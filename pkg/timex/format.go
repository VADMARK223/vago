package timex

import "time"

const Layout = "02.01.2006 15:04:05"

func Parse(s string) (time.Time, error) {
	return time.Parse(Layout, s)
}

func Format(t time.Time) string {
	return t.Format(Layout)
}

func ParseLocal(s string) (time.Time, error) {
	return time.ParseInLocation(Layout, s, time.Local)
}
