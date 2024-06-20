package pattern

import "fmt"

/*
	Реализовать паттерн «состояние».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/State_pattern
*/

type GettingChange struct {
	vendingMachine *VendingMachine
}

func (i *GettingChange) insertMoney(money int) error {

	return fmt.Errorf("Error")
}

func (i *GettingChange) startCoffee() error {
	return fmt.Errorf("Error")
}
func (i *GettingChange) setChange() error {
	if i.vendingMachine.change > 0 {
		fmt.Printf("Please, take your change: %d rub\n", i.vendingMachine.change)
		i.vendingMachine.change = 0
		i.vendingMachine.setState(i.vendingMachine.waiting)
	} else {
		i.vendingMachine.setState(i.vendingMachine.waiting)
	}
	return nil
}
func (i *GettingChange) setCheck() error {
	return fmt.Errorf("Error")
}
func (i *GettingChange) finishWork() error {
	return fmt.Errorf("Error")
}

type PrintingCheck struct {
	vendingMachine *VendingMachine
}

func (i *PrintingCheck) insertMoney(money int) error {

	return fmt.Errorf("Error")
}

func (i *PrintingCheck) startCoffee() error {
	return fmt.Errorf("Error")
}
func (i *PrintingCheck) setChange() error {
	return fmt.Errorf("Error")
}
func (i *PrintingCheck) setCheck() error {
	if i.vendingMachine.checkPaper > 0 {
		i.vendingMachine.checkPaper--
		fmt.Println("Check is ready")
		i.vendingMachine.setState(i.vendingMachine.gettingChange)
	} else {
		fmt.Println("No paper for check")
		i.vendingMachine.isBroken = true
		i.vendingMachine.setState(i.vendingMachine.waiting)
	}
	return nil
}
func (i *PrintingCheck) finishWork() error {
	return fmt.Errorf("Error")
}

type GettingMoney struct {
	vendingMachine *VendingMachine
}

func (i *GettingMoney) insertMoney(money int) error {

	if money >= i.vendingMachine.coffeePrice && i.vendingMachine.capsuleCoffeeCnt > 0 {
		i.vendingMachine.change = money - i.vendingMachine.coffeePrice
		fmt.Println("Start coffee")
		i.vendingMachine.setState(i.vendingMachine.makingCoffee)
	} else if money < i.vendingMachine.coffeePrice {
		i.vendingMachine.change = money
		i.vendingMachine.setState(i.vendingMachine.gettingChange)
	} else {
		i.vendingMachine.isBroken = true
		i.vendingMachine.change = money
		i.vendingMachine.setState(i.vendingMachine.gettingChange)
	}
	return nil
}

func (i *GettingMoney) startCoffee() error {
	return fmt.Errorf("Error")
}
func (i *GettingMoney) setChange() error {
	return fmt.Errorf("Error")
}
func (i *GettingMoney) setCheck() error {
	return fmt.Errorf("Error")
}
func (i *GettingMoney) finishWork() error {
	return fmt.Errorf("Error")
}

type MakingCoffee struct {
	vendingMachine *VendingMachine
}

func (i *MakingCoffee) insertMoney(money int) error {

	return fmt.Errorf("Error")
}

func (i *MakingCoffee) startCoffee() error {
	i.vendingMachine.decrementCoffeeCnt(1)
	fmt.Println("Coffee is ready")
	i.vendingMachine.setState(i.vendingMachine.printingCheck)
	return nil
}
func (i *MakingCoffee) setChange() error {
	return fmt.Errorf("Error")
}
func (i *MakingCoffee) setCheck() error {
	return fmt.Errorf("Error")
}
func (i *MakingCoffee) finishWork() error {
	return fmt.Errorf("Error")
}

type State interface {
	insertMoney(money int) error
	startCoffee() error
	setChange() error
	setCheck() error
	finishWork() error
}

type VendingMachine struct {
	gettingMoney  State
	gettingChange State
	makingCoffee  State
	printingCheck State
	waiting       State

	currentState State

	capsuleCoffeeCnt int
	coffeePrice      int
	change           int
	checkPaper       int
	isBroken         bool
}

func newVendingMachine(capsuleCoffeeCnt int, coffeePrice int, checkPaper int) *VendingMachine {
	v := VendingMachine{
		capsuleCoffeeCnt: capsuleCoffeeCnt,
		coffeePrice:      coffeePrice,
		checkPaper:       checkPaper,
	}
	gettingMoneyVal := GettingMoney{vendingMachine: &v}
	makingCoffeeVal := MakingCoffee{vendingMachine: &v}
	printingCheckVal := PrintingCheck{vendingMachine: &v}
	gettingChangeVal := GettingChange{vendingMachine: &v}
	waitingVal := Waiting{vendingMachine: &v}

	v.setState(&gettingMoneyVal)
	v.gettingMoney = &gettingMoneyVal
	v.makingCoffee = &makingCoffeeVal
	v.printingCheck = &printingCheckVal
	v.gettingChange = &gettingChangeVal
	v.waiting = &waitingVal
	return &v
}

func (v *VendingMachine) insertMoney(money int) error {
	return v.currentState.insertMoney(money)
}

func (v *VendingMachine) startCoffee() error {
	return v.currentState.startCoffee()
}

func (v *VendingMachine) setChange() error {
	return v.currentState.setChange()
}

func (v *VendingMachine) setCheck() error {
	return v.currentState.setCheck()
}

func (v *VendingMachine) finishWork() error {
	return v.currentState.finishWork()
}

func (v *VendingMachine) setState(s State) {
	v.currentState = s
}

func (v *VendingMachine) decrementCoffeeCnt(count int) {
	v.capsuleCoffeeCnt -= count
}

type Waiting struct {
	vendingMachine *VendingMachine
}

func (i *Waiting) insertMoney(money int) error {
	fmt.Println("contact technical support number")
	return fmt.Errorf("vending Machine is broken")
}

func (i *Waiting) startCoffee() error {
	fmt.Println("contact technical support number")
	return fmt.Errorf("vending Machine is broken")
}
func (i *Waiting) setChange() error {
	fmt.Println("contact technical support number")
	return fmt.Errorf("vending Machine is broken")
}
func (i *Waiting) setCheck() error {
	fmt.Println("contact technical support number")
	return fmt.Errorf("vending Machine is broken")
}
func (i *Waiting) finishWork() error {
	if i.vendingMachine.isBroken {
		fmt.Println("contact technical support number")
		return fmt.Errorf("vending Machine is broken")
	}
	fmt.Println("Have a good day!")
	i.vendingMachine.setState(i.vendingMachine.gettingMoney)
	return nil
}

// func main() {
// 	vendingMachine := newVendingMachine(1, 10, 3)

// 	err := vendingMachine.insertMoney(15)
// 	if err != nil {
// 		log.Fatalf(err.Error())
// 	}

// 	err = vendingMachine.startCoffee()
// 	if err != nil {
// 		log.Fatalf(err.Error())
// 	}

// 	err = vendingMachine.setCheck()
// 	if err != nil {
// 		log.Fatalf(err.Error())
// 	}

// 	err = vendingMachine.setChange()
// 	if err != nil {
// 		log.Fatalf(err.Error())
// 	}

// 	err = vendingMachine.finishWork()
// 	if err != nil {
// 		log.Fatalf(err.Error())
// 	}

// 	fmt.Println()
// 	err = vendingMachine.insertMoney(15)
// 	if err != nil {
// 		log.Fatalf(err.Error())
// 	}

// 	err = vendingMachine.setChange()
// 	if err != nil {
// 		log.Fatalf(err.Error())
// 	}

// 	err = vendingMachine.finishWork()
// 	if err != nil {
// 		log.Fatalf(err.Error())
// 	}

// }
