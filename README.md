# go-random-names

A random human names generator written in Golang.

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
