/**
 * Definition for _Node.
 * class _Node {
 *     val: number
 *     next: _Node | null
 *     random: _Node | null
 * 
 *     constructor(val?: number, next?: _Node, random?: _Node) {
 *         this.val = (val===undefined ? 0 : val)
 *         this.next = (next===undefined ? null : next)
 *         this.random = (random===undefined ? null : random)
 *     }
 * }
 */

function copyRandomList(head: _Node | null): _Node | null {
    if(head == null) {
        return head;
    }

    let current : _Node | null = head;
    let values = new Map<_Node, _Node>();

    while(current != null) {
        let node = new _Node(current.val, null, null);
        values.set(current, node);
        current = current.next;
    }

    for (const [key, val] of values.entries()) {
        val.next = values.get(key.next) || null;
        val.random = values.get(key.random) || null;
    }
    
    return values.get(head);
};