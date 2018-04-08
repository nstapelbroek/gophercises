package main

import (
	"flag"
	"os"
	"encoding/csv"
	"errors"
	"bufio"
	"fmt"
	"strings"
)

type Game struct {
	questions                []CsvQuestionRow
	questionsAnsweredCorrect int
}

type CsvQuestionRow struct {
	question string
	answer   string
}

type Options struct {
	filename *string
}

func main() {
	programOptions := initOptions()
	questions, err := loadQuestionsFromFile(programOptions.filename)
	if (err != nil) {
		panic("could not load questions from file: " + err.Error())
	}

	game := Game{
		questions,
		0,
	}
	game = runGame(game)
	fmt.Println(fmt.Sprintf("You answered %d of %d questions correctly", game.questionsAnsweredCorrect, len(game.questions)))
}

func runGame(game Game) (Game) {
	for index := range game.questions {
		fmt.Print(game.questions[index].question + "= ")
		givenAnswer, _ := getInputFromCli()
		correctAnswer := game.questions[index].answer

		if strings.Compare(givenAnswer, correctAnswer) == 0 {
			fmt.Println("Correct!")
			game.questionsAnsweredCorrect++
			continue
		}

		fmt.Println("Wrong, correct answer was: " + correctAnswer)
	}

	return game
}

func getInputFromCli() (givenAnswer string, err error) {
	reader := bufio.NewReader(os.Stdin)
	givenAnswer, err = reader.ReadString('\n')
	if err != nil {
		return
	}

	givenAnswer = strings.Replace(givenAnswer, "\n", "", -1)
	return
}

func initOptions() (options Options) {
	options.filename = flag.String("filename", "./quiz/problems.csv", "The locations of the problems csv")
	flag.Parse()

	return
}

func loadQuestionsFromFile(filename *string) (questions []CsvQuestionRow, err error) {
	csvFile, err := os.Open(*filename)
	if (err != nil) {
		return
	}

	reader := csv.NewReader(csvFile)
	records, err := reader.ReadAll()
	if (err != nil) {
		return
	}

	questions = make([]CsvQuestionRow, len(records))

	for index, question := range records {
		if len(question) != 2 {
			err = errors.New("incompatible CSV format with more or less than 2 fields")
			return
		}

		questions[index] = CsvQuestionRow{
			question[0],
			question[1],
		}
	}

	return
}
