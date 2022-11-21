package main

import (
	"fmt"
	"math/big"
	"sort"
	"strings"
)

/*
=== Поиск анаграмм по словарю ===

Напишите функцию поиска всех множеств анаграмм по словарю.
Например:
'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
'листок', 'слиток' и 'столик' - другому.

Входные данные для функции: ссылка на массив - каждый элемент которого - слово на русском языке в кодировке utf8.
Выходные данные: Ссылка на мапу множеств анаграмм.
Ключ - первое встретившееся в словаре слово из множества
Значение - ссылка на массив, каждый элемент которого, слово из множества. Массив должен быть отсортирован по возрастанию.
Множества из одного элемента не должны попасть в результат.
Все слова должны быть приведены к нижнему регистру.
В результате каждое слово должно встречаться только один раз.

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	arr := []string{"тяпка", "пятак", "пЯтка", "пятка", "листок", "слиток", "столик", "", "", "fsdfds"}
	result := getSetAnagrams(arr)

	for k, v := range result {
		fmt.Printf("Key:%s Values: %v \n", k, v)
	}

}

func getSetAnagrams(wordsSlice []string) map[string][]string {
	wordsSet := map[string]struct{}{}   // ключ: слово из слайса wordSlice
	result := map[string][]string{}     // ключ: первое слово анаграммы, значение: множество слов данной анаграммы
	firstWordMap := map[string]string{} // ключ число big.Int, значение слово
	for _, word := range wordsSlice {
		if word == "" {
			continue
		}

		w := strings.ToLower(word)

		if _, ok := wordsSet[w]; !ok {
			wordsSet[w] = struct{}{}
			wRunes := []rune(w)
			lettersMultiplicationNum := big.NewInt(int64(wRunes[0]))

			for i := 1; i < len(wRunes); i++ { // подсчет уникального номера анаграммы
				runeInBigInt := big.NewInt(int64(wRunes[i]))
				lettersMultiplicationNum = lettersMultiplicationNum.Mul(lettersMultiplicationNum, runeInBigInt)
			}

			numOfAnagram := lettersMultiplicationNum.String() // уникальный номер анаграмыы

			if resultKeyName, ok := firstWordMap[numOfAnagram]; !ok { // если уникальный номер анаграммы еще не был добавлен ,то она добавляется в множество как ключ ,а значение в себе хранит первое слово данной анаграмы
				firstWordMap[numOfAnagram] = w
				result[w] = []string{w}
			} else {
				result[resultKeyName] = append(result[resultKeyName], w) // добавляем в слайс результатов
			}

		}

	}

	// убираем множества где только 1 результат и сортируем где больше 1
	for k, v := range result {
		if len(v) <= 1 {
			delete(result, k)
		} else {
			sort.Strings(v)
		}
	}

	return result
}
