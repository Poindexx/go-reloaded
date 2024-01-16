package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func replaceHex(input string) string {
	re := regexp.MustCompile(`^\(hex\)(\s|$)`)
	if re.MatchString(input) {
		return re.ReplaceAllString(input, "")
	}
	reHex := regexp.MustCompile(`(\b[0-9A-Fa-f]+)\s*\(hex\)`)
	return reHex.ReplaceAllStringFunc(input, func(match string) string {
		hexNumber := reHex.FindStringSubmatch(match)[1]
		decimalNumber, _ := strconv.ParseInt(hexNumber, 16, 64)
		return strconv.FormatInt(decimalNumber, 10)
	})
}

func replaceBin(input string) string {
	re := regexp.MustCompile(`^\(bin\)(\s|$)`)
	if re.MatchString(input) {
		return re.ReplaceAllString(input, "")
	}
	reBin := regexp.MustCompile(`(\b[01]+)\s*\(bin\)`)
	return reBin.ReplaceAllStringFunc(input, func(match string) string {
		binNumber := reBin.FindStringSubmatch(match)[1]
		decimalNumber, _ := strconv.ParseInt(binNumber, 2, 64)
		return strconv.FormatInt(decimalNumber, 10)
	})
}

func capitalize(input string) string {
	words := strings.Fields(input)
	for i, word := range words {
		if len(word) > 0 {
			words[i] = strings.ToUpper(word[:1]) + strings.ToLower(word[1:])
		}
	}
	result := strings.Join(words, " ")
	return result
}

func modifyCase(input string) string {
	re := regexp.MustCompile(`^\((up|low|cap)(?:, (\d+))?\)$`)
	if re.MatchString(input) {
		return re.ReplaceAllString(input, "")
	}
	saan := make([]int, 0)
	reModifiers1 := regexp.MustCompile(`(up|low|cap)(?:, (\d+))?`)
	kkol := ""
	reModifiers1.ReplaceAllStringFunc(input, func(match string) string {
		kkol = reModifiers1.FindStringSubmatch(match)[2]
		kkkol, _ := strconv.Atoi(kkol)
		saan = append(saan, kkkol)
		return kkol
	})
	oz := 	input
	for {
		modified := oz
		for i := 0; i < len(saan); i++ {
			if saan[i] >= 0 {
				b := 0
				if saan[i] > 0 {
					b = saan[i] - 1
				}
				
				reModifiers := regexp.MustCompile(fmt.Sprintf(`((?:\S*\b[\w']+\b[^\w']*){0,%d}\b[\w']+\b)\s*\((up|low|cap)(?:, (\d+))?\)`, b))

				// reModifiers := regexp.MustCompile(fmt.Sprintf(`((?:[^\w']*\b[\w']+\b[^\w']*){0,%d}\b[\w']+\b)\s*\((up|low|cap)(?:, (\d+))?\)`, b))
				oz = reModifiers.ReplaceAllStringFunc(oz, func(match string) string {
					soz := reModifiers.FindStringSubmatch(match)[1]
					mod := reModifiers.FindStringSubmatch(match)[2]
					kolll := reModifiers.FindStringSubmatch(match)[3]
					kollll, _ := strconv.Atoi(kolll)
					var ss string
					if kollll == saan[i] {
						switch mod {
						case "up":
							ss = strings.ToUpper(soz)
						case "low":
							ss = strings.ToLower(soz)
						case "cap":
							ss = capitalize(soz)
						}
						return ss
					}
					return match
				})
			}
		}
		if modified == oz {
			break
		}
	}
	return oz
}

func processA(input string) string {
	re := regexp.MustCompile(`\b([aA])\s+(\b\w+\b)?`)
	return re.ReplaceAllStringFunc(input, func(match string) string {
		soz := re.FindStringSubmatch(match)[1]
		nextWord := re.FindStringSubmatch(match)[2]
		if len(nextWord) > 0 && (strings.Contains("aeiouhAEIOUH", string(nextWord[0]))) {
			return soz + "n" + " " + nextWord
		}
		return soz + " " + nextWord
	})
}

func formatPunctuation(input string) string {
	rePunctuation := regexp.MustCompile(`(?: [.,!?:;])`)
	input = rePunctuation.ReplaceAllStringFunc(input, func(match string) string {
		soz := rePunctuation.FindStringSubmatch(match)[0]
		return strings.TrimSpace(soz)
	})
	reComma := regexp.MustCompile(`,([^ ])`)
	input = reComma.ReplaceAllString(input, ", $1")
	reSingleQuotes := regexp.MustCompile(`' \s*([^']+)\s* '`)
	input = reSingleQuotes.ReplaceAllString(input, "'$1'")
	reSingleQuotes = regexp.MustCompile(`'\s*([^']+)\s* '`)
	input = reSingleQuotes.ReplaceAllString(input, "'$1'")
	reSingleQuotes = regexp.MustCompile(`' \s*([^']+)\s*'`)
	input = reSingleQuotes.ReplaceAllString(input, "'$1'")
	return input
}



func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run main.go input.txt output.txt")
		os.Exit(1)
	}

	inputFileName := os.Args[1]
	outputFileName := os.Args[2]

	inputFile, err := os.Open(inputFileName)
	if err != nil {
		fmt.Println("Error opening input file:", err)
		os.Exit(1)
	}
	defer inputFile.Close()
	scanner := bufio.NewScanner(inputFile)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	var modifiedLines []string
	for _, line := range lines {
		line = replaceHex(line)
		line = replaceBin(line)
		line = modifyCase(line)
		line = processA(line)
		line = formatPunctuation(line)
		modifiedLines = append(modifiedLines, line)
	}

	outputFile, err := os.Create(outputFileName)
	if err != nil {
		fmt.Println("Error creating output file:", err)
		os.Exit(1)
	}
	defer outputFile.Close()

	for _, modifiedLine := range modifiedLines {
		fmt.Fprintln(outputFile, modifiedLine)
	}


}