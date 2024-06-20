package pattern

/*
	Реализовать паттерн «строитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Builder_pattern
*/

type carBuilder interface {
	setRoof()
	setDoors()
	setTrunk()
	setDrive()
	getCar() car
}

func createCar(builderType string) car {
	if builderType == "pickup truck" {
		return buildCar(&pickupTruckBuilder{})
	}
	if builderType == "roadster" {
		return buildCar(&roadsterBuilder{})
	}
	if builderType == "sedan" {
		return buildCar(&sedanBuilder{})
	}
	return car{}
}

func buildCar(b carBuilder) car {
	b.setRoof()
	b.setDoors()
	b.setTrunk()
	b.setDrive()
	return b.getCar()
}

type car struct {
	isRoofOpened bool
	doors        int
	trunkAmount  float32
	isTruncOpend bool
	drive        int
}

type pickupTruckBuilder struct {
	isRoofOpened bool
	doors        int
	trunkAmount  float32
	isTruncOpend bool
	drive        int
}

func (b *pickupTruckBuilder) setRoof() {
	b.isRoofOpened = false
}
func (b *pickupTruckBuilder) setDoors() {
	b.doors = 4
}
func (b *pickupTruckBuilder) setTrunk() {
	b.trunkAmount = 500
	b.isTruncOpend = true
}
func (b *pickupTruckBuilder) setDrive() {
	b.drive = 500
}

func (b *pickupTruckBuilder) getCar() car {
	return car{
		isRoofOpened: b.isRoofOpened,
		doors:        b.doors,
		trunkAmount:  b.trunkAmount,
		isTruncOpend: b.isTruncOpend,
		drive:        b.drive,
	}
}

type roadsterBuilder struct {
	isRoofOpened bool
	doors        int
	trunkAmount  float32
	isTruncOpend bool
	drive        int
}

func (b *roadsterBuilder) setRoof() {
	b.isRoofOpened = true
}
func (b *roadsterBuilder) setDoors() {
	b.doors = 2
}
func (b *roadsterBuilder) setTrunk() {
	b.trunkAmount = 100
	b.isTruncOpend = false
}
func (b *roadsterBuilder) setDrive() {
	b.drive = 400
}

func (b *roadsterBuilder) getCar() car {
	return car{
		isRoofOpened: b.isRoofOpened,
		doors:        b.doors,
		trunkAmount:  b.trunkAmount,
		isTruncOpend: b.isTruncOpend,
		drive:        b.drive,
	}
}

type sedanBuilder struct {
	isRoofOpened bool
	doors        int
	trunkAmount  float32
	isTruncOpend bool
	drive        int
}

func (b *sedanBuilder) setRoof() {
	b.isRoofOpened = false
}
func (b *sedanBuilder) setDoors() {
	b.doors = 4
}
func (b *sedanBuilder) setTrunk() {
	b.trunkAmount = 300
	b.isTruncOpend = false
}
func (b *sedanBuilder) setDrive() {
	b.drive = 200
}

func (b *sedanBuilder) getCar() car {
	return car{
		isRoofOpened: b.isRoofOpened,
		doors:        b.doors,
		trunkAmount:  b.trunkAmount,
		isTruncOpend: b.isTruncOpend,
		drive:        b.drive,
	}
}

// func main() {
// 	pickupTruck := createCar("pickup truck")
// 	roadster := createCar("roadster")
// 	sedan := createCar("sedan")

// 	fmt.Printf("pickupTruck: %v\n", pickupTruck)
// 	fmt.Printf("roadster: %v\n", roadster)
// 	fmt.Printf("sedan: %v\n", sedan)

// }
