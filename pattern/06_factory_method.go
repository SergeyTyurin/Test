package pattern

import (
	"fmt"
	"strings"
)

/*
	Реализовать паттерн «фабричный метод».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Factory_method_pattern
*/

func getDocument(docType string, size int, name string) (IDocument, error) {
	if strings.EqualFold(docType, "pdf") {
		return newPDF(name, size), nil
	} else if strings.EqualFold(docType, "image") {
		return newImage(name, size), nil
	} else if strings.EqualFold(docType, "excel") {
		return newExcel(name, size), nil
	}

	return nil, fmt.Errorf(" Document type is not found")
}

type Document struct {
	name         string
	maxSize      int
	isChangeable bool
	size         int
	typeDoc      string
	sendBy       []string
}

func (d *Document) getType(typeDoc string, maxSize int, isChangeable bool) {
	d.typeDoc = typeDoc
	d.maxSize = maxSize
	d.isChangeable = isChangeable
}

func (d *Document) getDocument(size int, name string) {
	if d.maxSize < size {
		fmt.Println("error size")
		return
	}
	d.size = size
	d.name = name
}

func (d *Document) sendDocument(sendBy []string) {
	copy(d.sendBy, sendBy)
}

func (d *Document) printDocument() {
	fmt.Printf("Document '%s.%s' (%d bytes) printed\n", d.name, d.typeDoc, d.size)
}

type Excel struct {
	Document
}

func newExcel(name string, size int) IDocument {
	return &Excel{
		Document: Document{
			name:         name,
			size:         size,
			typeDoc:      "xlsx",
			maxSize:      1024,
			isChangeable: true,
		},
	}
}

type IDocument interface {
	getType(typeDoc string, maxSize int, isChangeable bool)
	getDocument(size int, name string)
	sendDocument(sendBy []string)
	printDocument()
}

type Image struct {
	Document
}

func newImage(name string, size int) IDocument {
	return &Image{
		Document: Document{
			name:         name,
			size:         size,
			typeDoc:      "img",
			maxSize:      512,
			isChangeable: false,
		},
	}
}

type PDF struct {
	Document
}

func newPDF(name string, size int) IDocument {
	return &PDF{
		Document: Document{
			name:         name,
			size:         size,
			typeDoc:      "pdf",
			maxSize:      128,
			isChangeable: false,
		},
	}
}

// func main() {
// 	excel1, err := getDocument("excel", 300, "loadcards")
// 	if err != nil {
// 		fmt.Println(err)
// 	} else {
// 		excel1.printDocument()
// 	}

// 	pdf1, err := getDocument("pdf", 100, "passport_scan")
// 	if err != nil {
// 		fmt.Println(err)
// 	} else {
// 		pdf1.printDocument()
// 	}

// 	image1, err := getDocument("image", 15, "photo")
// 	if err != nil {
// 		fmt.Println(err)
// 	} else {
// 		image1.printDocument()
// 	}

// }
