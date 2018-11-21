

/**
 * @param {string} str
 * @return {string}
 */
var toLowerCase = function(str) {
    "use strict";
    let c = [];
    [...str].forEach(x=>c.push(x.toLowerCase()));
    return c.join("");
};

console.log(toLowerCase("toLowerCase"));
console.log(toLowerCase("fFffFF"));
console.log(toLowerCase(""));
console.log(toLowerCase("AcM"));