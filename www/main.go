// Handles some amount of voidedtech.com hosting
package main

import (
	"bytes"
	_ "embed"
	"errors"
	"flag"
	"fmt"
	"html/template"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/gorilla/feeds"
)

var (
	//go:embed index.html
	indexHTML string
	//go:embed main.css
	mainCSS []byte
	linkSet = []string{newRecord("https://github.com/enckse", "Github"),
		newRecord("https://goodreads.com/enckse", "Goodreads"),
		newRecord("https://instagram.com/seanenck", "Instagram"),
		newRecord("https://www.linkedin.com/in/sean-enck-22420314", "LinkedIn")}
)

const (
	delimiter = ","
	rootURL   = "https://voidedtech.com"
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
)

func newRecord(href, disp string) string {
	return fmt.Sprintf("%s%s%s", href, delimiter, disp)
}

func build(sub, dest string) error {
	sorted := linkSet
	for _, s := range strings.Split(sub, " ") {
		l := len(strings.TrimSpace(s))
		switch l {
		case 0:
			continue
		case 1:
			return errors.New("invalid subsite")
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
	obj := SiteData{}
	obj.Date = time.Now().Format("2006-01-02")
	obj.Links = links
	tmpl, err := template.New("t").Parse(indexHTML)
	if err != nil {
		return err
	}
	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, obj); err != nil {
		return err
	}
	if err := os.WriteFile(filepath.Join(dest, "index.html"), buf.Bytes(), 0644); err != nil {
		return err
	}
	if err := os.WriteFile(filepath.Join(dest, "main.css"), mainCSS, 0644); err != nil {
		return err
	}
	return genFeed(dest)
}

func main() {
	target := flag.String("target", "", "target output")
	subsites := flag.String("sites", "", "subsites")
	flag.Parse()
	if err := build(*subsites, *target); err != nil {
		fmt.Printf("build failed: %v", err)
		os.Exit(1)
	}
}

func genFeed(dest string) error {
	now := time.Now()
	feed := &feeds.Feed{
		Title:       "voidedtech.com updates",
		Link:        &feeds.Link{Href: rootURL},
		Description: "various updates from voidedtech",
		Created:     now,
	}
	output, err := exec.Command("git", "-C", "notebook", "log", "-n", "25", "--format=%ai %f").Output()
	if err != nil {
		return err
	}
	for _, e := range strings.Split(string(output), "\n") {
		line := strings.TrimSpace(e)
		if line == "" {
			continue
		}
		parts := strings.Split(line, " ")
		if len(parts) < 4 {
			return errors.New("invalid log entry from git")
		}
		dt, err := time.Parse("2006-01-02 15:04:05 -0700", strings.Join(parts[0:3], " "))
		if err != nil {
			return err
		}
		title := parts[3]
		feed.Items = append(feed.Items, &feeds.Item{
			Title:       title,
			Description: title,
			Created:     dt,
			Link:        &feeds.Link{Href: rootURL},
		})
	}

	rss, err := feed.ToRss()
	if err != nil {
		return err
	}

	return os.WriteFile(filepath.Join(dest, "rss-notebook.xml"), []byte(rss), 0644)
}
