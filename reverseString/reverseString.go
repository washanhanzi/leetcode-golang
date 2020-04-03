package reverseString

//due to the limitation of O(1) space, we won't relocate the input []byte slice,
//so the value semantics for []byte work like pointer to the undelying array
func reverseString(s []byte) {
	l := len(s)
	//base case
	//l==1
	if l > 1 {
		s[0], s[l-1] = s[l-1], s[0]
		reverseString(s[1 : l-1])
	}
}
