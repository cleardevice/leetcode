function swapPairs(head: ListNode | null): ListNode | null {
    if (head == null || head.next == null) {
        return head;
    }
    const t = swapPairs(head.next.next);
    const p = head.next;
    p.next = head;
    head.next = t;
    return p;
};
