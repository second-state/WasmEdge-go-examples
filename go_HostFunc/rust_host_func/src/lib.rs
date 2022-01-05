extern "C" {
	fn fetch(url_pointer: *const u8, url_length: i32) -> i32;
	fn write_mem(pointer: *const u8);
}

#[no_mangle]
pub unsafe extern fn run() -> i32 {
	let url = "https://www.google.com";
	let pointer = url.as_bytes().as_ptr();

	// call host function to fetch the source code, return the result length
	let res_len = fetch(pointer, url.len() as i32) as usize;

	// malloc memory
	let mut buffer = Vec::with_capacity(res_len);
	let pointer = buffer.as_mut_ptr();

	// call host function to write source code to the memory
	write_mem(pointer);

	// find occurrences from source code
	buffer.set_len(res_len);
	let str = std::str::from_utf8(&buffer).unwrap();
	str.matches("google").count() as i32
}