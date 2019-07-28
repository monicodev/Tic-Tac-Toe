package main

import (
	//"strings"
	"strconv"
	"bufio"
	"net"
	"log"
	"fmt"
	//"time"
)

type partida struct {
	Player net.Conn
	Tablero [3][3]string
}
/*
*	metodos privados inicial en mayuscula
*	metodos publicos inicial en minuscula
*	de igual forma con los atributos de la clase
*/

func (p partida) PedirValor(i int, j int) string {
	switch i {
		case 0:
			switch j {
				case 0: return "1"
				case 1: return "2"
				case 2: return "3"	
			}
		case 1:
			switch j {
				case 0: return "4"
				case 1: return "5"
				case 2: return "6"	
			}
		case 2:
			switch j {
				case 0: return "7"
				case 1: return "8"
				case 2: return "9"	
			}
	}
	return ""
}
func (p partida) GetPlayer() net.Conn{
	//log.Println("partida -> GetPlayer -> p.Player = ",p.Player.RemoteAddr()) //ESTO ES PARA PRUEBAS
	return p.Player
}

func (p *partida) SetPlayer(j net.Conn){
	//log.Println("partida -> SetPlayer -> p.Player = ",j.RemoteAddr()) //ESTO ES PARA PRUEBAS
	p.Player = j
}

func (p *partida) iniTabla(){
	for i := 0 ; i<3; i++ {
		for j := 0 ; j<3; j++ {
			p.Tablero[i][j] = strconv.Itoa(((i*3)+j+1))
		}
	}
}

//en espera de la jugabilidad del cliente
func (p *partida)printTabla(){
	fmt.Println("  ",""," ",""," ","\n")
	fmt.Println(p.Tablero[0][0], "|", p.Tablero[0][1] ,"|", p.Tablero[0][2])
	fmt.Println("─","+","─","+","─")
	fmt.Println(p.Tablero[1][0], "|", p.Tablero[1][1] ,"|", p.Tablero[1][2])
	fmt.Println("─","+","─","+","─")
	fmt.Println(p.Tablero[2][0], "|", p.Tablero[2][1] ,"|", p.Tablero[2][2])
}

func (p partida) enviar(s string){
	w  := bufio.NewWriter(p.GetPlayer())
	w.Write([]byte(s))			
	w.Flush()
}

func (p *partida) PlayerTurn() int{
	var (
		buf = make([]byte, 1024)
		r   = bufio.NewReader(p.GetPlayer()) 
	)

	n,_ := r.Read(buf)
	
	switch string(buf[:n]) {
		case "1":
			
			p.Tablero[0][0] = "X";
			break
		case "2":
			
			p.Tablero[0][1] = "X";
			break
		case "3":
			
			p.Tablero[0][2] = "X";
			break
		case "4":
			
			p.Tablero[1][0] = "X";
			break
		case "5":
			
			p.Tablero[1][1] = "X";
			break
		case "6":
			
			p.Tablero[1][2] = "X";
			break
		case "7":
			
			p.Tablero[2][0] = "X";
			break
		case "8":
			
			p.Tablero[2][1] = "X";
			break
		case "9":
			
			p.Tablero[2][2] = "X";
			break
		default:
			log.Println("<-- Se ha desconectado el cliente",p.Player.RemoteAddr())
			
			panic("ALV Wey")
			return 1
	}
	return 0
}
func Fila(i int, e string, p *partida) bool{
	log.Println("Fila ",i)
	switch p.Tablero[i][0] {
	case e:
		switch p.Tablero[i][1] {
		case e:
			switch p.Tablero[i][2] {
			case "3","6","9":
				log.Println(p.Tablero[i][2])
				p.Tablero[i][2] = "O"
				log.Println(e," ",e," 3-6-9")
				p.Jugue(p.PedirValor(i,2))
				p.enviar(p.PedirValor(i,2))
				return true
			}
		case "2","5","8":
			switch p.Tablero[i][2] {
			case e:
				p.Tablero[i][1] = "O"
				log.Println(e," 2-5-8 ",e)
				p.Jugue(p.PedirValor(i,1))
				p.enviar(p.PedirValor(i,1))
				
				return true
			}
			break
		}
		break
	case "1","4","7":
		switch p.Tablero[i][1] {
		case e:
			switch p.Tablero[i][2] {
			case e:
				p.Tablero[i][0] = "O"
				log.Println(" 1-4-7",e," ",e)
				p.Jugue(p.PedirValor(i,0))
				p.enviar(p.PedirValor(i,0))
				return true
			}
			break
		}
		break
	}
	return false
}

func Columna(i int,e string, p *partida) bool{
	log.Println("Columna ",i)
	switch p.Tablero[0][i] {
	case e:
		switch p.Tablero[1][i] {
		case e:
			switch p.Tablero[2][i] {
			case "7","8","9":
				log.Println(p.Tablero[i][2])
				p.Tablero[2][i] = "O"
				log.Println(e," ",e," 7-8-9")
				p.Jugue(p.PedirValor(2,i))
				p.enviar(p.PedirValor(2,i))
				return true
			}
		case "4","5","6":
			switch p.Tablero[2][i] {
			case e:
				log.Println(e," 4-5-6 ",e)
				p.Jugue(p.PedirValor(1,i))
				p.enviar(p.PedirValor(1,i))
				p.Tablero[1][i] = "O"
				return true
			}
			break
		}
		break
	case "1","2","3":
		switch p.Tablero[1][i] {
		case e:
			switch p.Tablero[2][i] {
			case e:
				p.Tablero[0][i] = "O"
				log.Println(" 1-2-3",e," ",e)
				p.Jugue(p.PedirValor(0,i))
				p.enviar(p.PedirValor(0,i))
				return true
			}
			break
		}
		break
	}
	return false
}
func DiagonalPrincipal(e string, p *partida) bool {
	log.Println("DiagonalPrincipal")
	switch p.Tablero[0][0] {
		case e:
			switch p.Tablero[1][1] {
			case e:
				switch p.Tablero[2][2] {
				case "9":
					p.Tablero[2][2] = "O"
					log.Println("00-11-9 ")
					p.Jugue(p.PedirValor(2,2))
					p.enviar(p.PedirValor(2,2))
					return true
				}
			case "5":
				switch p.Tablero[2][2] {
				case e:
					p.Tablero[1][1] = "O"
					log.Println("00-5-22 ")
					p.Jugue(p.PedirValor(1,1))
					p.enviar(p.PedirValor(1,1))
					return true
				}
				break
			}
			break
		case "1":
			switch p.Tablero[1][1] {
			case e:
				switch p.Tablero[2][2] {
				case e:
					p.Tablero[0][0] = "O"
					log.Println("1-11-22 ")
					p.Jugue(p.PedirValor(0,0))
					p.enviar(p.PedirValor(0,0))
					return true
				}
				break
			}
			break
		}
		return false
}
func DiagonalSequndaria(e string, p *partida) bool {
	log.Println("DiagonalSequndaria")
	switch p.Tablero[2][0] {
		case e:
			switch p.Tablero[1][1] {
			case e:
				switch p.Tablero[0][2] {
				case "3":
					p.Tablero[0][2] = "O"
					log.Println("20-11-3 ")
					p.Jugue(p.PedirValor(0,2))
					p.enviar(p.PedirValor(0,2))
					return true
				}
			case "5":
				switch p.Tablero[0][2] {
				case e:
					p.Tablero[1][1] = "O"
					log.Println("20-5-02 ")
					p.Jugue(p.PedirValor(1,1))
					p.enviar(p.PedirValor(1,1))
					return true
				}
				break
			}
			break
		case "7":
			switch p.Tablero[1][1] {
			case e:
				switch p.Tablero[0][2] {
				case e:
					p.Tablero[2][0] = "O"
					log.Println("7-11-02 ")
					p.Jugue(p.PedirValor(2,0))
					p.enviar(p.PedirValor(2,0))
					return true
				}
				break
			}
			break
		}
		return false
}
func HorizontalWin(p *partida) bool{ //para ganar busqueda horizontal

	if Fila(0,"O",p) { ;return true }
	if Fila(1,"O",p) { return true }
	if Fila(2,"O",p) { return true }
	return false
}
func Horizontal(p *partida) bool{ //para evitar perder busqueda horizontal
	if Fila(0,"X",p) { return true }
	if Fila(1,"X",p) { return true }
	if Fila(2,"X",p) { return true }
	return false
}
func VerticalWin(p *partida) bool { // para ganar busqueda vertical
	if Columna(0,"O",p) { return true }
	if Columna(1,"O",p) { return true }
	if Columna(2,"O",p) { return true }
	return false
}
func Vertical(p *partida) bool { //para evitar perder busqueda vertical
	if Columna(0,"X",p) { return true }
	if Columna(1,"X",p) { return true }
	if Columna(2,"X",p) { return true }
	return false
}
func Urgencia(p *partida) bool {
	if DiagonalPrincipal("O",p) {  return true }//para ganar
	if DiagonalSequndaria("O",p) { return true }//para ganar
	if HorizontalWin(p) { return true }//para ganar
	if VerticalWin(p) { return true }//para ganar
	if Horizontal(p) { return true }//para evitar perder
	if Vertical(p) { return true }//para evitar perder
	if DiagonalPrincipal("X",p) { return true }//para evitar perder
	if DiagonalSequndaria("X",p) { return true }//para evitar perder
	return false
}

//jugabilidad de la computadora
func ComputerTurn(p *partida){
	log.Println("mi turno")
	if Urgencia(p) != true {
		for i := 0 ; i<3; i++ {
			for j := 0 ; j<3; j++ {
				switch p.Tablero[i][j] {
				case "1":
					p.Jugue(p.Tablero[i][j])
					p.Tablero[i][j] = "O";
					p.enviar("1")
					return
				case "2":
					p.Jugue(p.Tablero[i][j])
					p.Tablero[i][j] = "O";
					p.enviar("2")
					return
				case "3":
					p.Jugue(p.Tablero[i][j])
					p.Tablero[i][j] = "O";
					p.enviar("3")
					return
				case "4":
					p.Jugue(p.Tablero[i][j])
					p.Tablero[i][j] = "O";
					p.enviar("4")
					return
				case "5":
					p.Jugue(p.Tablero[i][j])
					p.Tablero[i][j] = "O";
					p.enviar("5")
					return
				case "6":
					p.Jugue(p.Tablero[i][j])
					p.Tablero[i][j] = "O";
					p.enviar("6")
					return
				case "7":
					p.Jugue(p.Tablero[i][j])
					p.Tablero[i][j] = "O";
					p.enviar("7")
					return
				case "8":
					p.Jugue(p.Tablero[i][j])
					p.Tablero[i][j] = "O";
					p.enviar("8")
					return
				case "9":
					p.Jugue(p.Tablero[i][j])
					p.Tablero[i][j] = "O";
					p.enviar("9")
					return
				default:

				}
			}
		}
	}

}
func (p partida) Jugue(i string){
	log.Println("yo jugue con ",p.GetPlayer().RemoteAddr(),"en la posicion ",i)
}


func (p partida) Start(){
	p.iniTabla()
	for{
		p.PlayerTurn()
		ComputerTurn(&p)
		p.printTabla()
	}
}
