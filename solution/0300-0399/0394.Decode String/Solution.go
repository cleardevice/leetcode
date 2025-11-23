func decodeString(s string) string {
	numStack := []int{}
	strStack := []string{}

	num := 0
	res := ""

	for _, c := range s {
		switch {
		case c >= '0' && c <= '9':
			digit, _ := strconv.Atoi(string(c))
			num = num*10 + digit

		case c == '[':
			numStack = append(numStack, num)
			strStack = append(strStack, res)
			num = 0
			res = ""

		case c == ']':
			n := numStack[len(numStack)-1]
			numStack = numStack[:len(numStack)-1]

			prevStr := strStack[len(strStack)-1]
			strStack = strStack[:len(strStack)-1]

			res = prevStr + strings.Repeat(res, n)

		default:
			res += string(c)
		}
	}

	return res
}
