package main

/*
#cgo CFLAGS: -I/usr/include/libxml2
#cgo LDFLAGS: -L/usr/lib -l xml2 -l xslt
#include <stdlib.h>
#include <libxml/parser.h>
#include <libxslt/xslt.h>
#include <libxslt/xsltInternals.h>
#include <libxslt/transform.h>
#include <libxslt/xsltutils.h>

static void
print_element_names(xmlNode * a_node)
{
    xmlNode *cur_node = NULL;

    for (cur_node = a_node; cur_node; cur_node = cur_node->next) {
        if (cur_node->type == XML_ELEMENT_NODE) {
            printf("node type: Element, name: %s\n", cur_node->name);
        }

        print_element_names(cur_node->children);
    }
}

static void
print_stylesheet_result(xmlDocPtr res, xsltStylesheetPtr styleSheet) {
	xsltSaveResultToFile(stdout, res, styleSheet);
}
*/
import "C"

import (
	"fmt"
	"unsafe"
)

func main() {
	fmt.Println("grab xml bytes")
	xmlDoc := `
<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:urn="urn:schemas-xtrac-foo-com:b2b">
   <soapenv:Header>
	<urn:Cookie>261765034988290296725527451071864505174</urn:Cookie>
   </soapenv:Header>
   <soapenv:Body>
      <urn:AddNote>
         <urn:WorkItemNo>W019039-27NOV01</urn:WorkItemNo>
         <urn:Name></urn:Name>
         <urn:ControlNo>123</urn:ControlNo>
         <urn:Memo>some notes</urn:Memo>
         <urn:Note>c e-flat f-sharp a</urn:Note>
         <urn:Foo a='1' b='2'/>
      </urn:AddNote>
   </soapenv:Body>
</soapenv:Envelope>
	`

	fmt.Println("initialize xml library")

	fmt.Println("read xml from memory")
	docChars := C.CString(xmlDoc)
	baseName := C.CString("noname.xml")
	defer C.free(unsafe.Pointer(baseName))
	defer C.free(unsafe.Pointer(docChars))
	doc := C.xmlReadMemory(docChars, C.int(len(xmlDoc)), baseName, nil, 0)

	fmt.Println("null doc?")
	root := C.xmlDocGetRootElement(doc)

	C.print_element_names(root)

	//XSLT
	fmt.Println("Read style sheet from file")
	xsltFileName := C.CString("./xml2json.xslt")
	defer C.free(unsafe.Pointer(xsltFileName))
	styleSheet := C.xsltParseStylesheetFile((*C.xmlChar)(unsafe.Pointer(xsltFileName)))
	res := C.xsltApplyStylesheet(styleSheet, doc, nil)
	C.print_stylesheet_result(res, styleSheet)

	fmt.Println("free C stuff")
	C.xsltFreeStylesheet(styleSheet)
	C.xmlFreeDoc(res)
	C.xmlFreeDoc(doc)
	fmt.Println("done")
}
