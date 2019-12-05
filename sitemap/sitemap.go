package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/gophercises/link"
)

func main() {

	webpage := flag.String("url", "https://gophercises.com", "url you d like to build sitemap for")
	maxDepth := flag.Int("depth", 3, "max depth you d like to reach down in sitemap traverse")

	pages := bfs(*webpage, *maxDepth)

	toXml(pages)
}

func bfs(urlString string, maxDepth int) []string {
	seen := make(map[string]struct{})

	var queue = make(map[string]struct{})
	nextQueue := map[string]struct{}{
		urlString: struct{}{},
	}

	for i := 0; i <= maxDepth; i++ {
		queue, nextQueue = nextQueue, make(map[string]struct{})

		for key, _ := range queue {
			if _, ok := seen[key]; ok {
				continue
			}
			seen[key] = struct{}{}

			for _, link := range get(key) {
				nextQueue[link] = struct{}{}
			}
		}
	}

	all := make([]string, 0, len(seen))
	for url, _ := range seen {
		all = append(all, url)
	}
	return all
}

func get(webpage string) []string {
	// 1) fetch HTML
	resp, err := http.Get(webpage)
	if err != nil {
		panic("no no no")
	}
	defer resp.Body.Close()

	// resolve baseUrl from request
	reqUrl := resp.Request.URL
	baseUrl := &url.URL{
		Scheme: reqUrl.Scheme,
		Host:   reqUrl.Host,
	}
	base := baseUrl.String()

	return hrefs(resp.Body, base)
}

// 2) parse links
func hrefs(body io.Reader, base string) []string {
	var filteredLinks []string
	links, _ := link.Parse(body)

	for _, curLink := range links {
		switch {
		case strings.HasPrefix(curLink.Href, "/"):
			filteredLinks = append(filteredLinks, fmt.Sprintf("%v%v", base, curLink.Href))
		case strings.HasPrefix(curLink.Href, base):
			filteredLinks = append(filteredLinks, curLink.Href)
		}
	}
	return filteredLinks
}

type loc struct {
	Value string `xml:"loc"`
}

type urlset struct {
	Urls []loc `xml:"url"`
}

func toXml(pages []string) {
	var xmlContent urlset

	for _, page := range pages {
		xmlContent.Urls = append(xmlContent.Urls, loc{page})
	}

	enc := xml.NewEncoder(os.Stdout)
	enc.Indent("", " ")
	err := enc.Encode(xmlContent)
	if err != nil {
		panic("nonono")
	}
}
