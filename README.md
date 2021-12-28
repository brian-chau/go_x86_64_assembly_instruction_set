## x86_64 assembly instruction set dictionary

This is a Go application that fetches system call function definitions from the [ChromiumOS documentation](https://chromium.googlesource.com/chromiumos/docs/+/HEAD/constants/syscalls.md#x86_64-64_bit).

I copied the definitions into a CSV, and run this Go application to search for the values based on the CSV column headers.

Example usage:

To find the syscall definition for the "open" syscall:

> ./syscalls_x86 open

To find the syscall definition for the "write" syscall:

> ./syscalls_x86_64 write
