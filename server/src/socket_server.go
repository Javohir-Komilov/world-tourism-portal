package main

import (
	"log"
	"net"
	"world-tourism-portal/queries"
)

func startSocketServer(db *queries.DB, address string) {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Failed to start socket server: %v", err)
	}
	defer listener.Close()

	log.Printf("Socket server started on %s", address)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Failed to accept connection: %v", err)
			continue
		}

		go handleSocketConnection(db, conn)
	}
}

func handleSocketConnection(db *queries.DB, conn net.Conn) {
	defer conn.Close()
	// Реализуйте логику общения с клиентом через сокет
	// Например, обработку запросов на создание турпакетов, бронирование отелей и так далее

	log.Printf("New connection established from %s", conn.RemoteAddr())

	// Пример ответа
	conn.Write([]byte("Welcome to the World Tourism Portal socket server!\n"))
}
