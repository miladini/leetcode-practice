'use strict';
/**
 * @param {number[]} nums
 * @return {void} Do not return anything, modify nums in-place instead.
 */ 
var moveZeroes = function(nums) {
    let q = [];
    for (let i = 0; i < nums.length; ) {
        if (nums[i]==0) {
            q.push(i);
            i++;
        } else if (q.length>0) {
            let temp = nums[i];
            let z_i = q.shift();
            nums[i] = nums[z_i];
            nums[z_i] = temp;
        }
        else {
            i++;
        }
    }
};


var nums = [2,1,0,0,0,0,0,22,2,-1,0];
moveZeroes(nums);
console.log(nums);