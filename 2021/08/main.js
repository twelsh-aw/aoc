const fs = require('fs')
const path = require('path')

let inputPath = path.resolve(__dirname, "input.txt");
let input = fs.readFileSync(inputPath).toString().split("\n");
let board = {
    0: ['a', 'b', 'c', 'e', 'f', 'g'],
    1: ['c', 'f'],
    2: ['a', 'c', 'd', 'e', 'g'],
    3: ['a', 'c', 'd', 'f', 'g'],
    4: ['b', 'c', 'd', 'f'],
    5: ['a', 'b', 'd', 'f', 'g'],
    6: ['a', 'b', 'd', 'e', 'f', 'g'],
    7: ['a', 'c', 'f'],
    8: ['a', 'b', 'c', 'd', 'e', 'f', 'g'],
    9: ['a', 'b', 'c', 'd', 'f', 'g']
}
let parsed = parseInput(input);

let p1 = puzzle1()
let p2 = puzzle2()
console.log(p1, p2)

function puzzle1() {
    // let uniqueLengths = Object.keys(board).filter(k => Object.keys(board).filter(k2 => board[k2].length === board[k].length).length == 1)
    // 1, 4, 7, 8 means correct
    let count = 0
    for (let i of [1, 7, 4, 8]) {
        let length = board[i].length
        for (let row of parsed) {
            count += row.output.filter(r => r.length === length).length
        }
    }

    return count;
}

function puzzle2() {
    let sum = 0
    for (let row of parsed) {
        let nums = Array.from(Array(10).keys());
        let map = Object.fromEntries(nums.map(n => [n, '']))

        map[1] = sortString(row.input.find(r => r.length === 2))
        map[7] = sortString(row.input.find(r => r.length === 3))
        map[4] = sortString(row.input.find(r => r.length === 4))
        map[8] = sortString(row.input.find(r => r.length === 7))
        //
        map[9] = sortString(row.input.find(r => intersect(r, map[4]).length === 4 && r.length===6))
        map[0] = sortString(row.input.find(r => intersect(r, map[7]).length === 3 && intersect(r, map[4]).length === 3 && intersect(r, map[1]).length === 2 && r.length===6))
        map[3] = sortString(row.input.find(r => intersect(r, map[7]).length === 3 && r.length===5))
        map[2] = sortString(row.input.find(r => r.length===5 && intersect(r, map[7]).length === 2 && intersect(r, map[4]).length === 2))
        map[5] = sortString(row.input.find(r => r.length===5 && intersect(r, map[7]).length === 2 && intersect(r, map[4]).length === 3))
        map[6] = sortString(row.input.find(r => r.length===6 && intersect(r, map[1]).length === 1 && intersect(r, map[4]).length === 3 && intersect(r, map[7]).length === 2))

        let final = ''
        for (let out of row.output) {
            let sorted = sortString(out)
            let key = Object.keys(map).find(k => map[k] === sorted)
            if (!key) {
                console.log(sum, row, map, sorted)
                throw 'a'
            }
            final += key.toString()
        }
        if (final.length != 4) {
            throw 'bad'
        }

        sum += parseInt(final)
    }

    return sum;
}



function sortString(text) {
    return text.split('').sort().join('');
}

function allLetters() {
    return [...'a', 'b', 'c', 'd', 'e', 'f', 'g']
}

function intersect(arr1, arr2) {
    let i = []
    for (let a1 of Array.from(arr1)) {
        for (let a2 of arr2) {
            if (a2 === a1) {
                i.push(a2)
                break
            }
        }
    }

    return i
}

function parseInput() {
    let parsed = []
    for (let i of input) {
        let io = i.split(' | ')
        let out = io[1].trim().split(' ')
        let inp = io[0].trim().split(' ')
        parsed.push({input: inp, output: out})
    }

    return parsed
}
