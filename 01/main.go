package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
		fmt.Print("Введите числа через пробел: ")

		// получим строку с входными данными с утсройства ввода
			str := readString()

		// Распарсим строку в []int64
			arrInts, err := strToArrInt(str)

		// При отстутствии ошибок произведем сортировку
		// при ошибках парсинга введем собщение о не валидности ввходных данных
			if err == nil {
				// Произведем сортировку по возрастанию полученного массива чисел
					arrInts = sortInt64Arr(arrInts)

				// Корвертируем []int64 в string
					s := fmt.Sprint(arrInts)

				// Обрежем "лишние" символы в начале и конце строки вывода
					s = strings.Trim(s, "[]")

				// Ввыведем отсортированные входдные данные
				fmt.Printf("Сортировка закончена: %s", s)
			} else {
				fmt.Printf("Ошибка валидации входных данных, ERROR: %s", err)
			}
}

/**
* Читает стандартный ввод как строку.
* @return str {string} полученная строка с устройства ввода.
*/
func readString() string {
	in := bufio.NewReader(os.Stdin)
	str, err := in.ReadString('\n')
	if err != nil {
		fmt.Printf("Ошибка ввода: %s\n", err)
	}
	return str
}

/**
* Производит разбор строки по словам в []int64.
* @param input {string} входная строка для разбора.
* @return arr {[]int64} рультирующий "массив".
* @return err {error} ошибка разбора строки.
*/
func strToArrInt(input string) (arr []int64, err error) {
	// получим строку из источника ввода
		scanner := bufio.NewScanner(strings.NewReader(input))

	// Создаим  функцию разделения, обернув функцию bufio.ScanWords.
	split := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		// Инициализируем временную переменную для хранения результата конвертации
		// string to int64
			var num int64

		advance, token, err = bufio.ScanWords(data, atEOF)
		if err == nil && token != nil {
			// Конвертируем полученный token в int64
				num, err = strconv.ParseInt(string(token), 10, 64)

			// Если нет ошибки конвертации добавим число в резльтатирующий массив
				if err == nil {
					arr = append(arr, num)
				}
		}

		return
	}

	// Установим функцию разделения для операции сканирования.
		scanner.Split(split)

	// Скаинирование входной строки
		for scanner.Scan() {
		}

	// Полижим ошибку сканирования в выходную переменную
		err = scanner.Err()

	return
}

/**
* Производит сортировку []int64 алгоритмом вставками по возрастанию.
* @param arr {[]int64} сортируемый "массив".
* @return arr {[]int64} отсортированный "массив" по возрастанию.
*/
func sortInt64Arr(arr []int64) []int64 {
	for i := 1; i < len(arr); i++ {
		for j := i; j > 0 && arr[j-1] > arr[j]; j-- {
			t := arr[j]
			arr[j] = arr[j-1]
			arr[j-1] = t
		}
	}

	return arr
}
