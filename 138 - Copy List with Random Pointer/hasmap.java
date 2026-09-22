class Solution {
    public Node copyRandomList(Node head) {
        
        if(head == null){ 
            return head;
        }

        Node current = head;
        Map<Node, Node> nodes = new HashMap<Node, Node>();

        while(current != null) {
            Node node = new Node(current.val, null, null);
            nodes.put(current, node);
            current = current.next;
        }

        nodes.forEach((key, value) -> {
            value.next = nodes.get(key.next);
            value.random = nodes.get(key.random);
        });

        return nodes.get(head);
    }
}