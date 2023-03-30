package main 

import (
"fmt"


)

func method1(n int){

   fmt.Println("Ahhh the good old days of computing when all there was, was step instructions ;) \n " + 
               "Try to think back to your first computer program ever ....if it was in the 80s (BBC Micro, Spectrum, Commodore, Amstrad etc) this'll make sense ! ")
   
   fmt.Println(1 * n)
   fmt.Println(2 * n)
   fmt.Println(3 * n)
   fmt.Println(4 * n)
   fmt.Println(5 * n)
   fmt.Println(6 * n)
   fmt.Println(7 * n)
   fmt.Println(8 * n)
   fmt.Println(9 * n)
   fmt.Println(10 * n)

}


func method2(n int){

   fmt.Println("\nWhere would we be without procedural code ?")
   
   for i := 1; i <= 10; i++ {
      fmt.Println(n * i)
      
   }
   
   fmt.Println("....this could also constitute an object orientated approach.")

}

func method3(n int, i int){

  
  if i == 1 {
   fmt.Println("\nAnd now folks for some recursive programming, this style popular in functional programming circles")
   }

  if i > 10 {
   return 
  }

  fmt.Println(i * n)

  method3(n, i + 1)
      

}


func main(){
	
   method1(2)
   method2(2)
   method3(2, 1)


}
