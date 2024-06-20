package pattern

import "fmt"

/*
	Реализовать паттерн «посетитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Visitor_pattern
*/

type Book struct {
	name            string
	author          string
	publishingHouse string
	binding         string
	pageCnt         int
	price           float32
}

func (b *Book) accept(v Visitor) {
	v.visitForBook(b)
}

func (b *Book) getType() string {
	return "book"
}

type InfoSbj struct {
	infoMsg string
}

func (i *InfoSbj) visitForBook(b *Book) {
	i.infoMsg = fmt.Sprintf("name: %s, author:%s, publishingHouse: %s, binding: %s, pageCnt: %d, price: %f", b.name, b.author, b.publishingHouse, b.binding, b.pageCnt, b.price)
	fmt.Println(i.infoMsg)
}

func (i *InfoSbj) visitForJeans(b *Jeans) {
	i.infoMsg = fmt.Sprintf("model: %s, size:%d, color: %s, length: %d, price: %f", b.model, b.size, b.color, b.length, b.price)
	fmt.Println(i.infoMsg)
}

func (i *InfoSbj) visitForLaptop(b *Laptop) {
	i.infoMsg = fmt.Sprintf("model: %s, cpu:%s, diagonal:%f, color: %s, diskType: %s, diskGb: %d, os: %s, price: %f", b.model, b.cpu, b.diagonal, b.color, b.diskType, b.diskGb, b.os, b.price)
	fmt.Println(i.infoMsg)
}

type Jeans struct {
	model  string
	size   int
	color  string
	length int
	price  float32
}

func (j *Jeans) accept(v Visitor) {
	v.visitForJeans(j)
}

func (j *Jeans) getType() string {
	return "jeans"
}

type Laptop struct {
	model    string
	cpu      string
	diagonal float32
	color    string
	diskType string
	diskGb   int
	os       string
	price    float32
}

func (l *Laptop) accept(v Visitor) {
	v.visitForLaptop(l)
}

func (l *Laptop) getType() string {
	return "laptop"
}

type PackageSbj struct {
	typePack      string
	sizePack      int
	isFragilePack bool
}

func (p *PackageSbj) visitForBook(b *Book) {
	p.typePack = "plastic bag"
	p.sizePack = 1
	fmt.Println(p)
}

func (p *PackageSbj) visitForJeans(j *Jeans) {
	p.typePack = "plastic bag"
	p.sizePack = 2
	fmt.Println(p)
}

func (p *PackageSbj) visitForLaptop(l *Laptop) {
	p.typePack = "box"
	p.sizePack = 0
	p.isFragilePack = true
	fmt.Println(p)
}

type Sbj interface {
	getType() string
	accept(Visitor)
}

type Visitor interface {
	visitForBook(*Book)
	visitForJeans(*Jeans)
	visitForLaptop(*Laptop)
}

// func main() {
// 	book1 := &Book{name: "Hobbit", author: "J. R. R. Tolkien", publishingHouse: "asti", binding: "hard", pageCnt: 312, price: 80.99}
// 	jeans1 := &Jeans{model: "H&M skiny", size: 38, color: "blue", length: 32, price: 34.95}
// 	laptop1 := &Laptop{model: "Lenovo Yoga", cpu: "Intel Celeron N4020", diagonal: 15.6, color: "pink", diskType: "SSD", diskGb: 128, os: "Windows 11 Home", price: 415}

// 	info := &InfoSbj{}

// 	book1.accept(info)
// 	jeans1.accept(info)
// 	laptop1.accept(info)

// 	packageSubject := &PackageSbj{}

// 	book1.accept(packageSubject)
// 	jeans1.accept(packageSubject)
// 	laptop1.accept(packageSubject)
// }
