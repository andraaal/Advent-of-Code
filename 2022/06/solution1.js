const fs = require('fs');

const data = fs.readFileSync('./input.txt', 'utf8').split('');

let lastFour = new Array();
let i = 0;

for(i = 0; i < data.length; i++) {
    const char = data[i];
    lastFour.push(char);
    if(lastFour.length > 4) lastFour.shift();

    if(new Set(lastFour).size === 4) break;
}

console.log(i + 1); 