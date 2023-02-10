package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
	"unicode"
	"unicode/utf8"
)

/*
+---+
    |
    |
    |
   ===

Secret Word : ______
Incorrect Guesses :

Guess a Letter : a

Sorry Your Dead! The word is ZOMBIE
Yes the Secret Word is ZOMBIE

Please Enter Only One Letter
Please Enter a Letter
Please Enter a Letter you Haven't Guessed
*/

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

// Stores the random word to be guessed
var randWord string

// Stores all letters guessed
var guessedLetters []string

// Stores correct guesses
var correctLetters []string

// Letters guessed that aren't in the randWord
var wrongGuesses []string

func generateRandomWord() string {
	seedSecs := time.Now().Unix()
	rand.Seed(seedSecs)

	return wordArr[rand.Intn(7)+1]
}
func isLetter(s string) bool {
	for _, r := range s {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return true
}
func isGuessedBefore(guessedLetters []string, letter string) bool {

	for _, value := range guessedLetters {
		if value == letter {
			return true
		}
	}
	return false
}

func main() {
	randWord = generateRandomWord()
	fmt.Println("ADAM ASMACA OYUNUNA HOŞGELDİNİZ")

	fmt.Println(hmArr[0])
	reader := bufio.NewReader(os.Stdin)

	for true {
		fmt.Print("Harf tahmini giriniz : ")
		letter, err := reader.ReadString('\n')
		letter = strings.Trim(letter, " ")

		if err == nil {
			if utf8.RuneCountInString(letter) > 1 {
				fmt.Println(utf8.RuneCountInString(letter))
				fmt.Println("Tek karakter giriniz")
			} else if !isLetter(letter) {
				fmt.Println("Harf girmeniz gerekiyor")
			} else if isGuessedBefore(guessedLetters, letter) {
				fmt.Println("Bu karakteri daha önce girdiniz")
			} else {
				guessedLetters = append(guessedLetters, letter)
			}
		} else {
			fmt.Println(err)
		}
	}

}
