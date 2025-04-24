function removeDuplicates(nums: number[]): number {
    let k: number = 0;
    for (const v of nums) {
        if (k === 0 || v !== nums[k-1]) {
            nums[k++] = v;
        }
    }
    return k;
};
