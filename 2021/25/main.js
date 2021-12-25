const fs = require('fs')
const path = require('path')

let inputPath = path.resolve(__dirname, "input.txt");
let input = fs.readFileSync(inputPath).toString().split('\n').map(row => row.split(''))

let p1 = puzzle1()
let p2 = puzzle2()
console.log(p1, p2)

function puzzle1() {
    let board = JSON.parse(JSON.stringify(input))
    let numMoves = Infinity
    let turn = 0
    while(numMoves > 0) {
        let res = update(board)
        numMoves = res.numMoves
        board = res.newBoard
        turn++
    }

    return turn
}

function puzzle2() {
    // nothing to do
}

function update(board) {
    let numMoves = 0
    let tempBoard = JSON.parse(JSON.stringify(board))
    // left moves
    for (let row = 0; row < board.length; row++) {
        for (let col = 0; col < board[row].length; col++) {
            let val = board[row][col]
            let right = getRightCol(board, row, col)
            if (val === ">" && board[row][right] === ".") {
                tempBoard[row][right] = ">"
                tempBoard[row][col] = "."
                numMoves++
            }
        }
    }

    // down moves
    let newBoard = JSON.parse(JSON.stringify(tempBoard))
    for (let row = 0; row < tempBoard.length; row++) {
        for (let col = 0; col < tempBoard[row].length; col++) {
            let val = tempBoard[row][col]
            let down = getDownRow(tempBoard, row, col)
            if (val === "v" && tempBoard[down][col] === ".") {
                newBoard[down][col] = "v"
                newBoard[row][col] = "."
                numMoves++
            }
        }
    }

    return {newBoard, numMoves}
}

function getRightCol(board, row, col) {
    if (board[row][col + 1]) {
        return col + 1
    } else {
        return 0
    }
}

function getDownRow(board, row, col) {
    if (board[row + 1]) {
        return row + 1
    } else {
        return 0
    }
}
