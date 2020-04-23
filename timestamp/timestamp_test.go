package timestamp

import (
	"testing"
)

func TestTimestampNormalize(t *testing.T) {
	t1 := Timestamp{hour: 1, minute: 0, second: 0, millisec: 0}
	if t1.Normalize() != 3600000 {
		t.Error(`t1 - 1 hour normalize failed`)
	}
	t1 = Timestamp{hour: 0, minute: 1, second: 0, millisec: 0}
	if t1.Normalize() != 60000 {
		t.Error(`t1 - 1 minute normalize failed`)
	}
	t1 = Timestamp{hour: 0, minute: 0, second: 1, millisec: 0}
	if t1.Normalize() != 1000 {
		t.Error(`t1 - 1 second normalize failed`)
	}
	t1 = Timestamp{hour: 0, minute: 0, second: 0, millisec: 1}
	if t1.Normalize() != 1 {
		t.Error(`t1 - 1 milisecond normalize failed`)
	}
	t1 = Timestamp{hour: 0, minute: 0, second: 0, millisec: 0}
	if t1.Normalize() != 0 {
		t.Error(`t1 - 0 millisec normalize failed`)
	}
}

func TestAsString(t *testing.T) {
	t1 := Timestamp{1, 5, 5, 207}
	if t1.AsString() != "1:05:05.207" {
		t.Error(`t1 - 1:05:05.207`)
	}

	t1 = Timestamp{99, 0, 0, 7}
	if t1.AsString() != "99:00:00.007" {
		t.Error(`t1 - 99:00:00.007`)
	}

	t1 = Timestamp{0, 0, 0, 7}
	if t1.AsString() != "0:00:00.007" {
		t.Error(`t1 - 0:00:00.007`)
	}
}

func TestTimestampReadFromString(t *testing.T) {

	// random value test
	t1 := ReadFromString("2:10:19.920")
	if t1.hour != 2 {
		t.Error(`t1 - hour failed`)
	}
	if t1.minute != 10 {
		t.Error(`t1 - minute failed`)
	}
	if t1.second != 19 {
		t.Error(`t1 - second failed`)
	}
	if t1.millisec != 920 {
		t.Error(`t1 - millisc failed`)
	}

	// null test
	t1 = ReadFromString("0:0:0.000")
	t2 := Timestamp{0, 0, 0, 0}
	if t1 != t2 {
		t.Error(`t1 - null val failed`)
	}

	// null test short millisecon string
	t1 = ReadFromString("0:0:0.0")
	t2 = Timestamp{0, 0, 0, 0}
	if t1 != t2 {
		t.Error(`t1 - null val failed`)
	}
}

func TestReadTwoFromString(t *testing.T) {
	testcase := "0:00:21.720,10:00:29.930"
	t1, t2 := ReadTwoFromString(testcase)

	if t1.Normalize() != 21720 {
		t.Error(`t1 - failed`)
	}
	if t2.Normalize() != 36029930 {
		t.Error(`t2 - failed`)
	}

}

func TestReadFromInt(t *testing.T) {
	t1 := ReadFromInt(1)

	if t1.Normalize() != 1 {
		t.Error(t1.Normalize())
	}

	t1 = ReadFromInt(1000)
	if t1.Normalize() != 1000 {
		t.Error(`t1 - failed`)
	}

	t1 = ReadFromInt(60000)
	if t1.Normalize() != 60000 {
		t.Error(`t1 - failed`)
	}

	t1 = ReadFromInt(60000 * 60)
	if t1.Normalize() != 60000*60 {
		t.Error(`t1 - failed`)
	}

}
