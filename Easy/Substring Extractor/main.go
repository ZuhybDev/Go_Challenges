func extractSubstring(str string, startIdx, destiIdx int) string  {

 var result string

 for i, v := range str {

    if startIdx < 0 {
      result += ""
      break
    } 
    
    for i >= startIdx && i <= destiIdx {
        result += string(v)
        break
    }
 }

  return result
}