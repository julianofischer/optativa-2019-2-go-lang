package main

import (
	"fmt"
)

type datastructure interface {
	add(num int32)
	remove() int32
	imprime()
}

type fila struct {
	container[] int32
}
type pilha struct {
	container[] int32
}

func (f *fila) add(num int32) {
	slice := []int32{num}
	f.container = append(slice, f.container...)
}

func (f *fila) remove() int32 {
	r := f.container[len(f.container)-1]
	f.container = f.container[:len(f.container)-1]
	return r
}

func (f fila) imprime() {
	for i := len(f.container) - 1; i >=0; i-- {
		fmt.Print(f.container[i], " ")
	}
	fmt.Print("\n")
}

func (p *pilha) add(num int32) {
	p.container = append(p.container, num)
}

func (p *pilha) remove() int32 {
	r := p.container[len(p.container)-1]
	p.container = p.container[:len(p.container)-1]
	return r
}

func (p pilha) imprime() {
	for i := len(p.container)-1; i >= 0; i-- {
		fmt.Print(p.container[i], " ")
	}
	fmt.Print("\n")
}

func test(s datastructure) {
	fmt.Println("Início")
	s.add(1)
	s.add(2)
	s.add(3)
	fmt.Println(s)
	s.remove()
	s.imprime()
	s.add(4)
	s.add(5)
	s.add(6)
	s.add(7)
	s.imprime()
	fmt.Println("Fim")
}

func main() {
	p := pilha{}
	f := fila{}
	test(&p)
	test(&f)
}