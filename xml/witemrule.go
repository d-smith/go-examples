package main

var workItemRuleResponse = `
<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:soapenc="http://schemas.xmlsoap.org/soap/encoding/" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance">
   <soapenv:Header>
      <xt:AppServer xmlns:xt="http://xmlns.foo.com/systems/dev/xtrac/2004/06/">server1</xt:AppServer>
   </soapenv:Header>
   <soapenv:Body>
      <p347:retrieveRuntimeWorkItemRuleResponse xmlns:p347="http://xmlns.foo.com/systems/dev/xtrac/2004/06/">
         <p347:retrieveRuntimeWorkItemRuleReturn xmlns:p972="http://xmlns.foo.com/systems/dev/xtrac/2004/06/types">
            <p972:node>XWBViewer</p972:node>
            <p972:itemType>DSPLT1</p972:itemType>
            <p972:requireResolutionNote>0</p972:requireResolutionNote>
            <p972:workItemClassification>UserDefined</p972:workItemClassification>
            <p972:active>1</p972:active>
            <p972:size>
               <p972:height>379</p972:height>
               <p972:width>342</p972:width>
            </p972:size>
            <p972:duplicateSearchRules/>
            <p972:runtimeFieldRules>
               <p972:runtimeFieldRule>
                  <p972:fieldId>100</p972:fieldId>
                  <p972:visible>1</p972:visible>
                  <p972:description>Item Type</p972:description>
                  <p972:name>ItemType</p972:name>
                  <p972:dataType>Value Group</p972:dataType>
                  <p972:displayMaskName>NoCheck</p972:displayMaskName>
                  <p972:length>30</p972:length>
                  <p972:updateState>Required</p972:updateState>
                  <p972:resolutionState>Required</p972:resolutionState>
                  <p972:createState>Required</p972:createState>
                  <p972:location>
                     <p972:xCoordinate>0</p972:xCoordinate>
                     <p972:yCoordinate>0</p972:yCoordinate>
                  </p972:location>
                  <p972:size>
                     <p972:height>24</p972:height>
                     <p972:width>65</p972:width>
                  </p972:size>
                  <p972:defaultValue xsi:nil="true"/>
                  <p972:runtimeValidValues>
                     <p972:runtimeValidValue>
                        <p972:value>DSPLT1</p972:value>
                        <p972:description>Demo Split1</p972:description>
                     </p972:runtimeValidValue>
                  </p972:runtimeValidValues>
                  <p972:displayLabel>ReplaceWithDescription</p972:displayLabel>
                  <p972:externalField>0</p972:externalField>
                  <p972:globalField>0</p972:globalField>
               </p972:runtimeFieldRule>
               <p972:runtimeFieldRule>
                  <p972:fieldId>200</p972:fieldId>
                  <p972:visible>1</p972:visible>
                  <p972:description>Subtype</p972:description>
                  <p972:name>Subtype</p972:name>
                  <p972:dataType>Value Group</p972:dataType>
                  <p972:displayMaskName>NoCheck</p972:displayMaskName>
                  <p972:length>30</p972:length>
                  <p972:updateState>Optional</p972:updateState>
                  <p972:resolutionState>Optional</p972:resolutionState>
                  <p972:createState>Optional</p972:createState>
                  <p972:location>
                     <p972:xCoordinate>70</p972:xCoordinate>
                     <p972:yCoordinate>0</p972:yCoordinate>
                  </p972:location>
                  <p972:size>
                     <p972:height>24</p972:height>
                     <p972:width>65</p972:width>
                  </p972:size>
                  <p972:defaultValue xsi:nil="true"/>
                  <p972:runtimeValidValues>
                     <p972:runtimeValidValue>
                        <p972:value>407 Letter</p972:value>
                        <p972:description>407 Letter</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>ACAT</p972:value>
                        <p972:description>ACAT</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>Account Agreement</p972:value>
                        <p972:description>Account Agreement</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>ACH Authorization</p972:value>
                        <p972:description>ACH Authorization</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>Address Change</p972:value>
                        <p972:description>ECP Address Change</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>Advisor Invoices</p972:value>
                        <p972:description>Advisor Invoices</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>AMEND</p972:value>
                        <p972:description>Account Amendments</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>Annual Statements</p972:value>
                        <p972:description xsi:nil="true"/>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>APP</p972:value>
                        <p972:description>Application</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>Auto</p972:value>
                        <p972:description>Auto</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>Auto Loans</p972:value>
                        <p972:description>Auto Loans</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>Bankwire</p972:value>
                        <p972:description>Bankwire</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>Batch</p972:value>
                        <p972:description>ECP Batch</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>BCOA</p972:value>
                        <p972:description>Broker Change of Addressr</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>Beneficiary Update</p972:value>
                        <p972:description>Beneficiary Update</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>BF</p972:value>
                        <p972:description>Bond Funds</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>Billing</p972:value>
                        <p972:description>Order Billing Information</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>Bills</p972:value>
                        <p972:description>Bills</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>BROKEO</p972:value>
                        <p972:description>Broker Offshore</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>BROKER</p972:value>
                        <p972:description>Broker transfer W/I</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>Business Loans</p972:value>
                        <p972:description>Business Loans</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>BUY</p972:value>
                        <p972:description>Buy Deal</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>CANPEP</p972:value>
                        <p972:description>PEP Cancellations</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>CANUT</p972:value>
                        <p972:description>Unit Trust Cancellations</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>Cash Management Feature</p972:value>
                        <p972:description>Cash Management Feature</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>CBDBSU</p972:value>
                        <p972:description>Customer Change of Address</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>CCOA</p972:value>
                        <p972:description>Customer Change of Address</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>CEDEL</p972:value>
                        <p972:description>Cedel Deal Instructions</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>CF</p972:value>
                        <p972:description>Cash Funds</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>CFMF</p972:value>
                        <p972:description>Cash Funds Money Funds</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>Check</p972:value>
                        <p972:description>Check</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>Check Authorization</p972:value>
                        <p972:description>Check Authorization</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>Check Reorder</p972:value>
                        <p972:description>Check Reorder</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>Civil</p972:value>
                        <p972:description>Civil Work</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>CLARAP</p972:value>
                        <p972:description>Explanation of admin. procedur</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>CLARML</p972:value>
                        <p972:description>Explanation of Marketing Lit.</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>CLARNI</p972:value>
                        <p972:description>Explanation of nature of inves</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>CLARRB</p972:value>
                        <p972:description>Other</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>CLARRE</p972:value>
                        <p972:description>Explanation of regulations</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>CLARTI</p972:value>
                        <p972:description>Explanation of Taxation Issues</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>CLARUB</p972:value>
                        <p972:description>UK Products Clarification BRK</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>CLARUR</p972:value>
                        <p972:description>UK Products Clarification Ret</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>Client Update</p972:value>
                        <p972:description>Client Update</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>CMF</p972:value>
                        <p972:description>Cash Management Feature</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>COA</p972:value>
                        <p972:description>Change of Address</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>COAA</p972:value>
                        <p972:description>Change of Address Agent</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>COAB</p972:value>
                        <p972:description>Agent Change of Address</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>Collectibles</p972:value>
                        <p972:description>Collectibles</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>Contribution Authorization</p972:value>
                        <p972:description>Contribution Authorization</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>Core Fund Change</p972:value>
                        <p972:description>Core Fund Change</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>Correspondence</p972:value>
                        <p972:description>Correspondence</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>Cost Basis</p972:value>
                        <p972:description>Cost Basis</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>CPBRCC</p972:value>
                        <p972:description>Clarification of comm. terms</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>CPBRCP</p972:value>
                        <p972:description>Non receipt of commission</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>CPBRIC</p972:value>
                        <p972:description>Incorrect commission</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>CPDBCD</p972:value>
                        <p972:description>Incorrect commission</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>CPDBCM</p972:value>
                        <p972:description>Rec.of another client's mail</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>CPDBRB</p972:value>
                        <p972:description>Non receipt of documentation</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>CPDBUE</p972:value>
                        <p972:description>Unsealed evelopes received</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>CPDLBD</p972:value>
                        <p972:description>Buy Deal Error</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>CPDLIS</p972:value>
                        <p972:description>Interoffice Switch Deal Error</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>CPDLNI</p972:value>
                        <p972:description>Deal NIGO</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>CPDLSD</p972:value>
                        <p972:description>Sell Deal Error</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>CPDLSS</p972:value>
                        <p972:description>Switch Deal Error</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>CPDVCD</p972:value>
                        <p972:description>Change distribution status</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>CPDVDR</p972:value>
                        <p972:description>Explanation of dividends requi</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>CPDVRD</p972:value>
                        <p972:description>Reissue of divi/income payment</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>CPDVTE</p972:value>
                        <p972:description>Tax Voucher Explanation</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>CPDVTR</p972:value>
                        <p972:description>Tax Voucher Request</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>CPEROD</p972:value>
                        <p972:description>Clarification of ivmt decision</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>CPEXNI</p972:value>
                        <p972:description>Exits NIGO</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>CPHBCT</p972:value>
                        <p972:description>Certificate to be forwarded</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>CPHBFR</p972:value>
                        <p972:description>Fulfilment request</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>CPHBMM</p972:value>
                        <p972:description>Marketing mailings</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>CPHBRA</p972:value>
                        <p972:description>Confirm client records ag. reg</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>CPHBSM</p972:value>
                        <p972:description>Broker Seminar Material</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>CPHBUH</p972:value>
                        <p972:description>Confirm uncertificated holding</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>CPLBFR</p972:value>
                        <p972:description>Fulfilment request</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>CPLBMM</p972:value>
                        <p972:description>Marketing mailings</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>CPLBSM</p972:value>
                        <p972:description>Broker Seminar Material</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>CPNDCN</p972:value>
                        <p972:description>Duplicate Contract Note</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>CPNPCV</p972:value>
                        <p972:description>Reissue Chg &amp; Voucher</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>CPNPM</p972:value>
                        <p972:description>Reissue Chg &amp; Voucher</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>CPNPRC</p972:value>
                        <p972:description>Cheque Book Request</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>CPNPRI</p972:value>
                        <p972:description>Re Invest Dividend</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>CPNPRS</p972:value>
                        <p972:description>Statement &amp; Valuation Request</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>CPNPRT</p972:value>
                        <p972:description>Transaction Book Request</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>CPNRD</p972:value>
                        <p972:description>Re Issue Dividend</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>CPNREG</p972:value>
                        <p972:description>Customer Order Registratio</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>CPPDCR</p972:value>
                        <p972:description>Clarification required at settl</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>CPPDNG</p972:value>
                        <p972:description>NIGO</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>CPRERP</p972:value>
                        <p972:description>Expectation exceeds performanc</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>CPRETC</p972:value>
                        <p972:description>Redemption Returns</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>CPRETR</p972:value>
                        <p972:description>Commission Returns</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>CPRETT</p972:value>
                        <p972:description>Treasury</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>CPRGCH</p972:value>
                        <p972:description>Change to holders of account</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>CPRGFJ</p972:value>
                        <p972:description>FAXA/JTRE to be amended</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>CPRGMA</p972:value>
                        <p972:description>Minor amendment to reg. of ac.</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>CPRGMH</p972:value>
                        <p972:description>Mail Held to be added</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>CPRGRI</p972:value>
                        <p972:description>Registration Issue - deceased</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>CPRGST</p972:value>
                        <p972:description>Stock Transfer form to be issu</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>CPSBMD</p972:value>
                        <p972:description>Mandate Problem</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>CPSBMN</p972:value>
                        <p972:description>Sell sett. - mandate problems</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>CPSBNF</p972:value>
                        <p972:description>Non Receipt of Funds</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>CPSBNP</p972:value>
                        <p972:description>Buy Sett. - non receipt of pay</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>CPSBNR</p972:value>
                        <p972:description>Buy Sett. - non receipt of reg</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>CPWKSV</p972:value>
                        <p972:description>Statement &amp; Valuation</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>CPWKWC</p972:value>
                        <p972:description>Written confirmation of other</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>CPWKWR</p972:value>
                        <p972:description>Written reiteration of telephone</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>Customer Complaint</p972:value>
                        <p972:description>Customer Complaint</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>DALL</p972:value>
                        <p972:description>Demo, one TSF of each type &amp; all other fields</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>DBOOL</p972:value>
                        <p972:description>Demo, all typs of Boolean TSF's</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>DBRKR</p972:value>
                        <p972:description>Demo Broker</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>DCHAR</p972:value>
                        <p972:description>Demo, all types of Character TSF's</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>DCOR1</p972:value>
                        <p972:description>Demo Correspondence Use Sub Type 1</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>DDATE</p972:value>
                        <p972:description>Demo, all types of Date TSF's</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>DDEC</p972:value>
                        <p972:description>Demo, all types of Decimal TSF's</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>DEAL</p972:value>
                        <p972:description>Switch, Sell or Special</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>DEALF1</p972:value>
                        <p972:description>deal form one</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>Deeds</p972:value>
                        <p972:description>Deeds</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>DIDEN</p972:value>
                        <p972:description>Demo, all types of Identifier TSF's</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>DINTE</p972:value>
                        <p972:description>Demo, all types of Integer TSF's</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>DIR</p972:value>
                        <p972:description>DIR</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>DIRA</p972:value>
                        <p972:description>Demo Individual Retirement Acc</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>DIVID</p972:value>
                        <p972:description>Dividend Changes</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>Dividend</p972:value>
                        <p972:description>Dividend Options</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>Dividend Options</p972:value>
                        <p972:description>Dividend Options</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>DMFND</p972:value>
                        <p972:description>Demo Mutual Fund</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>DMONE</p972:value>
                        <p972:description>Demo, all types of Money TSF's</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>Donations</p972:value>
                        <p972:description>Donations</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>DOP</p972:value>
                        <p972:description>Deed of Pledge</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>DPHONE</p972:value>
                        <p972:description>Demo, all types of Phone TSF's</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>DUSA</p972:value>
                        <p972:description>Demo Ultra Service Account</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>DVALG</p972:value>
                        <p972:description>Demo, all types of Value Group TSF's</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>Educational Loans</p972:value>
                        <p972:description>Educational Loans</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>EFT</p972:value>
                        <p972:description>EFT</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>FF</p972:value>
                        <p972:description>Fidelity Funds</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>FFBF</p972:value>
                        <p972:description>Fidelity Funds Bond Funds</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>Financial Plans</p972:value>
                        <p972:description>Financial Plans</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>FSWP</p972:value>
                        <p972:description>FSWP</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>G Number</p972:value>
                        <p972:description>G Number Change</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>G Number Change</p972:value>
                        <p972:description>G Number Change</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>General Inquiry</p972:value>
                        <p972:description>General Inquiry</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>General Maintenance</p972:value>
                        <p972:description>General Maintenance</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>Hardware</p972:value>
                        <p972:description>COLT Hardware</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>Health CP</p972:value>
                        <p972:description>Health Care Proxies/DNR</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>Home Equity Loans</p972:value>
                        <p972:description>Home Equity Loans</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>Householding</p972:value>
                        <p972:description>Householding</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>IBP Authorization</p972:value>
                        <p972:description>IBP Authorization</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>International Bankwire</p972:value>
                        <p972:description>International Bankwire</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>IOS</p972:value>
                        <p972:description>InterOffice Switch Deal</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>Journal</p972:value>
                        <p972:description>Journal</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>Journal Authorization</p972:value>
                        <p972:description>Journal Authorization</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>LEGAL</p972:value>
                        <p972:description>Legal Issues</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>Life</p972:value>
                        <p972:description>Life</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>Long Term Care</p972:value>
                        <p972:description>Long Term Care</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>MAINT</p972:value>
                        <p972:description>Maintenance</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>Managed</p972:value>
                        <p972:description>Managed</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>Margin</p972:value>
                        <p972:description>Margin</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>Master Order</p972:value>
                        <p972:description>Master VPN Order Item</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>Medical</p972:value>
                        <p972:description>Medical</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>MF</p972:value>
                        <p972:description>Money Funds</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>MISC</p972:value>
                        <p972:description>Miscellaneous</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>Miscellaneous</p972:value>
                        <p972:description>Miscellaneous</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>Monthly Statements</p972:value>
                        <p972:description>Monthly Statements</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>MRD</p972:value>
                        <p972:description>MRD</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>MRT</p972:value>
                        <p972:description>MRT</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>MSP</p972:value>
                        <p972:description>Monthly Savings Plan</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>MSPM</p972:value>
                        <p972:description>Monthly Savings Plan Maintenance</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>MSPUNP</p972:value>
                        <p972:description>MSP Unpaid</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>Multi Acc Maint</p972:value>
                        <p972:description>Multiple Account Maintenance</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>Multiple Account Maintenance</p972:value>
                        <p972:description>Multiple Account Maintenance</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>Multiple Margin</p972:value>
                        <p972:description>Multiple Margin</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>NEWB</p972:value>
                        <p972:description>New Business</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>Non-ACAT</p972:value>
                        <p972:description>Non-ACAT</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>Non-Retirement</p972:value>
                        <p972:description>Non-Retirement</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>OMAPP</p972:value>
                        <p972:description>Open Application</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>OMBUY</p972:value>
                        <p972:description>Open Buy Deal</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>OMMSP</p972:value>
                        <p972:description>Open MSP Application</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>Omnibus Form</p972:value>
                        <p972:description>Omnibus Form</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>OMSPMN</p972:value>
                        <p972:description>Offshore Savings Plan Maintena</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>OPAPP</p972:value>
                        <p972:description>Open Application</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>OPBUY</p972:value>
                        <p972:description>Open Buy Deal</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>OPMSP</p972:value>
                        <p972:description>Open MSP Application</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>Options</p972:value>
                        <p972:description>Options</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>Other</p972:value>
                        <p972:description>Other</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>PDLBA</p972:value>
                        <p972:description>UK PEP Automated Buy</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>Periodic Investment Plan</p972:value>
                        <p972:description>Periodic Investment Plan</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>Periodic Withdrawal Plan</p972:value>
                        <p972:description>Periodic Withdrawal Plan</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>Personal</p972:value>
                        <p972:description>Personal</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>Personal Loans</p972:value>
                        <p972:description>Personal Loans</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>Portfolio Reviews</p972:value>
                        <p972:description xsi:nil="true"/>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>Power of Attorney</p972:value>
                        <p972:description>Power of Attorney</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>Prime Brokerage</p972:value>
                        <p972:description>Prime Brokerage</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>Project Mgmt</p972:value>
                        <p972:description>Order Proj Management Request</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>Property Insurance</p972:value>
                        <p972:description>Property Insurance</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>Proxy Mailing</p972:value>
                        <p972:description>Proxy Mailing</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>PWP</p972:value>
                        <p972:description>Periodic Withdrawal Plan</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>PWP Change</p972:value>
                        <p972:description>PWP Change</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>Quarterly Statements</p972:value>
                        <p972:description>Quarterly Statements</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>RBAC Sub type1</p972:value>
                        <p972:description>RBAC testing</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>RBAC Sub type2</p972:value>
                        <p972:description>RBAC Testing</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>RBAC Sub type3</p972:value>
                        <p972:description>RBAC testing</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>RBAC Sub type4</p972:value>
                        <p972:description>RBAC Testing</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>RBAC Sub type5</p972:value>
                        <p972:description>RBAC Testing</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>RBAC Sub type6</p972:value>
                        <p972:description>RBAC Testing</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>Real Estate</p972:value>
                        <p972:description>Real Estate</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>Receipts</p972:value>
                        <p972:description>Receipts</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>Retirement</p972:value>
                        <p972:description>Retirement</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>ROE</p972:value>
                        <p972:description>Return of Excess</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>Roth Conversion</p972:value>
                        <p972:description>Roth Conversion</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>Roth Rechar</p972:value>
                        <p972:description>Roth Recharacterization</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>Roth Recharacterization</p972:value>
                        <p972:description>Roth Recharacterization</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>Short Name Change</p972:value>
                        <p972:description>Short Name Change</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>Software</p972:value>
                        <p972:description>COLT Software</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>Solution</p972:value>
                        <p972:description>Order Solution</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>SpecMax</p972:value>
                        <p972:description>Special And International Characters and Max Length Testing.</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>SWP Maint</p972:value>
                        <p972:description>SWP Maintenance</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>Tax</p972:value>
                        <p972:description>Tax</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>Tax Lot Accounting</p972:value>
                        <p972:description>Tax Lot Accounting</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>Tax Returns</p972:value>
                        <p972:description>Tax Returns</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>Technical Detail</p972:value>
                        <p972:description>Order Tech Details</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>Trading Complaint</p972:value>
                        <p972:description>Trading Complaint</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>Trusts</p972:value>
                        <p972:description>Trusts</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>U4/U5</p972:value>
                        <p972:description>U4/U5</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>Upload</p972:value>
                        <p972:description>Upload</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>W8</p972:value>
                        <p972:description>W8</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>W9</p972:value>
                        <p972:description>W9</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>Wills</p972:value>
                        <p972:description>Wills</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>Wire</p972:value>
                        <p972:description>Wire</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>Wire Authorization</p972:value>
                        <p972:description>Wire Authorization</p972:description>
                     </p972:runtimeValidValue>
                     <p972:runtimeValidValue>
                        <p972:value>Wiring</p972:value>
                        <p972:description>COLT Wiring</p972:description>
                     </p972:runtimeValidValue>
                  </p972:runtimeValidValues>
                  <p972:displayLabel>ReplaceWithDescription</p972:displayLabel>
                  <p972:externalField>0</p972:externalField>
                  <p972:globalField>0</p972:globalField>
               </p972:runtimeFieldRule>
               <p972:runtimeFieldRule>
                  <p972:fieldId>400</p972:fieldId>
                  <p972:visible>1</p972:visible>
                  <p972:description>Status</p972:description>
                  <p972:name>Status</p972:name>
                  <p972:dataType>Value Group</p972:dataType>
                  <p972:displayMaskName>NoCheck</p972:displayMaskName>
                  <p972:length>30</p972:length>
                  <p972:updateState>Required</p972:updateState>
                  <p972:resolutionState>Required</p972:resolutionState>
                  <p972:createState>Required</p972:createState>
                  <p972:location>
                     <p972:xCoordinate>140</p972:xCoordinate>
                     <p972:yCoordinate>0</p972:yCoordinate>
                  </p972:location>
                  <p972:size>
                     <p972:height>24</p972:height>
                     <p972:width>65</p972:width>
                  </p972:size>
                  <p972:defaultValue xsi:nil="true"/>
                  <p972:runtimeValidValues xsi:nil="true"/>
                  <p972:displayLabel>ReplaceWithDescription</p972:displayLabel>
                  <p972:externalField>0</p972:externalField>
                  <p972:globalField>0</p972:globalField>
               </p972:runtimeFieldRule>
               <p972:runtimeFieldRule>
                  <p972:fieldId>830</p972:fieldId>
                  <p972:visible>1</p972:visible>
                  <p972:description>Work Item Memo</p972:description>
                  <p972:name>Memo</p972:name>
                  <p972:dataType>Character</p972:dataType>
                  <p972:displayMaskName>NoCheck</p972:displayMaskName>
                  <p972:length>60</p972:length>
                  <p972:updateState>Optional</p972:updateState>
                  <p972:resolutionState>Optional</p972:resolutionState>
                  <p972:createState>Optional</p972:createState>
                  <p972:location>
                     <p972:xCoordinate>0</p972:xCoordinate>
                     <p972:yCoordinate>28</p972:yCoordinate>
                  </p972:location>
                  <p972:size>
                     <p972:height>24</p972:height>
                     <p972:width>215</p972:width>
                  </p972:size>
                  <p972:defaultValue xsi:nil="true"/>
                  <p972:runtimeValidValues xsi:nil="true"/>
                  <p972:displayLabel>ReplaceWithDescription</p972:displayLabel>
                  <p972:externalField>0</p972:externalField>
                  <p972:globalField>0</p972:globalField>
               </p972:runtimeFieldRule>
            </p972:runtimeFieldRules>
            <p972:runtimeFieldGroups/>
            <p972:runtimePartyRules>
               <p972:runtimePartyRule>
                  <p972:name>ORIG</p972:name>
                  <p972:description>Originator</p972:description>
                  <p972:requiredOnCreate>1</p972:requiredOnCreate>
                  <p972:requiredOnResolution>1</p972:requiredOnResolution>
               </p972:runtimePartyRule>
            </p972:runtimePartyRules>
            <p972:statusValidValues>
               <p972:statusValidValue>
                  <p972:value>ACCREQ</p972:value>
                  <p972:description>Account Requird for Dealing</p972:description>
                  <p972:resolutionValue>InProcess</p972:resolutionValue>
               </p972:statusValidValue>
               <p972:statusValidValue>
                  <p972:value>ANLP</p972:value>
                  <p972:description>Anlage Plan</p972:description>
                  <p972:resolutionValue>InProcess</p972:resolutionValue>
               </p972:statusValidValue>
               <p972:statusValidValue>
                  <p972:value>APRQ</p972:value>
                  <p972:description>Approval required.</p972:description>
                  <p972:resolutionValue>Pending</p972:resolutionValue>
               </p972:statusValidValue>
               <p972:statusValidValue>
                  <p972:value>Assign Project Mgr</p972:value>
                  <p972:description>Assign a Project Mgr</p972:description>
                  <p972:resolutionValue>InProcess</p972:resolutionValue>
               </p972:statusValidValue>
               <p972:statusValidValue>
                  <p972:value>Bill Customer</p972:value>
                  <p972:description>Bill the Customer</p972:description>
                  <p972:resolutionValue>InProcess</p972:resolutionValue>
               </p972:statusValidValue>
               <p972:statusValidValue>
                  <p972:value>BROKER</p972:value>
                  <p972:description>Agent Required</p972:description>
                  <p972:resolutionValue>InProcess</p972:resolutionValue>
               </p972:statusValidValue>
               <p972:statusValidValue>
                  <p972:value>Calculate Charges</p972:value>
                  <p972:description>Calculate Customer Charges</p972:description>
                  <p972:resolutionValue>InProcess</p972:resolutionValue>
               </p972:statusValidValue>
               <p972:statusValidValue>
                  <p972:value>Canceled</p972:value>
                  <p972:description>Canceled</p972:description>
                  <p972:resolutionValue>Resolved</p972:resolutionValue>
               </p972:statusValidValue>
               <p972:statusValidValue>
                  <p972:value>CANCLD</p972:value>
                  <p972:description>Cancelled correspondence.</p972:description>
                  <p972:resolutionValue>Resolved</p972:resolutionValue>
               </p972:statusValidValue>
               <p972:statusValidValue>
                  <p972:value>Close Order</p972:value>
                  <p972:description>Close Out Order</p972:description>
                  <p972:resolutionValue>InProcess</p972:resolutionValue>
               </p972:statusValidValue>
               <p972:statusValidValue>
                  <p972:value>CMPLTD</p972:value>
                  <p972:description>Completed.</p972:description>
                  <p972:resolutionValue>Resolved</p972:resolutionValue>
               </p972:statusValidValue>
               <p972:statusValidValue>
                  <p972:value>Configure Solution</p972:value>
                  <p972:description>Configure VPN Hardware</p972:description>
                  <p972:resolutionValue>InProcess</p972:resolutionValue>
               </p972:statusValidValue>
               <p972:statusValidValue>
                  <p972:value>D01012</p972:value>
                  <p972:description>Status update code for 01/01/2000</p972:description>
                  <p972:resolutionValue>InProcess</p972:resolutionValue>
               </p972:statusValidValue>
               <p972:statusValidValue>
                  <p972:value>D01032</p972:value>
                  <p972:description>Status update code for 01/03/2000</p972:description>
                  <p972:resolutionValue>InProcess</p972:resolutionValue>
               </p972:statusValidValue>
               <p972:statusValidValue>
                  <p972:value>D01042</p972:value>
                  <p972:description>Status update code for 01/04/2000</p972:description>
                  <p972:resolutionValue>InProcess</p972:resolutionValue>
               </p972:statusValidValue>
               <p972:statusValidValue>
                  <p972:value>D01052</p972:value>
                  <p972:description>Status update code for 01/05/2000</p972:description>
                  <p972:resolutionValue>InProcess</p972:resolutionValue>
               </p972:statusValidValue>
               <p972:statusValidValue>
                  <p972:value>D01312</p972:value>
                  <p972:description>Status update code for 01/31/2000</p972:description>
                  <p972:resolutionValue>InProcess</p972:resolutionValue>
               </p972:statusValidValue>
               <p972:statusValidValue>
                  <p972:value>D02292</p972:value>
                  <p972:description>Status update code for 02/29/2000</p972:description>
                  <p972:resolutionValue>InProcess</p972:resolutionValue>
               </p972:statusValidValue>
               <p972:statusValidValue>
                  <p972:value>D09099</p972:value>
                  <p972:description>Status update code for 09/09/1999</p972:description>
                  <p972:resolutionValue>InProcess</p972:resolutionValue>
               </p972:statusValidValue>
               <p972:statusValidValue>
                  <p972:value>D12312</p972:value>
                  <p972:description>Status update code for 12/31/2000</p972:description>
                  <p972:resolutionValue>InProcess</p972:resolutionValue>
               </p972:statusValidValue>
               <p972:statusValidValue>
                  <p972:value>D12319</p972:value>
                  <p972:description>Status update code for 12/31/1999</p972:description>
                  <p972:resolutionValue>InProcess</p972:resolutionValue>
               </p972:statusValidValue>
               <p972:statusValidValue>
                  <p972:value>DAPRF</p972:value>
                  <p972:description>Failed Approval</p972:description>
                  <p972:resolutionValue>InProcess</p972:resolutionValue>
               </p972:statusValidValue>
               <p972:statusValidValue>
                  <p972:value>DAPRP</p972:value>
                  <p972:description>Passed Approval</p972:description>
                  <p972:resolutionValue>InProcess</p972:resolutionValue>
               </p972:statusValidValue>
               <p972:statusValidValue>
                  <p972:value>Data Failure</p972:value>
                  <p972:description>Data Failure</p972:description>
                  <p972:resolutionValue>InProcess</p972:resolutionValue>
               </p972:statusValidValue>
               <p972:statusValidValue>
                  <p972:value>DBIMG</p972:value>
                  <p972:description>Bad Image</p972:description>
                  <p972:resolutionValue>Resolved</p972:resolutionValue>
               </p972:statusValidValue>
               <p972:statusValidValue>
                  <p972:value>DDATE</p972:value>
                  <p972:description>Status for Date TSF</p972:description>
                  <p972:resolutionValue>InProcess</p972:resolutionValue>
               </p972:statusValidValue>
               <p972:statusValidValue>
                  <p972:value>DDECML</p972:value>
                  <p972:description>Status for Decimal TSF</p972:description>
                  <p972:resolutionValue>InProcess</p972:resolutionValue>
               </p972:statusValidValue>
               <p972:statusValidValue>
                  <p972:value>Determine Billing</p972:value>
                  <p972:description>Determine How to Bill Customer</p972:description>
                  <p972:resolutionValue>InProcess</p972:resolutionValue>
               </p972:statusValidValue>
               <p972:statusValidValue>
                  <p972:value>DFRONT</p972:value>
                  <p972:description>Front Office</p972:description>
                  <p972:resolutionValue>InProcess</p972:resolutionValue>
               </p972:statusValidValue>
               <p972:statusValidValue>
                  <p972:value>DIGO</p972:value>
                  <p972:description>In Good Order</p972:description>
                  <p972:resolutionValue>InProcess</p972:resolutionValue>
               </p972:statusValidValue>
               <p972:statusValidValue>
                  <p972:value>DMONEY</p972:value>
                  <p972:description>Status for Money TSF</p972:description>
                  <p972:resolutionValue>InProcess</p972:resolutionValue>
               </p972:statusValidValue>
               <p972:statusValidValue>
                  <p972:value>DNIGO</p972:value>
                  <p972:description>Not In Good Order</p972:description>
                  <p972:resolutionValue>Pending</p972:resolutionValue>
               </p972:statusValidValue>
               <p972:statusValidValue>
                  <p972:value>DOPEN</p972:value>
                  <p972:description>Open Work Item</p972:description>
                  <p972:resolutionValue>InProcess</p972:resolutionValue>
               </p972:statusValidValue>
               <p972:statusValidValue>
                  <p972:value>DPHONE</p972:value>
                  <p972:description>Status for Phone TSF</p972:description>
                  <p972:resolutionValue>InProcess</p972:resolutionValue>
               </p972:statusValidValue>
               <p972:statusValidValue>
                  <p972:value>DQC</p972:value>
                  <p972:description>Quality Review</p972:description>
                  <p972:resolutionValue>InProcess</p972:resolutionValue>
               </p972:statusValidValue>
               <p972:statusValidValue>
                  <p972:value>DQCF</p972:value>
                  <p972:description>Failed QC</p972:description>
                  <p972:resolutionValue>InProcess</p972:resolutionValue>
               </p972:statusValidValue>
               <p972:statusValidValue>
                  <p972:value>DQCP</p972:value>
                  <p972:description>Passed QC</p972:description>
                  <p972:resolutionValue>InProcess</p972:resolutionValue>
               </p972:statusValidValue>
               <p972:statusValidValue>
                  <p972:value>DREJCT</p972:value>
                  <p972:description>Rejected Image</p972:description>
                  <p972:resolutionValue>InProcess</p972:resolutionValue>
               </p972:statusValidValue>
               <p972:statusValidValue>
                  <p972:value>DRSVLD</p972:value>
                  <p972:description>Resolved Work Item</p972:description>
                  <p972:resolutionValue>Resolved</p972:resolutionValue>
               </p972:statusValidValue>
               <p972:statusValidValue>
                  <p972:value>DURGT</p972:value>
                  <p972:description>Urgent</p972:description>
                  <p972:resolutionValue>InProcess</p972:resolutionValue>
               </p972:statusValidValue>
               <p972:statusValidValue>
                  <p972:value>ENTRD</p972:value>
                  <p972:description>Entered.</p972:description>
                  <p972:resolutionValue>InProcess</p972:resolutionValue>
               </p972:statusValidValue>
               <p972:statusValidValue>
                  <p972:value>Error</p972:value>
                  <p972:description>Error</p972:description>
                  <p972:resolutionValue>InProcess</p972:resolutionValue>
               </p972:statusValidValue>
               <p972:statusValidValue>
                  <p972:value>Escalate Complete</p972:value>
                  <p972:description>Escalate Complete</p972:description>
                  <p972:resolutionValue>Resolved</p972:resolutionValue>
               </p972:statusValidValue>
               <p972:statusValidValue>
                  <p972:value>Escalate Continue</p972:value>
                  <p972:description>Escalate Continue</p972:description>
                  <p972:resolutionValue>InProcess</p972:resolutionValue>
               </p972:statusValidValue>
               <p972:statusValidValue>
                  <p972:value>Failure</p972:value>
                  <p972:description>Failure</p972:description>
                  <p972:resolutionValue>InProcess</p972:resolutionValue>
               </p972:statusValidValue>
               <p972:statusValidValue>
                  <p972:value>Fill Order</p972:value>
                  <p972:description>Fill the Order for VPN</p972:description>
                  <p972:resolutionValue>InProcess</p972:resolutionValue>
               </p972:statusValidValue>
               <p972:statusValidValue>
                  <p972:value>Finished</p972:value>
                  <p972:description>Finished</p972:description>
                  <p972:resolutionValue>Resolved</p972:resolutionValue>
               </p972:statusValidValue>
               <p972:statusValidValue>
                  <p972:value>Gather Tech Details</p972:value>
                  <p972:description>Gather VPN Tech Details</p972:description>
                  <p972:resolutionValue>InProcess</p972:resolutionValue>
               </p972:statusValidValue>
               <p972:statusValidValue>
                  <p972:value>IGO</p972:value>
                  <p972:description>IGO</p972:description>
                  <p972:resolutionValue>InProcess</p972:resolutionValue>
               </p972:statusValidValue>
               <p972:statusValidValue>
                  <p972:value>Indexed</p972:value>
                  <p972:description>Indexed</p972:description>
                  <p972:resolutionValue>InProcess</p972:resolutionValue>
               </p972:statusValidValue>
               <p972:statusValidValue>
                  <p972:value>INTMSG</p972:value>
                  <p972:description>Int. message.</p972:description>
                  <p972:resolutionValue>InProcess</p972:resolutionValue>
               </p972:statusValidValue>
               <p972:statusValidValue>
                  <p972:value>NEWCO</p972:value>
                  <p972:description>New correspondence.</p972:description>
                  <p972:resolutionValue>InProcess</p972:resolutionValue>
               </p972:statusValidValue>
               <p972:statusValidValue>
                  <p972:value>Osg IGO</p972:value>
                  <p972:description>Osg IGO</p972:description>
                  <p972:resolutionValue>InProcess</p972:resolutionValue>
               </p972:statusValidValue>
               <p972:statusValidValue>
                  <p972:value>Osg NIGO</p972:value>
                  <p972:description>Osg NIGO</p972:description>
                  <p972:resolutionValue>InProcess</p972:resolutionValue>
               </p972:statusValidValue>
               <p972:statusValidValue>
                  <p972:value>OSG Send</p972:value>
                  <p972:description>OSG Send</p972:description>
                  <p972:resolutionValue>InProcess</p972:resolutionValue>
               </p972:statusValidValue>
               <p972:statusValidValue>
                  <p972:value>Partial</p972:value>
                  <p972:description>Partial</p972:description>
                  <p972:resolutionValue>InProcess</p972:resolutionValue>
               </p972:statusValidValue>
               <p972:statusValidValue>
                  <p972:value>Pre-Configure System</p972:value>
                  <p972:description>Pre-Configure VPN System</p972:description>
                  <p972:resolutionValue>InProcess</p972:resolutionValue>
               </p972:statusValidValue>
               <p972:statusValidValue>
                  <p972:value>PRFAIL</p972:value>
                  <p972:description>Printing failed.</p972:description>
                  <p972:resolutionValue>InProcess</p972:resolutionValue>
               </p972:statusValidValue>
               <p972:statusValidValue>
                  <p972:value>PRNTD</p972:value>
                  <p972:description>Printed.</p972:description>
                  <p972:resolutionValue>Resolved</p972:resolutionValue>
               </p972:statusValidValue>
               <p972:statusValidValue>
                  <p972:value>Provision System</p972:value>
                  <p972:description>Provision VPN Solution</p972:description>
                  <p972:resolutionValue>InProcess</p972:resolutionValue>
               </p972:statusValidValue>
               <p972:statusValidValue>
                  <p972:value>PRREQD</p972:value>
                  <p972:description>Printing requested</p972:description>
                  <p972:resolutionValue>InProcess</p972:resolutionValue>
               </p972:statusValidValue>
               <p972:statusValidValue>
                  <p972:value>PRTRDY</p972:value>
                  <p972:description>Print ready.</p972:description>
                  <p972:resolutionValue>InProcess</p972:resolutionValue>
               </p972:statusValidValue>
               <p972:statusValidValue>
                  <p972:value>Resolved</p972:value>
                  <p972:description>Resolved Work Item</p972:description>
                  <p972:resolutionValue>Resolved</p972:resolutionValue>
               </p972:statusValidValue>
               <p972:statusValidValue>
                  <p972:value>Review Tech Requirements</p972:value>
                  <p972:description>Review Technical Requirements</p972:description>
                  <p972:resolutionValue>InProcess</p972:resolutionValue>
               </p972:statusValidValue>
               <p972:statusValidValue>
                  <p972:value>RJCT</p972:value>
                  <p972:description>Rejected.</p972:description>
                  <p972:resolutionValue>InProcess</p972:resolutionValue>
               </p972:statusValidValue>
               <p972:statusValidValue>
                  <p972:value>RQPR</p972:value>
                  <p972:description>Requires printing.</p972:description>
                  <p972:resolutionValue>InProcess</p972:resolutionValue>
               </p972:statusValidValue>
               <p972:statusValidValue>
                  <p972:value>RSLVD</p972:value>
                  <p972:description>Resolved.</p972:description>
                  <p972:resolutionValue>Resolved</p972:resolutionValue>
               </p972:statusValidValue>
               <p972:statusValidValue>
                  <p972:value>Send Email</p972:value>
                  <p972:description>Send Email</p972:description>
                  <p972:resolutionValue>InProcess</p972:resolutionValue>
               </p972:statusValidValue>
               <p972:statusValidValue>
                  <p972:value>Send Fax</p972:value>
                  <p972:description>Send Fax</p972:description>
                  <p972:resolutionValue>InProcess</p972:resolutionValue>
               </p972:statusValidValue>
               <p972:statusValidValue>
                  <p972:value>Success</p972:value>
                  <p972:description>Success</p972:description>
                  <p972:resolutionValue>InProcess</p972:resolutionValue>
               </p972:statusValidValue>
               <p972:statusValidValue>
                  <p972:value>Successful Transmission</p972:value>
                  <p972:description>Successful Transmission</p972:description>
                  <p972:resolutionValue>InProcess</p972:resolutionValue>
               </p972:statusValidValue>
               <p972:statusValidValue>
                  <p972:value>Suspended</p972:value>
                  <p972:description>Suspended</p972:description>
                  <p972:resolutionValue>InProcess</p972:resolutionValue>
               </p972:statusValidValue>
               <p972:statusValidValue>
                  <p972:value>Task Complete</p972:value>
                  <p972:description>Task Complete</p972:description>
                  <p972:resolutionValue>Resolved</p972:resolutionValue>
               </p972:statusValidValue>
               <p972:statusValidValue>
                  <p972:value>Task Ready</p972:value>
                  <p972:description>Task Ready</p972:description>
                  <p972:resolutionValue>InProcess</p972:resolutionValue>
               </p972:statusValidValue>
               <p972:statusValidValue>
                  <p972:value>Test Solution</p972:value>
                  <p972:description>Test VPN Solution</p972:description>
                  <p972:resolutionValue>InProcess</p972:resolutionValue>
               </p972:statusValidValue>
               <p972:statusValidValue>
                  <p972:value>Transmit to IWS</p972:value>
                  <p972:description>Transmit to IWS</p972:description>
                  <p972:resolutionValue>InProcess</p972:resolutionValue>
               </p972:statusValidValue>
               <p972:statusValidValue>
                  <p972:value>Unindexed</p972:value>
                  <p972:description>Unindexed</p972:description>
                  <p972:resolutionValue>InProcess</p972:resolutionValue>
               </p972:statusValidValue>
               <p972:statusValidValue>
                  <p972:value>Update System Config</p972:value>
                  <p972:description>Update VPN Configuration</p972:description>
                  <p972:resolutionValue>InProcess</p972:resolutionValue>
               </p972:statusValidValue>
               <p972:statusValidValue>
                  <p972:value>UPDCO</p972:value>
                  <p972:description>Updated correspondence.</p972:description>
                  <p972:resolutionValue>InProcess</p972:resolutionValue>
               </p972:statusValidValue>
               <p972:statusValidValue>
                  <p972:value>Validate Charges</p972:value>
                  <p972:description>Validate Customer Charges</p972:description>
                  <p972:resolutionValue>InProcess</p972:resolutionValue>
               </p972:statusValidValue>
            </p972:statusValidValues>
            <p972:markedForDelete>0</p972:markedForDelete>
            <p972:workItemCreateAllowed>1</p972:workItemCreateAllowed>
            <p972:runtimeDynamicFieldRules/>
         </p347:retrieveRuntimeWorkItemRuleReturn>
      </p347:retrieveRuntimeWorkItemRuleResponse>
   </soapenv:Body>
</soapenv:Envelope>
`
