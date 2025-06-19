function getIntersectionNode(headA: ListNode | null, headB: ListNode | null): ListNode | null {
    let [a, b] = [headA, headB];
    while (a !== b) {
        a = a ? a.next : headB;
        b = b ? b.next : headA;
    }
    return a;
}
