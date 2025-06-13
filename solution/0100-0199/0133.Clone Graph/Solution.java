class Solution {
    private Map<Node, Node> g = new HashMap<>();

    public Node cloneGraph(Node node) {
        return dfs(node);
    }

    private Node dfs(Node node) {
        if (node == null) {
            return null;
        }
        Node cloned = g.get(node);
        if (cloned == null) {
            cloned = new Node(node.val);
            g.put(node, cloned);
            for (Node nxt: node.neighbors) {
                cloned.neighbors.add(dfs(nxt));
            }
        }
        return cloned;
    }
}
