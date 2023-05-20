package main

import (
	"fmt"
	"github.com/projectdiscovery/goflags"
	"github.com/projectdiscovery/gologger"
	fUtils "github.com/projectdiscovery/utils/file"
	sUtils "github.com/projectdiscovery/utils/slice"
)

type options struct {
	result string
	file1  string
	file2  string
	silent bool
}

var opt = &options{}

func main() {
	flagSet := goflags.NewFlagSet()
	flagSet.SetDescription(`comx is a tool to compare the contents of two files and return the unique data`)
	flagSet.StringVarP(&opt.file1, "file1", "f1", "", "File1 for compare")
	flagSet.StringVarP(&opt.file2, "file2", "f2", "", "File2 for compare")
	flagSet.StringVarP(&opt.result, "result", "r", "", "Print only one Diff content")
	flagSet.BoolVarP(&opt.silent, "silent", "s", false, "Silent mode")
	if err := flagSet.Parse(); err != nil {
		gologger.Error().Msgf("Could not parse flags: %s\n", err)
	}

	if opt.file1 == "" && opt.file2 == "" {
		Banner()
		gologger.Print().Msgf("Flags:\n")
		gologger.Print().Msgf("\t-file1,   -f1 	File1 for compare")
		gologger.Print().Msgf("\t-file2,   -f2 	File2 for compare")
		gologger.Print().Msgf("\t-result,  -r 	Print unique content for each file (1,2)")
		gologger.Print().Msgf("\t-silent,  -s 	Silent Mode")
		return
	}

	if opt.result != "" && opt.result != "1" && opt.result != "2" && opt.result != "1,2" {
		gologger.Print().Msgf("Result flag must be 1 or 2:")
		gologger.Print().Msgf("\t-result,   -r    Print only one Diff content (1,2)")
		return
	}

	if !opt.silent {
		Banner()
	}

	// Get input from users
	data, data2 := Input(opt.file1, opt.file2)

	// Switch case for print results
	switch opt.result {
	case "1":
		for _, line := range data {
			fmt.Println(line)
		}
	case "2":
		for _, line := range data2 {
			fmt.Println(line)
		}
	case "1,2":
		for _, line := range data {
			fmt.Println(line)
		}

		for _, line := range data2 {
			fmt.Println(line)
		}

	default:
		fmt.Printf("[%s] Unique Content:\n", opt.file1)
		for _, line := range data {
			fmt.Println(line)
		}

		fmt.Printf("\n[%s] Unique Content:\n", opt.file2)
		for _, line := range data2 {
			fmt.Println(line)
		}
	}
}

// Input function for get input from users
func Input(input1, input2 string) (file1 []string, file2 []string) {
	var File1 []string
	var File2 []string

	f1, err := fUtils.ReadFile(input1)
	if err != nil {
		panic(err)
	}
	for v := range f1 {
		File1 = append(File1, v)
	}

	f2, err := fUtils.ReadFile(input2)
	if err != nil {
		panic(err)
	}

	for v := range f2 {
		File2 = append(File2, v)
	}

	content1, content2 := sUtils.Diff(File1, File2)

	return content1, content2
}

func Banner() {
	gologger.Print().Msgf(`
 ██████╗ ██████╗ ███╗   ███╗██╗  ██╗
██╔════╝██╔═══██╗████╗ ████║╚██╗██╔╝
██║     ██║   ██║██╔████╔██║ ╚███╔╝ 
██║     ██║   ██║██║╚██╔╝██║ ██╔██╗ 
╚██████╗╚██████╔╝██║ ╚═╝ ██║██╔╝ ██╗
 ╚═════╝ ╚═════╝ ╚═╝     ╚═╝╚═╝  ╚═╝
`)
	gologger.Print().Msgf("\t\t\t\tCreated by WatchDogs :)\n\n")
}
