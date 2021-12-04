package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	wipe "github.com/0x9ef/go-wiper/wipe"
)

var mappedPolicy = map[int]*wipe.Policy{
	1: &wipe.Policy{"Fast", "Data will be overwrited with zeroes (1 passes)", wipe.RuleFast},
	2: &wipe.Policy{"VSITR", "German VSITR (7 passes)", wipe.RuleVSITR},
	3: &wipe.Policy{"UsDod5220_22_M", "US Department of Defense DoD 5220.22-M (3 passes)", wipe.RuleUsDod5220_22_M},
	4: &wipe.Policy{"Gutmann", "Peter Gutmann Secure Method (35 passes)", wipe.RuleGutmann},
}

func runmethod(rule int, file_args []string) {
	head()

	policy, ok := mappedPolicy[rule]
	if !ok {
		panic("wiper: provided unknown wipe rule")
	}

	wrule := policy.Rule
	fmt.Printf("Selected rule:\n%s\n", policy.String())

	var files []string

	switch len(file_args) {
	case 0:
		os.Exit(0)
	default:
		files = run_files(file_args)
	}

	for _, file := range files {
		err := wipe.Wipe(file, wrule)
		if err != nil {
			panic(err)
		}
		if !keep {
			base_name := filepath.Base(file)
			base_name = strings.Repeat("0", len(base_name))
			os.Rename(file, filepath.Join(filepath.Dir(file), base_name))
			err = os.Remove(filepath.Join(filepath.Dir(file), base_name))
			if err != nil {
				os.Exit(1)
			}
		}
	}

}

func run_files(files []string) []string {
	var ret_file []string
	for _, file := range files {

		fi, err := os.Stat(file)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		if fi.Mode().IsDir() {
			filepath.Walk(file, func(path string, info os.FileInfo, err error) error {
				if err != nil {
					return err
				}
				if info.Mode().IsRegular() {
					ret_file = append(ret_file, path)
				}

				return nil
			})
			if err != nil {
				log.Println(err)
			}
		} else {
			path, _ := filepath.Abs(file)
			ret_file = append(ret_file, path)
		}

	}
	return ret_file
}
