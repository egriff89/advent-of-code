use std::fs;

fn main() {
    let contents = fs::read_to_string("input.txt").expect("Couldn't load file");
    let mut floor = 0;

    // Part 1
    which_floor(&contents, &mut floor);

    // Part 2
    basement_position(contents, floor);
}

fn which_floor(contents: &str, floor: &mut i32) {
    let mut floor = *floor;

    for c in contents.chars() {
        match c {
            '(' => floor += 1,
            ')' => floor -= 1,
            _ => {}
        }
    }

    println!("---- PART 1 ----");
    println!("Santa's floor: {}", floor);
    println!("----------------\n");
}

fn basement_position(contents: String, mut floor: i32) {
    for (i, c) in contents.char_indices() {
        match c {
            '(' => floor += 1,
            ')' => {
                floor -= 1;

                if floor == -1 {
                    println!("---- PART 2 ----");
                    println!("Basement Position: {}", i + 1);
                    println!("----------------\n");
                    break;
                }
            }
            _ => {}
        }
    }
}
