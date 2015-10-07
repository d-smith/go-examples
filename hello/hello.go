package main

import (
	"fmt"
	"github.com/d-smith/go-examples/stringutil"
)

type KVStore interface {
	Put(string, string)
	Get(string) string
}

type A struct {
	Store KVStore
}

func NewA(store KVStore) *A {
	return &A{
		Store: store,
	}
}

func (a *A) doFoo(k string, v string) {
	a.Store.Put(k, v)
}

type KV1 struct{}

func (*KV1) Put(key string, vsl string) {
	fmt.Println("kv1 put")
}

func (*KV1) Get(key string) string {
	return "kv1 get"
}

type KV2 struct{}

func (*KV2) Put(key string, vsl string) {
	fmt.Println("kv2 put")
}

func (*KV2) Get(key string) string {
	return "kv2 get"
}

func main() {
	greeting := "Hello, world"
	doubleReveresed := stringutil.Reverse(stringutil.Reverse(greeting))
	fmt.Printf("%s\n", doubleReveresed)

	kvs := new(KV1)
	a := NewA(kvs)
	a.doFoo("a", "b")
}
