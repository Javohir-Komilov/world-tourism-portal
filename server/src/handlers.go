// server/src/handlers.go

package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()
	log.Println("Новое соединение:", conn.RemoteAddr())

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		request := scanner.Text()
		log.Printf("Получен запрос: %s", request)

		if request == "поиск" {
			response := "Доступные туры: Тур1, Тур2, Тур3"
			fmt.Fprintln(conn, response)
		} else if request == "выход" {
			log.Println("Клиент отключился")
			break
		} else {
			fmt.Fprintln(conn, "Неизвестная команда")
		}
	}

	if err := scanner.Err(); err != nil {
		log.Printf("Ошибка чтения: %v", err)
	}
}
