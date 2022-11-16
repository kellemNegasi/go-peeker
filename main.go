package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/rodaine/table"
)

func main() {
	// cmd.Execute()
	file_path := "./data/effects-of-covid-19-on-trade-at-15-december-2021-provisional.csv"
	err := read(file_path)
	if err != nil {
		panic(err)
	}
}
func read(file string) error {
	in, err := os.Open(file)
	if err != nil {
		fmt.Println("error opening file")
	}
	defer in.Close()
	scanner := bufio.NewScanner(in)
	if err != nil {
		return err
	}
	headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
	columnFmt := color.New(color.FgYellow).SprintfFunc()
	scanner.Scan()
	h := strings.Split(scanner.Text(), ",")
	header := convert(h)
	tbl := table.New(header...)
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)
	for i := 1; i < 10; i++ {
		scanner.Scan()
		text := scanner.Text()
		r := strings.Split(text, ",")
		row := convert(r)
		tbl.AddRow(row...)
	}
	tbl.Print()
	return nil
}

func convert(list []string) []interface{} {
	s := make([]interface{}, len(list))
	for i, v := range list {
		s[i] = v
	}
	return s
}
