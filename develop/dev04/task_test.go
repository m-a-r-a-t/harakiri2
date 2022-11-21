package main

import (
	"strings"
	"testing"
)

func TestGetSetAnagrams(t *testing.T) {

	t.Run("Проверка правильности данных и их порядка", func(t *testing.T) {

		expectAnagrams := map[string][]string{
			"тяпка":  {"пятак", "пятка", "тяпка"},
			"листок": {"листок", "слиток", "столик"},
		}
		anagrams := getSetAnagrams([]string{"тяпка", "пятак", "пятка", "пятка", "листок", "слиток", "столик", "репка"})

		for k, gotSlice := range anagrams {
			if expectedSlice, ok := expectAnagrams[k]; ok {
				if len(expectedSlice) != len(gotSlice) {
					t.Errorf("Длина не соответствует, ожидалось :%v\nполучено:%v\n", expectedSlice, gotSlice)
				}
				for i := 0; i < len(gotSlice); i++ {
					if gotSlice[i] != expectedSlice[i] {
						t.Errorf("Не соответствует, ожидалось:%s\nполучено:%s\n", expectedSlice[i], gotSlice[i])
					}
				}
			} else {
				t.Errorf("Ключа %s нет в  expectAnagrams", k)
			}
		}

	})

	t.Run("Проверка уникальности", func(t *testing.T) {

		anagrams := getSetAnagrams([]string{"тяпка", "пятка", "пятка", "пятак", "пятка", "пятка", "листок", "слиток", "столик"})
		uniqueWordsMap := map[string]struct{}{}
		for _, gotSlice := range anagrams {
			for _, v := range gotSlice {
				if _, ok := uniqueWordsMap[v]; !ok {
					uniqueWordsMap[v] = struct{}{}
				} else {
					t.Errorf("Слово '%s' не уникально", v)
				}
			}
		}

	})

	t.Run("Проверка всех слов в нижнем регистре", func(t *testing.T) {

		anagrams := getSetAnagrams([]string{"ТЯПКА", "пЯтка", "пяТКа", "пЯтак", "ПяткА", "пяТка", "Листок", "слИток", "стоЛИК"})
		for _, gotSlice := range anagrams {
			for _, v := range gotSlice {
				if v != strings.ToLower(v) {
					t.Errorf("Слово не в нижнем регистре\nожидалось:%s\nполучено:%s", strings.ToLower(v), v)
				}
			}
		}

	})

}
