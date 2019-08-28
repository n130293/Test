package main

import "fmt"

func main(){
   b:=255
   var a *int =&b
   fmt.Printf("Type of %T",a)
   fmt.Println(a)
   }