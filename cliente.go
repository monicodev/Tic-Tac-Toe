package main

import (
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"
	"bufio"
	"os"
)

var Tabla [3][3]string

func esperarTurnoPC(conn net.Conn){
		buf := make([]byte, 1024)
		r := bufio.NewReader(conn)
		fmt.Println("\nEsperando jugada del servidor...")
		n, _ := r.Read(buf)

		switch string(buf[:n]) {
			case "1":
				log.Println("PC jugo en 1")
				Tabla[0][0] = "O";
				break
			case "2":
				log.Println("PC jugo en 2")
				Tabla[0][1] = "O";
				break
			case "3":
				log.Println("PC jugo en 3")
				Tabla[0][2] = "O";
				break
			case "4":
				log.Println("PC jugo en 4")
				Tabla[1][0] = "O";
				break
			case "5":
				log.Println("PC jugo en 5")
				Tabla[1][1] = "O";
				break
			case "6":
				log.Println("PC jugo en 6")
				Tabla[1][2] = "O";
				break
			case "7":
				log.Println("PC jugo en 7")
				Tabla[2][0] = "O";
				break
			case "8":
				log.Println("PC jugo en 8")
				Tabla[2][1] = "O";
				break
			case "9":
				log.Println("PC jugo en 9")
				Tabla[2][2] = "O";
				break
			default:
				log.Println("PC NO jugo su turno (?)")
			}
}

func send(conn net.Conn, a string){
	conn.Write([]byte(a))
}

func SocketClient(ip string, port int) {
	addr := strings.Join([]string{ip, strconv.Itoa(port)}, ":")
	conn, err := net.Dial("tcp", addr)
	defer func() {
		conn.Close()
		recover()
	}()
	if err != nil {
		log.Fatalln(err)
		panic("error de conexion")
	}
	
	iniTabla()
	for {
		iniciarPartida(conn)
		printTabla()
		if (buscarGanador("E")) {fmt.Println("\n Empataste"); return}
		if (buscarGanador("X")) {
			fmt.Println("\n\n     __  ()_                          _/)  __ ");
			fmt.Println("    (_ ) ( '>                        <' ) / _)");
			fmt.Println("      ) )/_)=| FELICIDADES, GANASTE |=(_)/ (");
			fmt.Println("      (_(_ )_|                      |_(_ )_)");
			return
		}
		esperarTurnoPC(conn)
		clearScreen()
		printTabla()
		if (buscarGanador("O")) {fmt.Println("\n Perdiste"); return}
		if (buscarGanador("E")) {fmt.Println("\n Empataste"); return}
	}
}

func buscarGanador(s string)  bool{
	switch s {
	case "E": ///EMPATE
		if (	(Tabla[0][0] == "1") || ( Tabla[0][1] == "2") || (Tabla[0][2] == "3")  ||	//primera fila
				(Tabla[1][0] == "4") || ( Tabla[1][1] == "5") || (Tabla[1][2] == "6")  ||	//segunda fila
				(Tabla[2][0] == "7") || ( Tabla[2][1] == "8") || (Tabla[2][2] == "9") ) {	//segunda diagonal
				return false
		}
		return true;
		break
	default: //GANADOR
		if (	((Tabla[0][0] == s) && ( Tabla[0][1] == s) && (Tabla[0][2] == s) ) ||	//primera fila
				((Tabla[1][0] == s) && ( Tabla[1][1] == s) && (Tabla[1][2] == s) ) ||	//segunda fila
				((Tabla[2][0] == s) && ( Tabla[2][1] == s) && (Tabla[2][2] == s) ) ||	//tercera fila
				((Tabla[0][0] == s) && ( Tabla[1][0] == s) && (Tabla[2][0] == s) ) ||	//primera columna
				((Tabla[0][1] == s) && ( Tabla[1][1] == s) && (Tabla[2][1] == s) ) ||	//segunda columna
				((Tabla[0][2] == s) && ( Tabla[1][2] == s) && (Tabla[2][2] == s) ) ||	//tercera columna
				((Tabla[0][0] == s) && ( Tabla[1][1] == s) && (Tabla[2][2] == s) ) ||	//primera diagonal
				((Tabla[0][2] == s) && ( Tabla[1][1] == s) && (Tabla[2][0] == s) ) ){	//segunda diagonal
					return true;
		}
		return false;
	}
	return false
}

func printTabla(){
	fmt.Println("\n\n        ",Tabla[0][0],"|", Tabla[0][1] ,"|", Tabla[0][2])
	fmt.Println("         ─","+","─","+","─")
	fmt.Println("        ",Tabla[1][0], "|", Tabla[1][1] ,"|", Tabla[1][2])
	fmt.Println("         ─","+","─","+","─")
	fmt.Println("        ",Tabla[2][0], "|", Tabla[2][1] ,"|", Tabla[2][2])
}

func iniciarPartida(conn net.Conn){
	clearScreen()

	var scanner = bufio.NewScanner(os.Stdin)
	var input string
	for {
		printTabla()
		fmt.Println("\nInserte la posicion [1~9] en que desea jugar:")
		scanner.Scan()
		input = scanner.Text()
		clearScreen()
		switch input {
			case "1":	if (Tabla[0][0] == "1") {Tabla[0][0] = "X"; send(conn,input); return  } else {fmt.Println("Posicion ocupada...\n")}
			case "2":	if (Tabla[0][1] == "2") {Tabla[0][1] = "X"; send(conn,input); return  } else {fmt.Println("Posicion ocupada...\n")}
			case "3":	if (Tabla[0][2] == "3") {Tabla[0][2] = "X"; send(conn,input); return  } else {fmt.Println("Posicion ocupada...\n")}
			case "4":	if (Tabla[1][0] == "4") {Tabla[1][0] = "X"; send(conn,input); return  } else {fmt.Println("Posicion ocupada...\n")}
			case "5":	if (Tabla[1][1] == "5") {Tabla[1][1] = "X"; send(conn,input); return  } else {fmt.Println("Posicion ocupada...\n")}
			case "6":	if (Tabla[1][2] == "6") {Tabla[1][2] = "X"; send(conn,input); return  } else {fmt.Println("Posicion ocupada...\n")}
			case "7":	if (Tabla[2][0] == "7") {Tabla[2][0] = "X"; send(conn,input); return  } else {fmt.Println("Posicion ocupada...\n")}
			case "8":	if (Tabla[2][1] == "8") {Tabla[2][1] = "X"; send(conn,input); return  } else {fmt.Println("Posicion ocupada...\n")}
			case "9":	if (Tabla[2][2] == "9") {Tabla[2][2] = "X"; send(conn,input); return  } else {fmt.Println("Posicion ocupada...\n")}
			default:
				fmt.Println("Posicion no valida...\n")
		}
	}
}

func iniTabla(){
	for i := 0 ; i<3; i++ {
		for j := 0 ; j<3; j++ {
			Tabla[i][j] = strconv.Itoa(((i*3)+j+1))
		}
	}
}

func main() {
	clearScreen()
	fmt.Println("")
	menu()
}
 func comoJugar(){
	clearScreen()
	fmt.Println("Para jugar, es sencillo, lo unico de debes hacer es escribir el numero corespondiente a la casilla\nen la que usted jugara...\n")
 }
func menu()  {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		
		fmt.Println("\n			Bienvenido a 3 en Linea")
		fmt.Println("			1- Buscar Partida")
		fmt.Println("			2- ¿como jugar?")
		fmt.Println("			3- Salir\n")
		scanner.Scan()
	  	input := scanner.Text()
		input = strings.ToUpper(input)
		switch input {
			case "1": 
				buscarPartida()
				break
			case "2":
				comoJugar()
				break
			case "3":
				os.Exit(3)
				break
			default:
				clearScreen()
				fmt.Println("Opcion invalida")
				break
		}
	}
}

func buscarPartida()  {
	clearScreen()
	ip := "127.0.0.1"
	port := 3333
	SocketClient(ip, port)
}
