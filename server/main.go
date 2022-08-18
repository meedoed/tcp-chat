package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"tcp-chat/config"
)

var (
	connections []net.Conn
)

func main() {

	listener, err := net.Listen("tcp", config.Conf.IP)

	if err != nil {
		log.Fatal(err)
	}

	defer listener.Close()

	fmt.Printf("Server is listening on %s\n\n", config.Conf.IP)
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		// добавление подключения
		connections = append(connections, conn)

		// обработка поключений в горутине
		go handleRequest(conn)
	}

}

// handleRequest - обработка подключения
func handleRequest(conn net.Conn) {
	for {
		// считывание полученных в запросе данных
		input := make([]byte, (1024 * 4))
		n, err := conn.Read(input)
		if n == 0 || err != nil {
			if err == io.EOF {
				removeConn(conn)
				conn.Close()
				return
			}
			log.Println("Read error:", err)
			return
		}
		msg := string(input[0:n])
		fmt.Print(msg)
		broadcast(conn, msg)
	}
}

// removeConn - удаление отключенного пользователя
func removeConn(conn net.Conn) {
	var i int
	for i = range connections {
		if connections[i] == conn {
			break
		}
	}
	connections = append(connections[:i], connections[i+1:]...)
}

// broadcast - отправка сообщения остальным пользователям
func broadcast(conn net.Conn, msg string) {
	for i := range connections {
		if connections[i] != conn {
			n, err := connections[i].Write([]byte(msg))
			if err != nil || n == 0 {
				log.Println(err)
			}
		}
	}
}
