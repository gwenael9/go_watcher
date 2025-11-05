package checker

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/gwenael9/go_watcher/internal/config"
)

type CheckResult struct {
	InputTarget config.InputTarget
	Status      string
	Err         error
}

type ReportEntry struct {
	Name   string
	URL    string
	Owner  string
	Status string // "OK", "Unreachable" ou "Error"
	ErrMsg string // Message d'erreur, omit si vide
}

func CheckUrl(target config.InputTarget) CheckResult {
	client := http.Client{
		Timeout: 3 * time.Second,
	}

	resp, err := client.Get(target.URL)
	if err != nil {
		return CheckResult{
			InputTarget: target,
			Err: &UnreachableError{
				URL: target.URL,
				Err: err,
			},
		}
	}

	defer resp.Body.Close()

	return CheckResult{
		InputTarget: target,
		Status:      resp.Status,
	}
}

func ConvertToReportEntry(results CheckResult) ReportEntry {
	report := ReportEntry{
		Name:   results.InputTarget.Name,
		URL:    results.InputTarget.URL,
		Owner:  results.InputTarget.Owner,
		Status: results.Status,
	}

	if results.Err != nil {
		var unreachable *UnreachableError
		if errors.As(results.Err, &unreachable) {
			report.Status = "Unreachable"
			report.ErrMsg = fmt.Sprintf("Unreachable URL: %v", unreachable.URL)
		} else {
			report.Status = "Error"
			report.ErrMsg = fmt.Sprintf("Error URL: %v", results.Err)

		}
	}
	return report
}
