function sortColors(nums: number[]): void {
    let [i, j, k] = [-1, nums.length, 0];
    while (k < j) {
        if (nums[k] === 0) {
            i++;
            [nums[i], nums[k]] = [nums[k], nums[i]];
            k++;
        } else if (nums[k] === 2) {
            j--;
            [nums[j], nums[k]] = [nums[k], nums[j]];
        } else {
            k++;
        }
    }
};
