package aoc2015

import "core:fmt"
import "core:os"

// Part 1
count_floors :: proc(data: []byte) -> int {
    current_floor: int
    for c in string(data) {
        if c == '(' {
            current_floor += 1
        } else {
            current_floor -= 1
        }
    }

    return current_floor
}

// Part 2
basement_visit :: proc(data: []byte) -> int {
    position: int
    current_floor: int

    for c in string(data) {
        if current_floor == -1 {
            return position
        }

        if c == '(' {
            current_floor += 1
        } else {
            current_floor -= 1
        }
        position += 1
    }

    return current_floor
}

main :: proc() {
    data, _ := os.read_entire_file("input.txt")
    fmt.println(count_floors(data))
    fmt.println(basement_visit(data))
}