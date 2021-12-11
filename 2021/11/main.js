const fs = require('fs')
const path = require('path')

let inputPath = path.resolve(__dirname, "input.txt");
let input = fs.readFileSync(inputPath).toString().split("\n").map(row => row.split('').map(Number));

let p1 = puzzle1()
let p2 = puzzle2()
console.log(p1, p2)

function puzzle1() {
    let numFlashes = 0
    for (let i = 0; i < 100; i++) {
        for (let row = 0; row < 10; row++) {
            for (let col = 0; col < 10; col++) {
                Step({row, col})
            }
        }

        for (let row = 0; row < 10; row++) {
            for (let col = 0; col < 10; col++) {
                if (input[row][col] >= 10) {
                    numFlashes++
                    input[row][col] = 0
                }
            }
        }
    }

    return numFlashes
}

function puzzle2() {
    input = fs.readFileSync(inputPath).toString().split("\n").map(row => row.split('').map(Number));
    let i = 0
    while (true) {
        i++
        for (let row = 0; row < 10; row++) {
            for (let col = 0; col < 10; col++) {
                Step({row, col})
            }
        }

        let numFlashes = 0
        for (let row = 0; row < 10; row++) {
            for (let col = 0; col < 10; col++) {
                if (input[row][col] >= 10) {
                    numFlashes++
                    input[row][col] = 0
                }
            }
        }

        if (numFlashes === 100) {
            return i
        }
    }
}

function Step(point) {
    let val = ++input[point.row][point.col]
    if (val === 10) {
        // console.log(point)
        let adj = getAdjacentPoints(point.row, point.col)
        for (let a of adj) {
            Step(a)
        }
    }
}

function getAdjacentPoints(row, col) {
    let possible = [
        {row: row - 1, col},
        {row: row + 1, col},
        {row, col: col -1},
        {row, col: col + 1},
        {row: row - 1, col: col - 1},
        {row: row - 1, col: col + 1},
        {row: row + 1, col: col - 1},
        {row: row + 1, col: col + 1}
    ]

    return possible.filter(p => input[p.row] && input[p.row][p.col] !== undefined)
}
