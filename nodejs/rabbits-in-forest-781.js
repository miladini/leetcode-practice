'use strict';
/**
 * @param {number[]} answers
 * @return {number}
 */
var numRabbits = function(answers) {
    let map = new Map();

    for (let v of answers) {
        if (map.has(v)) map.set(v,map.get(v)+1);
        else map.set(v,1);
    }
    let count = 0;
    for (var [k,v] of map) {
        if (k+1>=v) {
            count += k+1
        } else {
            count += (k+1)*Math.ceil(v/(k+1));
        }
    }
    return count;
};

console.log(numRabbits([3, 3, 3, 3,3]));