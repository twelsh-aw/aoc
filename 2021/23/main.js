const fs = require('fs')
const path = require('path')

let inputPath = path.resolve(__dirname, "input.txt");
let input = fs.readFileSync(inputPath).toString().split('\n')
let initialLocations = parseInput(input)

const scoreMap = {
    "A": 1,
    "B": 10,
    "C": 100,
    "D": 1000
}

let p1 = puzzle1(initialLocations, 15000)
let p2 = puzzle2()
console.log(p1, p2)

function puzzle1(board, initialMinScore) {
    let puzzleBoard = {board: board, score: 0}
    let stack = [puzzleBoard]
    let minScore = initialMinScore
    let minScoreByBoard = {}
    while (stack.length) {
        let obj = stack.shift()
        if (obj.score >= minScore) {
            continue
        }

        if (minScoreByBoard[JSON.stringify(obj.board)] !== undefined
            && obj.score >= minScoreByBoard[JSON.stringify(obj.board)]) {
            continue
        } else {
            minScoreByBoard[JSON.stringify(obj.board)] = obj.score
        }

        if (isComplete(obj.board) && obj.score < minScore) {
            minScore = obj.score
            continue
        }

        let newObjs = getNextBoards(obj)
        for (let newObj of newObjs) {
            stack.push(newObj)
        }
    }

    return minScore
}

function puzzle2() {
    let puzzle2Board = getPuzzle2Board()
    return puzzle1(puzzle2Board, 60000)
}

function getNextBoards(obj) {
    let board = obj.board
    let nextObjs = []
    for (let col = 0; col < board[1].length; col++) {
        let val = board[1][col]
        if (isLetter(val)) {
            let room = getRoomColumnFor(val)
            let roomDest = roomReady(board, val, room)
            if (roomDest.isReady && canMoveInHallwayTo(board, col, room)) {
                let newObj = deepClone(obj)
                newObj.board[1][col] = "."
                newObj.board[roomDest.row][room] = val
                newObj.score += ((roomDest.row - 1) + Math.abs(room - col)) * scoreMap[val.charAt(0)]
                nextObjs.push(newObj)
            }
        }
    }

    for (let room of [3, 5, 7, 9]) {
        if (roomComplete(board, room)) {
            continue
        }

        for (let row = 2; row < board.length - 1; row++) {
            let val = board[row][room]
            if (isLetter(val)) {
                for (let col of [1, 2, 4, 6, 8, 10, 11]) {
                    if (canMoveInHallwayTo(board, room, col)) {
                        // console.log(row, room, val, 1, col)
                        let newObj = deepClone(obj)
                        newObj.board[row][room] = "."
                        newObj.board[1][col] = val
                        newObj.score += ((row - 1) + Math.abs(room - col)) * scoreMap[val.charAt(0)]
                        nextObjs.push(newObj)
                    }
                }
                break
            }
        }
    }


    return nextObjs
}

function canMoveInHallwayTo(board, col, dest) {
    if (col > dest) {
        for (let i = col - 1; i >= dest; i--) {
            if (board[1][i] !== ".") {
                return false
            }
        }

        return true
    } else if (col < dest) {
        for (let i = col + 1; i <= dest; i++) {
            if (board[1][i] !== ".") {
                return false
            }
        }

        return true
    } else {
        throw 'column already on dest'
    }
}

function roomReady(board, val, room) {
    for (let row = board.length - 2; row >= 1; row--) {
        if (board[row][room] === ".") {
            return {isReady: true, row}
        } else if (board[row][room].charAt(0) === val.charAt(0)) {
            continue;
        } else {
            return {isReady: false, row: -1}
        }
    }

    throw 'room already filled correctly'
}

function roomComplete(board, room) {
    let letter = getLetterForRoom(room)
    for (let row = 2; row < board.length - 1; row++) {
        if (board[row][room].charAt(0) !== letter) {
            return false
        }
    }

    return true
}

function isLetter(val) {
    return !!scoreMap[val.charAt(0)]
}

function isComplete(board) {
    return roomComplete(board, 3) && roomComplete(board, 5) && roomComplete(board,7) && roomComplete(board, 9)
}

function getRoomColumnFor(val) {
    switch (val.charAt(0)) {
        case "A":
            return 3
        case "B":
            return 5
        case "C":
            return 7
        case "D":
            return 9
        default:
            throw 'bad val'
    }
}

function getLetterForRoom(room) {
    switch (room) {
        case 3:
            return "A"
        case 5:
            return "B"
        case 7:
            return "C"
        case 9:
            return "D"
        default:
            throw 'bad room'
    }
}

function getPuzzle2Board() {
    let board = deepClone(initialLocations)
    board.splice(3, 0, ["#", "#", "#", "D", "#", "C", "#", "B", "#", "A", "#", "#", "#"])
    board.splice(4, 0, ["#", "#", "#", "D", "#", "B", "#", "A", "#", "C", "#", "#", "#"])
    return board
}

function deepClone(obj) {
    return JSON.parse(JSON.stringify(obj))
}

function parseInput(input) {
    let location =  input.map(row => row.split('').map(v => {
        if (v.trim().length === 0) {
            return '#'
        }

        return v
    }))

    for (let row of location) {
        for (let col = row.length; col < 13; col++) {
            row.push('#')
        }
    }

    return location
}
