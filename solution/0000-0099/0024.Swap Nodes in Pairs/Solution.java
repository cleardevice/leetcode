class Solution {
    public ListNode swapPairs(ListNode head) {
        if (head == null || head.next == null) {
            return head;
        }
        ListNode t = swapPairs(head.next.next);
        ListNode curr = head.next;
        curr.next = head;
        head.next = t;
        return curr;
    }
}
