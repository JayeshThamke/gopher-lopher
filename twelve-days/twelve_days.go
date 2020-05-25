package twelve

import (
	"bytes"
	"fmt"
)

// [...] is a syntactic sugar for [12] - may be compiler determines
// the size of slice at runtime and converts into golang array construct
var days = [...]string{"first", "second", "third", "fourth", "fifth", "sixth",
	"seventh", "eighth", "ninth", "tenth", "eleventh", "twelfth"}

var gifts = [...]string{
	"a Partridge in a Pear Tree.",
	"two Turtle Doves, ",
	"three French Hens, ",
	"four Calling Birds, ",
	"five Gold Rings, ",
	"six Geese-a-Laying, ",
	"seven Swans-a-Swimming, ",
	"eight Maids-a-Milking, ",
	"nine Ladies Dancing, ",
	"ten Lords-a-Leaping, ",
	"eleven Pipers Piping, ",
	"twelve Drummers Drumming, "}

const startWith = "On the %s day of Christmas my true love gave to me: "

// Verse returns a single verse of poem
func Verse(day int) string {
	// strings are inherently immutable hence
	// bytes.Buffer offer performance boost
	// and does not put load on garbage collector
	var b bytes.Buffer

	// slices are 0 based
	day--
	b.WriteString(fmt.Sprintf(startWith, days[day]))
	for i := day; i >= 0; i-- {
		if i == 0 && day > 0 {
			b.WriteString("and ")
		}
		b.WriteString(gifts[i])
	}

	return b.String()
}

// Song returns a complete song
func Song() string {
	var b bytes.Buffer
	for i := 1; i <= len(days); i++ {
		b.WriteString(Verse(i) + "\n")
	}

	// return without last `new line` char.
	// hmm, looks somewhat complex but avoids
	// if condition in above for loop
	return b.String()[:len(b.String())-1]
}
