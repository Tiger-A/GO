package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

type Text struct {
	Content string
}

func (t *Text) textModifier() {

	t.Content = strings.ReplaceAll(t.Content, "  ", " ")
	for strings.Contains(t.Content, "  ") {
		t.Content = strings.ReplaceAll(t.Content, "  ", " ")
	}
	runes := []rune(t.Content)
	result := []rune{}
	i := 0

	for i < len(runes) {
		if runes[i] == '-' {
			if i > 0 && i < len(runes)-1 {
				// Меняем местами от -
				resultLen := len(result)
				if resultLen > 0 { // не пуст ли?
					// del последний символ из result
					lastRune := result[resultLen-1]
					result = result[:resultLen-1]
					// Добавляем символ справа от '-' сначала, затем символ слева
					result = append(result, runes[i+1], lastRune)
				} else {
					// if result пустой, app символы справа и слева от -
					result = append(result, runes[i+1], runes[i-1])
				}
				// Пропустить символ справа от '-' и сам '-'
				i += 2
			} else {
				// Если '-' в начале или конце строки,  пропустить его
				i++
			}
		} else {
			// Добавляем текущий символ в результат
			result = append(result, runes[i])
			i++
		}
	}

	t.Content = string(result)

	// Заменяем знак плюс (+) на восклицательный знак (!)
	t.Content = strings.ReplaceAll(t.Content, "+", "!")

	// Считаем сумму цифр и удаляем их из текста
	sum := 0
	runeText := []rune(t.Content)
	result = []rune{}
	for _, r := range runeText {
		if unicode.IsDigit(r) {
			sum += int(r - '0')
		} else {
			result = append(result, r)
		}
	}
	t.Content = string(result)

	if sum > 0 {
		t.Content = fmt.Sprintf("%s %d", t.Content, sum)
	}

	// Выводим результат
	fmt.Println(t.Content)
}

//   надо было один раз перегнать  все в руны и обрабатывать

func main() {
	text := &Text{}
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Введите строку:")

	for scanner.Scan() {
		text.Content = scanner.Text()
		text.textModifier()
		fmt.Println("Введите строку:")
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Ошибка чтения ввода:", err)
	}
}

//    https://github.com/Tiger-A/GO.git
