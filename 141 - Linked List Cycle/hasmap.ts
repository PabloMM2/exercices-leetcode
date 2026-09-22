function hasCycle(head: ListNode | null): boolean {
    let nodes : Map<ListNode, boolean> = new Map();
    while(head != null) {
        if(nodes.get(head)) {
            return true;
        }
        nodes.set(head, true);
        head = head.next;
    }
    return false;
};