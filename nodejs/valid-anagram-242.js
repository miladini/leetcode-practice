'use strict';

/**
 * @param {string} s
 * @param {string} t
 * @return {boolean}
 */
var isAnagram = function(s, t) {
    if (s.length!==t.length) return false;
    var map1 = new Map();
    var map2 = new Map();
    for (let v of s) {
        map1.set(v,(map1.get(v)||0)+1);
    }

    for (let v of t) {
        map2.set(v,(map2.get(v)||0)+1);
    }

    for (let k of map1) {
        if (k[1]!==map2.get(k[0])) return false;
    }
    return true;
};

console.log(isAnagram("afaf","fafa"));