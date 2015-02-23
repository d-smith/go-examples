package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"strings"
)

func printIndent(indent int) {
	for i := 0; i < indent; i++ {
		fmt.Print("  ")
	}
}

func printItem(indent int, key string, val interface{}) {
	switch t := val.(type) {
	default:
		printIndent(indent)
		fmt.Printf("Got a %T: %s -> %v\n", t, key, val)
	case []interface{}:
		printIndent(indent)
		fmt.Printf("%s is []interface\n", key)
		for _, v := range val.([]interface{}) {
			printItem(indent+1, key, v)
		}
	case map[string]interface{}:
		printIndent(indent)
		fmt.Printf("%s is map[string]interface{}\n", key)
		for k, v := range val.(map[string]interface{}) {
			printItem(indent+1, k, v)
		}

	}
}

func main() {
	jsonDoc := `
{  
   "tasks":[  
      {  
         "workItemNo":"W00001-25JUL14",
         "jeopardy":[  
            {  
               "field":"QCTD",
               "value":"2014-07-23T10:05:34.010Z",
               "status":"RED"
            }
         ],
         "fields":[  
            {  
               "field":"queue",
               "value":"HIREQ"
            },
            {  
               "field":"status",
               "value":"APPROVE"
            },
            {  
               "field":"memo",
               "value":"Ready to go"
            },
            {  
               "field":"Priority",
               "value":"high"
            }
         ]
      },
      {  
         "workItemNo":"W02002-12JUL14",
         "jeopardy":[  
            {  
               "field":"QCTD",
               "value":"2014-07-23T10:05:34.010Z",
               "status":"RED"
            }
         ],
         "fields":[  
            {  
               "field":"queue",
               "value":"PROMOTEQ"
            },
            {  
               "field":"status",
               "value":"REJECT"
            },
            {  
               "field":"memo",
               "value":"Surely you jest"
            },
            {  
               "field":"Priority",
               "value":"high"
            }
         ]
      }
   ],
   "totalCount":456,
   "moreData":true
}    
	`

	dec := json.NewDecoder(strings.NewReader(jsonDoc))
	data := make(map[string]interface{})

	for {
		if err := dec.Decode(&data); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
	}

	for k, v := range data {
		printItem(0, k, v)
	}

}
