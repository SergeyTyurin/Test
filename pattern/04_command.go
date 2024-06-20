package pattern

import "fmt"

/*
	Реализовать паттерн «комманда».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Command_pattern
*/

type Command interface {
	execute()
}

type Button struct {
	command Command
}

func (b *Button) press() {
	b.command.execute()
}

type Device interface {
	on()
	off()
}

type OffCommand struct {
	device Device
}

func (c *OffCommand) execute() {
	c.device.off()
}

type OnCommand struct {
	device Device
}

func (c *OnCommand) execute() {
	c.device.on()
}

type Sound struct {
	isRunning bool
}

func (s *Sound) on() {
	s.isRunning = true
	fmt.Println("Turning sound on")
}

func (s *Sound) off() {
	s.isRunning = false
	fmt.Println("Turning sound off")
}

type Tv struct {
	isRunning bool
}

func (t *Tv) on() {
	t.isRunning = true
	fmt.Println("Turning tv on")
}

func (t *Tv) off() {
	t.isRunning = false
	fmt.Println("Turning tv off")
}

// func main() {
// 	tv := Tv{}
// 	sound := Sound{}

// 	onTV := OnCommand{device: &tv}
// 	onSound := OnCommand{device: &sound}

// 	offTV := OffCommand{device: &tv}
// 	offSound := OffCommand{device: &sound}

// 	onBtnTV := Button{
// 		command: &onTV,
// 	}

// 	offBtnTV := Button{
// 		command: &offTV,
// 	}

// 	onBtnSound := Button{
// 		command: &onSound,
// 	}

// 	offBtnSound := Button{
// 		command: &offSound,
// 	}

// 	onBtnTV.press()
// 	onBtnSound.press()
// 	offBtnSound.press()
// 	offBtnTV.press()

// }
