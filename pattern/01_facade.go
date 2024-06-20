package pattern

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

/*
	Реализовать паттерн «фасад».
Объяснить применимость паттерна, его плюсы и минусы,а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Facade_pattern
*/

type Account struct {
	name      string
	accountId int
	isBlocked bool
}

func newAccount(accountName string) *Account {
	fmt.Println("Starting create account")
	account := &Account{
		name:      accountName,
		accountId: rand.Int(),
		isBlocked: false,
	}
	fmt.Println("Account created")
	return account
}

func (acc *Account) getName() string {
	return acc.name
}

func (acc *Account) checkAccount(accountId int) bool {
	return !acc.isBlocked && accountId == acc.accountId
}

type Address struct {
	addressId int
	address   string
	price     float32
	isValid   bool
}

func newAddress(addressName string) *Address {
	fmt.Println("Starting create address")
	address := &Address{
		addressId: rand.Int(),
		address:   addressName,
		price:     rand.Float32() * 10,
		isValid:   true,
	}
	fmt.Println("Address created")
	return address
}

func (a *Address) checkAddress(addressId int) bool {
	return a.isValid && addressId == a.addressId
}

func (a *Address) getPrice() float32 {
	return a.price
}

func (a *Address) getAddress() string {
	return a.address
}

type Bucket struct {
	bucketId int
	products []*Product
}

func newBucket() *Bucket {
	fmt.Println("Starting create bucket")
	bucket := &Bucket{
		bucketId: rand.Int(),
	}
	fmt.Println("Bucket created")
	return bucket
}

func (b *Bucket) addProduct(product ...*Product) {
	b.products = append(b.products, product...)
}

func (b *Bucket) getOrder(bucketId int) (float32, error) {
	var order float32
	if bucketId != b.bucketId {
		return 0, errors.New("not found")
	}
	for _, el := range b.products {
		order += el.productPrice
	}

	return order, nil
}

type BuyFacade struct {
	account      Account
	pincode      PinCode
	bucket       Bucket
	address      Address
	notification Notification
}

func newBuyFacade(accountName string) *BuyFacade {
	fmt.Println("Starting create account")
	buyFacade := &BuyFacade{
		account:      *newAccount(accountName),
		pincode:      PinCode{},
		bucket:       *newBucket(),
		address:      Address{},
		notification: Notification{},
	}
	fmt.Println("Account created")
	return buyFacade
}

func (b *BuyFacade) setAddress(address *Address) {
	b.address = *address
}

func (b *BuyFacade) setPinCode(pin *PinCode) {
	b.pincode = *pin
}

func (b *BuyFacade) setNotification(msg *Notification) {
	b.notification = *msg
}

func (b *BuyFacade) getBucket() *Bucket {
	return &b.bucket
}

func (b *BuyFacade) buyProducts(accountId int, pinCode int, bucketId int, addressId int) string {
	status := true
	if !b.account.checkAccount(accountId) || !b.pincode.check(pinCode) || !b.address.checkAddress(addressId) {
		status = false
	}

	order, err := b.bucket.getOrder(bucketId)
	if err != nil {
		status = false
	}

	name := b.account.getName()

	addressStr := b.address.getAddress()
	order += b.address.getPrice()

	b.notification = *newNotification(name, order, addressStr, status)

	return b.notification.getMsg()
}

type Notification struct {
	message string
}

func newNotification(name string, order float32, address string, status bool) *Notification {
	var message string
	if !status {
		message = "Payment failed. Please try again"
	} else {

		message = fmt.Sprintf("Dear %s! You bought products for the %f dollars. Products will be delivered to your address: %s", name, order, address)
	}
	notification := Notification{
		message: message,
	}

	return &notification
}

func (n *Notification) getMsg() string {
	return n.message
}

type PinCode struct {
	pin        int
	createDate time.Time
}

func newPinCode() *PinCode {
	fmt.Println("Starting create pinCode")
	pinCode := &PinCode{
		pin:        rand.Intn(9999),
		createDate: time.Now(),
	}
	fmt.Println("PinCode created")
	return pinCode
}

func (p *PinCode) check(pin int) bool {
	return pin == p.pin && time.Now().Before(p.createDate.Add(5*time.Minute))
}

type Product struct {
	productId    int
	productName  string
	productPrice float32
}

func newProduct(name string, price float32) *Product {
	fmt.Println("Starting create product")
	product := &Product{
		productId:    rand.Int(),
		productName:  name,
		productPrice: price,
	}
	fmt.Println("Product created")
	return product
}

// func main() {

// 	product1 := newProduct("Apple", 5)
// 	product2 := newProduct("Milk", 3.5)
// 	product3 := newProduct("Chips", 2)

// 	addr1 := newAddress("Ivanovskaya street, 12")

// 	pin := newPinCode()

// 	buyEl := newBuyFacade("Oleg")

// 	buyEl.bucket.addProduct(product1, product2, product3)

// 	buyEl.setAddress(addr1)
// 	buyEl.setPinCode(pin)

// 	fmt.Println()
// 	msg := buyEl.buyProducts(buyEl.account.accountId, pin.pin, buyEl.bucket.bucketId, addr1.addressId)
// 	fmt.Println(msg)
// }
