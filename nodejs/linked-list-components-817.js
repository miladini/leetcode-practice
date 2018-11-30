/**
 * Definition for singly-linked list.
 * function ListNode(val) {
 *     this.val = val;
 *     this.next = null;
 * }
 */
'use strict';
/**
 * @param {ListNode} head
 * @param {number[]} G
 * @return {number}
 */
var numComponents = function (head, G) {
    let set = new Set();
    for (let v of G) {
        set.add(v);
    }

    let connected = false;
    let count = 0;
    while (head !== null) {
        if (set.has(head.val)) {
            if (!connected) {
                count++;
                connected = true;
            }
        }
        else {
            connected = false;
        }
        head = head.next;
    }
    return count;
};

function ListNode(val) {
    this.val = val;
    this.next = null;
}


let node = new ListNode(44);
node.next = new ListNode(33);
node.next.next = new ListNode(22);
node.next.next.next = new ListNode(11);

numComponents(node,[]);