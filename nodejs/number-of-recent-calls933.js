
var RecentCounter = function() {
    arr = [];
    arr.push(0);
};

/** 
 * @param {number} t
 * @return {number}
 */
RecentCounter.prototype.ping = function(t) {
    // console.log(arr);
    arr.push(t);
    if (t <= 3000) return arr.length-1;
    else {
        // console.log(t);
        let hi = arr.length;
        let lo = 0;
        while (lo!==hi) {
            
            let mid = Math.floor((lo+hi+1)/2);
            // console.log(mid + ' ' + lo + ' ' + hi);
            
            // if (arr[mid] <= t-3000) lo = mid;
            // else hi = mid - 1;
            if (arr[mid]>t-3000) hi = mid - 1;
            else lo = mid;
        }
        // console.log(arr);
        // console.log(`for ${t} lo is ${lo} ${t-3000 === arr[lo]}   ${arr.length} ${arr.length-lo+1}`);
        return (t-3000 === arr[lo]) ? arr.length-lo: arr.length - lo - 1;
    }

};




/** 
 * Your RecentCounter object will be instantiated and called as such:
 * var obj = Object.create(RecentCounter).createNew()
 * var param_1 = obj.ping(t)
 */


 var obj = new RecentCounter();
//  obj.ping(1);
//  obj.ping(2);

 console.log('res is: ' + obj.ping(1));
 console.log('res is: ' + obj.ping(2));
 console.log('res is: ' + obj.ping(100));
 console.log('res is: ' + obj.ping(3000));
 console.log('res is: ' + obj.ping(3001));
 console.log('res is: ' + obj.ping(3002));
 console.log('res is: ' + obj.ping(3003));
 console.log('res is: ' + obj.ping(4003));
 console.log('res is: ' + obj.ping(9002));
 