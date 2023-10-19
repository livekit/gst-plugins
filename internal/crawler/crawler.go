// Copyright 2023 LiveKit, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package crawler

import (
	"bytes"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly/v2"
	"github.com/iancoleman/strcase"

	"github.com/livekit/gstplugins/internal/types"
)

const (
	baseClassURL = "https://gstreamer.freedesktop.org/documentation/base/index.html?gi-language=c"
	pluginsURL   = "https://gstreamer.freedesktop.org/documentation/plugins_doc.html?gi-language=c"
)

func ReadElements() map[string]*types.Element {
	elements := make(map[string]*types.Element)

	baseCrawler := colly.NewCollector(
		colly.AllowedDomains("gstreamer.freedesktop.org"),
		colly.CacheDir("./internal/crawler/cache"),
	)
	pluginCrawler := colly.NewCollector(
		colly.AllowedDomains("gstreamer.freedesktop.org"),
		colly.CacheDir("./internal/crawler/cache"),
	)
	elementCrawler := pluginCrawler.Clone()

	baseCrawler.OnHTML("#subpages", func(div *colly.HTMLElement) {
		div.ForEach("a", func(_ int, e *colly.HTMLElement) {
			if strings.HasPrefix(e.Text, "Gst") {
				elementCrawler.Visit(e.Request.AbsoluteURL(e.Attr("href")) + "?gi-language=c")
			}
		})
	})

	pluginCrawler.OnHTML("tbody tr td a", func(e *colly.HTMLElement) {
		elementCrawler.Visit(e.Request.AbsoluteURL(e.Attr("href")) + "?gi-language=c")
	})

	elementCrawler.OnResponse(func(resp *colly.Response) {
		doc, _ := goquery.NewDocumentFromReader(bytes.NewBuffer(resp.Body))

		isBase := false
		baseClass := ""
		shouldCrawl := false
		if strings.HasPrefix(resp.Request.URL.String(), "https://gstreamer.freedesktop.org/documentation/base/") {
			isBase = true
			shouldCrawl = true
		} else {
			doc.Find("h2").Each(func(_ int, s *goquery.Selection) {
				if s.Text() == "Factory details" {
					s.Next().Children().Each(func(_ int, child *goquery.Selection) {
						if strings.HasPrefix(child.Text(), "Package – GStreamer") {
							shouldCrawl = true
						}
					})
				} else if s.Text() == "Hierarchy" {
					s.Next().Find("a").Each(func(_ int, child *goquery.Selection) {
						href, _ := child.Attr("href")
						if strings.HasPrefix(href, "base/") {
							baseClass, _ = child.Attr("title")
						}
					})
				}
			})
		}
		if !shouldCrawl {
			return
		}

		factoryName := doc.Find("title").Text()
		elementName := normalizeElementName(factoryName, isBase)
		element := &types.Element{
			Name:        elementName,
			IsBase:      isBase,
			BaseClass:   baseClass,
			FactoryName: factoryName,
			Properties:  make([]*types.Property, 0),
			Constants:   make(map[string]*types.Constant),
		}

		i := 0
		existingProperties := make(map[string]bool)

		propSelector := "pre.property_signature"
		if isBase {
			propSelector = "div.gi-symbol-c pre.property_signature"
		}
		doc.Find(propSelector).Each(func(_ int, s *goquery.Selection) {
			for _, n := range s.Nodes {
				e := colly.NewHTMLElementFromSelectionNode(resp, s, n, i)
				i++

				text := strings.Split(strings.TrimSpace(e.Text), " ")
				name := strings.Trim(text[0], "“”")
				dataType := toGoType(text[1])
				description := e.DOM.Next().Text()
				flags := e.DOM.Next().Next().Text()
				if strings.Contains(description, "Flags :") {
					flags = description
					description = ""
				}
				readable := strings.Contains(flags, "Read")
				writable := strings.Contains(flags, "Write")

				if !existingProperties[name] {
					element.Properties = append(element.Properties, &types.Property{
						Name:        name,
						DataType:    dataType,
						Description: description,
						Readable:    readable,
						Writable:    writable,
					})
					existingProperties[name] = true
				}
			}
		})

		doc.Find("div.member_details").Each(func(_ int, s *goquery.Selection) {
			for _, n := range s.Nodes {
				e := colly.NewHTMLElementFromSelectionNode(resp, s, n, i)
				i++

				id := strings.Split(e.Attr("data-hotdoc-id"), "::")
				name := strcase.ToCamel(id[0])
				member := id[1]
				desc := strings.TrimSpace(e.Text)

				constant := element.Constants[name]
				if constant == nil {
					constant = &types.Constant{
						Name:         name,
						Members:      make([]string, 0),
						Descriptions: make([]string, 0),
					}
					element.Constants[name] = constant
				}

				constant.Members = append(constant.Members, member)
				constant.Descriptions = append(constant.Descriptions, desc)
			}
		})

		elements[elementName] = element
	})

	// Start scraping
	baseCrawler.Visit(baseClassURL)
	pluginCrawler.Visit(pluginsURL)
	// elementCrawler.Visit("https://gstreamer.freedesktop.org/documentation/coreelements/filesink.html?gi-language=c")

	return elements
}

func normalizeElementName(name string, isBase bool) string {
	if !isBase && strings.HasPrefix(name, "Gst") {
		name = name[3:]
	}
	if strings.HasPrefix(name, "3") {
		name = "_" + name
	}
	if !isBase {
		name = strings.ToLower(name)
	}
	name = strings.ReplaceAll(name, "-", "_")
	return name
}

func toGoType(gType string) string {
	switch gType {
	case "gboolean":
		return "bool"
	case "gint":
		return "int"
	case "guint":
		return "uint"
	case "gint8":
		return "int8"
	case "guint8":
		return "uint8"
	case "gint16":
		return "int16"
	case "guint16":
		return "uint16"
	case "gint32":
		return "int32"
	case "guint32":
		return "uint32"
	case "gint64":
		return "int64"
	case "guint64":
		return "uint64"
	case "gfloat":
		return "float32"
	case "gdouble":
		return "float64"
	case "gchararray":
		return "string"
	default:
		if strings.HasPrefix(gType, "Gst") {
			return gType
		}
		return "Gst" + strcase.ToCamel(gType)
	}
}
