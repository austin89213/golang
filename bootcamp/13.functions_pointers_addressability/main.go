package main

import "fmt"

// 建立 person type
// type person struct {
//     firstName string
//     lastName  string
// }

// // 建立 person 的 function receiver
// func (p person) updateName(newFirstName string) {
//     fmt.Printf("Before update: %+v\n", p)
//     p.firstName = newFirstName
//     fmt.Printf("After update: %+v\n", p)
// }

// func (p person) print() {
//     fmt.Printf("Current person is: %+v\n", p)
// }

// func main() {
//     jim := person{
//         firstName: "Jim",
//         lastName:  "Party",
//     }

//   // Before update: {firstName:Jim lastName:Party}
//     jim.updateName("Aaron")
//   // After update: {firstName:Aaron lastName:Party}

//     jim.print() // Current person is: {firstName:Jim lastName:Party}
// }

// func main() {
//     jim := "Jim"
//     fmt.Println(jim)        // Jim
//     changeName(jim)
//     fmt.Println(jim)        // Jim
// }

// func changeName(person string) {
//     person = "Bob"
// }

type person struct {
	firstName string
	lastName  string
}

// 直接傳入參數時，thePerson 會複製一份新的
func updateFirstName(thePerson person) {
	thePerson.firstName = "Aaron"
}

// 透過 *type 傳入 pointer，會參照到原本的 thePerson
func updateFirstNameWithPointer(thePerson *person) {
	(*thePerson).firstName = "Aaron"
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Anderson",
	}

	fmt.Println(jim) // {Jim Anderson}
	updateFirstName(jim)
	fmt.Println(jim) // {Jim Anderson}

	jimPointer := &jim
	updateFirstNameWithPointer(jimPointer)
	fmt.Println(jim) // {Aaron Anderson}
}
