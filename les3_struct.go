package main

import "fmt"

type storage interface {
	insert(e employee) error
	delete(id int) error
	get(id int) (employee, error)
}

type memoryStorage struct {
	data map[int]employee
}

func newMemoryStorage() *memoryStorage {
	return &memoryStorage{
		data: make(map[int]employee),
	}
}

func (ms *memoryStorage) insert(e employee) error {
	ms.data[e.id] = e
	return nil
}

func (ms *memoryStorage) delete(id int) error {
	delete(ms.data, id)
	return nil
}
func (ms *memoryStorage) get(id int) (employee, error) {
	return ms.data[id], nil
}

type dumbStorage struct{}

func newDumbStorage() *dumbStorage {
	return &dumbStorage{}
}

func (ds *dumbStorage) insert(e employee) error {
	fmt.Printf("dumbStorage insert: %v\n", e.id)
	return nil
}

func (ds *dumbStorage) delete(id int) error {
	fmt.Printf("dumbStorage delete: %v\n", id)
	return nil
}

func (ds *dumbStorage) get(id int) (employee, error) {
	e := employee{
		id: id,
	}

	return e, nil
}

type employee struct {
	id     int
	name   string
	sex    string
	age    int
	salary int
}

func newEmployee(id int, name, sex string, age, salary int) employee {
	return employee{
		id:     id,
		name:   name,
		sex:    sex,
		age:    age,
		salary: salary,
	}
}

func (e employee) getInfo() string {
	return fmt.Sprintf("Сотрудник: %s\nId: %d\nВозраст: %d\nЗарплата: %d", e.name, e.id, e.age, e.salary)
}
func (e *employee) printInfo() {
	fmt.Printf("%s\n\n", e.getInfo())
}

func (e *employee) setName(newName string) {
	e.name = newName
}
func (e *employee) upSalary(addSalary int) {
	e.salary += addSalary
}

func main() {
	employee1 := newEmployee(1, "James", "M", 40, 1000)
	employee2 := newEmployee(2, "Bart", "M", 12, 0)
	employee3 := newEmployee(3, "Liza", "W", 10, 0)
	employee1.printInfo()

	employee1.setName("Gomer")
	employee1.printInfo()

	employee1.upSalary(100)
	employee1.printInfo()

	var memStor storage

	memStor = newMemoryStorage()
	_ = memStor.insert(employee1)
	_ = memStor.insert(employee2)
	_ = memStor.insert(employee3)

	_ = memStor.delete(2)

	fmt.Println(memStor)
	fmt.Println()

	liza, _ := memStor.get(3)
	liza.printInfo()

	memStor = newDumbStorage()
	_ = memStor.insert(employee1)
	_ = memStor.insert(employee2)
	_ = memStor.insert(employee3)

	_ = memStor.delete(2)

	fmt.Println(memStor)
	fmt.Println()

	gomer, _ := memStor.get(1)
	gomer.printInfo()
}
