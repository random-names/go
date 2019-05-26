package randomnames

import (
	"bufio"
	"go/build"
	"math/rand"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"
)

const (
	DATA_BASE_PATH = "github.com/random-names/names"
)

const (
	INDEX_NAME = iota
	INDEX_WEIGHT
	INDEX_CUMMULATIVE
	INDEX_RANK
)

type name struct {
	name        string
	weight      float64
	cummulative float64
	rank        int
}

type Options struct {
	Number int
	Real   bool
	Max    float64
}

// GetFromDatabase returns random names from the given database or file
func GetRandomNames(path string, opt *Options) ([]string, error) {
	names := []string{}

	file, err := getFile(path)
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

func getFile(path string) (*os.File, error) {
	gopath := os.Getenv("GOPATH")
	if gopath == "" {
		gopath = build.Default.GOPATH
	}

	file, err := os.Open(filepath.Join(gopath, "src", DATA_BASE_PATH, path))
	if err == nil {
		return file, nil
	}

	path, err = filepath.Abs(path)
	if err != nil {
		return nil, err
	}

	file, err = os.Open(path)
	if err != nil {
		return nil, err
	}

	return file, err
}

func getNameStruct(data string) (n *name) {
	args := strings.Fields(data)
	if len(args) < INDEX_RANK+1 {
		diff := INDEX_RANK + 1 - len(args)
		for diff > 0 {
			args = append(args, "0")
			diff--
		}
	}

	n = &name{}
	n.name = args[INDEX_NAME]

	var value float64
	var intValue int
	var err error
	value, err = strconv.ParseFloat(args[INDEX_WEIGHT], 32)
	if err != nil {
		value = 0
	}
	n.weight = value
	value, err = strconv.ParseFloat(args[INDEX_CUMMULATIVE], 32)
	if err != nil {
		value = 0
	}
	n.cummulative = value
	intValue, err = strconv.Atoi(args[INDEX_RANK])
	if err != nil {
		intValue = 0
	}
	n.rank = intValue
	return
}

func getRandomNames(data []*name, opt *Options) (names []string) {
	rand.Seed(int64(time.Now().UnixNano()))

	len := len(data)
	max := opt.Max
	maxCummulative := data[len-1].cummulative
	if max <= 0 || max > maxCummulative {
		max = maxCummulative
	}
	if len <= 0 {
		names = append(names, "Silly Data")
		return
	}

	var index int
	if opt.Number <= 0 {
		opt.Number = 1
	}
	for opt.Number > 0 {
		if opt.Real {
			random := rand.Float64() * max
			index = sort.Search(len, func(i int) bool {
				return data[i].cummulative > random
			})
		} else {
			index = rand.Intn(len)
		}
		if index < 0 || index >= len {
			index = 0
		}
		names = append(names, data[index].name)
		opt.Number--
	}
	return
}
