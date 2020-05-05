package main

import (
	"errors"
	"fmt"
	"hash/fnv"
	"os"
)

func main() {
	name := os.Args[1]
	h, _ := hashtoport(name, 0, 0)
	fmt.Println(h)
}

func hashtoport(name string, from, to uint32) (uint32, error) {
	if from == 0 {
		from = 1024
	}
	if to == 0 {
		to = 65535
	}

	if from >= to {
		return 0, errors.New("'to' must be greater or equal to than 'from'")
	}

	if from < 1024 || to > 65535 {
		return 0, errors.New("port numbers must be between 1024 and 65535")
	}

	return hashindex(name, to-from) + from, nil
}

func hashindex(s string, max uint32) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32() % max
}
