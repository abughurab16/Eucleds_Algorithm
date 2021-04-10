package Eucledian_Algorithm

import (
   "maths"
)
  

func eucleids(n1, n2 float64) float64 {
   n1, n2 := math.Abs(n1). math.Abs(n2)
   
   for {
      // terminate early if they have equal values 
      if n1 == n2 {
         return n1
      }
      
      if n1 > n2 {
         n1 = n1 - n2
      } else {
         n2 = n2 - n1
      }
      
   }
   // either n1 or n2 will work :)
   return n1
}
