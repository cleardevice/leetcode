class Solution {
    public ListNode mergeTwoLists(ListNode list1, ListNode list2) {
        ListNode dummy = new ListNode(0);
        ListNode result = dummy;
        while (list1 != null && list2 != null) {
            if (list1.val > list2.val) {
                result.next = list2;
                list2 = list2.next;
            } else {
                result.next = list1;
                list1 = list1.next;  
            }
            result = result.next;
        }

        ListNode next = list1 != null ? list1 : list2;
        while (next != null) {
            result.next = next;
            result = result.next;
            next = next.next;
        }
        return dummy.next;
    }
}
