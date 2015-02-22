package main 

import (

)



func main() {
	
	xmlLiteral := `
<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:urn="urn:schemas-xtrac-fmr-com:b2b">
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
	
	xml2Literal := `
	<a><b><c><d><e></e></d></c></b></a>
	`
	
	xmlLiteral3 := `
	<docRoot>
		<collection>
			<thing name='a'>a</thing>
			<thing name='b'>b</thing>
			<thing name='c'>c</thing>
		</collection>
	</docRoot>
	`
	
	xml1()
	xml2()
	streaming(xmlLiteral)
	streaming(xml2Literal)
	xpathSampleFindOne(xmlLiteral,"/Envelope/Body/AddNote/Memo")
	xpathSampleFindOne(xmlLiteral3,"/docRoot/collection/thing[2]") 
	xpathSampleFindMany(xmlLiteral3,"/docRoot/collection/thing")
}

