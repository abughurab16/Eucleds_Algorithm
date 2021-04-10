package Eucledian_Algorithm

  
// defining our own abs function since go's default works with float types

func abs(a int) {
   if a < 0 {
      a = -a
   }
    
   return a
}

func eucleidSub(n1, n2 int) int {
   n1 := abs(n1)
   n2 := abs(n2)
   
  // loop until both values are equal
  for n1 != n2 {
    
    if n1 > n2 {
      n1 = n1 - n2
    } else {
      n2 = n2 - n1
    }
    
  }
   // either n1 or n2 will work :)
   return n1
}

func eucleidMod(n1, n2 int) int {
  n1 := abs(n1)
  n2 := abs(n2)
  
  for n2 != 0 {
    a,b = b, a%b
  }
  
  return a
}
