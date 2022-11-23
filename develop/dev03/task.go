package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

/*
=== Утилита sort ===

Отсортировать строки (man sort)
Основное

Поддержать ключи

-k — указание колонки для сортировки
-n — сортировать по числовому значению
-r — сортировать в обратном порядке
-u — не выводить повторяющиеся строки

Дополнительное

Поддержать ключи

-M — сортировать по названию месяца
-b — игнорировать хвостовые пробелы
-c — проверять отсортированы ли данные
-h — сортировать по числовому значению с учётом суффиксов

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func permutateArgs(args []string) int {
	args = args[1:]
	optind := 0

	for i := range args {
		if args[i][0] == '-' {
			tmp := args[i]
			args[i] = args[optind]
			args[optind] = tmp
			optind++
		}
	}

	return optind + 1
}

func main() {
	var numOfSortColumn int
	var isNumericSort, isReverseSort, notPrintRepetitions bool
	flag.IntVar(&numOfSortColumn, "k", 1, "указание колонки для сортировки")
	flag.BoolVar(&isNumericSort, "n", false, "сортировать по числовому значению")
	flag.BoolVar(&isReverseSort, "r", false, "сортировать в обратном порядке")
	flag.BoolVar(&notPrintRepetitions, "u", false, "не выводить повторяющиеся строки")
	permutateArgs(os.Args)
	flag.Parse()
	numOfSortColumn -= 1
	fileName := flag.Arg(0)

	sortUtil(fileName, numOfSortColumn, isNumericSort, isReverseSort, notPrintRepetitions)

}

func sortUtil(fileName string, numOfSortColumn int, isNumericSort bool, isReverseSort bool, notPrintRepetitions bool) error {
	sortingObj := SortingObj{
		rowsSlices: make([][]string, 0, 1024),
		Sortoptions: Sortoptions{
			numOfSortColumn,
			isNumericSort,
			isReverseSort,
			notPrintRepetitions,
		},
	}

	err := readFileAndGetRowsSlice(fileName, &sortingObj)

	if err != nil {
		return err
	}

	for _, v := range sortingObj.rowsSlices {
		fmt.Println(strings.Join(v, " "))
	}

	sort.Sort(sortingObj)

	fmt.Println()

	for _, v := range sortingObj.rowsSlices {
		fmt.Println(strings.Join(v, " "))
	}

	return nil
}

type SortingObj struct {
	rowsSlices [][]string
	Sortoptions
}

func (s SortingObj) getSortingColumn(i, j int) (col int, skipElement1 bool, skipElement2 bool) {
	col = 0

	if s.Sortoptions.numOfSortColumn > len(s.rowsSlices[i])-1 {
		skipElement1 = true
		return
	}

	if s.Sortoptions.numOfSortColumn > len(s.rowsSlices[j])-1 {
		skipElement2 = true
		return
	}

	col = s.Sortoptions.numOfSortColumn

	return
}

func (s SortingObj) reverseResultIfReverseSort(b bool) bool {
	if s.Sortoptions.isReverseSort {
		return !b
	}
	return b
}

// Метод
func (s SortingObj) skipSolution(s1, s2 bool) bool {

	if s1 {
		return s.reverseResultIfReverseSort(true)
	} else if s2 {
		return s.reverseResultIfReverseSort(false)
	}

	return false

}

// Метод сортировки
func (s SortingObj) Sort() {
	sort.SliceStable(s.rowsSlices, func(i, j int) bool {
		col, skip1, skip2 := s.getSortingColumn(i, j)
		if skip1 || skip2 {
			return s.skipSolution(skip1, skip2)
		}

		

		return s.rowsSlices[i][col] < s.rowsSlices[j][col]
	})
}

type Sortoptions struct {
	numOfSortColumn     int
	isNumericSort       bool
	isReverseSort       bool // передавать в defaultSort
	notPrintRepetitions bool // при чтении файла
}

func (o SortingObj) Len() int { return len(o.rowsSlices) }
func (o SortingObj) Swap(i, j int) {
	o.rowsSlices[i], o.rowsSlices[j] = o.rowsSlices[j], o.rowsSlices[i]
}
func (o SortingObj) Less(i, j int) bool {

	c1 := o.numOfSortColumn - 1
	c2 := c1
	// fmt.Println("fsdfds", c1, c2, o.numOfSortColumn)
	if c1 > len(o.rowsSlices[i])-1 {
		c1 = 0
	}

	if c2 > len(o.rowsSlices[j])-1 {
		// Если индекс колонки больше чем индексов ,то сортируем по первому
		c2 = 0
	}
	return o.rowsSlices[i][c1] < o.rowsSlices[j][c2]
}

func readFileAndGetRowsSlice(fileName string, sortingObj *SortingObj) error {

	if fileName == "" {
		fmt.Println("Argument not passed: file name")
		return errors.New("Argument not passed: file name")
	}

	repetitions := map[string]struct{}{}
	f, err := os.Open(fileName)
	defer f.Close()
	if err != nil {
		return err
	}

	sc := bufio.NewScanner(f)
	for sc.Scan() {
		txt := sc.Text()
		row := strings.Split(txt, " ")
		if sortingObj.notPrintRepetitions { // Если есть повторения ,то добавляем только 1 раз
			if _, ok := repetitions[txt]; !ok {
				sortingObj.rowsSlices = append(sortingObj.rowsSlices, row)
				repetitions[txt] = struct{}{}
			}
		} else {
			sortingObj.rowsSlices = append(sortingObj.rowsSlices, row)
		}
	}

	if err := sc.Err(); err != nil {
		log.Fatalf("scan file error: %v", err)
		return err
	}

	return nil
}
