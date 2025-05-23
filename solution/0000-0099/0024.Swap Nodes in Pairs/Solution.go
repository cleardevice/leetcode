func swapPairs(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    t := swapPairs(head.Next.Next)
    p := head.Next
    p.Next = head
    head.Next = t
    return p
}
