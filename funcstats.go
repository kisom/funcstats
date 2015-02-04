package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"regexp"
	"sort"
	"strings"
)

const (
	funcName   = "funcname"
	function   = "function"
	funcStart  = "funcstart"
	funcEnd    = "funcend"
	ignoreLine = "ignoreline"
)

var cRE = map[string]*regexp.Regexp{
	function:   regexp.MustCompile(`^(?:static )?\w+(?: \*)?$`),
	funcName:   regexp.MustCompile(`^(\w+)\(`),
	funcStart:  regexp.MustCompile(`^{$`),
	funcEnd:    regexp.MustCompile(`^}$`),
	ignoreLine: regexp.MustCompile(`^\s*((\/\*.+%)|(\*.*)|(\*\/))\s*$`),
}

var goRE = map[string]*regexp.Regexp{
	function:   regexp.MustCompile(`^func .+{$`),
	funcName:   regexp.MustCompile(`^func (?:\(\w+ \*?\w+\) )?(\w+).+{$`),
	funcStart:  regexp.MustCompile(`^func .+{$`),
	funcEnd:    regexp.MustCompile(`^}$`),
	ignoreLine: regexp.MustCompile(`^\s*((\/\/.+)|(\))|(}))\s*$`),
}

type langMap map[string]*regexp.Regexp

type stats struct {
	Max     int
	MaxName string
	Min     int
	MinName string
	Mean    int
	Median  int
	N       int
}

func computeStats(fm map[string]int) stats {
	if len(fm) == 0 {
		return stats{}
	}

	s := stats{N: len(fm)}
	var counts = make([]int, 0, s.N)
	for k, v := range fm {
		counts = append(counts, v)
		if v > s.Max {
			s.Max = v
			s.MaxName = k
		}

		if s.Min == 0 {
			s.Min = v
			s.MinName = k
		} else if v < s.Min {
			s.Min = v
			s.MinName = k
		}

		s.Mean += v
	}
	sort.Ints(counts)
	s.Median = counts[s.N/2]
	s.Mean /= s.N
	return s
}

func displayFuncStats(name string, s stats) {
	if s.N == 0 {
		fmt.Println("No functions counted for", name, ".")
		return
	}

	fmt.Printf(`Function stats for %s:
	%d functions counted
	%s is the longest function with %d lines
	%s is the shortest function with %d lines
	The mean number of lines per function is %d
	The median number of lines per function is %d
`, name, s.N, s.MaxName, s.Max, s.MinName, s.Min, s.Mean, s.Median)
}

func displayLineStats(name string, s stats) {
	if s.N == 0 {
		fmt.Println("No functions counted for", name, ".")
		return
	}

	fmt.Printf(`Line count stats for %s:
	%d lines counted
	%s has the longest line with %d characters
	%s has shortest line with %d characters
	The mean lines length in this file is %d
	The median line length in this file is %d
`, name, s.N, s.MaxName, s.Max, s.MinName, s.Min, s.Mean, s.Median)
}

func validLangMap(lm langMap) bool {
	if _, ok := lm[function]; !ok {
		return false
	} else if _, ok = lm[funcName]; !ok {
		return false
	} else if _, ok = lm[funcStart]; !ok {
		return false
	} else if _, ok = lm[funcEnd]; !ok {
		return false
	} else if _, ok = lm[ignoreLine]; !ok {
		return false
	}
	return true
}

func countLineStats(s *stats, currentFunction, line string, lines *[]int) {
	ll := len(line)
	if ll > 1 {
		*lines = append(*lines, ll)
		s.N++
		if ll > s.Max {
			if currentFunction == "" {
				s.MaxName = "(global)"
			} else {
				s.MaxName = currentFunction
			}
			s.Max = ll
		}

		if ll < s.Min || s.Min == 0 {
			if currentFunction == "" {
				s.MinName = "(global)"
			} else {
				s.MinName = currentFunction
			}
			s.Min = ll
		}
		s.Mean += ll
	}
}

func scanFuncs(lm langMap, r io.Reader) (stats, map[string]int) {
	var (
		inFunc          bool
		s               stats
		start           bool
		currentFunction string
		lines           = []int{}
		funcmap         = map[string]int{}
	)

	if !validLangMap(lm) {
		return stats{}, nil
	}

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := strings.TrimRight(scanner.Text(), `\n`)
		if !inFunc {
			if !lm[function].MatchString(line) {
				countLineStats(&s, currentFunction, line, &lines)
				continue
			}

			inFunc = true
		}

		if inFunc && !start {
			groups := lm[funcName].FindStringSubmatch(line)
			if len(groups) > 1 {
				countLineStats(&s, currentFunction, line, &lines)
				currentFunction = groups[1]
			}
		}

		if inFunc && lm[funcStart].MatchString(line) {
			start = true
			continue
		}

		if !lm[ignoreLine].MatchString(line) {
			countLineStats(&s, currentFunction, line, &lines)
		}
		if !inFunc {
			continue
		}

		if lm[funcEnd].MatchString(line) {
			start = false
			inFunc = false
		}

		if start {
			funcmap[currentFunction]++
		}
	}

	sort.Ints(lines)
	s.Mean /= s.N
	s.Median = lines[s.N/2]
	return s, funcmap
}

var langMaps = map[string]langMap{
	"c":  cRE,
	"go": goRE,
}

var extRegexp = regexp.MustCompile(`.+\.(\w+)$`)

func main() {
	flag.Parse()

	fileCount := flag.NArg() - 1
	for i, fileName := range flag.Args() {
		if !extRegexp.MatchString(fileName) {
			fmt.Println("Skipping", fileName)
			continue
		}

		groups := extRegexp.FindStringSubmatch(fileName)
		if len(groups) < 2 {
			fmt.Println("Skipping", fileName)
			continue
		}

		lm := langMaps[groups[1]]
		if lm == nil {
			fmt.Println("Skipping", fileName)
			continue
		}

		file, err := os.Open(fileName)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			fmt.Println("Skipping", fileName)
			continue
		}

		ls, fm := scanFuncs(lm, file)
		file.Close()

		if fm == nil {
			fmt.Println("No results from", fileName)
			continue
		}

		fs := computeStats(fm)
		displayFuncStats(fileName, fs)
		displayLineStats(fileName, ls)

		fmt.Printf("Functions in %s:\n", fileName)
		for k, v := range fm {
			fmt.Printf("\t%s: %d lines\n", k, v)
		}

		if i < fileCount {
			fmt.Printf("\n---\n")
		}
	}
}
