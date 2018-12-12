var isOneBitCharacter = function (bits) {
    for (let i = 0, L = bits.length; i < L; i++)
        if (i == bits.length - 1) return true;
        else if (bits[i]===1)i++;
    return false;
}

console.log(isOneBitCharacter([1,0,0,0,1,1,1,1,0,1,0,0,1,1,0,1,1,1,1,1,0,0]));