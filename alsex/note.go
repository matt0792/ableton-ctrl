package alsex

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var noteValues = map[string]int{
	"C":  0,
	"C#": 1,
	"Db": 1,
	"D":  2,
	"D#": 3,
	"Eb": 3,
	"E":  4,
	"F":  5,
	"F#": 6,
	"Gb": 6,
	"G":  7,
	"G#": 8,
	"Ab": 8,
	"A":  9,
	"A#": 10,
	"Bb": 10,
	"B":  11,
}

func ToMidi(note string) (int32, error) {
	if len(note) < 2 || len(note) > 4 {
		return 0, errors.New("invalid note format")
	}

	octaveStart := -1
	for i, c := range note {
		if c >= '0' && c <= '9' || c == '-' {
			octaveStart = i
			break
		}
	}

	if octaveStart == -1 {
		return 0, errors.New("no octave number found")
	}

	noteName := strings.TrimSpace(note[:octaveStart])
	octaveStr := note[octaveStart:]

	baseValue, exists := noteValues[noteName]
	if !exists {
		return 0, fmt.Errorf("unknown note: %s", noteName)
	}

	octave, err := strconv.Atoi(octaveStr)
	if err != nil {
		return 0, fmt.Errorf("invalid octave: %w", err)
	}

	midiValue := int32((octave+1)*12 + baseValue)

	if midiValue < 0 || midiValue > 127 {
		return 0, fmt.Errorf("MIDI value %d out of range (0-127)", midiValue)
	}

	return midiValue, nil
}
