package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"time"

	wipe "github.com/0x9ef/go-wiper/wipe"
	"github.com/cheggaaa/pb/v3"
	log "github.com/sirupsen/logrus"
)

var mappedPolicy = map[int]*wipe.Policy{
	1: {Name: "Fast", Description: "Data will be overwritten with zeroes (1 pass)", Rule: wipe.RuleFast},
	2: {Name: "VSITR", Description: "German VSITR (7 passes)", Rule: wipe.RuleVSITR},
	3: {Name: "UsDod5220_22_M", Description: "US Department of Defense DoD 5220.22-M (3 passes)", Rule: wipe.RuleUsDod5220_22_M},
	4: {Name: "Gutmann", Description: "Peter Gutmann Secure Method (35 passes)", Rule: wipe.RuleGutmann},
}

func runmethod(rule int, file_args []string) error {
	head()

	// Log-Level setzen
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

	if rule == 0 {
		rule = 3 // Default to UsDod5220_22_M
	}

	policy, ok := mappedPolicy[rule]
	if !ok {
		var validRules []string
		for k, v := range mappedPolicy {
			validRules = append(validRules, fmt.Sprintf("%d: %s", k, v.Name))
		}
		return fmt.Errorf("wiper: provided unknown wipe rule: %d. Valid rules are: %v", rule, validRules)
	}

	wrule := policy.Rule
	log.Infoln("Selected rule: " + policy.String())

	var files []string
	var err error

	log.Debugln(file_args)

	if len(file_args) > 0 {
		files, err = run_files(file_args)
		if err != nil {
			return fmt.Errorf("error processing files: %v", err)
		}
	} else {
		return fmt.Errorf("no files or directories provided. Use 'shred --help' for usage information")

	}

	// Fortschrittsleiste erstellen
	bar := pb.StartNew(len(files))
	bar.SetTemplateString(`{{counters . }} {{bar . }} {{percent . }} {{etime . }}`)

	wg := sync.WaitGroup{}
	errCh := make(chan error, len(files)) // Fehlerkanal

	// Shred all files with wipe
	for _, file := range files {
		log.Debugln("delete file: " + file)

		wg.Add(1)
		go func(file_ string, wg *sync.WaitGroup) {
			defer wg.Done()
			err := wipe.Wipe(file_, wrule)
			if err != nil {
				errCh <- fmt.Errorf("failed to wipe file %s: %v", file_, err)
				return
			}
			if !keep {
				base_name := filepath.Base(file_)
				newName := strings.Repeat("0", len(base_name))
				new_path := filepath.Join(filepath.Dir(file_), newName)
				log.Debugf("Renaming file %s to %s", file_, new_path)
				err = os.Rename(file_, new_path)
				if err != nil {
					log.Errorf("Failed to rename file %s: %v", file_, err)
					return
				}
				time.Sleep(100 * time.Millisecond) // Kurze Verzögerung
				log.Debugf("Removing file %s", new_path)
				err = os.Remove(new_path)
				if err != nil {
					log.Errorf("Failed to remove file %s: %v", new_path, err)
					return
				}
			}
			bar.Increment() // Fortschrittsleiste aktualisieren
		}(file, &wg)
	}

	// Warte auf alle Goroutines
	go func() {
		wg.Wait()
		close(errCh) // Schließe den Fehlerkanal, wenn alle Goroutines fertig sind
	}()

	// Sammle Fehler
	var errors []error
	for err := range errCh {
		errors = append(errors, err)
	}

	bar.Finish() // Fortschrittsleiste beenden

	// Gib alle gesammelten Fehler zurück
	if len(errors) > 0 {
		return fmt.Errorf("encountered %d errors during processing: %v", len(errors), errors)
	}

	return nil
}

func run_files(files []string) ([]string, error) {
	var ret_file []string
	for _, file := range files {
		log.Infoln("check file: " + file)
		fi, err := os.Stat(file)
		if err != nil {
			return nil, fmt.Errorf("failed to stat file %s: %v", file, err)
		}

		if fi.Mode().IsDir() {
			err := filepath.Walk(file, func(path string, info os.FileInfo, err error) error {
				if err != nil {
					return err
				}
				if info.Mode().IsRegular() {
					absPath, err := filepath.Abs(path)
					if err != nil {
						return fmt.Errorf("failed to get absolute path for %s: %v", path, err)
					}
					ret_file = append(ret_file, absPath)
				}
				return nil
			})
			if err != nil {
				return nil, fmt.Errorf("failed to walk directory %s: %v", file, err)
			}
		} else {
			absPath, err := filepath.Abs(file)
			if err != nil {
				return nil, fmt.Errorf("failed to get absolute path for %s: %v", file, err)
			}
			ret_file = append(ret_file, absPath)
		}
	}
	return ret_file, nil
}
