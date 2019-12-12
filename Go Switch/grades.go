package main

import (
	"fmt"
	"time"
	"math/rand"
	"sort"
)

const totalDays = 192

type student struct {
	name string
	grade rune
	percentage float32
	attDays int
	courses course
}

type course struct {
	dataLogic float32
	algorithm float32
	calculus float32
	finalProject int	
}

//type byPercent [12]student

type studentList [12]student

func (s *studentList) generateMarks() {
	for i:=0 ; i<12 ; i++ {
		name := fmt.Sprintf("Student%v",i+1)
		s[i] = student{name ,' ' ,0 ,82+rand.Intn(111) ,course{50 + float32(rand.Intn(51)),50 + float32(rand.Intn(51)),50 + float32(rand.Intn(51)), 8 + rand.Intn(18) } } 
	}

}

func printStudentDetails(s []student) {
	fmt.Println("\nStudent Name\tGrade\tPercentage\tAttendance\tProject")
	for i:=0 ; i<12 ; i++ {
		fmt.Printf("%-10v\t%c\t%8.3v\t%10v\t%7v\n",s[i].name,s[i].grade,s[i].percentage,s[i].attDays,s[i].courses.finalProject)
	}


}

func calculate(s student) (float32 , rune) {
	var percent float32
	grade := 'F'
	
	/*
	if (percent >= 90.0) {
		grade = 'A'
	} else if (percent>=80.0 && percent < 90.0) {
		grade = 'B'
	} else if (percent>=70.0 && percent < 80.0) {
		grade = 'C'
	} else if (percent>=60.0 && percent < 70.0) {
		grade = 'D'
	} else {
		grade = 'F'
	} 
	*/

	switch percent = (s.courses.algorithm + s.courses.calculus + s.courses.dataLogic) /3 ; {
		case percent > 90.0:
			grade = 'A'
		case percent>=80.0 && percent < 90.0:
			grade = 'B'
		case percent>70.0 && percent < 80.0:
			grade = 'C'
		case percent>=60.0 && percent < 70.0:
			grade = 'D'
		default:
			grade = 'F'	
	
		}


	return percent , grade
}

func runBusinessLogic(s []student) {
	var grade rune
	for i:=0; i<12 ; i++ {
		//grade := 'T'
		// Calculate the percentage and the grade.
		s[i].percentage , grade = calculate(s[i])
	
		// Check Attendance
		if grade != 'F' {
			calAttendance(s[i].attDays , &grade)
		}
		// Check Final Project
		grade = calProject(s[i].courses.finalProject , grade)
		
		// Assign the grade to the student grade
		s[i].grade = grade
	}

}

func calProject(marks int, grade rune) rune {
	//if >= 80% then grade+1 
	// 40-80 no change
	// < 40 grade-1
	percent := float32(marks)/25 * 100
	//fmt.Println(percent)

	if percent >= 80 && grade != 'A' {
		grade--
	} else if percent < 40 && grade != 'F' {
		grade++
	}

	return grade
}

func calAttendance(attDays int , grade *rune ) {
//	fmt.Printf("\t%v: %c",s.name,s.grade)
	if float32(attDays)/(totalDays) < 0.75 {
		*grade++
//		fmt.Printf("\t%v: %c\n",s.name,s.grade)
		
	}
	
}

func main() {

rand.Seed(time.Now().UnixNano())

var students studentList
students.generateMarks()

ss := students[:]

runBusinessLogic(ss)

sort.Slice(ss , func(i,j int) bool {
	return ss[i].grade < ss[j].grade 
})

printStudentDetails(ss)

}

/*										
Grades				Attendance						Final Project
										
Avg marks from 3 subjects										
				if greater than 75% no change						
				If less than 75% , lower grade by 1						If Final Project marks out of 25… if >80% then grade+1 
										40 -80 % then no change
				Can't go lower than F						if less than 40% then lower grade by 1
*/