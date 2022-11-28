package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

/*
=== Утилита cut ===

Принимает STDIN, разбивает по разделителю (TAB) на колонки, выводит запрошенные

Поддержать флаги:
-f - "fields" - выбрать поля (колонки)
-d - "delimiter" - использовать другой разделитель
-s - "separated" - только строки с разделителем

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	var fields, delimiter string
	var separated bool
	flag.StringVar(&fields, "f", "", "выбор колонок. Поумолчанию выводятся все")
	flag.StringVar(&delimiter, "d", "\t", "использовать другой разделитель. Поумолчанию TAB")
	flag.BoolVar(&separated, "s", false, "только строки с разделителем")
	flag.Parse()

	fieldsInNumArr, err := parseFields(fields)
	if err != nil {
		errStr := fmt.Sprintf("error: %v.\n", err.Error())
		_, _ = os.Stderr.WriteString(errStr)
		os.Exit(1)
	}

	fmt.Println("input text:")
	reader := bufio.NewReader(os.Stdin)

	var strRows []string
	for {
		// read line from stdin using newline as separator
		line, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		// if line is empty, break the loop
		if len(strings.TrimSpace(line)) == 0 {
			break
		}

		//append the line to a slice
		strRows = append(strRows, line)
	}

	cutUtil := NewCutUtil(strRows, fieldsInNumArr, delimiter, separated)
	// var output bytes.Buffer
	cutUtil.cut(os.Stdout)
}

func parseFields(fieldsStr string) ([]int, error) {
	if fieldsStr == "" {
		return nil, errors.New("вы должны задать список полей")
	}

	var fields []int
	var set = make(map[int]struct{})
	diapazones := strings.Split(fieldsStr, ",")
	for _, zone := range diapazones {
		c, err := strconv.Atoi(zone)
		if err != nil {
			dz := strings.Split(zone, "-")
			errTxt := "неправильный диапазон поля"
			// fmt.Println(zone, dz)
			if len(dz) != 2 {
				return []int{}, errors.New(errTxt)
			}

			num1, err1 := strconv.Atoi(dz[0])
			num2, err2 := strconv.Atoi(dz[1])

			if err1 != nil || err2 != nil {
				return nil, errors.New(errTxt)
			}

			if num2 < num1 {
				return nil, errors.New("неверный уменьшающийся диапазон")
			}

			for i := num1 - 1; i < num2; i++ {
				if _, ok := set[i]; !ok {
					fields = append(fields, i)
					set[i] = struct{}{}
				}
			}

		} else {
			if _, ok := set[c-1]; !ok {
				fields = append(fields, c-1)
				set[c-1] = struct{}{}
			}
		}

	}

	sort.Ints(fields)

	return fields, nil
}

type CutUtil struct {
	rowSlice []string
	opts
}

type opts struct {
	fields    []int
	delimiter string
	separated bool
}

func NewCutUtil(rows []string, fields []int, delimiter string, separated bool) CutUtil {
	return CutUtil{rows, opts{fields, delimiter, separated}}
}

func (c CutUtil) cut(w io.Writer) {
	for _, v := range c.rowSlice {
		result := []string{}
		splitedStr := strings.Split(v, c.delimiter)

		if len(splitedStr) <= 1 {
			if !c.separated {
				fmt.Fprintln(w, strings.Trim(v, "\n"))
			}
			continue
		}

		for _, v := range c.fields {
			if v > len(splitedStr)-1 {
				break
			}

			result = append(result, splitedStr[v])
		}

		// fmt.Println(strings.Join(result, c.delimiter))
		fmt.Fprintln(w, strings.TrimSpace(strings.Join(result, c.delimiter)))

	}
}
