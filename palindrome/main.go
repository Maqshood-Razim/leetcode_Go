func isPalindrome(x int) bool {
	if x < 0 {
		   return false
	   }
   
	   s := strconv.Itoa(x)
	   for i := 0; i < len(s)/2; i++ {
		   if s[i] != s[len(s)-1-i] {
			   return false
		   }
	   }
	   return true
   }    
   