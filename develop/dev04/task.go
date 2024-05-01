package main

import (
	"fmt"
	"slices"
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
	d := []string{"столик", "листок", "Пятак", "пятак", "Листок", "Слиток", "пятка", "тяпка", "слово", "волос"}
	fmt.Println(findAnagram(d))
}

func findAnagram(dict []string) map[string][]string {
	seen := make(map[string]string)
	anagrams := make(map[string][]string)

	for _, word := range dict {
		lowerWord := strings.ToLower(word)
		runeWord := []rune(lowerWord)
		slices.Sort(runeWord)
		sortedWord := string(runeWord)

		if key, exists := seen[sortedWord]; exists {
			if key != lowerWord {
				duplicate := false
				for _, v := range anagrams[key] {
					if v == lowerWord {
						duplicate = true
						break
					}
				}
				if !duplicate {
					anagrams[key] = append(anagrams[key], lowerWord)
				}
			}
		} else {
			seen[sortedWord] = lowerWord
		}
	}

	for key := range anagrams {
		if len(anagrams[key]) <= 1 {
			delete(anagrams, key)
		} else {
			sort.Strings(anagrams[key])
		}
	}

	return anagrams
}
