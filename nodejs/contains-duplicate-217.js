'use strict';
/**
 * @param {number[]} nums
 * @return {boolean}
 */
var containsDuplicate = function(nums) {
    var set = new Set();
    for (let num of nums) set.add(num);    
    return nums.length == set.size;
};

console.log(containsDuplicate([]));

