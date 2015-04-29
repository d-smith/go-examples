package main

/*
#include <stdlib.h>
#include <stdio.h>
#include <string.h>

void do_memset(char **dest) {
  const char src[50] = "http://www.tutorialspoint.com";
  *dest = malloc(50 * sizeof(char));
 //char dest[50];

 printf("Before memcpy dest = %s\n", *dest);
 memcpy(*dest, src, strlen(src)+1);
 printf("After memcpy dest = %s\n", *dest);
}

*/
import "C"
import "unsafe"

func main() {
	println("alive")
	var data *C.char
	C.do_memset(&data)
	println(*data)
	C.free(unsafe.Pointer(data))
}
