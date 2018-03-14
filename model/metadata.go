package model

import (
	"encoding/xml"
	"fmt"
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
	props.add("collection", tnsArticle.Collection)
	props.add("title", tnsArticle.TnsArticleInfo.OnlinetTitle)
	props.add("main_cc", tnsArticle.TnsArticleInfo.CountryList.Main)
	props.add("author_initials", tnsArticle.TnsArticleInfo.Author.Initials)
	props.add("correspondent", tnsArticle.TnsArticleInfo.Correspondent)
	xmlMeta, err := xml.MarshalIndent(&props, "", "    ")
	if err != nil {
		panic(err)
	}
	return []byte(xml.Header + nowAsComment() + metaDataDoctype + string(xmlMeta))
}

func (p *Properties) add(key string, value string) {
	p.Entries = append(p.Entries, Entry{key, value})
}

func nowAsComment() string {
	return fmt.Sprintf("<!-- Generated %s -->\n", time.Now().Format(time.RFC3339))
}
