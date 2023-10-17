package pgc

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func UpdatePgHba(pgHbaConf []struct {
	Type     string `yaml:"type"`
	Database string `yaml:"database"`
	User     string `yaml:"user"`
	Address  string `yaml:"address"`
	Method   string `yaml:"method"`
}, pgHbaConfPath string) error {

	// Reading existing lines
	existingLines, err := os.ReadFile(pgHbaConfPath)
	if err != nil {
		return err
	}

	// Opening file in append mode
	file, err := os.OpenFile(pgHbaConfPath, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatalf("could not close file: %v", err)
		}
	}(file)

	// Looping over each new rule
	for _, rule := range pgHbaConf {
		newRule := fmt.Sprintf("%s\t%s\t%s\t%s\t%s", rule.Type, rule.Database, rule.User, rule.Address, rule.Method)

		// Check if the rule already exists
		if !containsRule(string(existingLines), newRule) {
			_, err := file.WriteString(newRule + "\n")
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// Helper function to check if a rule already exists in pg_hba.conf
func containsRule(existingLines, newRule string) bool {
	lines := strings.Split(existingLines, "\n")
	for _, line := range lines {
		if strings.TrimSpace(line) == strings.TrimSpace(newRule) {
			return true
		}
	}
	return false
}
