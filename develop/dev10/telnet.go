package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

/*
=== Утилита telnet ===

Реализовать примитивный telnet клиент:
Примеры вызовов:
go-telnet --timeout=10s host port go-telnet mysite.ru 8080 go-telnet --timeout=3s 1.1.1.1 123

Программа должна подключаться к указанному хосту (ip или доменное имя) и порту по протоколу TCP.
После подключения STDIN программы должен записываться в сокет, а данные полученные и сокета должны выводиться в STDOUT
Опционально в программу можно передать таймаут на подключение к серверу (через аргумент --timeout, по умолчанию 10s).

При нажатии Ctrl+D программа должна закрывать сокет и завершаться.
Если сокет закрывается со стороны сервера, программа должна также завершаться.
При подключении к несуществующему сервер, программа должна завершаться через timeout.
*/

func main() {
	// Разбор аргументов командной строки
	timeout := flag.Duration("timeout", 10*time.Second, "timeout for connection")
	flag.Parse()
	args := flag.Args()
	if len(args) != 2 {
		fmt.Println("Usage: go-telnet --timeout=10s host port")
		os.Exit(1)
	}
	host := args[0]
	port := args[1]

	// Установка обработчика сигнала SIGINT для закрытия сокета при нажатии Ctrl+C
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-sigCh
		fmt.Println("Interrupt connection")
		os.Exit(0)
	}()

	// Установка таймаута для подключения
	conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%s", host, port), *timeout)
	if err != nil {
		log.Fatal("Connection error:", err)
	}

	// Закрытие сокета при завершении работы программы
	defer conn.Close()

	// Чтение данных из STDIN и запись их в сокет
	go func() {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			fmt.Fprintf(conn, "%s\n", scanner.Text())
		}
	}()

	// Чтение данных из сокета и вывод их в STDOUT
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal("Error reading from socket:", err)
	}
}

/*
Примеры использования:
запустить в отдельном терминале:
go run server/serv.go
запустить в отдельном терминале:
go run telnet.go --timeout=3s 127.0.0.1 5555
прерывание - Ctrl + C
Пример отправки данных в Stdin и чтение ответа из сокета onn
telnet.go
1
response: 1
serv.go
2023/11/08 12:34:23 received: 1
*/
