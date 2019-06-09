# go-random-names

[![GoDoc](https://godoc.org/github.com/random-names/go?status.svg)](https://godoc.org/github.com/random-names/go)
![License](https://img.shields.io/badge/License-MIT-blue.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/random-names/go)](https://goreportcard.com/report/github.com/random-names/go)

A random names generator written in Golang.

### Installation

```
$ go get github.com/random-names/go
```

### Usage

#### Using in your code

```go
package main

import (
	"fmt"
	rn "github.com/random-names/go"
)

func main() {
	firstName, err := rn.GetRandomName("census-90/male.first", &rn.Options{})
	if err != nil {
		fmt.Println(err)
	}

	lastName, err := rn.GetRandomName("census-90/all.last", &rn.Options{})
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("random name: %v \n", firstName+" "+lastName)
}

```

The above example will return names from `census-90/male.first(also all.last)` database in [names](https://github.com/random-names/names/tree/master/census-90). If you wish to use the database from this repository, get it first!

```
go get github.com/random-names/names
```

You can also use database not from this repository. Just pass the relative path:

```go
name, err := rn.GetRandomName("relative/path/to/database", &rn.Options{})
```

To get multiple names, switch to **`GetRandomNames`** and specify the option **`Number`**:

```go
package main

import (
	"fmt"
	rn "github.com/random-names/go"
	"strings"
)

func main() {
	lastNames, err := rn.GetRandomNames("census-90/all.last", &rn.Options{Number: 5})
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("random last-names: %v \n", strings.Join(lastNames, ", "))
}

```

Want names to be more realistic? You only need to enable the **`Real`** option!
*Note: Only database provides cumulative percentage of each name support this feature. See [here](https://github.com/random-names/names/tree/master/README.md#database-structure).*

```go
name, err := rn.GetRandomName("relative/path/to/database", &rn.Options{Real: true})
```

The last option is **`Max`**, which is used along with `Real`. For example, if you set `Max` to 50 while `Real` enabled, you will always get names those are the most popular 50% ones.
*Note: You cannot set the value of `Max` larger than the maximum percentage in the database.*

#### Running in command line

You should install the `go-rn` binary first. In your `$GOPATH/src/github.com/random-names/go` folder, run the following command:

```
$ go install ./...
```

Then you can run it. The example will pick a name from `census-90/male.first` randomly:

```
$ go-rn census-90/male.first
```

Check the usage here:

```
USAGE:
   go-rn [global options] command [command options] [arguments...]

COMMANDS:
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --max value, -m value     the maximum of the random number (default: 0)
   --number value, -n value  how many names to generate (default: 1)
   --real, -r                use the real percentage
   --help, -h                show help
   --version, -v             print the version
```
