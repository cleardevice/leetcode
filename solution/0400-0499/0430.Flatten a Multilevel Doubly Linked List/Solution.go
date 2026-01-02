/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Prev *Node
 *     Next *Node
 *     Child *Node
 * }
 */

func flatten(root *Node) *Node {
    if root == nil {
        return nil
    }

    dummy := &Node{}
    dummy.Next = root

    preorder(dummy, root)

    dummy.Next.Prev = nil
    return dummy.Next
}

func preorder(pre, cur *Node) *Node {
    if cur == nil {
        return pre
    }

    cur.Prev = pre
    pre.Next = cur

    next := cur.Next

    tail := preorder(cur, cur.Child)
    cur.Child = nil
    
    return preorder(tail, next)
}
