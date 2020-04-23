package timestamp

import (
	"strconv"
	"strings"
)

// Timestamp struct to hold timestamp from sbv file
type Timestamp struct {
	hour     int
	minute   int
	second   int
	millisec int
}

// Normalize returns timestamp in miliseconds as integer
func (t *Timestamp) Normalize() int {
	return t.millisec + t.second*1000 + t.minute*60*1000 + t.hour*60*60*1000
}

// Before Tests if this timestamp is before t2 time stamp
func (t *Timestamp) Before(t2 Timestamp) bool {
	return t.Normalize() < t2.Normalize()
}

// AsString returns timestamp as string
func (t *Timestamp) AsString() string {
	// we might need to add two trailing zeros to miliseconds
	millis := "00" + strconv.Itoa(t.millisec)
	millis = millis[len(millis)-3:]

	// we might need to add one trailing zero to miliseconds
	seconds := "0" + strconv.Itoa(t.second)
	seconds = seconds[len(seconds)-2:]

	minutes := "0" + strconv.Itoa(t.minute)
	minutes = minutes[len(minutes)-2:]

	// hours have no trailing zeros
	return strconv.Itoa(t.hour) + ":" + minutes + ":" + seconds + "." + millis
}

// ReadFromString parses hours, minutes and millisconds into struct
func ReadFromString(init string) Timestamp {
	var t Timestamp

	// initially split hours, minutes and remainder
	parts := strings.Split(init, ":")
	// TODO error handling
	t.hour, _ = strconv.Atoi(parts[0])
	t.minute, _ = strconv.Atoi(parts[1])

	// split remainder into seconds and milliseconds
	parts = strings.Split(parts[2], ".")
	// TODO error handling
	t.second, _ = strconv.Atoi(parts[0])
	t.millisec, _ = strconv.Atoi(parts[1])

	return t
}

// ReadFromInt initializes a timestamp from an integer
func ReadFromInt(i int) Timestamp {
	var t Timestamp

	t.hour = i / 3600000
	i %= (60 * 60 * 1000)
	t.minute = i / (60 * 1000)
	i %= (60 * 1000)
	t.second = i / 1000
	i %= 1000
	t.millisec = i
	return t
}

// ReadTwoFromString parses two timestamps from single string
func ReadTwoFromString(init string) (Timestamp, Timestamp) {
	parts := strings.Split(init, ",")
	t1 := ReadFromString(parts[0])
	t2 := ReadFromString(parts[1])
	return t1, t2
}
