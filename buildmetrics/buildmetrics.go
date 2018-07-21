// buildmetrics reads a list of Adobe glyph names as well as a set of
// AFM files and outputs a go source file containing font definitions
// and glyph metrics for the corresponding fonts.
//
// Adobe glyph names can be found in the github project
// github.com/adobe-type-tools/agl-aglfn
// in the files glyphlist.txt and zapfdingbats.txt.
//
// Font metrics for the PDF core 14 fonts can be downloaded from
// http://www.adobe.com/devnet/font.html
//
// ./buildmetrics --afm_dir=. --glyph_names=./glyphlist.txt
//   --dingbat_names=./zapfdingbats.txt --output=./metrics.go
// gofmt -w ./metrics.go
package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strconv"
	"strings"
)

var glyphToRune = map[string][]string{}

// Populate glyphToRune with data from provided file.
func readNames(r io.Reader) error {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "#") {
			continue
		}
		parts := strings.Split(line, ";")
		if len(parts) != 2 {
			return fmt.Errorf("Want glyph name and code point, got %v.", parts)
		}
		glyphName := parts[0]
		runes, ok := glyphToRune[glyphName]
		if !ok {
			runes = []string{}
		}
		for _, codePoint := range strings.Split(parts[1], " ") {
			if r, err := strconv.ParseInt(codePoint, 16, 32); err != nil {
				return fmt.Errorf("Could not parse code point %v.", codePoint)
			} else {
				runes = append(runes, strconv.QuoteRuneToASCII(rune(r)))
			}
		}
		glyphToRune[glyphName] = runes
	}
	return scanner.Err()
}

func readAFM(r io.Reader, w io.Writer) error {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "#") {
			continue
		}
		parts := strings.Split(line, " ")
		if parts[0] == "FontName" {
			fmt.Fprintf(w, "\t%#v: []glyphInfo{\n", parts[1])
		}
		if parts[0] == "C" {
			encoded := parts[1]
			width := parts[4]
			glyph := parts[7]
			unicode, ok := glyphToRune[glyph]
			if !ok {
				unicode = []string{"-1"}
			}
			for _, u := range unicode {
				fmt.Fprintf(w, "\t\t{%v, %v, %v}, // %v\n", u, encoded, width, glyph)
			}
		}
	}
	fmt.Fprint(w, "\t},\n")
	return scanner.Err()
}

func main() {
	afmDir := flag.String("afm_dir", "", "Directory with AFM files.")
	glyphNames := flag.String("glyph_names", "", "File with Adobe glyph names.")
	dingbatNames := flag.String("dingbat_names", "", "File with Adobe glyph names for Zapf Dingbats.")
	metricsFile := flag.String("output", "", "Output is written to this file.")
	flag.Parse()

	if *glyphNames != "" {
		if file, err := os.Open(*glyphNames); err != nil {
			log.Fatal(err)
		} else {
			defer file.Close()
			if err := readNames(file); err != nil {
				log.Fatal(err)
			}
		}
	}

	if *dingbatNames != "" {
		if file, err := os.Open(*dingbatNames); err != nil {
			log.Fatal(err)
		} else {
			defer file.Close()
			if err := readNames(file); err != nil {
				log.Fatal(err)
			}
		}
	}

	file, err := os.Create(*metricsFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	out := bufio.NewWriter(file)
	defer out.Flush()

	fmt.Fprint(out, goHeader)
	if afmNames, err := ioutil.ReadDir(*afmDir); err != nil {
		log.Fatal(err)
	} else {
		for _, fileInfo := range afmNames {
			if !strings.HasSuffix(fileInfo.Name(), ".afm") {
				continue
			}
			filePath := path.Join(*afmDir, fileInfo.Name())
			if file, err := os.Open(filePath); err != nil {
				log.Fatal(err)
			} else {
				defer file.Close()
				if err := readAFM(file, out); err != nil {
					log.Fatal(err)
				}
			}
		}
	}
	fmt.Fprint(out, goFooter)
}

var goHeader = `// Metrics and glyph code points for the 14 built-in PDF fonts. The AFM
// files from which this data was compiled are copyright © Adobe Systems
// Incorporated and are available at http://www.adobe.com/devnet/font.html
package pdf

type glyphInfo struct {
	rune rune

	// Code point of the rune in the font in its default encoding.
	// -1 if the rune is in the font but not encoded in the default encoding.
	codePoint int

	// Width of the glyph, as per-mille of the font size.
	width int
}

var fontMetrics = map[name][]glyphInfo{
`

var goFooter = "}\n"
