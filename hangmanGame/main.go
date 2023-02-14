package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"
)

var hmArr = [7]string{
	" +---+\n" +
		"     |\n" +
		"     |\n" +
		"     |\n" +
		"    ===\n",
	" +---+\n" +
		" 0   |\n" +
		"     |\n" +
		"     |\n" +
		"    ===\n",
	" +---+\n" +
		" 0   |\n" +
		" |   |\n" +
		"     |\n" +
		"    ===\n",
	" +---+\n" +
		" 0   |\n" +
		"/|   |\n" +
		"     |\n" +
		"    ===\n",
	" +---+\n" +
		" 0   |\n" +
		"/|\\  |\n" +
		"     |\n" +
		"    ===\n",
	" +---+\n" +
		" 0   |\n" +
		"/|\\  |\n" +
		"/    |\n" +
		"    ===\n",
	" +---+\n" +
		" 0   |\n" +
		"/|\\  |\n" +
		"/ \\  |\n" +
		"    ===\n",
}

var wordArr = [7]string{
	"JAZZ", "ZIGZAG", "ZILCH", "ZIPPER",
	"ZODIAC", "ZOMBIE", "FLUFF",
}

var randWord string
var guessedLetters []string
var correctLetters []string
var correctIndexes []int
var wrongGuesses []string

var isLetter = regexp.MustCompile(`^[a-zA-Z]+$`).MatchString

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func generateRandomWord() string {
	seedSecs := time.Now().Unix()
	rand.Seed(seedSecs)

	return wordArr[rand.Intn(7)]
}

func isGuessedBefore(guessedLetters []string, letter string) bool {

	for _, value := range guessedLetters {
		if value == letter {
			return true
		}
	}
	return false
}

func printGame() {
	fmt.Print("\n\n")
	fmt.Println(hmArr[len(wrongGuesses)])

	for i := 0; i < len(randWord); i++ {
		if contains(correctIndexes, i) {
			fmt.Print(string(randWord[i]))
		} else {
			fmt.Print("_")
		}
	}
	fmt.Print("\n")
}

func getUserGuess() {
	reader := bufio.NewReader(os.Stdin)
	for true {
		fmt.Print("Harf tahmini giriniz : ")
		letter, err := reader.ReadString('\n')

		if err != nil {
			fmt.Println(err)
		} else {
			letter = strings.ToUpper(letter)
			letter = strings.TrimSpace(letter)

			if utf8.RuneCountInString(letter) > 1 {
				fmt.Println("Tek karakter giriniz")
			} else if !isLetter(letter) {
				fmt.Println("Harf girmeniz gerekiyor")
			} else if isGuessedBefore(guessedLetters, letter) {
				fmt.Println("Bu karakteri daha önce girdiniz")
			} else {
				guessedLetters = append(guessedLetters, letter)
				if strings.Contains(randWord, letter) {
					correctLetters = append(correctLetters, letter)

					for i := 0; i < len(randWord); i++ {
						if string(randWord[i]) == letter {
							correctIndexes = append(correctIndexes, i)
						}
					}
				} else {
					wrongGuesses = append(wrongGuesses, letter)
				}
				break
			}
		}
	}
}

func main() {
	randWord = generateRandomWord()
	fmt.Println("ADAM ASMACA OYUNUNA HOŞGELDİNİZ")
	//fmt.Println(randWord)

	for true {
		printGame()
		getUserGuess()

		if len(wrongGuesses) == 7 {
			fmt.Println("Kaybettiniz")
			break
		}
		if len(correctIndexes) == len(randWord) {
			fmt.Println("Tebrikler")
			break
		}
	}
}
