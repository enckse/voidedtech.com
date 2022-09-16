// Package main has indexing needs for links.
package main

type (
	// SiteData is the necessary data to render the site.
	SiteData struct {
		Date     string
		IsPublic bool
		Title    string
		Header   string
		Links    []Link
		Mail     string
		Places   []Place
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
)
