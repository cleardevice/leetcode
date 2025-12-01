func findNthDigit(n int) int {
    k, cnt := 1, 9
    for n > k*cnt {
        n -= k * cnt
        k++
        cnt *= 10
    }

    start := 1
    for i := 1; i < k; i++ {
        start *= 10
    }

    num := start + (n-1)/k
    idx := (n - 1) % k

    for i := 0; i < k-idx-1; i++ {
        num /= 10
    }
    return num % 10
}
