// go run b2w.go < input file > output file    - encode
// go run b2w.go -d < output file > input file - decode

package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

func find(s string, WORDS []string) int {
	for i, v := range WORDS {
		if v == s {
			return i
		}
	}
	return -1
}

type Bintowords struct {
	in  io.Reader
	out io.Writer
}

func (btw Bintowords) read(p []string) (int, error) {
	inBuffer := make([]byte, len(p))
	n, err := btw.in.Read(inBuffer)
	if err != nil {
		return n, err
	}

	c := 0
	for i := 0; i < n; i++ {
		p[i] = WORDS[inBuffer[i]]
		c++
	}
	return c, nil
}

func (btw Bintowords) decode() {
	scanner := bufio.NewScanner(btw.in)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		v := find(scanner.Text(), WORDS)
		if v < 0 {
			panic("word not found at dict")
		}
		btw.out.Write([]byte{byte(v)})
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

}

func (btw Bintowords) encode() {
	final := []string{}
	for {
		encoded := make([]string, 1024)
		n, err := btw.read(encoded)
		if err != nil {
			final = append(final, encoded[0:n]...)
			if err == io.EOF {
				break
			} else {
				panic(err)
			}
		}

		final = append(final, encoded[0:n]...)
	}
	for i := 0; i < len(final); i += 10 {
		if i+10 < len(final) {
			// win32 ln
			fmt.Fprintf(btw.out, "%s\r\n", strings.Join(final[i:(i+10)], " "))
		} else {
			fmt.Fprintf(btw.out, "%s", strings.Join(final[i:len(final)], " "))
		}
	}
	fmt.Println()
}

func main() {
	action := flag.Bool("d", false, "use -d to specify decode")
	flag.Parse()

	inFD := os.Stdin
	outFD := os.Stdout

	if *action == false {
		Bintowords := Bintowords{inFD, outFD}
		Bintowords.encode()
	} else {
		Bintowords := Bintowords{inFD, outFD}
		Bintowords.decode()
	}
}

// WORDS is our dict to convert binary to string
var WORDS = []string{
	"about",
	"above",
	"acorn",
	"admit",
	"adore",
	"again",
	"agree",
	"ahead",
	"album",
	"alias",
	"allow",
	"alone",
	"among",
	"apple",
	"apron",
	"arise",
	"aroma",
	"audio",
	"award",
	"bacon",
	"baker",
	"basic",
	"batch",
	"begin",
	"below",
	"bikes",
	"bingo",
	"blame",
	"bless",
	"blink",
	"blitz",
	"board",
	"bogus",
	"boost",
	"bored",
	"boxer",
	"brain",
	"bread",
	"bride",
	"bring",
	"brown",
	"brush",
	"buddy",
	"build",
	"bulge",
	"bunny",
	"buyer",
	"cabin",
	"cadet",
	"cameo",
	"canal",
	"candy",
	"carry",
	"cause",
	"cease",
	"cedar",
	"chair",
	"chest",
	"chief",
	"cigar",
	"cinch",
	"civil",
	"claim",
	"class",
	"clear",
	"cliff",
	"climb",
	"clock",
	"clown",
	"coast",
	"condo",
	"could",
	"cover",
	"crawl",
	"cream",
	"crust",
	"curve",
	"daily",
	"dance",
	"debut",
	"defer",
	"delay",
	"delve",
	"dense",
	"depot",
	"depth",
	"derby",
	"disco",
	"ditto",
	"doing",
	"donor",
	"dowel",
	"dozen",
	"draft",
	"dress",
	"drink",
	"drive",
	"early",
	"elect",
	"elite",
	"empty",
	"enjoy",
	"equal",
	"error",
	"event",
	"exact",
	"faith",
	"fancy",
	"favor",
	"fence",
	"floor",
	"focus",
	"force",
	"frame",
	"fresh",
	"fruit",
	"given",
	"glove",
	"grass",
	"green",
	"group",
	"guard",
	"guest",
	"guide",
	"happy",
	"hedge",
	"hello",
	"honey",
	"hotel",
	"house",
	"hurry",
	"icing",
	"ideal",
	"image",
	"input",
	"juice",
	"knock",
	"known",
	"learn",
	"level",
	"limit",
	"local",
	"lower",
	"lucky",
	"lunch",
	"madam",
	"march",
	"mason",
	"metal",
	"model",
	"morse",
	"motor",
	"movie",
	"music",
	"never",
	"night",
	"noise",
	"north",
	"nylon",
	"often",
	"onion",
	"other",
	"owner",
	"paint",
	"paper",
	"party",
	"penny",
	"phone",
	"piano",
	"pilot",
	"place",
	"plaza",
	"prime",
	"proud",
	"quick",
	"quite",
	"radio",
	"react",
	"ready",
	"relax",
	"remix",
	"reply",
	"reset",
	"retry",
	"rigid",
	"river",
	"round",
	"route",
	"salad",
	"scale",
	"scene",
	"scope",
	"seven",
	"share",
	"shell",
	"shirt",
	"silly",
	"since",
	"skill",
	"sleep",
	"small",
	"smart",
	"smile",
	"smoke",
	"snowy",
	"solar",
	"solid",
	"sorry",
	"speed",
	"sport",
	"stamp",
	"stand",
	"stick",
	"store",
	"style",
	"sugar",
	"sweet",
	"swing",
	"table",
	"taken",
	"taste",
	"teach",
	"thank",
	"there",
	"those",
	"thumb",
	"tired",
	"title",
	"today",
	"touch",
	"trade",
	"trend",
	"truck",
	"tulip",
	"twice",
	"under",
	"unite",
	"until",
	"upper",
	"urban",
	"usual",
	"valid",
	"value",
	"venus",
	"vibes",
	"video",
	"voice",
	"voter",
	"water",
	"wheel",
	"world",
	"wrote",
	"xerox",
	"young",
	"youth",
	"zeros",
}
