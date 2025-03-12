// go mod init github.com/xcheng85/blind75-cpp-golang/golang/sandbox
// go mod tidy
package main

import (
	"encoding/csv"
	_ "encoding/csv"
	"encoding/json"
	"fmt"
	"maps"
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

func testRTII(x interface{}) {
	// must within switch
	// T := x.(type)

	switch T := x.(type) {
	case Player:
		fmt.Println("Player type")
		// default: break c++
		// fallthrough: not supported in type switch
	default:
		fmt.Println("Unknown type: ", T)
	}
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
	// sort player by jersey number
	sort.Sort(players)
	// rtii
	testRTII(players[0])

	writeCSV("bucks.csv")

	readCSV("bucks.csv")

	// list
	l := core.NewList[string]()
	l.PushFront("Giannis")
	l.PushFront("Dame")
	l.Debug()

	// shallow copy
	// different memory spaces
	playersCopy := slices.Clone(players)
	players[0].Number = 99
	fmt.Println(playersCopy)

	// search, very similar to std::find
	fmt.Println(slices.ContainsFunc(playersCopy, func(p Player) bool {
		return p.Number == 34
	}))

	// erase algo in go
	playerJerseyMap := map[string]int{
		"Giannis": 34,
		"Dame":    0,
	}

	maps.DeleteFunc(playerJerseyMap, func(name string, jersey int) bool {
		return name == "Giannis"
	})

	fmt.Println(playerJerseyMap)

	// map copy with separate memory
	playerJerseyMapCopy := maps.Clone(playerJerseyMap)
	playerJerseyMap["Bobby"] = 9
	fmt.Println(playerJerseyMapCopy)

	// json parser
	js := make(map[string]interface{})
	//js := map[string]interface{}{}
	err := json.Unmarshal([]byte(core.TestJson), &js)
	if err != nil {
		fmt.Println(err)
	} else {
		core.Probe(js)
		core.TypeSwitch(js)
	}

}
