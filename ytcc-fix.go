package main

import (
	"bufio"
	"flag"
	"os"

	"./timestamp"
)

// NEWLINE neue Zeile
const NEWLINE = "\n"

// OFFSET Zeilenabstand zwischen zwei Untertiteln
const OFFSET = 3

var infilename = flag.String("infile", "captions.sbv",
	"YouTube captions input file name")
var outfilename = flag.String("outfile", "captions-fixed.sbv",
	"repaired output captions")

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func fixTimecode(this, next string) string {
	thisStart, thisEnd := timestamp.ReadTwoFromString(this)
	nextStart, _ := timestamp.ReadTwoFromString(next)

	var newEnd timestamp.Timestamp

	// test if we are already not overlapping, so keep it
	if thisEnd.Normalize() < nextStart.Normalize() {
		newEnd = thisEnd
	} else {
		// if we are overlapping, just cut it down.
		newEnd = timestamp.ReadFromInt(nextStart.Normalize() - 1)
	}

	return thisStart.AsString() + "," + newEnd.AsString()
}

func main() {
	flag.Parse()

	lines := make(map[int]string)

	infile, err := os.Open(*infilename)
	check(err)
	outfile, err := os.Create(*outfilename)
	check(err)

	input := bufio.NewScanner(infile)

	// slurp the full input file
	i := 0
	for input.Scan() {
		lines[i] = input.Text()
		i++
	}

	// gou through the input data and fix
	// timecodes if needed
	for y := 0; y < i; {
		if y+OFFSET < i {
			_, err = outfile.WriteString(fixTimecode(lines[y], lines[y+OFFSET]))
			check(err)
		} else {
			_, err = outfile.WriteString(lines[y])
			check(err)
		}
		_, err = outfile.WriteString(NEWLINE)
		check(err)
		_, err = outfile.WriteString(lines[y+1])
		check(err)
		_, err = outfile.WriteString(NEWLINE)
		check(err)
		_, err = outfile.WriteString(NEWLINE)
		check(err)

		y += OFFSET
	}

	infile.Close()
	outfile.Close()
}
