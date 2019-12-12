package main

import (
	"fmt"
	"math"
)

func main() {

	x := 0

	for {
		fmt.Println("\n**************Main Menu**********************")
		fmt.Println("Enter 1 for Area and 2 for Volume (0 to exit)")
		if _, err := fmt.Scan(&x); err == nil {
			switch x {
				case 1:
					y := 0
					for {
						fmt.Printf("\n*************Sub Menu******Calculate Area***************** ")
						fmt.Println("\n1 for Rectangle, 2 for Circle (9 to main menu, 0 to exit)")
						if _, err := fmt.Scan(&y); err == nil {
							switch y {
							case 1:
								areaRectangle()
							case 2:
								areaCircle()
							case 9:
								goto exitArea
							case 0:
								goto outerloop
							default:
								fmt.Println("Not a valid value. Enter Again")
							}
						}
		
					}
					exitArea:
						break
		
				case 2:
					y := 0
					for {
						fmt.Printf("\n*************Sub Menu******Calculate Volume***************** ")
						fmt.Println("\n1 for Cylinder, 2 for Cube (9 to main menu, 0 to exit)")
						if _, err := fmt.Scan(&y); err == nil {
							switch y {
							case 1:
								volumeCylinder()
							case 2:
								volumeCube()
							case 9:
								goto exitVolume
							case 0:
								goto outerloop
							default:
								fmt.Println("Not a valid value. Enter Again")
							}
						}
					
					}
					exitVolume:
						break
				case 0:
					goto outerloop
				default:
					fmt.Println("Not a valid value. Enter Again")
				
	
			}
			
		}

	}
outerloop:
}

func areaCircle() {
	var r float32
	fmt.Println("Enter the radius of the circle: ")
	if _, err := fmt.Scan(&r); err == nil {
		fmt.Printf("The Area of the circle of radius %5.2f is: %5.2f\n" ,r,math.Pi*r*r )
	}
}

func areaRectangle() {
	var l,w float32
	fmt.Println("Enter the length and width of rectangle: ")
	if _, err := fmt.Scan(&l,&w); err == nil {
		fmt.Printf("The Area of the rectangle %5.2f x %5.2f is: %5.2f\n" ,l,w,l*w)
	}

}


func volumeCylinder() {
	var r,h float32
	fmt.Println("Enter the radius and height of cylinder: ")
	if _, err := fmt.Scan(&r,&h); err == nil {
		fmt.Printf("The Volume of the cylinder is: %5.2f\n" ,math.Pi*r*r*h)
	}

}

func volumeCube() {
	var e float32
	fmt.Println("Enter the edge of the cube: ")
	if _, err := fmt.Scan(&e); err == nil {
		fmt.Printf("The Volume of the cube of edge %5.2f is: %5.2f\n" ,e,math.Pow(float64(e),3)) 
	}

}
