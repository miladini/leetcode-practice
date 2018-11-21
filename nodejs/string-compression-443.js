/**
 * @param {character[]} chars
 * @return {number}
 */
'use strict';
var compress = function(chars) {
    let writing = 0;
    let reading = 0;

    for (; reading < chars.length;) {
        let nextReading = reading+1;
        while(nextReading < chars.length && chars[nextReading]===chars[reading]) {
            nextReading++;
        }
        if (nextReading-reading===1) {
            chars[writing] = chars[reading];
            writing++;
        }
        else {
            chars[writing] = chars[reading];
            writing++;
            
            let toWrite = String(nextReading-reading);
            // console.log(toWrite);
            for (let nextWriting=writing;nextWriting-writing < toWrite.length; nextWriting++) {
                chars[nextWriting] = toWrite.charAt(nextWriting-writing);
                // console.log('hello!');
            }
            // chars[writing] = String(nextReading-reading);
            writing = writing + toWrite.length;
        }
        reading = nextReading;
    }
    return writing;
};

// let chars  = ["m","m","m","s","x","x"];
let chars  = ["m","m","m","m","m","m","m","m","m","m","m","m","m","m","m"];
console.log(chars);
let b  = compress(chars);
console.log(b,chars);

