/**
 * @param {string[]} emails
 * @return {number}
 */
var numUniqueEmails = function(emails) {
    'use strict';
    var set = new Set();
    for (let email of emails) {
        var sEmail = email.split('@');
        var c = sEmail[0].split('+');
        var arr = [];
        [...c[0]].forEach(a=>{if(a!=='.')arr.push(a);});
        set.add((arr.join(''))+'@'+sEmail[1]);
    }
    // console.log(set);
    return set.size;
};

numUniqueEmails(['fafa@uci.edu','f.a.fa+uni@uci.edu','fafa@uc+i.edu','fafa.a@u.ci.edu']);


