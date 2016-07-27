package main

import (
	"net/http"
	"time"
	. "github.com/gorilla/feeds"
)

func getFeed() *Feed {
	now := time.Now()

	feed := &Feed{
		Title:       "jmoiron.net blog",
		Link:        &Link{Href: "http://jmoiron.net/blog"},
		Description: "discussion about tech, footie, photos",
		Author:      &Author{Name: "Jason Moiron", Email: "jmoiron@jmoiron.net"},
		Created:     now,
		Copyright:   "This work is copyright © Benjamin Button",
	}

	feed.Items = []*Item{
		&Item{
			Title:       "Limiting Concurrency in Go",
			Link:        &Link{Href: "http://jmoiron.net/blog/limiting-concurrency-in-go/"},
			Description: "A discussion on controlled parallelism in golang",
			Author:      &Author{Name: "Jason Moiron", Email: "jmoiron@jmoiron.net"},
			Created:     now,
		},
		&Item{
			Title:       "Logic-less Template Redux",
			Link:        &Link{Href: "http://jmoiron.net/blog/logicless-template-redux/"},
			Description: "More thoughts on logicless templates",
			Created:     now,
		},
		&Item{
			Title:       "Idiomatic Code Reuse in Go",
			Link:        &Link{Href: "http://jmoiron.net/blog/idiomatic-code-reuse-in-go/"},
			Description: "How to use interfaces <em>effectively</em>",
			Created:     now,
		},
	}

	return feed
}

func feedHandler(rw http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		feed := getFeed()
		atom, _ := feed.ToAtom()
		rw.Write([]byte(atom))
		break
	default:
		http.Error(rw, "I only support GET and PUT\n", http.StatusMethodNotAllowed)
	}
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/feed", feedHandler)

	http.ListenAndServe(":4000", mux)
}
