package main

import (
	"fmt"
	"sort"
)

type Service struct {
	ServiceName string
	Host        string
	Port        int
	BaseUri     string
}

func (s Service) String() string {
	return fmt.Sprintf("{%s - @%s:%d - %s}",
		s.ServiceName, s.Host, s.Port, s.BaseUri)
}

type ByService []Service

func (s ByService) Len() int           { return len(s) }
func (s ByService) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s ByService) Less(i, j int) bool { return s[i].ServiceName < s[j].ServiceName }

func main() {
	services := []Service{
		{"z service", "z.com", 8080, "/z/service"},
		{"a service", "a.com", 8080, "/a/service"},
		{"crap service", "foo.bar.com", 3000, "/crap/part/of/the/site"},
	}

	fmt.Println(services)
	sort.Sort(ByService(services))
	fmt.Println(services)
}
