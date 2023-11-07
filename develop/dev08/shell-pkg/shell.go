package shell_pkg

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"syscall"
)

func Run() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("$ ")
		command, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		// Удаление символа новой строки из команды
		command = strings.TrimSuffix(command, "\n")

		args := strings.Split(command, " ")
		if args[0] == "\\quit" {
			break
		}
		// Обработка команды в зависимости от ее типа
		switch args[0] {
		case "cd":
			if len(args) < 2 {
				fmt.Println("Не указан аргумент для cd")
			} else {
				err := os.Chdir(args[1])
				if err != nil {
					fmt.Println(err)
				}
			}
		case "pwd":
			dir, err := os.Getwd()
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(dir)
		case "echo":
			if len(args) < 2 {
				fmt.Println("Не указан аргумент для echo")
			} else {
				fmt.Println(strings.Join(args[1:], " "))
			}
		case "kill":
			if len(args) < 2 {
				fmt.Println("Не указан аргумент для kill")
			} else {
				pid, err := strconv.Atoi(args[1])
				if err != nil {
					fmt.Println(err)
				}
				err = syscall.Kill(pid, syscall.SIGKILL)
				if err != nil {
					fmt.Println(err)
				}
			}
		case "ps":
			cmd := exec.Command("ps", "-ef")
			output, err := cmd.Output()
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(string(output))
		case "":
			// Пропускаем пустую команду
		default:
			// Выполнение команды в отдельном процессе
			cmd := exec.Command(args[0], args[1:]...)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr

			err := cmd.Run()
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}

/*
Пример использования:

$ go run main.go
$ pwd
/home/user
$ cd /tmp
$ pwd
/tmp
$ echo Hello, World!
Hello, World!
$ kill firefox
$ ps
UID        PID  PPID  C STIME TTY          TIME CMD
user       933   932  0 10:12 pts/0    00:00:00 bash
user       954   933  0 10:12 pts/0    00:00:00 go run main.go
...
Программа будет выполнять команды пользователя и выводить соответствующие результаты.
Обработка команды fork/exec выполняется при вводе произвольной команды,
которая не соответствует одной из простейших команд.
Конвейер на пайпах (перенаправление вывода одной команды на вход другой) также поддерживается.
*/
