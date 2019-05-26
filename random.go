package main

import (
	"bufio"
	"math/rand"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
)

const (
	DATA_BASE_URL = "github.com/random-names/names"
)

const (
	INDEX_NAME = iota
	INDEX_WEIGHT
	INDEX_CUMMULATIVE
	INDEX_RANK
)

type name struct {
	name        string
	weight      float32
	cummulative float32
	rank        int
}

type options struct {
	number int
	real   bool
	max    int
}

// GetFromDatabase returns random names from the given database
// func GetFromDatabase(data string, opt *options) ([]string, error) {
// 	// names = getRandomNames(data, opt)
// 	return
// }

// GetFromDatabase returns random names from the given file
func GetFromFile(filepath string, opt *options) ([]string, error) {
	names := []string{}
	wd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	file, err := os.Open(path.Join(wd, filepath))
	if err != nil {
		return nil, err
	}
	defer file.Close()

	data := []*name{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		n := getNameStruct(scanner.Text())
		data = append(data, n)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	names = getRandomNames(data, opt)
	return names, nil
}

func getNameStruct(data string) (n *name) {
	args := strings.Fields(data)
	// log.Printf("d: %#v,   %d", args, INDEX_NAME)
	n = &name{}
	n.name = args[INDEX_NAME]

	var value float64
	var intValue int
	var err error
	value, err = strconv.ParseFloat(args[INDEX_WEIGHT], 32)
	if err != nil {
		value = 0
	}
	n.weight = float32(value)
	value, err = strconv.ParseFloat(args[INDEX_CUMMULATIVE], 32)
	if err != nil {
		value = 0
	}
	n.cummulative = float32(value)
	intValue, err = strconv.Atoi(args[INDEX_RANK])
	if err != nil {
		intValue = 0
	}
	n.rank = intValue
	return
}

func getRandomNames(data []*name, opt *options) (names []string) {
	rand.Seed(int64(time.Now().UnixNano()))

	len := len(data)
	max := opt.max
	if max <= 0 || max > len {
		max = len - 1
	}
	if len <= 0 {
		names = append(names, "Silly Data")
		return
	}

	if opt.real == true {
		// for _, name := range data {

		// }
	} else {
		for opt.number > 0 {
			index := rand.Intn(max)
			names = append(names, data[index].name)
			opt.number--
		}
	}
	return
}
