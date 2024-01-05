package main

import (
	"bufio"
	"fmt"
	"log"
	"net"

	"github.com/google/uuid"
)

type Vec2 struct {
	X, Y float64
}

type UserModel struct {
	Position *Vec2
	Name     string
	Conn     net.Conn
}

var users = make(map[string]UserModel)

func StartServer() {
	l, err := net.Listen("tcp", "localhost:9090")
	if err != nil {
		return
	}

	fmt.Println("Server Started!")

	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Println(err)
			return
		}

		id := uuid.New().String()
		newModel := UserModel{Position: &Vec2{0, 0}, Name: "", Conn: conn}
		users[id] = newModel

		fmt.Println("User:", conn.LocalAddr().String(), "connected!")

		go handleUserConnection(newModel, id)
	}
}

func handleUserConnection(model UserModel, id string) {
	defer model.Conn.Close()
	for {
		data, _, err := bufio.NewReader(model.Conn).ReadLine()
		if err != nil {
			log.Println(model.Conn.LocalAddr().String(), "disconnected!")
			delete(users, id)
			return
		}

		strMessage := string(data)

		if strMessage == "1" {
			users[id].Position.X += 1

		} else if strMessage == "2" {
			users[id].Position.X -= 1
		}

		for _, user := range users {
			fmt.Println(user.Position)
		}
	}
}

func main() {
	StartServer()
}
