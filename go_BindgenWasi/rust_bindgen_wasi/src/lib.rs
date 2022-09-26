use rand::prelude::*;
use ssvm_wasi_helper::ssvm_wasi_helper::_initialize;
use std::env;
use std::fs;
use std::fs::File;
use std::io::{Read, Write};
#[allow(unused_imports)]
use wasmedge_bindgen::*;
use wasmedge_bindgen_macro::*;

#[wasmedge_bindgen]
pub fn get_random_i32() -> i32 {
    let x: i32 = random();
    return x;
}

#[wasmedge_bindgen]
pub fn get_random_bytes() -> Vec<u8> {
    let mut rng = thread_rng();
    let mut arr = [0u8; 128];
    rng.fill(&mut arr[..]);
    return arr.to_vec();
}

#[wasmedge_bindgen]
pub fn echo(content: String) -> String {
    println!("Printed from wasi: {}", content);
    return content;
}

#[wasmedge_bindgen]
pub fn print_env() {
    _initialize();
    println!("The env vars are as follows.");
    for (key, value) in env::vars() {
        println!("{}: {}", key, value);
    }

    println!("The args are as follows.");
    for argument in env::args() {
        println!("{}", argument);
    }

    /*
    match env::var("PATH") {
      Ok(path) => println!("PATH: {}", path),
      Err(e) => println!("Couldn't read PATH ({})", e),
    };
    */
}

#[wasmedge_bindgen]
pub fn create_file(path: String, content: String) {
    _initialize();
    let mut output = File::create(&path).unwrap();
    output.write_all(content.as_bytes()).unwrap();
}

#[wasmedge_bindgen]
pub fn read_file(path: String) -> String {
    _initialize();
    let mut f = File::open(&path).unwrap();
    let mut s = String::new();
    match f.read_to_string(&mut s) {
        Ok(_) => s,
        Err(e) => e.to_string(),
    }
}

#[wasmedge_bindgen]
pub fn del_file(path: String) {
    _initialize();
    fs::remove_file(&path).expect("Unable to delete");
}
