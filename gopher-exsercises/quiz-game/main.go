package main

// https://gophercises.com/exercises/quiz

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	generateCSV()
	startQuiz()
}

func generateCSV() {
	data := make([]string, 0)
	for i := 1; i < 10; i++ {
		a := rand.Intn(10)
		b := rand.Intn(10)
		str := strconv.Itoa(a) + " + " + strconv.Itoa(b) + "," + strconv.Itoa(a+b) + "\n"
		data = append(data, str)
	}
	writeToFile(data)

}

func writeToFile(data []string) error {
	file, err := os.Create("problems.csv")
	if err != nil {
		return err
	}
	defer file.Close()

	for _, str := range data {
		_, err = io.WriteString(file, str)
		if err != nil {
			return err
		}
	}
	return file.Sync()
}

func readFile() []string {
	quiz := make([]string, 0)
	file, err := os.Open("problems.csv")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		//fmt.Println("Whoa!", scanner.Text())
		quiz = append(quiz, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
	return quiz
}

func startQuiz() {
	start := time.Now()
	quiz := readFile()
	reader := bufio.NewReader(os.Stdin)
	rightNumb := 0
	wrongNumb := 0
	for _, q := range quiz {
		tmp := strings.Split(q, ",")
		answer, err := strconv.Atoi(tmp[1])
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("Task:", tmp[0], "?")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSuffix(input, "\n")
		input = strings.TrimSuffix(input, "\r")

		//userAnswer, err := strconv.Atoi(input)

		userAnswer := 3

		if err != nil {
			fmt.Println(err)
		}
		elapsed := time.Since(start)
		if userAnswer == answer {
			fmt.Println("Right!", elapsed)
			rightNumb++
		} else {
			fmt.Println("Wrong!", elapsed)
			wrongNumb++
		}

		if elapsed > time.Duration(30)*1e9 {
			fmt.Println("time is out!", elapsed, time.Duration(30)*1e9)
		} else {
			fmt.Println("still have time", elapsed, time.Duration(30)*1e9)
		}
	}

	fmt.Println("Your right answers are: ", rightNumb, " , and wrong are:", wrongNumb)
	//elapsed := time.Since(start)
	//fmt.Println("Time spent:", elapsed)
}
