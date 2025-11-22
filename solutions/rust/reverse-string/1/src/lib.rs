extern crate unicode_segmentation;

use unicode_segmentation::UnicodeSegmentation;

pub fn reverse(input: &str) -> String {
    let mut vec = Vec::new();

    for g in UnicodeSegmentation::graphemes(input, true) {
        vec.insert(0, g);
    }

    return vec.into_iter().collect();
}
