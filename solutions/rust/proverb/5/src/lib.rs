const MODERN_KEYWORD: &str = "pin";
const TRADITIONAL_KEYWORD: &str = "nail";

pub fn build_proverb(list: &[&str]) -> String {
    let mut verses = String::from("");

    let modern_period = is_modern_period(list);

    if !list.is_empty() {
        for &word in list.iter() {
            let verse = match word {
                "pin" => "",
                "nail" => "",
                "gun" => "pin",
                "shoe" => "nail",
                "horse" => "shoe",
                "rider" => "horse",
                "message" => "rider",
                "battle" => {
                    if modern_period {
                        "soldier"
                    } else {
                        "message"
                    }
                }
                "soldier" => "gun",
                "kingdom" => "battle",
                _ => "",
            };

            if word != TRADITIONAL_KEYWORD && word != MODERN_KEYWORD {
                verses.push_str(&regular_verse(verse.to_string(), word.to_string()));
            }
        }
        if modern_period {
            verses.push_str(&end_verse(MODERN_KEYWORD.to_string()));
        } else {
            verses.push_str(&end_verse(TRADITIONAL_KEYWORD.to_string()));
        }
    }

    verses
}

fn regular_verse(want: String, lost: String) -> String {
    format!("For want of a {} the {} was lost.\n", want, lost)
}

fn end_verse(word: String) -> String {
    format!("And all for the want of a {}.", word)
}

fn is_modern_period(list: &[&str]) -> bool {
    for &word in list.iter() {
        if word == MODERN_KEYWORD {
            return true;
        }
    }

    false
}
