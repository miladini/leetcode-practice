/**
 * @param {number[]} nums1
 * @param {number[]} nums2
 * @return {number[]}
 */
'use strict';
var nextGreaterElement = function(nums1, nums2) {
    let map = new Map();
    let stack = [];
    for (let i = nums2.length-1; i >= 0; i--) {
        // console.log(i);
        while(stack.length!==0 && stack[stack.length-1] < nums2[i]) {
            stack.pop();
        }
        if (stack.length===0) {
            map.set(nums2[i],-1);
        }
        else {
            map.set(nums2[i],stack[stack.length-1]);
        }
        stack.push(nums2[i]);
    }
    let res = [];
    for (let num of nums1) {
        console.log(num);
        res.push(map.get(num));
    }
    return res;
};


console.log(nextGreaterElement([4,1,2],[3,4,7,1,9,11,1,2]));