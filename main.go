package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	model "blixza.github.io/model/record"
	"blixza.github.io/pkg"
)

func main() {
	path := "library.json"

	file := pkg.Must(pkg.Load(path))

	var artist []string
	var album string
	var rating string
	var saved string
	var bestSongs []string
	var worstSongs []string
	var thoughts string

	fmt.Println("Enter artist(s) separated by comma and space.")
	pkg.Must("", Ask(&artist))
	fmt.Println()

	fmt.Println("Enter album title.")
	pkg.Must("", Ask(&album))
	fmt.Println()

	fmt.Println("Enter the album rating. Ex: 8/10")
	pkg.Must("", Ask(&rating))
	fmt.Println()

	fmt.Println("Enter how many tracks you saved.")
	pkg.Must("", Ask(&saved))
	fmt.Println()

	fmt.Println("Enter the best songs seperated by comma and space.")
	pkg.Must("", Ask(&bestSongs))
	fmt.Println()

	fmt.Println("Enter the worst songs seperated by comma and space.")
	pkg.Must("", Ask(&worstSongs))
	fmt.Println()

	fmt.Println("Enter your thoughts about the album.")
	pkg.Must("", Ask(&thoughts))
	fmt.Println()

	record := model.Record{
		Date:       time.Now().Format("2006-01-02"),
		Artist:     artist,
		Album:      album,
		Rating:     rating,
		Saved:      saved,
		BestSongs:  bestSongs,
		WorstSongs: worstSongs,
		Thoughts:   thoughts,
	}

	pkg.Must("", pkg.Write(file, record))
}

func Ask[T any](target T) error {
	scanner := bufio.NewScanner(os.Stdin)

	if !scanner.Scan() {
		err := scanner.Err()
		if err != nil {
			return err
		}
		return fmt.Errorf("no input provided")
	}

	input := scanner.Text()

	switch ptr := any(target).(type) {
	case *string:
		*ptr = input
	case *[]string:
		*ptr = strings.Split(input, ", ")
	default:
		return fmt.Errorf("unsupported type: %T", target)
	}

	return nil
}
