// so far spent50 minutes
package twelve

import (
	"fmt"
)

const (
	testVersion = 1
)

//Verse returns one verse of the song
func Verse(numGifts int) string {
	numGifts = numGifts - 1
	numGiftsToDays := []string{
		"first", "second", "third", "fourth", "fifth", "sixth",
		"seventh", "eighth", "ninth", "tenth", "eleventh", "twelfth",
	}
	var verse string
	gifts := map[string]string{
		"first":    "a Partridge in a Pear Tree.",
		"second":   "two Turtle Doves",
		"third":    "three French Hens",
		"fourth":   "four Calling Birds",
		"fifth":    "five Gold Rings",
		"sixth":    "six Geese-a-Laying",
		"seventh":  "seven Swans-a-Swimming",
		"eighth":   "eight Maids-a-Milking",
		"ninth":    "nine Ladies Dancing",
		"tenth":    "ten Lords-a-Leaping",
		"eleventh": "eleven Pipers Piping",
		"twelfth":  "twelve Drummers Drumming",
	}
	if numGifts == 0 {
		return fmt.Sprintf(
			"On the %v day of Christmas my true love gave to me, %v",
			numGiftsToDays[numGifts], gifts[numGiftsToDays[numGifts]],
		)
	}
	verse += fmt.Sprintf("On the %v day of Christmas my true love gave to me, ", numGiftsToDays[numGifts])
	for i := numGifts; i > 0; i-- {
		verse += fmt.Sprintf("%v, ", gifts[numGiftsToDays[i]])
	}
	verse += fmt.Sprintf("and %v", gifts["first"])
	return verse
}

//Song returns the entire song, complete with all 12 verses
func Song() string {
	song := ""
	for i := 1; i < 13; i++ {
		song += Verse(i)
		song += "\n"
	}
	return song
}
