function climbStairs(n: number): number {
    let [a, b] = [1, 1];
    for (let i = 1; i < n; i++) {
        [a, b] = [b, a + b];
    }
    return b;
};
