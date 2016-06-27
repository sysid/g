// vim: fdm=marker ts=4 sts=4 sw=4 fdl=0
// g jumps to the location according to its config-file defined in environment variable $twJUMPLIST.
// The twJUMPLIST file is CSV format with key,jumppath entries.
// When key is found with valid path g returns 0, else 1, i.e. when 0 then shell has got a valid path
//
// # JumpList_Example.csv
// h,/usr/home/foo
// xxx,/usr/log/xxx
package main

//// Imports {{{
import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/sysid/tw"
	//. "github.com/sysid/tw/basic"
	"gopkg.in/alecthomas/kingpin.v2"
)

////}}}

//// Variables/Constants {{{
var (
	debug    = func(v ...interface{}) {}
	filePath = kingpin.Flag("filepath", "path to config file").Required().Envar("twJUMPLIST").Short('f').ExistingFile()
	sKeys    = kingpin.Flag("skeys", "Show keys").Short('s').Bool() //for bash completion
	key      = kingpin.Arg("key", "key to identify path").String()
)

////}}}

//// Functions {{{
func check(e error) {
	if e != nil {
		panic(e)
	}
}
func printDirs(g map[string]string, sKeys bool) {
	var keys []string
	for k, _ := range g {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		if sKeys {
			fmt.Fprintf(os.Stdout, "%s\n", k)
		} else {
			fmt.Fprintf(os.Stderr, "%-10s%s\n", k, g[k])
		}
	}
}

////}}}

//// Main {{{
func main() {
	defer tw.HandleExit()
	//defer tw.End(time.Now())

	kingpin.Parse()

	csvfile, err := os.Open(*filePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Configfile: %s\n", err.Error())
		os.Exit(1)
	}
	defer csvfile.Close()

	reader := csv.NewReader(csvfile)
	reader.FieldsPerRecord = -1    // see the Reader struct information below
	reader.TrimLeadingSpace = true // see the Reader struct information below
	reader.Comment = '#'

	rawCSVdata, err := reader.ReadAll()
	check(err)

	//g []map[string]string
	g := make(map[string]string)

	// read into map
	for _, v := range rawCSVdata {
		//Debug(":%s:", v[1])
		g[v[0]] = filepath.Clean(os.ExpandEnv(strings.TrimSpace(v[1])))
	}
	printDirs(g, *sKeys)

	if v, ok := g[*key]; ok {
		//check whether jumppath exists
		if !tw.Exists(v) {
			fmt.Fprintf(os.Stderr, "%s does not exist.\nFix config: %s\n", v, *filePath)
			os.Exit(1)
		} else {
			fmt.Printf("%s\n", v)
			os.Exit(0)
		}
	} else {
		printDirs(g, *sKeys)
		os.Exit(1)
	}
}

////}}}
