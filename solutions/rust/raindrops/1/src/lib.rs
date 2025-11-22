pub fn raindrops(n: u32) -> String {

    match (n % 7, n % 5, n % 3) {
        (0, 0, 0) => "PlingPlangPlong".to_string(),
        (0, 0, _) => "PlangPlong".to_string(),
        (0, _, 0) => "PlingPlong".to_string(),
        (_, 0, 0) => "PlingPlang".to_string(),
        (_, _, 0) => "Pling".to_string(),
        (_, 0, _) => "Plang".to_string(),
        (0, _, _) => "Plong".to_string(),
        (_, _, _) => n.to_string(),
    }
}
