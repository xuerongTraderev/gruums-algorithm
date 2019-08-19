package leetcode

/**
 *  assume s is an alphabet string
 */
func lengthOfLongestSubstring(s string) int {
	// consider length of stirng is equal to 0 and 1
    if len(s) == 0 {
        return 0
    } else if len(s) == 1{
        return 1
    }
	begin, end := 0, 0
	max := 0
	for end < len(s) {
		// continue if alphabet not in the current sub-string
		for i := begin ; i < end ; i++ {
			if s[i] == s[end] {
				// record the length and compare with max
				if max < end - begin {
					max = end - begin 
				}
				//update begin point
				begin = i + 1
				break
			}
		}
		end++
	}
    if max < end - begin{
		max = end - begin
	}
	return max
}