#![allow(clippy::needless_return)]

pub fn nth(n: usize) -> usize {
    let mut sequence: Vec<usize> = Vec::new();
    let mut number: usize = 1;

    while sequence.len() <= n {
        if is_prime(number, 2) {
            sequence.push(number);
        }
        number += 1;
    }

    return sequence.pop().unwrap();
}

pub fn is_prime(n: usize, i: usize) -> bool {
    if n <= 2 {
        return n == 2;
    }

    if n % i == 0 {
        return false;
    }

    if i * i > n {
        return true;
    }

    return is_prime(n, i + 1);
}
