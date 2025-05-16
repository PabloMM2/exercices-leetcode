var hasCycle = function(head) {
    let nodes = new Set();
    let current = head

    while (current != null) {
        if (nodes.has(current)){
            return true
        }

        nodes.add(current)
        current = current.next
    }
    return false
};