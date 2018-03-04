package lib

import (
	"encoding/xml"
)

const meta_data_doctype = "<!DOCTYPE properties SYSTEM \"http://java.sun.com/dtd/properties.dtd\">\n"
const meta_file_preamble = xml.Header + meta_data_doctype

type Properties struct {
	XMLName xml.Name `xml:"properties"`
	Entries []Entry  `xml:"entry"`
}

type Entry struct {
	Key   string `xml:"key,attr"`
	Value string `xml:",innerxml"`
}

// Transform a TNS article into meta data XML
func NewMetaData(tnsArticle *TnsArticle) []byte {
	var props Properties
	props.add("type", "cm:content")
	props.add("id", tnsArticle.Guid)
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
	return []byte(meta_file_preamble + string(xmlMeta))
}

func (p *Properties) add(key string, value string) {
	p.Entries = append(p.Entries, Entry{key, value})
}
