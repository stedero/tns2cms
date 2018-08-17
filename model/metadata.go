package model

import (
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"strings"
	"time"
)

const metaDataDoctype = "<!DOCTYPE properties SYSTEM \"http://java.sun.com/dtd/properties.dtd\">\n"

// MetaData element for marshaling to XML
type MetaData struct {
	XMLName xml.Name `xml:"properties"`
	Entries []Entry  `xml:"entry"`
}

// Entry element for marshaling to XML
type Entry struct {
	Key   string `xml:"key,attr"`
	Value string `xml:",innerxml"`
}

const multiValueJoiner = ","
const metaNameSpace = "ibfd:"

// NewMetaData transforms a TNS article into a meta data structure.
func NewMetaData(tnsArticle *TnsArticle) *MetaData {
	metaData := &MetaData{}
	metaData.add("type", "cm:content")
	metaData.add("id", tnsArticle.GUID)
	metaData.add("created", tnsArticle.TnsArticleInfo.ArticleDate.IsoDate)
	metaData.add("report_type", tnsArticle.ReportType)
	metaData.add("collection", tnsArticle.Collection)
	metaData.add("title", tnsArticle.TnsArticleInfo.OnlinetTitle)
	metaData.add("author_initials", tnsArticle.TnsArticleInfo.Author.Initials)
	metaData.add("main_cc", tnsArticle.TnsArticleInfo.CountryList.Main)
	metaData.add("country_codes", mapJoin(countryCode(tnsArticle)))
	metaData.add("country_names", mapJoin(countryName(tnsArticle)))
	metaData.add("xrefs", mapJoin(reference(tnsArticle)))
	return metaData
}

// WriteXML writes the metadata as XML.
func (m *MetaData) WriteXML(w io.Writer) {
	w.Write([]byte(xml.Header))
	w.Write([]byte(nowAsComment()))
	w.Write([]byte(metaDataDoctype))
	encoder := xml.NewEncoder(w)
	encoder.Indent("", "   ")
	err := encoder.Encode(m)
	if err != nil {
		log.Fatalf("error encoding XML: %v", err)
	}
}

func (m *MetaData) add(key string, value string) {
	m.Entries = append(m.Entries, Entry{metaNameSpace + key, value})
}

func nowAsComment() string {
	return fmt.Sprintf("<!-- Generated %s -->\n", time.Now().Format(time.RFC3339))
}

func mapJoin(len int, get func(int) string) string {
	result := make([]string, len)
	for i := 0; i < len; i++ {
		result[i] = get(i)
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
