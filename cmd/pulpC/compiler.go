package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)


func main () {
    
    var file_data []byte
    var file *os.File
    var path string = "./tests/hello_world.pulp"

    path, err := filepath.Abs(path)
    if err != nil {
        log.Fatalf("Error converting to abs path: %s", err.Error())
    }

    file, err = os.Open(path)
    if err != nil {
        log.Fatalf("Error opening file: %s", err.Error())
    }

    defer file.Close()

    var scanner *bufio.Scanner = bufio.NewScanner(file)

    var line []byte
    var i int8 = 1
    for scanner.Scan() {
        line = []byte(strings.TrimSpace(scanner.Text()))
        if len(line) > 0 {
            // fmt.Printf("%d:%s\n", i, line)
            file_data = append(file_data, line...)
        }
        i++
    }

    fmt.Printf("%s\n", file_data)
}
