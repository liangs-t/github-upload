package main

/*
#include<stdio.h>
int sum(int x,int y)
{
	return x+y;
}
*/
import "C"

import "fmt"

func main() {
	k := int(C.sum(1, 2))
	fmt.Println(k)
}
