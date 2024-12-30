use std::fs;

pub fn solve() {
    let example = "AxBCDDCAxD";
    algo(&example);
    let contents = fs::read_to_string("./data/part-2.txt")
        .expect("Should have been able to read the file");
    algo(&contents);
}

fn algo(text: &str) {
    let mut chars = text.chars();
    let mut potions = 0;
    let mut done = false;
    for i in (0..chars.clone().count()).step_by(2) {
        let mut has_x = false;
        let mut cur_potions = 0;
        for j in 0..2 {
            let cur_char = chars.nth(0);
            match cur_char {
                Some('B') => cur_potions +=1,
                Some('C') => cur_potions +=3,
                Some('D') => cur_potions +=5,
                Some('x') => has_x = true,
                Some('\n') => done = true,
                _ => cur_potions += 0,
            }
            if done {
                break;
            }
            println!("{i}) {cur_char:?}: {potions},{cur_potions},{i},{j}");
        }
        if done {
            break;
        }
        if !has_x {
            cur_potions += 2
        }
        potions += cur_potions
    }

    println!("part_2: {potions}");
}
