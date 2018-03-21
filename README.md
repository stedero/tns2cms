# tns2cms
CMS meta data generator for Tax New Service articles

tns2cms adds metadata to TNS articles for bulkload in the Alfresco CMS.

```
Usage of tns2cms:
    tns2cms <input directory> <output directory>
```

An example of a TNS article:
```xml
<?xml version="1.0" encoding="utf-8"?>
<!-- Created 2018-02-22T15:32:15.176+01:00-->
<!DOCTYPE tnsarticle PUBLIC "-//IBFD//ELEMENTS TNS//EN" "http://dtd.ibfd.org/dtd/tns/tnsarticle.dtd">
<tnsarticle guid="tns_2018-02-22_am_1" suppress-in-treaty-news="n"
            reporttype="inc_tax_tt_neg" tns-report-hilites="no" tns-mli-hilites="no">
    <tnsarticleinfo>
        <countrylist main="am">
            <country cc="am">
                <countryname>Armenia</countryname>
            </country>
            <country cc="iq">
                <countryname>Iraq</countryname>
            </country>
        </countrylist>
        <topics>
            <topic tc="2_11_13" score="3">Bilateral Double Tax Relief [DTR]</topic>
            <topic tc="2_11_41" score="3">Withholding Taxes</topic>
        </topics>
        <onlinetitle>Treaty between Armenia and Iraq initialled</onlinetitle>
        <articledate isodate="20180222">22 February 2018</articledate>
        <author initials="BN"/>
        <correspondent>Report from IBFD Tax Treaties Unit</correspondent>
        <reference>
            <extxref target="tns_2018-02-19_am_1"
                     alttext="Armenia; Iraq - Treaty between Armenia and Iraq – negotiations ongoing (19 Feb. 2018), News IBFD.">Armenia-1, News 19 February 2018</extxref>
        </reference>
        <reference>
            <extxref target="tt_am-iq_01_eng_2018_tt__td1">Armenia - Iraq Income and Capital Tax Treaty</extxref>
        </reference>
        <source>see admin sheet</source>
    </tnsarticleinfo>
    <tnsarticletext>
        <p>On 22 February 2018, Armenia and Iraq initialled an income and capital tax treaty,
           following a successful round of negotiations held in Yerevan from 19 to 22 February 2018.
           Further developments will be reported as they occur.</p>
        <p>Treaty texts are published in the Treaties collection on
           <emph type="i">IBFD's Tax Research Platform</emph>, as available.</p>
			
    </tnsarticletext>
</tnsarticle>
```
And the generated properties file for this article:
```xml
<?xml version="1.0" encoding="UTF-8"?>
<!-- Generated 2018-03-21T16:36:13+01:00 -->
<!DOCTYPE properties SYSTEM "http://java.sun.com/dtd/properties.dtd">
<properties>
    <entry key="ibfd:type">cm:content</entry>
    <entry key="ibfd:id">tns_2018-02-22_am_1</entry>
    <entry key="ibfd:created">20180222</entry>
    <entry key="ibfd:report_type">inc_tax_tt_neg</entry>
    <entry key="ibfd:collection">tns</entry>
    <entry key="ibfd:title">Treaty between Armenia and Iraq initialled</entry>
    <entry key="ibfd:author_initials">BN</entry>
    <entry key="ibfd:main_cc">am</entry>
    <entry key="ibfd:country_codes">am,iq</entry>
    <entry key="ibfd:country_names">Armenia,Iraq</entry>
    <entry key="ibfd:xrefs">tns_2018-02-19_am_1,tt_am-iq_01_eng_2018_tt__td1</entry>
</properties>
```
Please note that this is not the final version of the meta data file. Waiting for specification.
