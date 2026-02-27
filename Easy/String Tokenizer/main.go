func insertSpace(str string , n int) string {

 var modified string

 for i, v := range str {
    if i % n == 0 && i != 0 {
        modified += " "
    }
    
        modified += string(v)
    
 }

     return modified
}