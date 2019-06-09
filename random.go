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
	dataBasePath = "github.com/random-names/names"
)

const (
	indexName = iota
	indexCumulative
)

type name struct {
	name       string
	cumulative float64
}

// Options is the option passed to public functions.
type Options struct {
	// Max represents the maximum cumulative percentage of names to be generated.
	Max float64
	// Number represents the number of names to be generated.
	Number int
	// Real indicates whether to use realistic percentage to generate names.
	Real bool
}

// GetRandomName returns a random name from the given database or file.
func GetRandomName(path string, opt *Options) (string, error) {
	if opt.Number != 1 {
		opt.Number = 1
	}

	names, err := GetRandomNames(path, opt)
	if err != nil {
		return "", err
	}
	return names[0], nil
}

// GetRandomNames returns random names from the given database or file.
func GetRandomNames(path string, opt *Options) ([]string, error) {
	if opt.Max > 100 {
		opt.Max = 100
	} else if opt.Max < 0 {
		opt.Max = 0
	}
	if opt.Number <= 0 {
		opt.Number = 1
	}

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

	names := getRandomNames(data, opt)
	return names, nil
}

func getFile(path string) (*os.File, error) {
	gopath := os.Getenv("GOPATH")
	if gopath == "" {
		gopath = build.Default.GOPATH
	}

	file, err := os.Open(filepath.Join(gopath, "src", dataBasePath, path))
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
	if len(args) < indexCumulative+1 {
		diff := indexCumulative + 1 - len(args)
		for diff > 0 {
			args = append(args, "0")
			diff--
		}
	}

	n = &name{}
	n.name = args[indexName]

	value, err := strconv.ParseFloat(args[indexCumulative], 32)
	if err != nil {
		value = 0
	}
	n.cumulative = value
	return
}

func getRandomNames(data []*name, opt *Options) (names []string) {
	rand.Seed(int64(time.Now().UnixNano()))

	len := len(data)
	max := opt.Max
	maxCumulative := data[len-1].cumulative
	if max <= 0 || max > maxCumulative {
		max = maxCumulative
	}
	if len <= 0 {
		names = append(names, "Silly Data")
		return
	}

	var index int
	for opt.Number > 0 {
		if opt.Real {
			random := rand.Float64() * max
			index = sort.Search(len, func(i int) bool {
				return data[i].cumulative > random
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
