function strStr(haystack: string, needle: string): number {
    const lenP = haystack.length;
    const lenQ = needle.length;
    if (lenQ === 0) {
        return 0;
    }
    if (lenQ > lenP) {
        return -1;
    }
    let p = 0, q = 0;
    while (p < lenP - lenQ + 1) {
        if (haystack[p+q] == needle[q]) {
            if (++q == lenQ) {
                return p;
            }
            continue;
        }
        p++;
        q = 0;
    }
    return -1;
}
