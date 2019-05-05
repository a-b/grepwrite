package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	// if err := run(os.Stdin); err != nil {
	// 	fmt.Fprintln(os.Stderr, err)
	// 	os.Exit(1)
	// }

	// input, output := io.Pipe()
	//
	// go func() {
	// 	defer output.Close()
	// 	for _, m := range []byte("123456") {
	// 		output.Write([]byte{m})
	// 		time.Sleep(time.Second)
	// 	}
	// }()

	input := os.Stdin

	scanner := bufio.NewScanner(input)
	if err := scanner.Err(); err != nil {
		log.Println(err)
	}

	for scanner.Scan() {
		text := scanner.Text()
		fmt.Printf("--> %v\n", text)

		fmt.Printf("input: %s\n", text)

		// {filename}|{lnum}[ col {col}][ {type} [{nr}]]| {text}
		data := strings.Split(text, " ")

		location, text := data[0], data[1]

		locationData := strings.Split(location, ":")

		file, lineNumber, _ := locationData[0], locationData[1], locationData[2]

		// fmt.Printf("--> %q\n", file)
		// fmt.Printf("--> %q\n", lineNumber)
		// fmt.Printf("--> %q\n", columnNumber)
		// fmt.Printf("--> %q\n", text)

		fi, err := os.Lstat(file)
		if err != nil {
			log.Fatal(err)
		}

		permissions := fi.Mode().Perm()

		fmt.Printf("permissions: %#o\n", permissions)

		input, err := ioutil.ReadFile(file)
		if err != nil {
			log.Fatalln(err)
		}

		lines := strings.Split(string(input), "\n")

		fmt.Println(lines)

		for i, line := range lines {

			fmt.Printf("%v FILE: %s\n", i, file)

			lNum := i + 1

			textLineNumber, err := strconv.ParseInt(lineNumber, 10, 32)
			if err != nil {
				log.Fatalln(err)
			}

			if lNum == int(textLineNumber) {
				fmt.Printf("File: %s\n", file)
				fmt.Printf("LineNumber: %s\n", lineNumber)
				fmt.Printf("Old: %s\n", line)
				fmt.Printf("New: %s\n", text)
				lines[i] = text
			}

		}

		output := strings.Join(lines, "\n")
		err = ioutil.WriteFile(file, []byte(output), permissions)
		if err != nil {
			log.Fatalln(err)
		}
	}

	// message := make([]byte, 5)
	// _, err := io.ReadFull(input, message)
	//
	// for err == nil {
	// 	fmt.Println(string(message))
	// 	_, err = io.ReadFull(input, message)
	// }
	// if err != io.EOF {
	// 	panic(err)
	// }
}

// func run(r io.Reader) {
// }
