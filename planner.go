package main

import (
	"fmt"
	"os"
	"bufio"
	"log"
	"strings"
	"regexp"
	"strconv"
)

func purgeCommas(targetString string) string{
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")

	if err != nil {
        log.Fatal(err)
	}
	
	processedString := reg.ReplaceAllString(targetString, " ")
	
	return processedString
}
func writeToFile(text []string){
	f, err := os.Create("notebook.txt")
    if err != nil {
        return
	}
	for i := 0; i < len(text); i++{
		l,err := f.WriteString(text[i] + "\n")
		if err != nil {
			return
		}
		l++; //useless function so i dont have to print
		
		
	}
	err = f.Close()
   		 if err != nil {
			return
	}
    
}
func load() []string{
	text := []string{};


	file, err := os.Open("notebook.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)

    for scanner.Scan() {         
		text = append(text, scanner.Text());
	}
	
	return text;
}

func main(){
	args := os.Args[1:]
	cmd := args[0]
	
	

	
	text := load();
	// upon loading , load notebook

	if cmd == "-h"{
		fmt.Println(" -a [task] = add \n -r [notebook index number] = remove \n -l = list \n -nuke = delete notebook \n -edit [notebook index number] [new task]")

	}else if cmd == "-a"{

		if len(args) > 1{
			att := args[1:]
			line := strings.Join(att[:], ",")
			line = purgeCommas(line)
			text = append(text,line)

			writeToFile(text);
		}else{
			fmt.Println("no task asigned")
		}
	}else if cmd == "-l"{
		for i := 1; i < len(text); i++{
			fmt.Printf("%d : %s \n", i, text[i])
		}
	}else if cmd == "-r"{
		index := args[1]
		index_i, err := strconv.Atoi(index)
		if err != nil {
			// handle error
			fmt.Println(err)
			os.Exit(2)
		}
		text = append(text[:index_i], text[index_i+1:]...)
		writeToFile(text)

	}else if cmd == "-nuke"{
		text = []string{}
		writeToFile(text)

	}else if cmd == "-edit"{
		index := args[1]
		att := args[2:]

		index_i, err := strconv.Atoi(index)
		if err != nil {
			// handle error
			fmt.Println(err)
			os.Exit(2)
		}
		
		line := strings.Join(att[:], ",")
		line = purgeCommas(line)

		text[index_i] = line
		writeToFile(text)
	}

}