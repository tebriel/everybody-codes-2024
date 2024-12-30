use std::fs;

pub fn solve() {
    let example = "xBxAAABCDxCC";
    algo(&example);
    let contents = fs::read_to_string("./data/part-3.txt")
        .expect("Should have been able to read the file");
    algo(&contents);
}

fn algo(text: &str) {
    let mut chars = text.chars();
    let mut potions = 0;
    let mut done = false;
    let stepper = 3;
    for i in (0..chars.clone().count()).step_by(stepper) {
        let mut num_x = 0;
        let mut cur_potions = 0;
        for j in 0..stepper {
            let cur_char = chars.nth(0);
            match cur_char {
                Some('B') => cur_potions +=1,
                Some('C') => cur_potions +=3,
                Some('D') => cur_potions +=5,
                Some('x') => num_x += 1,
                Some('\n') => done = true,
                _ => cur_potions += 0,
            }
            // println!("{i}) {cur_char:?}: {potions},{cur_potions},{i},{j}");
            if done {
                break;
            }
        }
        if done {
            break;
        }
        match num_x {
            0 => cur_potions += 6,
            1 => cur_potions += 2,
            _ => cur_potions += 0,
        }
        potions += cur_potions;
    }

    println!("part_3: {potions}");
}
