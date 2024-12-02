fn main() {
    let numbers = [2, -3, 7, 0, 8, -1, 5, -4, 6, 10];

    match average_positive(&numbers) {
        Some(average) => println!("average positive numbers {}", average),
        None => println!("no positive numbers"),
    }

    let product = product_of_evens(&numbers);
    println!("product of even numbers {}", product);
}

fn average_positive(arr: &[i32; 10]) -> Option<f64> {
    let positives: Vec<i32> = arr.iter().filter(|&&x| x > 0).cloned().collect();
    if positives.is_empty() {
        None
    } else {
        let sum: i32 = positives.iter().sum();
        let average = sum as f64 / positives.len() as f64;
        Some(average)
    }
}

fn product_of_evens(arr: &[i32; 10]) -> i32 {
    let evens: Vec<i32> = arr.iter().filter(|&&x| x % 2 == 0 && x != 0).cloned().collect();
    if evens.is_empty() {
        1
    } else {
        evens.iter().product()
    }
}
