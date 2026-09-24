package main

import "fmt"

type Engine struct {
	Horsepower int
	Type       string
}

func (e Engine) Start() {
	fmt.Printf("Starting %s engine with %d HP\n", e.Type, e.Horsepower)
}

func (e Engine) Stop() {
	fmt.Printf("Stopping %s engine\n", e.Type)
}

type Car struct {
	Make  string
	Model string
	Year  int
	Engine
}

func (c Car) Drive() {
	fmt.Printf("Driving %d %s %s\n", c.Year, c.Make, c.Model)
}

// Car's Start overrides the promoted Engine.Start method.
func (c Car) Start() {
	fmt.Printf("Starting %d %s %s\n", c.Year, c.Make, c.Model)
	c.Engine.Start()
}

type GPS struct {
	Brand   string
	HasMaps bool
}

func (g GPS) Navigate(destination string) {
	fmt.Printf("%s GPS navigating to %s\n", g.Brand, destination)
}

type Radio struct {
	Brand        string
	HasBluetooth bool
}

func (r Radio) PlayMusic() {
	fmt.Printf("Playing music on %s radio\n", r.Brand)
}

type LuxuryCar struct {
	LeatherSeats bool
	Car
	GPS
	Radio
}

type Person struct {
	Name string
	Age  int
}

type Company struct {
	Name      string
	Employees int
}

// Person and Company both have a Name field,
// so Name must be accessed explicitly through either one.
type Employee struct {
	ID       int
	Position string
	Person
	Company
}

func main() {
	car := Car{
		Make:  "Toyota",
		Model: "Camry",
		Year:  2023,
		Engine: Engine{
			Horsepower: 200,
			Type:       "V6",
		},
	}

	car.Start()
	car.Engine.Start()
	car.Drive()
	car.Stop()

	fmt.Println("---")

	luxuryCar := LuxuryCar{
		LeatherSeats: true,
		Car:          car,
		GPS: GPS{
			Brand:   "Garmin",
			HasMaps: true,
		},
		Radio: Radio{
			Brand:        "Sony",
			HasBluetooth: true,
		},
	}

	luxuryCar.Start()
	luxuryCar.Drive()
	luxuryCar.Navigate("Downtown")
	luxuryCar.Radio.PlayMusic()

	fmt.Println("---")

	emp := Employee{
		ID:       1,
		Position: "Software Engineer",
		Person: Person{
			Name: "Vijet",
			Age:  21,
		},
		Company: Company{
			Name:      "Example Corp",
			Employees: 100,
		},
	}

	fmt.Printf("Employee: %s (ID: %d)\n", emp.Person.Name, emp.ID)
	fmt.Printf("Works at: %s\n", emp.Company.Name)
	fmt.Printf("Position: %s\n", emp.Position)
	fmt.Printf("Age: %d\n", emp.Person.Age)

	// An embedded type's methods are promoted to the outer type.
	var engine Engine = car.Engine
	engine.Start()
}
