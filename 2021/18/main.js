const fs = require('fs')
const path = require('path')

let inputPath = path.resolve(__dirname, "input.txt");
let input = fs.readFileSync(inputPath).toString().split('\n')

let p1 = puzzle1()
let p2 = puzzle2()
console.log(p1, p2)

function puzzle1() {
    let current = JSON.parse(input[0])
    for (let i = 1; i < input.length; i++) {
        let next = JSON.parse(input[i])
        let summed = add(current, next)
        reduce(summed)
        current = JSON.parse(JSON.stringify(summed))
    }

    // print(current)
    return getMagnitude(current)
}

function puzzle2() {
    let maxMag = 0
    for (let i = 0; i < input.length; i++) {
        for (let j = 0; j < input.length; j++) {
            if (i === j) {
                continue
            }

            let current = JSON.parse(input[i])
            let next = JSON.parse(input[j])
            let summed = add(current, next)
            reduce(summed)
            let m = getMagnitude(summed)
            if (m > maxMag) {
                maxMag = m
            }
        }
    }

    return maxMag
}

function reduce(cur) {
    let explosion, split
    while (true) {
        explosion = getExplosion(cur)
        if (explosion) {
            explode(cur, explosion)
            continue
        }

        split = getSplit(cur)
        if (split) {
            splitUp(cur, split)
            continue
        }

        break
    }
}

function add(n, m) {
    return [n, m]
}

function explode(cur, position) {
    if (position.length !== 4) {
        throw position
    }

    let left = getValue(cur, position.concat(0))
    let right = getValue(cur, position.concat(1))

    setValue(cur, position, 0)
    addValueToNext(cur, position, right)
    addValueToPrevious(cur, position, left)
}

function splitUp(cur, position) {
    let v = getValue(cur, position)
    if (!isNumber(v) || v < 10) {
        throw position
    }

    let left = Math.floor(v / 2)
    let right = Math.ceil(v / 2)
    let prev = getValue(cur, position.slice(0, position.length - 1))
    prev[position[position.length - 1]] = [left, right]
}

function getSplit(cur, position) {
    if (!position) {
        position = []
    }

    let v = getValue(cur, position)
    if (isNumber(v)) {
        if (v >= 10) {
            return position
        }

        return
    }

    let p0 = [...position, 0]
    let s0 = getSplit(cur, p0)
    if (s0 !== undefined) {
        return s0
    }

    let p1 = [...position, 1]
    let s1 = getSplit(cur, p1)
    if (s1 !== undefined) {
        return s1
    }
}

function getExplosion(cur, position) {
    if (!position) {
        position = [0, 0, 0, 0]
    }

    let v = getValue(cur, position)
    if (v !== undefined && !isNumber(v)) {
        return position
    }

    let nextPosition = getNextPosition(position)
    if (nextPosition === undefined) {
        return
    }

    return getExplosion(cur, nextPosition)
}

function print(t) {
    console.log(JSON.stringify(t, null, 2))
}

function isNumber(t) {
    return typeof(t) === 'number';
}

function addValueToNext(cur, position, valueAdd, considerCur) {
    let nextPosition = getNextPosition(position)
    if (considerCur) {
        nextPosition = position
    }
    if (!nextPosition || nextPosition.length === 0) {
        return
    }

    let value = getValue(cur, nextPosition)
    if (isNumber(value)) {
        setValue(cur, nextPosition, value + valueAdd)
    } else if (value === undefined) {
        addValueToNext(cur, position.slice(0, position.length - 1), valueAdd)
    } else {
        addValueToNext(cur, [...nextPosition, 0], valueAdd, true)
    }
}

function addValueToPrevious(cur, position, valueAdd, considerCur) {
    let nextPosition = getPreviousPosition(position)
    if (considerCur) {
        nextPosition = position
    }

    if (nextPosition === undefined) {
        return
    }

    let value = getValue(cur, nextPosition)
    if (isNumber(value)) {
        setValue(cur, nextPosition, value + valueAdd)
    } else if (value === undefined) {
        addValueToPrevious(cur, position.slice(0, position.length - 1), valueAdd)
    } else {
        addValueToPrevious(cur, [...nextPosition, 0], valueAdd, true)
    }
}

function getValue(cur, position) {
    for (let p of position) {
        if (cur[p] === undefined) {
            return undefined
        }
        return getValue(cur[p], position.slice(1))
    }

    return cur
}

function setValue(cur, position, value) {
    let before = position.slice(0, position.length - 1)
    let arr = getValue(cur, before)
    arr[position[position.length - 1]] = value
}

function getNextPosition(position) {
    let firstOne = position.indexOf( 1)
    let lastOne = position.lastIndexOf(1)
    let lastZero = position.lastIndexOf(0)
    if (lastZero < 0) {
        // none found
        return
    }

    let nextPosition = [...position]
    nextPosition[lastZero] = 1
    if (lastZero < firstOne) {
        for (let i = lastZero + 1; i < position.length; i++) {
            nextPosition[i] = 0
        }
    } else if (lastZero < lastOne) {
        for (let i = lastZero + 1; i < position.length; i++) {
            nextPosition[i] = 0
        }
    }

    return nextPosition
}

function getPreviousPosition(position) {
    let firstZero = position.indexOf( 0)
    let lastZero = position.lastIndexOf( 0)
    let lastOne = position.lastIndexOf(1)
    if (lastOne < 0) {
        // none found
        return
    }

    let nextPosition = [...position]
    nextPosition[lastOne] = 0
    if (lastOne < firstZero) {
        for (let i = lastOne + 1; i < position.length; i++) {
            nextPosition[i] = 1
        }
    } else if (lastOne < lastZero) {
        for (let i = lastOne + 1; i < position.length; i++) {
            nextPosition[i] = 1
        }
    }

    return nextPosition
}

function getMagnitude(cur) {
    if (isNumber(cur)) {
        return cur
    }

    return 3 * getMagnitude(cur[0]) + 2 * getMagnitude(cur[1])
}
