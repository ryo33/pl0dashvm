# pl0dashvm
A VM for [simozono/pl0dash](https://github.com/simozono/pl0dash)

## Installation
First, install golang, prepare `$GOPATH`, and add `$GOPATH/bin` to the `$PATH`, if you have not done.  
Run the following commands:  
`$ go get -u github.com/ryo33/pl0dashvm`  
`$ go install github.com/ryo33/pl0dashvm`  

## Feature
See the following files:  
[features/run.feature](features/run.feature)  
[features/parse_fail.feature](features/parse_fail.feature)  
[features/process_fail.feature](features/process_fail.feature)  
[features/trace.feature](features/trace.feature)  

## Usage
Run:  
`$ pl0dashvm out.asm`  
Run with trace option:  
`$ pl0dashvm -t out.asm`  
Display help:  
`$ pl0dashvm -h`  

## Test
Run the following commands:  
`$ bundle --path "vendor/bundle" # installs the dependencies for testing`  
`$ bundle exec cucumber # runs test`  
