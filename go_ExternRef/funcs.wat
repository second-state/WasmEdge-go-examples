(module
  (type $t0 (func (param externref i32) (result i32)))
  (type $t1 (func (param externref i32 i32) (result i32)))
  (type $t2 (func (param externref externref i32 i32) (result i32)))
  (import "extern_module" "square" (func $square (type $t0)))
  (import "extern_module" "add" (func $add (type $t1)))
  (import "extern_module" "mul" (func $mul (type $t1)))
  (func $call_add (export "call_add") (type $t1) (param $p0 externref) (param $p1 i32) (param $p2 i32) (result i32)
    (call $add
      (local.get $p0)
      (local.get $p1)
      (local.get $p2)))
  (func $call_mul (export "call_mul") (type $t1) (param $p0 externref) (param $p1 i32) (param $p2 i32) (result i32)
    (call $mul
      (local.get $p0)
      (local.get $p1)
      (local.get $p2)))
  (func $call_square (export "call_square") (type $t0) (param $p0 externref) (param $p1 i32) (result i32)
    (call $square
      (local.get $p0)
      (local.get $p1)))
  (func $call_add_square (export "call_add_square") (type $t2) (param $p0 externref) (param $p1 externref) (param $p2 i32) (param $p3 i32) (result i32)
    (call $square
      (local.get $p1)
      (call $add
        (local.get $p0)
        (local.get $p2)
        (local.get $p3))))
  (memory $memory (export "memory") 1))
