package pattern

import "fmt"

/*
	Реализовать паттерн «цепочка вызовов».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Chain-of-responsibility_pattern
*/

type Assembler struct {
	next Equipment
}

func (a *Assembler) execute(d *Detail) {
	if d.isAssembleDone {
		fmt.Println("Assemble is already done")
		a.next.execute(d)
		return
	}
	if !d.evenlyColored {
		fmt.Println("Detail is not evenly colored")
		d.discontinued = true
		return
	}
	fmt.Println("Start assemble")
	d.isAssembleDone = true
	fmt.Println("End assemble")
	fmt.Printf("Detail '%s' is ready\n", d.name)
}

func (a *Assembler) setNext(next Equipment) {
	a.next = next
}

type Detail struct {
	name             string
	discontinued     bool
	isStampDone      bool
	isDefective      bool
	isWeldDone       bool
	isPaintDone      bool
	unweldedPartsCnt int
	isAssembleDone   bool
	evenlyColored    bool
}

type Equipment interface {
	execute(*Detail)
	setNext(Equipment)
}

type Painter struct {
	next Equipment
}

func (p *Painter) execute(d *Detail) {
	if d.isPaintDone {
		fmt.Println("Paint is already done")
		p.next.execute(d)
		return
	}
	if d.unweldedPartsCnt != 0 {
		fmt.Println("Several details are not welded")
		d.discontinued = true
		return
	}

	fmt.Println("Start paint")
	d.isPaintDone = true
	d.evenlyColored = true
	fmt.Println("End paint")
	p.next.execute(d)
}

func (p *Painter) setNext(next Equipment) {
	p.next = next
}

type Stamper struct {
	next Equipment
}

func (s *Stamper) execute(d *Detail) {
	if d.isStampDone {
		fmt.Println("Stamp is already done")
		s.next.execute(d)
		return
	}
	if d.isDefective {
		fmt.Println("Detail has defects")
		d.discontinued = true
		return
	}
	fmt.Println("Start stamp")
	d.isStampDone = true
	for i := d.unweldedPartsCnt; i > 0; i-- {
		d.unweldedPartsCnt--
	}
	fmt.Println("End stamp")
	s.next.execute(d)
}

func (s *Stamper) setNext(next Equipment) {
	s.next = next
}

type Welder struct {
	next Equipment
}

func (w *Welder) execute(d *Detail) {
	if d.isWeldDone {
		fmt.Println("Weld is already done")
		w.next.execute(d)
		return
	}
	fmt.Println("Start weld")
	d.isWeldDone = true
	fmt.Println("End weld")
	w.next.execute(d)
}

func (w *Welder) setNext(next Equipment) {
	w.next = next
}

// func main() {
// 	assembler := Assembler{}

// 	painter := Painter{}
// 	painter.setNext(&assembler)

// 	welder := Welder{}
// 	welder.setNext(&painter)

// 	stamper := Stamper{}
// 	stamper.setNext(&welder)

// 	detail1 := Detail{name: "car", unweldedPartsCnt: 10, isDefective: false}
// 	detail2 := Detail{name: "laptop", unweldedPartsCnt: 3, isDefective: true}

// 	stamper.execute(&detail1)
// 	if detail1.discontinued {
// 		fmt.Printf("Detail '%s' is broken\n", detail1.name)
// 	}
// 	stamper.execute(&detail2)
// 	if detail2.discontinued {
// 		fmt.Printf("Detail '%s' is broken\n", detail2.name)
// 	}
// }
