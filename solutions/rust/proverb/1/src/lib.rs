pub fn build_proverb(list: &[&str]) -> String {
    let mut verses = String::from("");

    let modern_period = is_modern_period(list);

    if !list.is_empty() {
        for &word in list.iter() {
            let verse = match word {
                "pin" => "".to_string(),
                "nail" => "".to_string(),
                "gun" => "pin".to_string(),
                "shoe" => "nail".to_string(),
                "horse" => "shoe".to_string(),
                "rider" => "horse".to_string(),
                "message" => "rider".to_string(),
                "battle" => {
                    if modern_period {
                        "soldier".to_string()
                    } else {
                        "message".to_string()
                    }
                }
                "soldier" => "gun".to_string(),
                "kingdom" => "battle".to_string(),
                _ => "".to_string(),
            };

            if word != "nail" && word != "pin" {
                verses.push_str(&format!(
                    "For want of a {} the {} was lost.\n",
                    &verse, word
                ));
            }
        }
        if modern_period {
            verses.push_str(&format!("And all for the want of a {}.{}", "pin", ""));
        } else {
            verses.push_str(&format!("And all for the want of a {}.{}", "nail", ""));
        }
    }

    verses
}

fn is_modern_period(list: &[&str]) -> bool {
    for &word in list.iter() {
        if word == "pin" {
            return true;
        }
    }

    false
}
