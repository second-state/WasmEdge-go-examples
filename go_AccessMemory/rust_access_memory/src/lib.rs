use std::mem;
use std::os::raw::{c_void, c_int};

#[no_mangle]
pub extern fn allocate(size: usize) -> *mut c_void {
    let mut buffer = Vec::with_capacity(size);
    let pointer = buffer.as_mut_ptr();
    mem::forget(buffer);

    pointer as *mut c_void
}

#[no_mangle]
pub extern fn deallocate(pointer: *mut c_void, capacity: usize) {
    unsafe {
        let _ = Vec::from_raw_parts(pointer, 0, capacity);
    }
}

#[no_mangle]
pub extern fn fib_array(n: i32, p: *mut c_int) -> i32 {
    unsafe {
        let mut arr = Vec::<i32>::from_raw_parts(p, 0, (4*n) as usize);
        for i in 0..n {
            if i < 2 {
                arr.push(i);
            } else {
                arr.push(arr[(i - 1) as usize] + arr[(i - 2) as usize]);
            }
        }
        let r = arr[(n - 1) as usize];
        mem::forget(arr);
        r
    }
}

#[no_mangle]
pub extern fn fib_array_return_memory(n: i32) -> *mut c_int {
    let mut arr = Vec::with_capacity((4 * n) as usize);
    let pointer = arr.as_mut_ptr();
    for i in 0..n {
        if i < 2 {
            arr.push(i);
        } else {
            arr.push(arr[(i - 1) as usize] + arr[(i - 2) as usize]);
        }
    }
    mem::forget(arr);
    pointer
}
