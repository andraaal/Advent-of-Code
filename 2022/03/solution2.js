const fs = require('fs');

const data = [];
fs.readFileSync('./input.txt', 'utf8').split('\r\n').forEach(value => data.push(new Set(value.split(''))));

let sum = 0;

data.forEach((value, index) => {
    if((index + 1) % 3 === 0) {
        for(badge of value) {
            if(data[index - 1].has(badge) && data[index - 2].has(badge)) {
                if (badge === badge.toLowerCase()) {
                    sum += badge.charCodeAt(0) - 96;
                } else {
                    sum += badge.charCodeAt(0) - 38;
                }
            }
        }
    }
});

console.log(sum);