const fs = require('fs')
const path = require('path')

let inputPath = path.resolve(__dirname, "input.txt");
let input = fs.readFileSync(inputPath).toString().split("\n");

let p1 = puzzle1()
let p2 = puzzle2()
console.log(p1, p2)

function puzzle1() {
    let scores = {
        ')': 3,
        ']': 57,
        '}': 1197,
        '>': 25137
    }
    let points = 0
    for (let line of input) {
        let chars = line.split('')
        let opens = []
        for (let char of chars) {
            if (char === '(' || char === '[' || char === '{' || char === '<') {
                opens.push(char)
            } else {
                let lastOpen = opens[opens.length - 1]
                if (lastOpen !== invertChar(char)) {
                    points += scores[char]
                    break
                } else {
                    opens.pop()
                }
            }
        }
    }

    return points
}

function puzzle2() {
    let scores = {
        ')': 1,
        ']': 2,
        '}': 3,
        '>': 4
    }
    let all = []
    for (let line of input) {
        let chars = line.split('')
        let opens = []
        let complete = true
        for (let char of chars) {
            if (char === '(' || char === '[' || char === '{' || char === '<') {
                opens.push(char)
            } else {
                let lastOpen = opens[opens.length - 1]
                if (lastOpen !== invertChar(char)) {
                    complete = false
                    break
                } else {
                    opens.pop()
                }
            }
        }

        if (!complete) {
            continue
        }

        let points = 0
        let closes = opens.reverse().map(o => invertChar(o))
        for (let close of closes) {
            points *= 5
            points += scores[close]
        }

        all.push(points)
    }

    let num = (all.length - 1) / 2
    let sorted = all.sort((a, b) => b - a)
    let mid = sorted[num]
    return mid
}

function invertChar(char) {
    if (char === ')') {
        return '('
    } else if (char === ']') {
        return '['
    } else if (char === '}') {
        return '{'
    } else if (char === '>') {
        return '<'
    }  else if (char === '(') {
        return ')'
    } else if (char === '{') {
        return '}'
    } else if (char === '<') {
        return '>'
    } else if (char === '[') {
        return ']'
    } else {
        throw char
    }
}