// go mod init github.com/xcheng85/blind75-cpp-golang/golang/sandbox
// go mod tidy
package main

import (
	"encoding/csv"
	_ "encoding/csv"
	"fmt"
	"math"
	"math/rand"
	"os"
	"slices"
	"sort"
	"strconv"
	"time"
	"github.com/xcheng85/blind75-cpp-golang/golang/sandbox/core"
)

const (
	ZERO = iota
	One
	Two
)

const (
	POW2_0 = 1 << iota
	_
	POW2_2
)

// slice is ref by nature
func printSlice(input []string) {
	for i, s := range input {
		fmt.Println(i, s)
	}
}

func copySlice(src []string, dst []string) {
	copy(dst, src)
}

func generateRandomInt() int {
	seed := time.Now().Unix()
	r := rand.New(rand.NewSource(seed))
	fmt.Println(r.Uint64())
	return r.Intn(100-0) + 0
}

// table 2d
func readCSV(path string) ([][]string, error) {
	_, err := os.Stat(path)
	if err != nil {
		return nil, err
	}
	fd, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer fd.Close()

	lines, err := csv.NewReader(fd).ReadAll()
	if err != nil {
		return [][]string{}, err
	}
	fmt.Println(lines)
	return lines, nil
}

func writeCSV(path string) error {
	fd, err := os.Create(path)
	if err != nil {
		return err
	}
	defer fd.Close()
	csvW := csv.NewWriter(fd)

	csvW.Comma = '\t'
	for _, p := range players {
		// string() is not right
		err = csvW.Write([]string{p.First, p.Last, strconv.FormatInt(int64(p.Number), 10)})
		if err != nil {
			return err
		}
	}
	csvW.Flush()
	return nil
}

func main() {

	fmt.Println(math.MaxInt)
	fmt.Println(time.Now().Unix())
	fmt.Println(time.Now().Local())

	fmt.Println(One)
	fmt.Println(POW2_2)

	input := []string{"giannis", "dame", "khris", "bobby"}
	printSlice(input)
	// giannis is deleted
	// [0, 1)
	input = slices.Delete(input, 0, 1)
	printSlice(input)

	dst2 := make([]string, len(input))
	copySlice(input, dst2)
	printSlice(dst2)

	// sort string
	sort.Strings(dst2)
	printSlice(dst2)

	// generic sort.Interface
	// reverse sort
	sort.Sort(sort.Reverse(sort.StringSlice(dst2)))
	printSlice(dst2)

	fmt.Println(generateRandomInt())

	// write to csv
	players = append(players, []Player{{"Giannis", "Antetokounmpo", 34}, {"Dame", "Lillard", 0}}...)
	writeCSV("bucks.csv")

	readCSV("bucks.csv")

	// list
	l := core.NewList[string]()
	l.PushFront("Giannis")
	l.PushFront("Dame")
	l.Debug()
}
