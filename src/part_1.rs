use std::fs;
pub fn solve() {
    let contents = fs::read_to_string("./data/part-1.txt")
        .expect("Should have been able to read the file");
    let chars = contents.chars();
    let mut potions = 0;
    for char in chars {
        match char {
            'B' => potions +=1,
            'C' => potions +=3,
            _ => potions += 0,
        }
    }

    println!("potions: {potions}");
}
