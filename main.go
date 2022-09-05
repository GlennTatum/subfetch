package main

import (
	"fmt"

	aic "github.com/TheZoraiz/ascii-image-converter/aic_package"
	"github.com/fatih/color"

	"bufio"
	"log"
	"strings"
)

// Use psuedo afs filesystem to store files

func main() {

	filePath := "images/A-resized.png"

	flags := aic.DefaultFlags()

	flags.Dimensions = []int{50, 20}
	flags.Colored = true

	asciiArt, err := aic.Convert(filePath, flags)
	if err != nil {
		log.Fatal(err)
	}

	scan := bufio.NewScanner(strings.NewReader(asciiArt))

	count := 0

	blue := color.New(color.FgBlue, color.Bold).SprintFunc()

	white := color.New(color.FgWhite).SprintFunc()

	vars := map[int]string{
		4: fmt.Sprintf("%v%v%v", blue("ACE"), white("@"), blue("Inwod-207 St")),
		5: "-----------------",
		6: fmt.Sprintf("%v%v%v", blue("Station"), white(": "), white("207th Street Station")),
		7: fmt.Sprintf("%v%v%v", blue("Uptown"), white(": "), white("3min, 10min, 12min")),
		8: fmt.Sprintf("%v%v%v", blue("Downton"), white(": "), white("5min, 13min, 15min")),
	}

	for scan.Scan() {
		count++

		fmt.Println(scan.Text(), vars[count])
	}
}
