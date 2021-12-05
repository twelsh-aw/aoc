const fs = require('fs')
const path = require('path')

let inputPath = path.resolve(__dirname, "input.txt");
let input = fs.readFileSync(inputPath).toString().split("\n");
let draws = input[0].split(',')
let boards = parseBoards(input)

let p1 = puzzle1(boards, draws)
let p2 = puzzle2(boards, draws)
console.log(p1, p2)

function puzzle1(boards, draws) {
    for (let num of draws) {
        for (let board of boards) {
            for (let i = 0; i < 5; i++) {
                for (let j = 0; j < 5; j++) {
                    if (board[i][j] == num) {
                        board[i][j] = -1
                    }

                    if (checkBoard(board)) {
                        return scoreBoard(board, num)
                    }
                }
            }
        }
    }
}

function puzzle2(boards, draws) {
    let winningBoards = []
    for (let num of draws) {
        for (let board of boards) {
            for (let i = 0; i < 5; i++) {
                for (let j = 0; j < 5; j++) {
                    if (board[i][j] == num) {
                        board[i][j] = -1
                    }

                    if (checkBoard(board) && !winningBoards.some(b => b == board)) {
                        winningBoards.push(board)
                        if (winningBoards.length === boards.length) {
                            return scoreBoard(board, num)
                        }
                    }
                }
            }
        }
    }
}

function scoreBoard(board, num) {
    let sum = 0
    for (let i = 0; i < 5; i++) {
        for (let j = 0; j < 5; j++) {
            if (board[i][j] !== -1) {
                sum += parseInt(board[i][j])
            }
        }
    }

    return sum * num
}

function checkBoard(board) {
    for (let i = 0; i < 5; i++) {
        if (board[i].every(v => v === -1)) {
            return true
        }
    }

    let b2 = transpose(board)
    for (let i = 0; i < 5; i++) {
        if (b2[i].every(v => v === -1)) {
            return true
        }
    }

    // let diag1 = [
    //     board[0][0],
    //     board[1][1],
    //     board[2][2],
    //     board[3][3],
    //     board[4][4]
    // ]
    // if (diag1.every(v => v === -1)) {
    //     return true
    // }
    //
    // let diag2 = [
    //     board[4][0],
    //     board[3][1],
    //     board[2][2],
    //     board[1][3],
    //     board[0][4]
    // ]
    // if (diag2.every(v => v === -1)) {
    //     return true
    // }

    return false
}

function transpose(board) {
    return board[0].map((_, colIndex) => board.map(row => row[colIndex]));
}

function parseBoards(input) {
    let boardsRaw = input.slice(2, input.len).filter(b => b !== '')
    let nums = []
    for (let i = 0; i < boardsRaw.length; i++) {
        let s = boardsRaw[i].split('\n')
        for (let j = 0; j < 5; j++) {
            let num = s.join(' ').trim().split(' ').filter((j => j != ''))[j]
            nums.push(num)
        }
    }

    let rows = []
    let row = []
    for (let i = 0; i < nums.length; i++) {
        row.push(nums[i])
        if ((i + 1) % 5 === 0) {
            rows.push([...row])
            row = []
        }
    }

    let boards = []
    let board = []
    for (let j = 0; j < rows.length; j++) {
        board.push(rows[j])
        if ((j+1) % 5 === 0) {
            boards.push([...board])
            board = []
        }
    }

    return boards
}