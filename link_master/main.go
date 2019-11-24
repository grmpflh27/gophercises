package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"golang.org/x/net/html"
)

var allLinkNodes []*html.Node
var curText []string

func main() {
	htmlFile := flag.String("file", "ex1.html", "file you'd wish to parse")
	flag.Parse()

	fmt.Printf("Parsing %v\n", *htmlFile)
	fp, err := os.Open(*htmlFile)

	if err != nil {
		log.Fatal("Could not open ", *htmlFile)
	}

	rootNode, err := html.Parse(fp)
	if err != nil {
		log.Fatal("Could not parse")
	}

	getAllLinkNodes(rootNode)
	fmt.Printf("Found %v links\n", len(allLinkNodes))

	getLinks()
}

func getAllLinkNodes(curNode *html.Node) {

	if isLink(curNode) {
		allLinkNodes = append(allLinkNodes, curNode)
	}

	for c := curNode.FirstChild; c != nil; c = c.NextSibling {
		getAllLinkNodes(c)
	}
}

func isLink(n *html.Node) bool {
	return n.Type == html.ElementNode && n.Data == "a"
}

func getLinks() {
	var links []Link
	for _, node := range allLinkNodes {
		href, err := getHref(node)
		if err != nil {
			log.Fatal("ERRRRR")
		}
		getAllText(node)
		links = append(links, Link{href, strings.Join(curText, "")})
		curText = curText[:0]
	}

	fmt.Printf("%+v\n", links)
}

func getHref(n *html.Node) (string, error) {
	for _, a := range n.Attr {
		if a.Key == "href" {
			return a.Val, nil
		}
	}
	return "", errors.New("no dice")
}

func getAllText(n *html.Node) {
	if n.Type == html.TextNode {
		curText = append(curText, n.Data)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		getAllText(c)
	}
}

type Link struct {
	Href string
	Text string
}
