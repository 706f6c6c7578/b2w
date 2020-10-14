// go run b2w.go -i input file -o output file		- encode
// go run b2w.go -d -i output file -o input file	- decode

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
}

func main() {
	inFP := flag.String("i", "", "input file")
	outFP := flag.String("o", "", "output file")
	action := flag.Bool("d", false, "use -d to specify decode")
	flag.Parse()

	if *inFP == "" || *outFP == "" {
		fmt.Printf("Usage: b2w [-d] -i input file -o output file")
		os.Exit(1)
	}

	inFD, err := os.Open(*inFP)
	if err != nil {
		panic(err)
	}
	defer inFD.Close()

	outFD, err := os.Create(*outFP)
	if err != nil {
		panic(err)
	}
	defer outFD.Close()

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
	"Atlas",
	"Balsa",
	"Bauch",
	"Bauer",
	"Benno",
	"Beruf",
	"Besen",
	"Biene",
	"Birne",
	"Blatt",
	"Blech",
	"Blitz",
	"Blume",
	"Bluse",
	"Boden",
	"Bonus",
	"Boxer",
	"Brief",
	"Couch",
	"Dachs",
	"Datum",
	"Dauer",
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
	"Figur",
	"Fisch",
	"Fleck",
	"Folie",
	"Forum",
	"Franz",
	"Fritz",
	"Frost",
	"Fuchs",
	"Gabel",
	"Galle",
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
	"Grube",
	"Gummi",
	"Gunst",
	"Gurke",
	"Hafen",
	"Hagel",
	"Halle",
	"Harfe",
	"Haube",
	"Hebel",
	"Hecht",
	"Heike",
	"Helga",
	"Henne",
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
	"Kabel",
	"Kalle",
	"Kanal",
	"Karin",
	"Kasse",
	"Katze",
	"Kerze",
	"Kiste",
	"Kleid",
	"Klotz",
	"Kluft",
	"Knauf",
	"Knick",
	"Knopf",
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
	"Laube",
	"Leder",
	"Leine",
	"Lende",
	"Leser",
	"Leute",
	"Linie",
	"Lippe",
	"Liter",
	"Lotto",
	"Lunge",
	"Lurch",
	"Luxus",
	"Magen",
	"Maler",
	"Manko",
	"Mappe",
	"Maria",
	"Maske",
	"Mathe",
	"Meile",
	"Mensa",
	"Meter",
	"Milch",
	"Mixer",
	"Modus",
	"Monat",
	"Motor",
	"Mulde",
	"Musik",
	"Nacht",
	"Nadel",
	"Nager",
	"Narbe",
	"Natur",
	"Nebel",
	"Neffe",
	"Nelke",
	"Notiz",
	"Onkel",
	"Orgel",
	"Osten",
	"Paket",
	"Panne",
	"Paste",
	"Pegel",
	"Peter",
	"Pferd",
	"Platz",
	"Pokal",
	"Porto",
	"Prinz",
	"Pudel",
	"Punkt",
	"Qualm",
	"Quote",
	"Radio",
	"Rampe",
	"Raupe",
	"Regal",
	"Rente",
	"Reuse",
	"Riese",
	"Rille",
	"Rinde",
	"Rippe",
	"Roman",
	"Ruder",
	"Sache",
	"Salto",
	"Segel",
	"Seide",
	"Serie",
	"Sicht",
	"Sinus",
	"Sirup",
	"Socke",
	"Sorte",
	"Spalt",
	"Sport",
	"Stadt",
	"Stein",
	"Stift",
	"Stoff",
	"Stube",
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
	"Wange",
	"Watte",
	"Wespe",
	"Woche",
	"Wolke",
	"Wonne",
	"Wurst",
	"Xenia",
	"Yacht",
	"Zeile",
	"Ziege",
	"Zunge",
	"Zweig",
}
