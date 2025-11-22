// REF: https://exercism.io/my/solutions/e891a16a07a24005816423c8700a88d2

pub fn verse(n: u32) -> String {
    match n {
        0 => none(),
        1 => one(n),
        2 => two(n),
        _ => multiple(n),
    }
}

pub fn sing(start: u32, end: u32) -> String {
    let mut lyrics = String::from("");
    let mut number: u32 = start;

    while number >= end {
        let line: String = verse(number);
        lyrics.push_str(&line);

        if number == end {
            break;
        }
        lyrics.push_str("\n");
        number -= 1;
    }
    lyrics
}

fn none() -> String {
    "No more bottles of beer on the wall, no more bottles of beer.\nGo to the store and buy some more, 99 bottles of beer on the wall.\n".to_string()
}

fn one(n: u32) -> String {
    format!("{}Take it down and pass it around, no more bottles of beer on the wall.\n", first_phrase_singular(n))
}

fn two(n: u32) -> String {
    format!("{}{}", first_phrase_plural(n), second_phrase_singular(n-1))
}

fn multiple(n: u32) -> String {
    format!("{}{}", first_phrase_plural(n), second_phrase_plural(n-1))
}

fn first_phrase_singular(n: u32) -> String {
    format!("{} bottle of beer on the wall, {} bottle of beer.\n", n, n)
}

fn first_phrase_plural(n: u32) -> String {
    format!("{} bottles of beer on the wall, {} bottles of beer.\n", n, n)
}

fn second_phrase_singular(n: u32) -> String {
    format!("Take one down and pass it around, {} bottle of beer on the wall.\n", n)
}

fn second_phrase_plural(n: u32) -> String {
    format!("Take one down and pass it around, {} bottles of beer on the wall.\n", n)
}
