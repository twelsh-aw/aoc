const fs = require('fs')
const path = require('path')

let inputPath = path.resolve(__dirname, "input.txt");
let input = fs.readFileSync(inputPath).toString().split('\n')

let p1 = puzzle1()
let p2 = puzzle2()
console.log(p1, p2)

function puzzle1() {
    let game = parseInput(input)
    let rolls = 0
    let turn = 0
    let curRoll = 1
    while (true) {
        let player = game[(turn % 2)]
        let rolled = 0
        for (let i = 0; i < 3; i++) {
            rolled += curRoll
            rolls++
            curRoll = (curRoll % 100) + 1
        }

        let position = (player.cur + rolled) % 10
        if (position === 0) {
            position = 10
        }

        player.score += position
        player.cur = position
        turn++
        if (player.score >= 1000) {
            break
        }
    }

    let minScore = Math.min(...game.map(p => p.score))
    return minScore * rolls
}

/// cache is map of:
/// curPos array + turn + curScore array => how many wins for each player array
/// use cache instead of traversing every universe
function puzzle2() {
    let game = parseInput(input)
    let wins = [0, 0]
    let cache = {}
    playGame(game, wins, cache, 0)
    let maxWins = Math.max(...wins)
    return maxWins
}

function playGame(game, wins, cache, turn) {
    if (turn > 21) {
        throw game
    }

    let key = `${game.map(g => g.cur).toString()}_${turn}_${game.map(g => g.score).toString()}`
    if (cache[key]) {
        wins[0] += cache[key][0]
        wins[1] += cache[key][1]
        return
    }

    let universes = []
    for (let i = 0; i < 3; i++) {
        for (let j = 0; j < 3; j++) {
            for (let k = 0; k < 3; k++) {
                universes.push(i+j+k+3)
            }
        }
    }

    let localWins = [0, 0]
    for (let rolled of universes) {
        let universeGame = [...game].map(p => { return {...p} })
        let player = universeGame[(turn % 2)]
        let position = (player.cur + rolled) % 10
        if (position === 0) {
            position = 10
        }

        player.score += position
        player.cur = position
        if (player.score >= 21) {
            localWins[(turn % 2)]++
        } else {
            playGame(universeGame, localWins, cache, turn + 1)
        }
    }

    cache[key] = [...localWins]
    wins[0] += localWins[0]
    wins[1] += localWins[1]
}

function parseInput(input) {
    let game = []
    game.push({cur: parseInt(input[0].split(': ')[1]), score: 0})
    game.push({cur: parseInt(input[1].split(': ')[1]), score: 0})
    return [...game]
}
