package main

import (
	"os"
	"log"
	"net"
	"strconv"
)

//funcion de control de fallos
func handler(conn net.Conn){
	defer func(){
		conn.Close()
		recover()
	}()
	p := new(partida)
	p.SetPlayer(conn)
	p.Start();
}

//buscador de jugadores
func SocketServer(port int) { 
	listen, err := net.Listen("tcp4", ":"+strconv.Itoa(port))
	defer listen.Close()

	if err != nil {
		log.Fatalf("Socket listen port %d failed,%s", port, err)
		os.Exit(1)
	}

	log.Printf("Begin listen port: %d", port)

	for {
		conn, err := listen.Accept()
		log.Println("--> Se ha conectado el cliente",conn.RemoteAddr())
		if err != nil {
			log.Fatalln(err)
			continue
		}
		go handler(conn) //proceso multicast
	}
}

func main() {
	clearScreen()
	port := 3333
	SocketServer(port)
}
