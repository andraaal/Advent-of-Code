const fs = require('fs');

const data = fs.readFileSync('./input.txt', 'utf8').split('\r\n');

const field = [
    'GBDCPR',
    'GVH',
    'MPJDQSN',
    'MNCDGLSP',
    'SLFPCNBJ',
    'STGVZDBQ',
    'QTFHMZB',
    'FBDMC',
    'GQCF'
];

let state = field.map(value => value.split('').reverse());

let result = '';
let exit = 0;

data.forEach(value => {
    //if (exit) return;
    let commandContents = value.split(' ');
    const origin = parseInt(commandContents[3]);
    const times = parseInt(commandContents[1]);
    const target = parseInt(commandContents[5]);
    const buffer = state[origin - 1].splice(state[origin - 1].length - times, times);
    state[target - 1] = state[target - 1].concat(buffer)
    if(exit < 10) console.log(state.map(value => value.join()), buffer);

    exit++;
});
state.forEach(value => result += value[value.length - 1]);

console.log(result);