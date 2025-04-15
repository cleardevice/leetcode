function reverse(x: number): number {
    let ans = 0;
    for (; x != 0; x = Math.trunc(x / 10)) {
        if (ans < -(2 ** 31) / 10 || ans > (2 ** 31 - 1) / 10) {
            return 0;
        }
        ans = ans * 10 + x % 10;
    }
    return ans;
};
