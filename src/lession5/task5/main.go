package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"structs"
)

const enoughArgsCount = 3 // filename pattern where -A n -B n

func main() {

	CommandLine := flag.NewFlagSet(os.Args[2], flag.ExitOnError)
	b := CommandLine.Int("B", 0, "Display n strings before the found string")
	a := CommandLine.Int("A", 0, "Display n strings after the found string")
	CommandLine.Parse(os.Args[3:])

	beforeCount := *b
	afterCount := *a

	//раз вариант упрощенный, - в одном файле будем искать и из stdin тоже не читаем
	if len(os.Args) < enoughArgsCount {

		log.Fatal("Please use the next pattern:\n >this_commend pattern filename [-A n] [-B n]")
	}

	pattern := os.Args[1]
	fileName := os.Args[2]

	if pattern == "" || fileName == "" {
		log.Fatal("Please specify the pattern and the filename")
	}

	// самодельные структурки. Первая содержит n последних считанных строк
	buffer := structs.NewCappedStringCollection(beforeCount + 1)
	// вторая содержит слайс совпадений и по ходу чтения файла дописывает им n строк,
	// лежащих после найденной строки. Так реализованы флаги -A и -B.
	matches := structs.NewMatchesStorage(afterCount)

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal("Can't open the file")
	}
	defer file.Close()

	lineReader := bufio.NewReader(file)

	for {
		line, err := lineReader.ReadString('\n')
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}

		isMatch := strings.Contains(line, pattern)
		before := buffer.GetSavedStrings()

		buffer.AddString(line)

		res := matches.Process(line, isMatch, before)
		displayMatched(res)
	}
	// На момент конца файла мы, скорее всего, имеем еще не завершенные Match,
	// у которых IsDone == false. Выключаем их и возвращаем, что они успели найти
	res := matches.Finish()
	displayMatched(res)

	println()
	log.Println("Successfully completed")
}

func displayMatched(m []*structs.Match) {

	for _, v := range m {
		res := v.Result()
		if len(res) > 1 {
			fmt.Println(strings.Repeat("------", 20))
		}
		for _, line := range res {
			fmt.Print(line)
		}
	}
}
