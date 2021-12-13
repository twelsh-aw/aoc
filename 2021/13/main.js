const fs = require('fs')
const path = require('path')

let inputPath = path.resolve(__dirname, "input.txt");
let input = fs.readFileSync(inputPath).toString().split("\n")
let parsed = parseInput(input);

let p1 = puzzle1()
let p2 = puzzle2()
console.log(p1, p2)

function puzzle1() {
    let maxX = Math.max(...parsed.pos.map(p => p.x))
    let maxY = Math.max(...parsed.pos.map(p => p.y))
    let board = makeArray(maxX + 1, maxY + 1)
    for (let row = 0; row < maxY + 1; row++) {
        for (let col = 0; col < maxX + 1; col++) {
            if (parsed.pos.some(p => p.y === row && p.x === col)) {
                board[row][col] = 1
            }
        }
    }

    doFold(board, parsed.folds[0])
    return countBoard(board)
}

function puzzle2() {
    let parsed = parseInput(input);
    let maxX = Math.max(...parsed.pos.map(p => p.x))
    let maxY = Math.max(...parsed.pos.map(p => p.y))
    let board = makeArray(maxX + 1, maxY + 1)
    for (let row = 0; row < maxY + 1; row++) {
        for (let col = 0; col < maxX + 1; col++) {
            if (parsed.pos.some(p => p.y === row && p.x === col)) {
                board[row][col] = 1
            }
        }
    }

    for (let fold of parsed.folds) {
        doFold(board, fold)
    }

    console.table(board)
}

function countBoard(board) {
    let count = 0
    for (let row of board) {
        for (let v of row) {
            if (v === 1) {
                count++
            }
        }
    }

    return count;
}

function doFold(board, fold) {
    if (fold.direction === 'y') {
        for (let row = fold.value + 1; row < board.length; row++) {
            let distance = row - fold.value
            let newRow = fold.value - distance
            for (let col = 0; col < board[0].length; col++) {
                if (board[row][col] === 1) {
                    board[newRow][col] = 1
                }
            }
        }

        board.splice(fold.value + 1, board.length - fold.value - 1)
    } else {
        for (let col = fold.value + 1; col < board[0].length; col++) {
            let distance = col - fold.value
            let newCol = fold.value - distance
            for (let row = 0; row < board.length; row++) {
                if (board[row][col] === 1) {
                    board[row][newCol] = 1
                }
            }
        }

        for (let row of board) {
            row.splice(fold.value + 1, row.length - fold.value - 1)
        }
    }
}

function parseInput() {
    let pos = []
    let folds = []
    for (let row of input) {
        let split = row.split(',')
        if (split.length === 2) {
            pos.push({x: parseInt(split[0]), y: parseInt(split[1])})
        } else if (row.includes("fold along")) {
            split = row.split('=')
            folds.push({direction: split[0].charAt(11), value: parseInt(split[1])})
        }
    }

    return {pos, folds}
}

function makeArray(d1, d2) {
    var arr = [];
    for(let i = 0; i < d2; i++) {
        arr.push(new Array(d1).fill(0));
    }
    return arr;
}