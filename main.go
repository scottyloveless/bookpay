package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Book Pay Calculator")
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("What grade reading level is your child?")
	scanner.Scan()
	childReadingLevel, err := strconv.Atoi(scanner.Text())
	if err != nil {
		fmt.Println("Please enter a number")
		return
	}

	fmt.Println("What reading level is the book?")
	scanner.Scan()
	bookReadingLevel, err := strconv.Atoi(scanner.Text())
	if err != nil {
		fmt.Println("Please enter a number")
		return
	}

	fmt.Println("How many pages are in the book?")
	scanner.Scan()
	pages, err := strconv.Atoi(scanner.Text())
	if err != nil {
		fmt.Println("Please enter a number")
		return
	}

	fmt.Println("How many chapters are in the book?")
	scanner.Scan()
	chapters, err := strconv.Atoi(scanner.Text())
	if err != nil {
		fmt.Println("Please enter a number")
		return
	}

	perChapter, wholeBook := calculatePay(childReadingLevel, bookReadingLevel, pages, chapters)

	fmt.Println("This book is worth:")
	fmt.Println(perChapter)
	fmt.Println(wholeBook)
}

func calculatePay(childRl, bookRl, numPages, numChapters int) (string, string) {
	var difficultyMultiplier float64
	readingLevelDiff := bookRl - childRl

	switch readingLevelDiff {
	case 0:
		difficultyMultiplier = 1.0
	case 1:
		difficultyMultiplier = 1.10
	case 2:
		difficultyMultiplier = 1.25
	case 3:
		difficultyMultiplier = 1.50
	default:
		difficultyMultiplier = -0.25
	}

	baseBookPrice := 5.00
	basePagePrice := 0.03
	priceForBook := baseBookPrice + (basePagePrice * difficultyMultiplier * float64(numPages))
	pricePerChapter := priceForBook / float64(numChapters)

	return fmt.Sprintf("Price per chapter: $%.2f", pricePerChapter), fmt.Sprintf("Price of entire book: $%.2f", priceForBook)
}
