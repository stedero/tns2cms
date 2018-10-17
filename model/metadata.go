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
	Value string `xml:",chardata"`
}

const multiValueJoiner = ","
const metaNameSpace = "ibfd:"

// NewMetaData transforms a TNS article into a meta data structure.
func NewMetaData(tnsArticle *TnsArticle) *MetaData {
	metaData := &MetaData{}
	metaData.add("type", "ibfd:onlinecontent")
	metaData.add("aspects", "ibfd:onlineContentProperties,cm:titled,ibfd:published")
	metaData.add("cm:title", tnsArticle.TnsArticleInfo.OnlineTitle)
	metaData.add("ibfd:collectionCode", "tns")
	metaData.add("ibfd:collectionCodeHumanReadable", "Tax News service")
	metaData.add("ibfd:publicationDate", formatDate(tnsArticle.TnsArticleInfo.ArticleDate.IsoDate))

	// metaData.add("id", tnsArticle.GUID)
	// metaData.add("report_type", tnsArticle.ReportType)
	// metaData.add("ibfd:collectionCode", tnsArticle.Collection)
	// metaData.add("title", tnsArticle.TnsArticleInfo.OnlinetTitle)
	// metaData.add("author_initials", tnsArticle.TnsArticleInfo.Author.Initials)
	// metaData.add("main_cc", tnsArticle.TnsArticleInfo.CountryList.Main)
	// metaData.add("country_codes", mapJoin(countryCode(tnsArticle)))
	// metaData.add("country_names", mapJoin(countryName(tnsArticle)))
	// metaData.add("xrefs", mapJoin(reference(tnsArticle)))
	return metaData
}

// WriteXML writes the metadata as XML.
func (m *MetaData) WriteXML(w io.Writer) {
	writeString(w, xml.Header)
	writeString(w, nowAsComment())
	writeString(w, metaDataDoctype)
	encoder := xml.NewEncoder(w)
	encoder.Indent("", "   ")
	err := encoder.Encode(m)
	if err != nil {
		log.Fatalf("error encoding XML: %v", err)
	}
}

func (m *MetaData) add(key string, value string) {
	m.Entries = append(m.Entries, Entry{key, value})
}

func nowAsComment() string {
	return fmt.Sprintf("<!-- Generated by tns2cms on %s -->\n", time.Now().Format(time.RFC3339))
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

func writeString(w io.Writer, s string) {
	_, err := io.WriteString(w, s)
	if err != nil {
		log.Fatalf("error writing %s: %v", s, err)
	}
}

func formatDate(date string) string {
	return date[:4] + "-" + date[4:6] + "-" + date[6:8] + "T00:00:00.000+02:00"
}
