package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
	"tcp-chat/config"
	"time"
)

func main() {
	tcpAddr, err := net.ResolveTCPAddr("tcp", config.Conf.IP)
	if err != nil {
		log.Fatal(err)
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	go printOutput(conn)
	writeInput(conn)
}

func writeInput(conn *net.TCPConn) {
	fmt.Print("Введите имя пользователя: ")
	reader := bufio.NewReader(os.Stdin)
	username, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	username = strings.Trim(username, "\n\r")
	fmt.Println("Введите текст: ")
	for {
		text, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		text = strings.Trim(text, "\r")
		message := time.Now().Format("02.01 15:04 ") + username + ": " + text
		if n, err := conn.Write([]byte(message)); n == 0 || err != nil {
			log.Fatal(err)
		}
	}
}

func printOutput(conn *net.TCPConn) {
	defer conn.Close()

	for {
		// получение ответа
		buff := make([]byte, 1024)
		n, err := conn.Read(buff)
		if err == io.EOF {
			conn.Close()
			fmt.Println("Connection Closed.")
			os.Exit(0)
		}
		if err != nil {
			log.Fatal(err)
		}

		fmt.Print(string(buff[0:n]))

	}
}
