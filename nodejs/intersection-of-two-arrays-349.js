/**
 * @param {number[]} nums1
 * @param {number[]} nums2
 * @return {number[]}
 */
var intersection = function(nums1, nums2) {
    set1 = new Set(nums1);
    set2 = new Set(nums2);    
    return [...set1].filter(x=>set2.has(x));
};