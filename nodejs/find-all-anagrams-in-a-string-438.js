/**
 * @param {string} s
 * @param {string} p
 * @return {number[]}
 */

'use strict';
var findAnagrams = function(s, p) {
    let res = [];
    let map = new Map();
    for(let i = 0; i < p.length; i++) {
        if (map.has(p.charAt(i))) {
            map.set(p.charAt(i),1 + map.get(p.charAt(i)));
        }
        else {
            map.set(p.charAt(i), 1);
        }
    }
    

    let dmap = new Map();

    for (let i = 0; i < p.length; i++) {
        if (dmap.has(s.charAt(i))) {
            dmap.set(s.charAt(i), 1 + dmap.get(s.charAt(i)));
        }
        else {
            dmap.set(s.charAt(i), 1);
        }
        
    }

    console.log(map);
    console.log(dmap);
    console.log('#####################');


    if (equal(map,dmap)) res.push(0);
    for (let i = 0; i+p.length < s.length; i++) {
        dmap.set(s.charAt(i), -1 + dmap.get(s.charAt(i)));
        if (dmap.get(s.charAt(i))===0) dmap.delete(s.charAt(i));

        if (dmap.has(s.charAt(i+p.length))) {
            dmap.set(s.charAt(i+p.length),  1 + dmap.get(s.charAt(i+p.length)));
        }
        else {
            dmap.set(s.charAt(i+p.length), 1);
        }
        console.log(map);
        console.log(dmap);
        console.log('----------------');

        if (equal(map,dmap) ) res.push(i+1);
    }

    console.log(res);
    return res;
};

var equal = function(map,dmap) {
    // console.log(map.size, dmap.size);
    if (map.size!==dmap.size) return false;
    for (let [k,v] of map) {
        console.log(k , v);
        if (!dmap.has(k)) return false;
        if (dmap.get(k)!==v) return false;
    }
    return true;
}
// findAnagrams("asdasdad","asd");
findAnagrams("fff","ff");