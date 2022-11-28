// Handles some amount of voidedtech.com hosting
package main

import (
	"bytes"
	_ "embed"
	"encoding/csv"
	"errors"
	"flag"
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

var (
	//go:embed index.html
	indexHTML string
	//go:embed main.css
	mainCSS []byte
)

const (
	delimiter = ","
)

type (
	// SiteData is the necessary data to render the site.
	SiteData struct {
		Date  string
		Links []Link
	}
	// Link represents page links.
	Link struct {
		Href    string
		Display string
	}
	hostData struct {
		mainIndex []byte
		mainCSS   []byte
	}
)

func newRecord(href, disp string) string {
	return fmt.Sprintf("%s%s%s", href, delimiter, disp)
}

func createHostData(file, sub string) (hostData, error) {
	f, err := os.Open(file)
	if err != nil {
		return hostData{}, err
	}
	defer f.Close()
	obj := SiteData{}
	r := csv.NewReader(f)
	records, err := r.ReadAll()
	if err != nil {
		return hostData{}, err
	}
	var sorted []string
	for idx, record := range records {
		if idx == 0 {
			continue
		}
		if len(record) != 2 {
			return hostData{}, errors.New("invalid record found")
		}
		sorted = append(sorted, newRecord(record[0], record[1]))
	}
	for _, s := range strings.Split(sub, " ") {
		l := len(strings.TrimSpace(s))
		switch l {
		case 0:
			continue
		case 1:
			return hostData{}, errors.New("invalid subsite")
		}
		title := s[1:]
		title = fmt.Sprintf("%s%s", strings.ToUpper(string(s[0])), title)
		sorted = append(sorted, newRecord(s+"/", title))
	}
	sort.Strings(sorted)
	var links []Link
	for _, raw := range sorted {
		record := strings.Split(raw, delimiter)
		l := Link{}
		l.Href = record[0]
		l.Display = record[1]
		links = append(links, l)
	}
	obj.Date = time.Now().Format("2006-01-02")
	obj.Links = links
	tmpl, err := template.New("t").Parse(indexHTML)
	if err != nil {
		return hostData{}, err
	}
	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, obj); err != nil {
		return hostData{}, err
	}
	return hostData{mainIndex: buf.Bytes(), mainCSS: mainCSS}, nil
}

func main() {
	config := flag.String("config", "", "site json definition")
	target := flag.String("target", "", "target output")
	subsites := flag.String("sites", "", "subsites")
	flag.Parse()
	data, err := createHostData(*config, *subsites)
	if err != nil {
		die("unable to create host data", err)
	}
	dest := *target
	if err := os.WriteFile(filepath.Join(dest, "index.html"), data.mainIndex, 0644); err != nil {
		die("failed to write index", err)
	}
	if err := os.WriteFile(filepath.Join(dest, "main.css"), data.mainCSS, 0644); err != nil {
		die("failed to write css", err)
	}
}

func die(message string, err error) {
	fmt.Printf("%s: %v", message, err)
	os.Exit(1)
}
