use project1::front_of_house::hosting;
use project1::back_of_house;

fn main() {
    hosting::add_to_wait_list();
    back_of_house::take_care_trash();
    println!("The value o PI is: {}",project1::PI);
}