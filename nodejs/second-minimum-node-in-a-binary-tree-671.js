'use strict';
/**
 * Definition for a binary tree node.
 * function TreeNode(val) {
 *     this.val = val;
 *     this.left = this.right = null;
 * }
 */
/**
 * @param {TreeNode} root
 * @return {number}
 */

let min = ~(1<<31);
let smin = -1;
var findSecondMinimumValue = function (root) {
    if (root===null||root===undefined) {
        return;
    }
    if (root.val < min) {
        smin = min;
        min = root.val;
    }
    else if (root.val < smin && root.val !== min) {
        smin = root.val;
    }
    findSecondMinimumValue(root.left);
    findSecondMinimumValue(root.right);

    return smin === ~(1<<31) ? -1 : smin;

};


function TreeNode(val) {
    this.val = val;
    this.left = this.right = null;
}

console.log("hello world!");

let root = new TreeNode(2);
root.left = new TreeNode(2);
root.right = new TreeNode(5);
// root.left.left = new TreeNode(3);
// root.left.right = new TreeNode(9);
root.right.left = new TreeNode(5);
root.right.right = new TreeNode(7);

console.log(findSecondMinimumValue(root));