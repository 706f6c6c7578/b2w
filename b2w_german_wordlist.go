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
			fmt.Fprintf(btw.out, "%s\r\n", strings.Join(final[i:len(final)], " "))
		}
	}
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
	"Achim",
	"Adler",
	"Ahorn",
	"Album",
	"Algen",
	"Alina",
	"Ampel",
	"Angel",
	"Anruf",
	"Anton",
	"Anzug",
	"Apfel",
	"April",
	"Aroma",
	"Asche",
	"Atlas",
	"Balsa",
	"Bauch",
	"Bauer",
	"Benno",
	"Beruf",
	"Besen",
	"Bibel",
	"Biene",
	"Birne",
	"Blatt",
	"Blech",
	"Blitz",
	"Block",
	"Bluse",
	"Boden",
	"Bonus",
	"Boxer",
	"Brief",
	"Couch",
	"Datum",
	"David",
	"Diana",
	"Docht",
	"Dogge",
	"Doris",
	"Duden",
	"Durst",
	"Edgar",
	"Eimer",
	"Eisen",
	"Engel",
	"Erbse",
	"Erika",
	"Erwin",
	"Essig",
	"Eulen",
	"Fabel",
	"Fahne",
	"Farbe",
	"Fasan",
	"Fauna",
	"Fazit",
	"Feder",
	"Feier",
	"Felge",
	"Felix",
	"Fiale",
	"Figur",
	"Fleck",
	"Folie",
	"Forum",
	"Franz",
	"Fritz",
	"Frost",
	"Fuchs",
	"Fulda",
	"Garde",
	"Gasse",
	"Gatte",
	"Geist",
	"Gerda",
	"Geste",
	"Gilde",
	"Gitta",
	"Glanz",
	"Gleis",
	"Golem",
	"Gramm",
	"Gummi",
	"Gurke",
	"Hafen",
	"Hagel",
	"Halle",
	"Harfe",
	"Haube",
	"Havel",
	"Heber",
	"Hecht",
	"Heike",
	"Helga",
	"Henne",
	"Hirse",
	"Hitze",
	"Honig",
	"Hotel",
	"Ilona",
	"Imker",
	"Index",
	"Insel",
	"Jacke",
	"Jolle",
	"Jubel",
	"Junge",
	"Jutta",
	"Kader",
	"Kamin",
	"Kanal",
	"Karin",
	"Kehle",
	"Kerze",
	"Keule",
	"Kiste",
	"Kleid",
	"Klotz",
	"Kluft",
	"Knauf",
	"Knick",
	"Knopf",
	"Knute",
	"Kohle",
	"Komma",
	"Konto",
	"Kraut",
	"Kreis",
	"Krimi",
	"Krone",
	"Kugel",
	"Kunst",
	"Lachs",
	"Lager",
	"Lampe",
	"Laura",
	"Lehre",
	"Leine",
	"Lende",
	"Leser",
	"Leute",
	"Licht",
	"Linie",
	"Lippe",
	"Liter",
	"Lotto",
	"Lurch",
	"Luxus",
	"Magie",
	"Maler",
	"Manko",
	"Mappe",
	"Maria",
	"Maske",
	"Mathe",
	"Mensa",
	"Milch",
	"Mixer",
	"Modus",
	"Monat",
	"Mosel",
	"Motor",
	"Mulde",
	"Musik",
	"Nadel",
	"Nager",
	"Natur",
	"Nebel",
	"Neffe",
	"Nelke",
	"Notiz",
	"Obhut",
	"Orgel",
	"Osten",
	"Paket",
	"Panne",
	"Paste",
	"Peter",
	"Pferd",
	"Pflug",
	"Platz",
	"Pokal",
	"Porto",
	"Prinz",
	"Pudel",
	"Punkt",
	"Qualm",
	"Quote",
	"Radio",
	"Rasen",
	"Raupe",
	"Regal",
	"Rente",
	"Reuse",
	"Rhein",
	"Riese",
	"Rille",
	"Rinde",
	"Rippe",
	"Robbe",
	"Roman",
	"Ruder",
	"Sache",
	"Salto",
	"Seide",
	"Serie",
	"Sicht",
	"Sinus",
	"Sirup",
	"Socke",
	"Spalt",
	"Spiel",
	"Sport",
	"Stadt",
	"Stein",
	"Stift",
	"Stoff",
	"Stube",
	"Sturm",
	"Suppe",
	"Tacho",
	"Tafel",
	"Tanne",
	"Tarif",
	"Teich",
	"Tempo",
	"Tisch",
	"Titel",
	"Tonne",
	"Torte",
	"Traum",
	"Tulpe",
	"Umbau",
	"Umweg",
	"Umzug",
	"Uschi",
	"Vater",
	"Venus",
	"Vogel",
	"Waage",
	"Wagen",
	"Waldi",
	"Wange",
	"Wespe",
	"Woche",
	"Wolke",
	"Wonne",
	"Wucht",
	"Wulst",
	"Xenia",
	"Yacht",
	"Zeile",
	"Ziege",
	"Zweig",
}
