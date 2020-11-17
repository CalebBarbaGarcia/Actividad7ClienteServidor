package main

import (
	"fmt"
	"time"
	"net"
	"encoding/gob"
)

var parar int64



type Procesos struct{
	SliceProcesos []Proceso
}

func (ps *Procesos) AgregarProceso(p Proceso){
	ps.SliceProcesos = append(ps.SliceProcesos,p)
}

func (ps *Procesos) MostrarProcesos(){
	for{
		for i:=0;i < len(ps.SliceProcesos);i = i +1{
			ps.SliceProcesos[i].MostrarProceso()
		}

		time.Sleep(time.Millisecond * 500)

		if parar == 1{
			return
		}
	}
}


func (ps *Procesos) EliminarProceso(b int64){
	var posicion int

	posicion = int(b)

	if posicion != -1{
		if posicion == len(ps.SliceProcesos)-1{
			ps.SliceProcesos = append(ps.SliceProcesos[:posicion])
		} else {
			ps.SliceProcesos = append(ps.SliceProcesos[:posicion],ps.SliceProcesos[posicion+1:]...)
		}
	}
}


type Proceso struct {
	Id int64
	I int64
}

func (p *Proceso) HacerProceso() {
	for {
		p.I = p.I + 1
		time.Sleep(time.Millisecond * 500)
	}
}

func (p *Proceso) MostrarProceso() {
	fmt.Println("id ", p.Id,": " ,p.I)
	p.I = p.I + 1
}


func servidor(procesos *Procesos)  {
	s, err := net.Listen("tcp", ":9999")
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		c, err := s.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}

		c, err = net.Dial("tcp", ":9997")
		if err != nil {
			fmt.Println(err)
			return
		}
		err = gob.NewEncoder(c).Encode(procesos.SliceProcesos[0])
		if err != nil {
			fmt.Println(err)
		}
		procesos.EliminarProceso(0)
		c.Close()
	}
}

func servidor1(procesos *Procesos)  {
	s, err := net.Listen("tcp", ":9998")
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		c, err := s.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}

		handleClient(c,procesos)
	}
}

func handleClient(c net.Conn, p* Procesos)  {
	var proceso Proceso
	err := gob.NewDecoder(c).Decode(&proceso)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		p.SliceProcesos = append(p.SliceProcesos,proceso)
	}
}


func main()  {
	procesos:= Procesos{}

	procesos.AgregarProceso(Proceso{Id:0,I:0})
	procesos.AgregarProceso(Proceso{Id:1,I:0})
	procesos.AgregarProceso(Proceso{Id:2,I:0})
	procesos.AgregarProceso(Proceso{Id:3,I:0})
	procesos.AgregarProceso(Proceso{Id:4,I:0})

	go servidor(&procesos)
	go servidor1(&procesos)
	go procesos.MostrarProcesos()

	fmt.Scanln(&parar)
} 