# WasmEdge-Go Fibonacci example

## Build

Before trying this example, the [WasmEdge installation](https://wasmedge.org/book/en/start/install.html) is required.

```bash
wget -qO- https://raw.githubusercontent.com/WasmEdge/WasmEdge/master/utils/install.sh | bash -s -- -v 0.11.2
```

Then you can build this example.

```bash
# In the current directory.
$ go get github.com/second-state/WasmEdge-go/wasmedge@v0.11.2
$ go build
```

## Run

```bash
# Run in interpreter mode
$ ./print_fibonacci
```

The output will be as the following:

```bash
registered modules:  [host wasi_snapshot_preview1 wasm]
 --- Exported instances of the anonymous module
     --- Functions ( 1 ) :  [print_val_and_fib]
     --- Tables    ( 0 ) :  []
     --- Memories  ( 0 ) :  []
     --- Globals   ( 0 ) :  []
 --- Exported instances of the module host
     --- Functions ( 1 ) :  [print_val_and_res]
     --- Tables    ( 0 ) :  []
     --- Memories  ( 0 ) :  []
     --- Globals   ( 0 ) :  []
 --- Exported instances of the module wasi_snapshot_preview1
     --- Functions ( 57 ) :  [args_get args_sizes_get clock_res_get clock_time_get environ_get environ_sizes_get fd_advise fd_allocate fd_close fd_datasync fd_fdstat_get fd_fdstat_set_flags fd_fdstat_set_rights fd_filestat_get fd_filestat_set_size fd_filestat_set_times fd_pread fd_prestat_dir_name fd_prestat_get fd_pwrite fd_read fd_readdir fd_renumber fd_seek fd_sync fd_tell fd_write path_create_directory path_filestat_get path_filestat_set_times path_link path_open path_readlink path_remove_directory path_rename path_symlink path_unlink_file poll_oneoff proc_exit proc_raise random_get sched_yield sock_accept sock_bind sock_connect sock_getaddrinfo sock_getlocaladdr sock_getpeeraddr sock_getsockopt sock_listen sock_open sock_recv sock_recv_from sock_send sock_send_to sock_setsockopt sock_shutdown]
     --- Tables    ( 0 ) :  []
     --- Memories  ( 0 ) :  []
     --- Globals   ( 0 ) :  []
 --- Exported instances of the module wasm
     --- Functions ( 1 ) :  [fib]
     --- Tables    ( 0 ) :  []
     --- Memories  ( 0 ) :  []
     --- Globals   ( 0 ) :  []
 ### Running print_val_and_fib with fib[ 20 ] ...
 [HostFunction] external value:  123456  , fibonacci number:  10946
 ### Running print_val_and_fib with fib[ 21 ] ...
 [HostFunction] external value:  876543210  , fibonacci number:  17711
 ### Running wasm::fib[ 22 ] ...
 Return value:  28657
```

If you want to try this example in AOT mode, please follow the [Wasm AOT example](https://github.com/second-state/WasmEdge-go-examples/tree/master/go_WasmAOT) to compile the WASM file.
