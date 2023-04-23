package util

import (
	"fmt"
	"github.com/df-mc/dragonfly/server"
	"math"
	"strings"
	"unicode"
)

var (
	Server *server.Server
)

func IsStringAlphanumeric(str string) bool {
	for _, char := range str {
		if !unicode.IsLetter(char) && !unicode.IsNumber(char) {
			return false
		}
	}
	return true
}

func InArray(val string, arr []string) bool {
	for _, item := range arr {
		if item == val {
			return true
		}
	}
	return false
}

func RemoveElementFromArray(arr []string, element string) []string {
	for i, v := range arr {
		if v == element {
			return append(arr[:i], arr[i+1:]...)
		}
	}
	return arr
}

// FormatSeconds 0 = LONG | 1 = SHORT
func FormatSeconds(seconds int, typeVal int) string {
	if seconds == -1 {
		return "Permanent"
	}

	d := seconds / (3600 * 24)
	h := (seconds % (3600 * 24)) / 3600
	m := (seconds % 3600) / 60
	s := seconds % 60

	var dDisplay, hDisplay, mDisplay, sDisplay string

	if d > 0 {
		if typeVal == 0 {
			if d == 1 {
				dDisplay = fmt.Sprintf("%d jour, ", d)
			} else {
				dDisplay = fmt.Sprintf("%d jours, ", d)
			}
		} else {
			dDisplay = fmt.Sprintf("%dj, ", d)
		}
	}
	if h > 0 {
		if typeVal == 0 {
			if h == 1 {
				hDisplay = fmt.Sprintf("%d heure, ", h)
			} else {
				hDisplay = fmt.Sprintf("%d heures, ", h)
			}
		} else {
			hDisplay = fmt.Sprintf("%dh, ", h)
		}
	}
	if m > 0 {
		if typeVal == 0 {
			if m == 1 {
				mDisplay = fmt.Sprintf("%d minute, ", m)
			} else {
				mDisplay = fmt.Sprintf("%d minutes, ", m)
			}
		} else {
			mDisplay = fmt.Sprintf("%dm, ", m)
		}
	}
	if s > 0 {
		if typeVal == 0 {
			if s == 1 {
				sDisplay = fmt.Sprintf("%d seconde, ", s)
			} else {
				sDisplay = fmt.Sprintf("%d secondes, ", s)
			}
		} else {
			sDisplay = fmt.Sprintf("%ds, ", s)
		}
	}

	format := strings.TrimRight(dDisplay+hDisplay+mDisplay+sDisplay, ", ")

	if strings.Count(format, ",") > 0 {
		return strings.Replace(format, ",", " et", 1)
	} else {
		return format
	}
}

func IndexOf(slice []string, value string) int {
	for i, val := range slice {
		if val == value {
			return i
		}
	}
	return -1
}

func SendTip(tip string) {
	for _, p := range Server.Players() {
		p.SendTip(tip)
	}
}

func FormatInt(n int) string {
	base := 1000
	prefixes := []string{"", "k", "M", "G", "T", "P", "E", "Z", "Y"}

	if n < base {
		return fmt.Sprintf("%d", n)
	}

	exp := int(math.Log(float64(n)) / math.Log(float64(base)))
	pre := prefixes[exp]

	scaled := float64(n) / math.Pow(float64(base), float64(exp))
	format := "%.1f"

	if scaled < 10.0 {
		format = "%.2f"
	}
	if scaled >= 100.0 {
		format = "%.0f"
	}

	return fmt.Sprintf(format+"%s", scaled, pre)
}
