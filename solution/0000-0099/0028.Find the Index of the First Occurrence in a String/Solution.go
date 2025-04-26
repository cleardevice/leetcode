func strStr(haystack string, needle string) int {
    lenP, lenQ := len(haystack), len(needle)
    if lenQ == 0 {
        return 0
    }
    if lenQ > lenP {
        return -1
    }
    p, q := 0, 0
    for p < lenP-lenQ+1 {
        if (haystack[p+q] == needle[q]) {
            if q == lenQ-1 {
                return p
            }
            q++
            continue
        }
        p++
        q = 0
    }
    return -1
}
