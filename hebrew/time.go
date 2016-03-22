package hebrewtime

type Weekday int

const (
	Rishon Weekday = 1 + iota
	Sheni
	Shlishi
	Revii
	Chamishi
	Shishi
	Shabbat
)

var weekdays = [...]string{
	"Yom Rishon",
	"Yom Sheni",
	"Yom Shlishi",
	"Yom Revi'i",
	"Yom Chamishi",
	"Yom Shishi",
	"Yom Shabbat",
}

func (w Weekday) String() string { return weekdays[w-1] }

type Month int

const (
	Nissan Month = 1 + iota
	Iyar
	Sivan
	Tammuz
	Av
	Elul
	Tishri
	Cheshvan
	Kislev
	Tevet
	Shvat
	Adar
)

var months = [...]string{
	"Nisan",
	"Iyar",
	"Sivan",
	"Tammuz",
	"Av",
	"Elul",
	"Tishrei",
	"Cheshvan",
	"Kislev",
	"Tevet",
	"Shvat",
	"Adar",
}

// String returns the English name of the month ("Nisan", "Iyyar", ...).
func (m Month) String() string { return months[m-1] }

type Constellation int

const (
	Taleh Constellation = 1 + iota
	Shor
	Teomim
	Sartan
	Arye
	Betulah
	Moznayim
	Akrab
	Keshet
	Gdi
	Dli
	Dagim
)

var constellations = [...]string{
	"Taleh",
	"Shor",
	"Teomim",
	"Sartan",
	"Arye",
	"Betulah",
	"Moznayim",
	"Akrab",
	"Keshet",
	"Gdi",
	"Dli",
	"Dagim",
}

// String returns the English name of the constellation ("Taleh", "Shor", ...).
func (c Constellation) String() string { return constellations[c-1] }

// cycleYear returns the first year for the current 19-year cycle.
// Adar I is added in the 3rd, 6th, 8th, 11th, 14th, 17th and 19th years of the cycle.
// This fixed calendar is based on mathematical and astronomical calculations,
// and aligns the lunar calendar with solar years.

func IsLeap(year int) bool {
	switch year % 19 {
	case 0, 3, 6, 8, 11, 14, 17:
		return true
	}
	return false
}
