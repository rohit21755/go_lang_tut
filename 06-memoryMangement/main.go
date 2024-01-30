package main

// new() => Allocate memory but not INIT(initialized)
// - you will get memory address
// - zeroed storage (means we cannot put data initially)

// make() => Allocate memory and INIT
// - you will get memory address
// - non-zeroed storage

// garbage collection happens automatically
// out of scope or nil
