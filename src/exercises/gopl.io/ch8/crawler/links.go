// Package links provides a link-extraction function.
package crawler

import (
	"fmt"
	"net/http"

	"golang.org/x/net/html"
)

// Extract makes an HTTP GET request to the specified URL url, parses
// the response as HTML, and returns the links in the HTML document.
// cancelChan (optional) is a channel which, when closed, indicates
//  that all active requests should be cancelled.
func Extract(url string, cancelChan <-chan struct{}) ([]string, error) {
	req, errReq := http.NewRequest("GET", url, nil)
	if errReq != nil {
		return nil, errReq
	}

	if cancelChan != nil {
		req.Cancel = cancelChan
	}

	resp, errResp := http.DefaultClient.Do(req)
	if errResp != nil {
		return nil, errResp
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("getting %s: %s", url, resp.Status)
	}

	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
	}

	var links []string
	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key != "href" {
					continue
				}
				link, err := resp.Request.URL.Parse(a.Val)
				if err != nil {
					continue // ignore bad URLs
				}
				links = append(links, link.String())
			}
		}
	}
	forEachNode(doc, visitNode, nil)
	return links, nil
}

//!-Extract

// Copied from gopl.io/ch5/outline2.
func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}
