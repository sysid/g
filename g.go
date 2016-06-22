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
	"sort"

	"github.com/sysid/tw"
	//. "github.com/thomas/tw/basic"
	"gopkg.in/caarlos0/env.v2"
)

////}}}

//// Variables/Constants {{{
var (
	debug = func(v ...interface{}) {}
	dbg   bool
)

type config struct {
	JumpList string `env:"twJUMPLIST"`
}

////}}}

//// Functions {{{
func check(e error) {
	if e != nil {
		panic(e)
	}
}
func printDirs(g map[string]string) {
	var keys []string
	for k, _ := range g {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	// To perform the opertion you want
	for _, k := range keys {
		fmt.Fprintf(os.Stderr, "%-10s%s\n", k, g[k])
	}
}

////}}}

//// Main {{{
func main() {
	defer tw.HandleExit()
	//defer tw.End(time.Now())

	cfg := config{}
	env.Parse(&cfg)

	if cfg.JumpList == "" {
		fmt.Fprintf(os.Stderr, "Error: Environmentvariable $twJUMPLIST not set.\n")
		os.Exit(1)
	}
	if !tw.Exists(cfg.JumpList) {
		fmt.Fprintf(os.Stderr, "%s does not exist.\n", cfg.JumpList)
		os.Exit(1)
	}

	//cfg.JumpList := "/Users/q187392/dev/go/src/github.com/thomas/g"

	csvfile, err := os.Open(cfg.JumpList)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Configfile: %s\n", err.Error())
		os.Exit(1)
	}
	defer csvfile.Close()

	reader := csv.NewReader(csvfile)
	reader.FieldsPerRecord = -1 // see the Reader struct information below
	reader.Comment = '#'

	rawCSVdata, err := reader.ReadAll()
	check(err)

	//g []map[string]string
	g := make(map[string]string)

	// sanity check, display to standard output
	for _, v := range rawCSVdata {
		g[v[0]] = v[1]
	}

	if len(os.Args) > 2 || len(os.Args) == 1 {
		printDirs(g)
		os.Exit(1)
	}

	if v, ok := g[os.Args[1]]; ok {
		if !tw.Exists(v) {
			fmt.Fprintf(os.Stderr, "%s does not exist.\nFix config: %s\n", v, cfg.JumpList)
			os.Exit(1)
		} else {
			fmt.Printf("%s\n", v)
			os.Exit(0)
		}
	} else {
		printDirs(g)
		os.Exit(1)
	}
}

////}}}
