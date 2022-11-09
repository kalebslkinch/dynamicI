package main

import (
	"github.com/rivo/tview"
	"github.com/gdamore/tcell/v2"
	"bufio"
	"os"
	"fmt"
	"strings"
	"log"
)
func menu() tview.Primitive { 
	list := tview.NewList(). 
		
	// 1. Convert Import to Dynamic Import
	AddItem("Convert Import to Dynamic Import", "", '1', nil).

	// 2. Exit 
	AddItem("Exit", "", '2', nil)
	return list
}

func importConverter (terminal *tview.Application) {
	
	// Stop the terminal
	terminal.Stop()

	// Print the message
	fmt.Println("Paste the Imports")
	
	// Create a new reader
	reader := bufio.NewReader(os.Stdin)

	var lines []string
	for {
		// read line from stdin using newline as separator
		line, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		// if line is empty, break the loop
		if len(strings.TrimSpace(line)) == 0 {
			break
		}

		//append the line to a slice
		lines = append(lines, line)
	}

	//print out all the lines
	println(len(lines))
	fmt.Println("output:")
	for _, eachLine := range lines { 
		replacedImport := strings.ReplaceAll(eachLine,"import","const")
		replacedImport = strings.ReplaceAll(replacedImport,"from ","= dynamic(() => import(")
		replacedImport = strings.ReplaceAll(replacedImport, ";","")
		replacedImport = replacedImport + "));"	
		fmt.Printf(replacedImport + "\n")
	}
}

func main() {
	terminal := tview.NewApplication()

	terminal.SetRoot(menu(), true)
	terminal.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {

		if event.Key() == tcell.KeyEscape {
			terminal.Stop()
		}

		if event.Key() == tcell.KeyRune {
			
			if event.Rune() == '1' {
				importConverter(terminal)
			}
				
			if event.Rune() == '2' {
				terminal.Stop()
			}
			
		}
		return event
	})

	if err := terminal.Run(); err != nil {
		panic(err)
	}
}
