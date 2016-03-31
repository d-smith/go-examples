## JWT-HS256

Generate a token signed using HS256 in node and decode in go... why not?

<pre>
go run $GOPATH/src/github.com/d-smith/go-examples/id-and-secret/genappcred.go
client id: 6b2ba18b-4490-4a69-56f7-6cb05dc8bdce, secret: 4JcfdPalEjdFej41Yia7tp0CYoozintlvpSF4c2fIyg=

node make-token.js -s 4JcfdPalEjdFej41Yia7tp0CYoozintlvpSF4c2fIyg= -c aud 6b2ba18b-4490-4a69-56f7-6cb05dc8bdce foo bar baz yeah
eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJhdWQiOiI2YjJiYTE4Yi00NDkwLTRhNjktNTZmNy02Y2IwNWRjOGJkY2UiLCJmb28iOiJiYXIiLCJiYXoiOiJ5ZWFoIiwiaWF0IjoxNDU5NDU5MDk3fQ.LAzTHxKJLpbNTk_M4L1nakDBIt9xZyvL1HZwDZaYfys

go run decoder.go 4JcfdPalEjdFej41Yia7tp0CYoozintlvpSF4c2fIyg= eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJhdWQiOiI2YjJiYTE4Yi00NDkwLTRhNjktNTZmNy02Y2IwNWRjOGJkY2UiLCJmb28iOiJiYXIiLCJiYXoiOiJ5ZWFoIiwiaWF0IjoxNDU5NDU5MDk3fQ.LAzTHxKJLpbNTk_M4L1nakDBIt9xZyvL1HZwDZaYfys
Token claims...
aud -> 6b2ba18b-4490-4a69-56f7-6cb05dc8bdce
foo -> bar
baz -> yeah
iat -> 1.459459097e+09
</pre>