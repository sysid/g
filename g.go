// vim: fdm=marker ts=4 sts=4 sw=4 fdl=0
package main

//// Imports {{{
import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/sysid/tw"
	//. "github.com/thomas/tw/basic"
)

////}}}

//// Variables/Constants {{{
var (
	debug = func(v ...interface{}) {}
	dbg   bool
)

////}}}

//// Functions {{{
func parseFlags() error {
	dbg = *flag.Bool("d", false, "debugging switch")
	flag.Usage = myUsage
	flag.Parse()
	//Log("%s>>ConfigFlags:%t, %v", tw.GetFN(), dbg, flag.Args())
	return nil
}
func myUsage() {
	fmt.Printf("Usage: %s [OPTIONS] argument ...\n", filepath.Base(os.Args[0]))
	flag.PrintDefaults()
}
func check(e error) {
	if e != nil {
		panic(e)
	}
}
func printDirs(g map[string]string) {
	for k, v := range g {
		fmt.Fprintf(os.Stderr, "%-10s%s\n", k, v)
	}
}

////}}}

//// Main {{{
func main() {
	defer tw.HandleExit()
	//defer tw.End(time.Now())
	_ = parseFlags()

	path := os.Getenv("twDev")
	if path == "" {
		fmt.Fprintf(os.Stderr, "Error: Environmentvariable $twDev not set.\n")
		os.Exit(1)
	}
	if !tw.Exists(path) {
		fmt.Fprintf(os.Stderr, "%s does not exist.\n", path)
		os.Exit(1)
	}

	//path := "/Users/q187392/dev/go/src/github.com/thomas/g"
	path = fmt.Sprintf("%s/cfg/g", path)
	hostname, _ := os.Hostname()
	path = fmt.Sprintf("%s/%s.csv", path, hostname)
	//Yellow(path)

	csvfile, err := os.Open(path)
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

	if len(flag.Args()) != 1 {
		//return errors.New("...")
		//flag.Usage()
		printDirs(g)
		os.Exit(1)
	}

	if v, ok := g[flag.Args()[0]]; ok {
		if !tw.Exists(v) {
			fmt.Fprintf(os.Stderr, "%s does not exist.\nFix config: %s\n", v, path)
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
