package model

import (
	"encoding/xml"
	"fmt"
	"log"
	"strings"
	"time"
)

const metaDataDoctype = "<!DOCTYPE properties SYSTEM \"http://java.sun.com/dtd/properties.dtd\">\n"

// Properties element for marshaling to XML
type Properties struct {
	XMLName xml.Name `xml:"properties"`
	Entries []Entry  `xml:"entry"`
}

// Entry element for marshaling to XML
type Entry struct {
	Key   string `xml:"key,attr"`
	Value string `xml:",innerxml"`
}

const multiValueJoiner = ","

// NewMetaData transforms a TNS article into meta data XML.
func NewMetaData(tnsArticle *TnsArticle) []byte {
	var props Properties
	props.add("type", "cm:content")
	props.add("id", tnsArticle.GUID)
	props.add("created", tnsArticle.TnsArticleInfo.ArticleDate.IsoDate)
	props.add("report_type", tnsArticle.ReportType)
	props.add("collection", tnsArticle.Collection)
	props.add("title", tnsArticle.TnsArticleInfo.OnlinetTitle)
	props.add("author_initials", tnsArticle.TnsArticleInfo.Author.Initials)
	props.add("main_cc", tnsArticle.TnsArticleInfo.CountryList.Main)
	props.add("country_codes", mapJoin(countryCode(tnsArticle)))
	props.add("country_names", mapJoin(countryName(tnsArticle)))
	props.add("xrefs", mapJoin(reference(tnsArticle)))
	xmlMeta, err := xml.MarshalIndent(&props, "", "    ")
	if err != nil {
		log.Fatalf("error marshaling TNS article %s to XML: %v", tnsArticle.GUID, err)
	}
	return []byte(xml.Header + nowAsComment() + metaDataDoctype + string(xmlMeta))
}

func (p *Properties) add(key string, value string) {
	p.Entries = append(p.Entries, Entry{key, value})
}

func nowAsComment() string {
	return fmt.Sprintf("<!-- Generated %s -->\n", time.Now().Format(time.RFC3339))
}

func mapJoin(max int, get func(int) string) string {
	var result []string
	for i := 0; i < max; i++ {
		result = append(result, get(i))
	}
	return strings.Join(result, multiValueJoiner)
}

func countryCode(tnsArticle *TnsArticle) (int, func(int) string) {
	return len(tnsArticle.TnsArticleInfo.CountryList.Countries), func(i int) string {
		return tnsArticle.TnsArticleInfo.CountryList.Countries[i].CC
	}
}

func countryName(tnsArticle *TnsArticle) (int, func(int) string) {
	return len(tnsArticle.TnsArticleInfo.CountryList.Countries), func(i int) string {
		return strings.Replace(tnsArticle.TnsArticleInfo.CountryList.Countries[i].Name, multiValueJoiner, " ", -1)
	}
}

func reference(tnsArticle *TnsArticle) (int, func(int) string) {
	return len(tnsArticle.TnsArticleInfo.Reference), func(i int) string {
		return tnsArticle.TnsArticleInfo.Reference[i].Target
	}
}
