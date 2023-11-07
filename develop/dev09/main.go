package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

/*
=== Утилита wget ===

Реализовать утилиту wget с возможностью скачивать сайты целиком

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	reader := bufio.NewReader(os.Stdin) // Введите URL сайта, который хотите скачать
	url, err := reader.ReadString('\n')
	url = strings.Trim(url, "\n")

	// Отправляем GET-запрос
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Не удалось выполнить запрос: %v\n", err)
		return
	}
	defer resp.Body.Close()

	// Проверяем статус ответа
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Сайт вернул неправильный статус: %v\n", resp.StatusCode)
		return
	}

	// Получаем имя файла из URL
	tokens := strings.Split(url, "/")
	fileName := tokens[len(tokens)-1]

	// Создаем файл для записи
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Printf("Не удалось создать файл: %v\n", err)
		return
	}
	defer file.Close()

	// Копируем содержимое ответа в файл
	_, err = io.Copy(file, resp.Body)
	if err != nil {
		fmt.Printf("Не удалось записать данные в файл: %v\n", err)
		return
	}

	fmt.Printf("Сайт успешно скачан и сохранен в файл: %s\n", fileName)
}

/*
В этом примере используется пакет net/http для выполнения GET-запроса к указанному URL сайта.
Затем полученное содержимое записывается в файл с помощью функции io.Copy.
Имя файла определяется путем разделения URL на отдельные токены и выбора последнего токена в качестве имени файла.
*/
