'use strict';
/**
 * @param {number[]} nums
 * @param {number} k
 * @return {number[]}
 */
var topKFrequent = function (nums, k) {
    let freqs = new Map();
    let buckets = [];
    for (let i = 0; i <= nums.length; i++) {
        buckets[i] = new Set();
    }

    for (let v of nums) {
        buckets[0].add(v);
        freqs[v] = 0;
    }

    for (let v of nums) {
        buckets[freqs[v]].delete(v);
        buckets[freqs[v] + 1].add(v);
        freqs[v]++;

    }

    console.log(buckets);
    let res = [];

    let f = buckets.length - 1;
    while (res.length < k && f >= 0) {
        if (buckets[f].size > 0) {
            for (let v of buckets[f]) {
                if (res.length < k) {
                    res.push(v);
                }
            }
        }
        f--;
    }
    return res;
};

console.log(topKFrequent([1,2,1,1,1,1], 3));