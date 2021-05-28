package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

var version string  = "1.1"

func main() {
	arguments := os.Args[1:]
	if (len(arguments) >0 ) {
		if (string([]rune(arguments[0])[0]) != "-") {
			log.Fatal("ERROR! The first argument should not be the file name, you've entered " + arguments[0] + "\n You might be missing a minus(-).")
		}
		switch(arguments[0]) {
		case "-f", "--follow":
			infiniteRead(arguments[1])
			break
		case "-n", "--lines":
			readLast(arguments[1], arguments[2]) // Read last n bytes from file
			break;
		case "-h", "--help":
			showHelp()
			break
		case "-v", "--version":
			showVersion()
			break
		default:
			log.Fatal("ERROR!\nNot enough arguments or missing file name.\n\nThe right syntax is tail [OPTION] [FILEN]]\n\nOptions\n\tOptions:\n\t-f - Outputs appended data while the file grows.\n\t-n NUM - Outputs the last NUM lines of the specified file.\n\t-v - Displays the current tail version.\n\t-h - Shows this help information.")
		}
	} else {
		log.Fatal("ERROR! - Incorrect syntax!\n\nThe right syntax is: \n\ttail [OPTION] [FILE]\n\nOptions:\n\t-f - Outputs appended data while the file grows.\n\t-n NUM - Outputs the last NUM lines of the specified file.\n\t-v - Displays the current tail version.\n\t-h - Shows this help information.")
	}
}

func showVersion() {
	fmt.Println("\n---=== TAIL for Windows v" + version + " ===--- \n\nDeveloper: Robert Grozea (robert.grozea@gmail.com)")
}

func showHelp() {
	fmt.Println("\n---=== HELP ===--- \n\nThis is a light implementation of TAIL Linux command for Windows users. Some of the options are missing since their use in Windows environment does not make lot of sense.\n\n\n SYNTAX tail -[option] filename.ext\n\nOptions:\n\t-f - Outputs appended data while the file grows.\n\t-n NUM - Outputs the last NUM lines of the specified file.\n\t-v - Displays the current tail version.\n\t-h - Shows this help information.")
}

func readLast(linesToShow string, filename string) {
	fmt.Println("Showing the last", linesToShow, "lines from file", filename)
	toRead, err := strconv.Atoi(linesToShow)
	if ( err != nil ) {
		log.Fatal("ERROR!\nThe amount of lines to read must be a number.\nYou entered: " + linesToShow)
	}
	file, err := os.Open(filename)
	check4errors(err)
	lines := 0;
	scanner := bufio.NewScanner(file)
	for scanner.Scan() { lines++ }
	showFrom := lines - toRead
	lines2 := 0
	file.Close()
	file2, err := os.Open(filename)
	defer file2.Close()
	scanner2 := bufio.NewScanner(file2)
	for scanner2.Scan() {
		if ( lines2 >= showFrom ) {
			fmt.Println(scanner2.Text())
		}
		lines2++
	}
}

func infiniteRead(filename string) {
	f, err := os.Stat(filename)
	if ( err != nil ) {
		fmt.Println("ERROR!\nUnable to read the specified file: " + filename)
	}
	file, err := os.Open(filename)
	defer file.Close()
	check4errors(err)
	lastSize := f.Size()
	readLast("5", filename)
	key := make(chan string)
	for true {
		go checkForEnter(key)
		f, err := os.Stat(filename)
		if (err != nil) {
			fmt.Println("ERROR!\nUnable to find specified file!\n" + err.Error())
		}

		if f.Size() > lastSize {
			increase := f.Size() - lastSize
			buffer := make([]byte, increase)
			_, err := file.ReadAt(buffer, lastSize)
			check4errors(err)
			fmt.Printf("%s\n ", buffer)
			lastSize = f.Size()
		}
	}
}

func checkForEnter(key chan string) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		key <- scanner.Text()
	}
}

func check4errors(e error) {
	if e != nil {
		panic(e)
	}
}