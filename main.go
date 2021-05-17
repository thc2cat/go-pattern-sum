package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

// Purpose :
// get rid of awk/sed/grep for logs calculations/extractions

// History
// v0.1 : functionnal
// v0.2 : -p can print non numerical values
// v0.21 : exit code when no pattern are founds

func main() {

	pat := flag.String("P", ", nrcpt=(?P<value>[0-9]+) ",
		"pattern with numeric value to sum")
	tag := flag.String("t", "",
		"tag to add after printing sum value")
	printFlag := flag.Bool("p", false,
		"only print values")
	statsFlag := flag.Bool("s", false,
		"Show sum count/min/max/avg instead of only sum")
	flag.Parse()

	r, err := regexp.Compile(*pat)
	if err != nil {
		fmt.Printf("Err %s with '%s'", err, *pat)
		os.Exit(1)
	}

	readandprint(r, *tag, *printFlag, *statsFlag)
}

func readandprint(p *regexp.Regexp, tag string, P, S bool) {

	var (
		line   []byte
		length int
		sum    int
		values []string
	)

	s := bufio.NewScanner(os.Stdin)

	for s.Scan() {
		line = s.Bytes()
		length = len(line)
		if length < 1 {
			continue
		}
		results := reSubMatchMap(p, s.Text())
		if results == nil {
			continue
		}

		values = append(values, results["value"])
	}

	if len(values) < 1 {
		fmt.Fprintf(os.Stderr, "pattern not found !\n")
		os.Exit(1)
	}
	// ref https://yourbasic.org/golang/max-min-int-uint/
	const UintSize = 32 << (^uint(0) >> 32 & 1) // 32 or 64
	min := 1<<(UintSize-1) - 1
	max := 0

	for _, valS := range values {
		val, err := strconv.Atoi(valS)
		if err != nil {
			continue
		}

		if P && !S {
			fmt.Println(val)
		}
		if max <= val {
			max = val
		}
		if min >= val {
			min = val
		}
		sum += val
	}

	if !S && !P {
		fmt.Printf("%d\t%s\n", sum, tag)
	}

	if S && !P {
		fmt.Printf("%d %d/%d/%d/%d\t%s\n", sum, len(values), min, max, int(sum/len(values)), tag)
	}

}

func reSubMatchMap(r *regexp.Regexp, str string) map[string]string {
	match := r.FindStringSubmatch(str)
	if len(match) == 0 {
		return nil
	}
	subMatchMap := make(map[string]string)
	for i, name := range r.SubexpNames() {
		if i != 0 {
			subMatchMap[name] = match[i]
		}
	}
	return subMatchMap
}
