package house

const testVersion = 1

var items = []string{"house that Jack built.", "malt", "rat", "cat",
	"dog", "cow with the crumpled horn", "maiden all forlorn",
	"man all tattered and torn", "priest all shaven and shorn",
	"rooster that crowed in the morn", "farmer sowing his corn",
	"horse and the hound and the horn"}
var actions = []string{"lay in", "ate", "killed", "worried", "tossed", "milked",
	"kissed", "married", "woke", "kept", "belonged to", ""}

// Verse returns only the nth verse
func Verse(n int) string {
	// zero indexed
	n = n - 1
	verse := "This is the "
	verse += items[n]
	if n > 0 {
		verse = line(n-1, verse)
	}
	return verse
}

func line(i int, verses string) string {
	verses += "\n"
	verses += "that "
	verses += actions[i]
	verses += " the "
	verses += items[i]
	if i == 0 {
		return verses
	}
	return line(i-1, verses)
}

// Song returns the entire Song
func Song() (song string) {
	for n := 1; n < 13; n++ {
		song += Verse(n)
		if n < 12 {
			song += "\n\n"
		}
	}
	return
}
