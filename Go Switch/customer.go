package main

import (
	"fmt"
	"os"
	"bufio"
	"log"
	"strings"
)

type rating float32

const (
	extraPositive rating = 1.2
	positive      rating = 0.6
	negative      rating = -0.6
	initial       rating = 5.0
	extraNegative rating = -1.2
)

type result struct {
	feedbackDate	string
	feedbackTotal	int
	feedbackPositive	int
	feedbackNegative	int
	feedbackNeutral		int
}

var customer []rating


func main() {

f,err := os.Open("data.csv")
if err!= nil {
	exitProgram(err.Error())
}

defer f.Close()

r :=bufio.NewReader(f)

str, err := r.ReadString('\n')
if err!= nil {
	exitProgram(err.Error())
}

var feedback result
feedback.feedbackDate = str
for {
	str, err := r.ReadString('\n')
	if err!= nil {
		break
	}

	if len(str) > 10 {
		feedback.feedbackTotal++

		var custRating rating
		custRating = 5.0
		text := strings.Split(str , " ")

		for _,word := range text {
			switch s:= strings.Trim(strings.ToLower(word)," ,.,!,?,\t,\n,\r") ; s {
				case "pleasure", "impressed", "wonderful", "fantastic", "splendid":
					custRating += extraPositive
				case "help", "helpful", "thanks", "happy":
					custRating += positive
				case "not helpful", "sad", "angry", "improve", "annoy":
					custRating += negative
				case "pathetic", "bad", "worse", "unfortunately", "agitated", "frustrated":
					custRating += extraNegative
			}
			
		}

		switch {
			case custRating > 8.0:
				feedback.feedbackPositive++
			case custRating >= 4.0 && custRating <= 8.0:
				feedback.feedbackNeutral++
			case custRating < 4.0:
				feedback.feedbackNegative++
		}

		
		customer = append(customer,custRating)
	}

}

feedbackTable(feedback,customer)


}

func exitProgram(s string) {
	log.Fatal("Exiting the program with: ", s)
}

func feedbackTable(f result , c []rating) {
	fmt.Printf("Report for : %s",f.feedbackDate)
	fmt.Printf("Total Customer Reviews: %d\n", f.feedbackTotal)
	fmt.Printf("Positive Reviews: %d\n", f.feedbackPositive)
	fmt.Printf("Negative Reviews: %d\n", f.feedbackNegative)
	fmt.Printf("Neutral Reviews: %d\n", f.feedbackNeutral)
	fmt.Println("Customer Ratings: ", c)

}