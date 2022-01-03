package cmd

import (
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"

	wipe "github.com/0x9ef/go-wiper/wipe"
	log "github.com/sirupsen/logrus"
)

var mappedPolicy = map[int]*wipe.Policy{
	1: &wipe.Policy{"Fast", "Data will be overwrited with zeroes (1 passes)", wipe.RuleFast},
	2: &wipe.Policy{"VSITR", "German VSITR (7 passes)", wipe.RuleVSITR},
	3: &wipe.Policy{"UsDod5220_22_M", "US Department of Defense DoD 5220.22-M (3 passes)", wipe.RuleUsDod5220_22_M},
	4: &wipe.Policy{"Gutmann", "Peter Gutmann Secure Method (35 passes)", wipe.RuleGutmann},
}

func runmethod(rule int, file_args []string) {
	head()

	switch verbose {
	case 0:
		log.SetLevel(log.PanicLevel)
	case 1:
		log.SetLevel(log.FatalLevel)
	case 2:
		log.SetLevel(log.ErrorLevel)
	case 3:
		log.SetLevel(log.WarnLevel)
	case 4:
		log.SetLevel(log.InfoLevel)
	case 5:
		log.SetLevel(log.DebugLevel)
	case 6:
		log.SetLevel(log.TraceLevel)
	default:
		log.SetLevel(log.InfoLevel)
	}
	log.Debugln("Log level is: " + strconv.Itoa(verbose))

	log.Infoln("log rule is: " + strconv.Itoa(rule))

	policy, ok := mappedPolicy[rule]
	if !ok {
		panic("wiper: provided unknown wipe rule")
	}

	wrule := policy.Rule
	log.Infoln("Selected rule: " + policy.String())

	var files []string

	log.Debugln(file_args)

	if len(file_args) == 0 {
		log.Debugln("Stop system no file set")
		os.Exit(0)
	} else {
		files = run_files(file_args)
	}

	wg := sync.WaitGroup{}

	// shread all files with wipe
	for _, file := range files {
		log.Debugln("delete file: " + file)

		wg.Add(1)
		go func(file_ string, wg *sync.WaitGroup) {
			err := wipe.Wipe(file_, wrule)
			if err != nil {
				panic(err)
			}
			if !keep {
				base_name := filepath.Base(file_)
				base_name = strings.Repeat("0", len(base_name))
				os.Rename(file_, filepath.Join(filepath.Dir(file_), base_name))
				err = os.Remove(filepath.Join(filepath.Dir(file_), base_name))
				if err != nil {
					os.Exit(1)
				}
			}
			wg.Done()
		}(file, &wg)
	}

	wg.Wait()

}

func run_files(files []string) []string {
	var ret_file []string
	for _, file := range files {

		log.Infoln("check file: " + file)
		fi, err := os.Stat(file)
		if err != nil {
			log.Println(err)
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
