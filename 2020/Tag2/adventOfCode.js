const fs = require('fs');
let input = fs.readFileSync('input.txt', 'utf-8').trim();

let pairs = input.split("\r\n");
let json = {};
let keys = [];
pairs.forEach(pair => {
    let items = pair.split(": ");
    let key = items[1];
    keys.push(key);
    let value = items[0].split("-");
    json[key] = [];
    json[key][0] = parseInt(value[0]);
    json[key][1] = parseInt(value[1].split(" ")[0]);
    json[key][2] = value[1].split(" ")[1]; 
});

console.log(json);
// console.log(keys);
/*validPassworts = 0;
for(let i = 0;i < /*keys.length* 5; i++) {
    let first = keys[i];
    let second = keys[i];
    while(second.indexOf(json[keys[i]][2])) {
        console.log(json[keys[i]][2]);
        second = second.replace(json[keys[i]][2], '');
        console.log(second);
    }
    let difrence = first.length - second.length;
    console.log(difrence);
}*/