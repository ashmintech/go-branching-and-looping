package main

import (
	"fmt"
)

type vehicle interface {
	printVehicleDetails(string)
}

type car struct {
	model	string
	make	string
	color	string
	typeVehicle	string
	cc	float32
	year int
	seating	int
	doors int
}

type truck struct {
	model	string
	make	string
	color	string
	typeVehicle	string
	cc	float32
	year int
	seating	int
	doors int
	loadingCapacity	float32
}

type bike struct {
	model	string
	make	string
	color	string
	cc	float32
	year int
}

var inventory []vehicle

func init() {

	var timeMachine vehicle

	inventory = []vehicle{
		bike{"FTR 1200", "Indian", "Black", 1203.0, 2019}, 
		bike{"Iron 1200", "Harley", "Blue", 1200.0, 2018}, 
		car{"Sonata", "Hyundai", "White", "Sedan", 2400.0, 2017, 5, 4},
		car{"SantaFe", "Hyundai", "Red", "SUV", 2400.0, 2016, 7, 4},
		car{"Civic", "Honda", "White", "Hatchback", 2000.0, 2017, 5, 4},
		car{"A5", "Audi", "Red", "Coupe", 3000.0, 2019, 2, 2},
		car{"Mazda6", "Mazda", "White", "Sedan", 2500.0, 2018, 5, 4},
		car{"CRV", "Honda", "Black", "SUV", 1500.0, 2017, 5, 4},
		car{"Camry", "Toyota", "Silver", "Sedan", 3500.0, 2018, 5, 4},
		truck{"F-150", "Ford", "Gray", "Truck", 3600.0, 2014, 7, 4, 13200.0},
		truck{"RAM1500", "Dodge", "White", "Truck", 5700.0, 2017, 5, 2, 12750.0},
		truck{"Silverado", "Chevrolet", "Black", "Truck", 6000.0, 2016, 7, 4, 14500.0},
		timeMachine }
		

}

func main() {

	carCount , bikeCount , truckCount := 0,0,0

	fmt.Printf("%-15v%-12v%-12v%-10v%-4v%12v%12v%8v%8v\n", "Type" , "Make" , "Model" , "Color" , "CC" , "Year" , "Seating", "Doors" , "Load(lbs)")
	fmt.Println("----------------------------------------------------------------------------------------------")

	for _,veh := range inventory {
		switch v := veh.(type) {
			case car:
				v.carDetails()
				carCount++
			case bike:
				v.bikeDetails()
				bikeCount++
			case truck:
				v.truckDetails()
				truckCount++
			default:
				fmt.Printf("Are you sure this Vehicle Type exists",)
		}
	}

}

func (c *car) carDetails() {
	msg :=	fmt.Sprintf("Car-%-11v%-12v%-12v%-10v%-6.2f%9v%9d%9d\t--\n", c.typeVehicle,c.make,c.model,c.color,c.cc,c.year,c.seating,c.doors) 
	c.printVehicleDetails(msg)
}

func (b *bike) bikeDetails() {
	msg := fmt.Sprintf("Bike           %-12v%-12v%-10v%-6.2f%9v\t--\t  --\t--\n", b.make,b.model,b.color,b.cc,b.year) 
	b.printVehicleDetails(msg)
}

func (t *truck) truckDetails() {
	msg := fmt.Sprintf("Truck-%-9v%-12v%-12v%-10v%-6.2f%9v%9d%9d%10.2f\n", t.typeVehicle,t.make,t.model,t.color,t.cc,t.year,t.seating,t.doors,t.loadingCapacity) 
	t.printVehicleDetails(msg)
}

func (c car) printVehicleDetails(s string) {
	fmt.Printf("%v", s)	
}

func (b bike) printVehicleDetails(s string) {
	fmt.Printf("%v", s)	
}

func (t truck) printVehicleDetails(s string) {
	fmt.Printf("%v", s)	
}



















