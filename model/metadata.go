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
	props.add("country_codes", joinCountryCodes(tnsArticle))
	props.add("country_names", joinCountryNames(tnsArticle))
	props.add("xrefs", joinXrefs(tnsArticle))
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

func joinCountryCodes(tnsArticle *TnsArticle) string {
	var countries []string
	for _, country := range tnsArticle.TnsArticleInfo.CountryList.Countries {
		countries = append(countries, country.CC)
	}
	return strings.Join(countries, ",")
}

func joinCountryNames(tnsArticle *TnsArticle) string {
	var countries []string
	for _, country := range tnsArticle.TnsArticleInfo.CountryList.Countries {
		countries = append(countries, country.Name)
	}
	return strings.Join(countries, ",")
}

func joinXrefs(tnsArticle *TnsArticle) string {
	var xrefs []string
	for _, ref := range tnsArticle.TnsArticleInfo.Reference {
		xrefs = append(xrefs, ref.Target)
	}
	return strings.Join(xrefs, ",")
}
