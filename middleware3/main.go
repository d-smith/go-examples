package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"time"
)

var soapStart = `
<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:ns="http://xmlns.fmr.com/systems/dev/xtrac/2004/06/" xmlns:ser="http://xmlns.fmr.com/common/headers/2005/12/ServiceProcessingDirectives" xmlns:ser1="http://xmlns.fmr.com/common/headers/2005/12/ServiceCallContext" xmlns:typ="http://xmlns.fmr.com/systems/dev/xtrac/2004/06/types">
   <soapenv:Body>
   	<ns:retrieve>
   		<ns:foo>`

var soapEnd = `</ns:foo>
	</ns:retreive>
    </soapenv:Body>
</soapenv:Envelope>`

func extractResource(uri string) (string, error) {
	parts := strings.Split(uri, "/")
	if len(parts) != 3 || parts[2] == "" {
		return "", fmt.Errorf("Expected URI format: /foo/<resource id>")
	}

	return parts[2], nil

}

func handleCall(w http.ResponseWriter, r *http.Request) {
	println("handleCall called...")
	println("\tmethod is ", r.Method)
	println("\turi is ", r.RequestURI)

	body, err := ioutil.ReadAll(r.Body)
	r.Body.Close()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Problem reading body"))
	}

	println("\tbody:\b\n", string(body))

	for key, value := range r.Header {
		fmt.Printf("key: %s value: %v\n", key, value)
	}

	w.Write([]byte("handleCall wrote this stuff\n"))
}

func timingWrapper(timingName string, fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		fn(w, r)
		elapsed := time.Since(start)
		fmt.Printf("time for %s: %d\n", timingName, elapsed.Nanoseconds())
	}
}

func restToSoapWrapper(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		resourceId, err := extractResource(r.RequestURI)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(err.Error()))
			return
		}

		println("resource id: ", resourceId)

		payload := fmt.Sprintf("%s%s%s", soapStart, resourceId, soapEnd)
		r.Body = ioutil.NopCloser(strings.NewReader(payload))

		rec := httptest.NewRecorder()
		println("wrap...")
		r.Method = "POST"
		r.RequestURI = "/soap/foo/service"
		r.Header.Add("SOAPAction", `"Some deal"`)

		fn(rec, r)

		w.Write(rec.Body.Bytes())
		w.Write([]byte("wrap wrote this\n"))
	}
}

func main() {
	println("yeah")
	wrapped := timingWrapper("soap wrapper", restToSoapWrapper(timingWrapper("handle call", handleCall)))
	http.Handle("/foo/", wrapped)
	http.ListenAndServe(":8080", nil)
}
