function sortedListToBST(head: ListNode | null): TreeNode | null {
    const nums: number[] = [];
    for (; head; head = head.next) {
        nums.push(head.val);
    }
    const dfs = (i: number, j: number): TreeNode | null => {
        if (i > j) {
            return null;
        }
        const mid = (i + j) >> 1;
        const left = dfs(i, mid - 1);
        const right = dfs(mid + 1, j);
        return new TreeNode(nums[mid], left, right);
    };
    return dfs(0, nums.length - 1);
};
