package dbminer

import (
	"fmt"
	"regexp"
)

type DatabaseMiner interface {
	GetSchema() (*Schema, error)
	GetDbClass() string
}

type Schema struct {
	Databases []Database
}

type Database struct {
	Name   string
	Tables []Table
}

type Table struct {
	Name    string
	Columns []string
}

func Search(m DatabaseMiner) error {
	s, err := m.GetSchema()
	if err != nil {
		return err
	}
	fmt.Println(m.GetDbClass() + " VICTIM!!! YEET")
	re := getRegex()
	for _, database := range s.Databases {
		for _, table := range database.Tables {
			for _, field := range table.Columns {
				for _, r := range re {
					if r.MatchString(field) {
						fmt.Println(database)
						fmt.Printf("[+] HIT A LICK!: %s\n", field)
					}
				}
			}
		}
	}
	return nil
}

func getRegex() []*regexp.Regexp {
	return []*regexp.Regexp{
		regexp.MustCompile(`(?i)social`),
		regexp.MustCompile(`(?i)security`),
		regexp.MustCompile(`(?i)pass(word)?`),
		regexp.MustCompile(`(?i)ccnum`),
		regexp.MustCompile(`(?i)cvv`),
		regexp.MustCompile(`(?i)exp`),
		regexp.MustCompile(`(?i)ssn`),
		regexp.MustCompile(`(?i)address`),
		regexp.MustCompile(`(?i)city`),
		regexp.MustCompile(`(?i)state`),
		regexp.MustCompile(`(?i)zip`),
	}
}

func (s Schema) String() string {
	var ret string
	for _, database := range s.Databases {
		ret += fmt.Sprint(database.String() + "\n")
	}
	return ret
}

func (d Database) String() string {
	ret := fmt.Sprintf("[DB] = %+s\n", d.Name)
	for _, table := range d.Tables {
		ret += table.String()
	}
	return ret
}

func (t Table) String() string {
	ret := fmt.Sprintf("	[TABLE] = %+s\n", t.Name)
	for _, field := range t.Columns {
		ret += fmt.Sprintf("		[COL] = %+s\n", field)
	}
	return ret
}
