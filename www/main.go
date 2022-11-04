// Handles some amount of voidedtech.com hosting
package main

import (
	"bytes"
	_ "embed"
	"encoding/json"
	"flag"
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"time"
)

var (
	//go:embed index.html
	indexHTML string
	//go:embed main.css
	mainCSS []byte
)

type (
	// SiteData is the necessary data to render the site.
	SiteData struct {
		Date   string
		Title  string
		Header string
		Links  []Link
		Mail   string
		Places []Place
	}
	// Link represents page links.
	Link struct {
		Href    string
		Display string
	}
	// Place indicate a place I have worked/learned at.
	Place struct {
		Name string
		From string
		To   string
	}
	hostData struct {
		mainIndex []byte
		mainCSS   []byte
	}
)

func createHostData(file string) (hostData, error) {
	d, err := os.ReadFile(file)
	if err != nil {
		return hostData{}, err
	}
	obj := SiteData{}
	if err := json.Unmarshal(d, &obj); err != nil {
		return hostData{}, err
	}
	obj.Date = time.Now().Format("2006-01-02")
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
	flag.Parse()
	data, err := createHostData(*config)
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
