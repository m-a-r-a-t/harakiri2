package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
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
	numOfSortColumn--
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

	sortingObj.Sort()

	fmt.Println("=======================")

	for _, v := range sortingObj.rowsSlices {
		fmt.Println(strings.Join(v, " "))
	}

	return nil
}

// SortingObj структура содержащая опции и сами строки
type SortingObj struct {
	rowsSlices [][]string
	Sortoptions
}

func (s SortingObj) getSortingColumn(i, j int) (skipElement1 bool, skipElement2 bool) {
	if s.Sortoptions.numOfSortColumn > len(s.rowsSlices[i])-1 {
		skipElement1 = true
	}

	if s.Sortoptions.numOfSortColumn > len(s.rowsSlices[j])-1 {
		skipElement2 = true
	}

	return
}

func (s SortingObj) reverseResultIfReverseSort(b bool) bool {
	if s.Sortoptions.isReverseSort {
		return !b
	}
	return b
}

// Метод
func (s SortingObj) skipSolution(i, j int, s1, s2 bool) bool {
	if s1 && s2 {
		s.defaultSort(0, i, j)
	} else if s1 {
		return s.reverseResultIfReverseSort(true)
	} else if s2 {
		return s.reverseResultIfReverseSort(false)
	}

	return false

}

func (s SortingObj) isNumeric(col, i, j int) (skipElement1 bool, skipElement2 bool) {
	_, err := strconv.ParseFloat(s.rowsSlices[i][col], 64)
	if err != nil {
		skipElement1 = true
	}
	_, err = strconv.ParseFloat(s.rowsSlices[j][col], 64)
	if err != nil {
		skipElement2 = true
	}

	return
}

func (s SortingObj) defaultSort(col, i, j int) bool {
	if s.Sortoptions.isReverseSort {
		return s.rowsSlices[i][col] > s.rowsSlices[j][col]
	}

	return s.rowsSlices[i][col] < s.rowsSlices[j][col]
}

func (s SortingObj) numericSort(col, i, j int) bool {
	n1, _ := strconv.ParseFloat(s.rowsSlices[i][col], 64)
	n2, _ := strconv.ParseFloat(s.rowsSlices[j][col], 64)

	if s.Sortoptions.isReverseSort {
		return n1 > n2
	}

	return n1 < n2
}

// Sort Метод сортировки
func (s SortingObj) Sort() {
	sort.SliceStable(s.rowsSlices, func(i, j int) bool {
		skip1, skip2 := s.getSortingColumn(i, j)
		if skip1 || skip2 {
			return s.skipSolution(i, j, skip1, skip2)
		}

		if s.Sortoptions.isNumericSort {
			skip1, skip2 := s.getSortingColumn(i, j)
			if skip1 || skip2 {
				return s.skipSolution(i, j, skip1, skip2)
			}
			return s.numericSort(s.numOfSortColumn, i, j)
		}

		return s.defaultSort(s.numOfSortColumn, i, j)
	})
}

// Sortoptions ключи передаваемые в программу
type Sortoptions struct {
	numOfSortColumn     int
	isNumericSort       bool
	isReverseSort       bool // передавать в defaultSort
	notPrintRepetitions bool // при чтении файла
}

// функция чтения файла построчно
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
