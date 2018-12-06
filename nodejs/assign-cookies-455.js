'use strict';
/**
 * @param {number[]} g
 * @param {number[]} s
 * @return {number}
 */
var findContentChildren = function(g, s) {
    g.sort((a,b)=>a-b);
    s.sort((a,b)=>a-b);
    let i = 0;
    let j = 0;
    let count = 0;
    while (i < s.length && s[i] < g[j]) i++;
    while (i < s.length && j < g.length) {
        count++;
        j++;
        i++;
        while (j < g.length && i < s.length && s[i] < g[j]) i++;
    }
    return count;
};


let g = [10,9,8,7];
let s = [5,6,7,8];

console.log(findContentChildren(g,s));

console.log(g)
console.log(s)