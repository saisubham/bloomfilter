# Bloom Filter in Go
This is an implementation of bloom filters in Go. According to Wikipedia:
> A Bloom filter is a space-efficient probabilistic data structure, conceived by Burton Howard Bloom in 1970, that is used to test whether an element is a member of a set. False positive matches are possible, but false negatives are not.

## Compiling
```
git clone https://github.com/saisubham/bloomfilter.git
cd ./bloomfilter
go mod tidy
```
## Examples
```
go run .\bloomfilter.go
> cat
[cat]
     
> mat
[cat mat]

> hello  
[cat mat hello]

> mat
mat may already be present
[cat mat hello]

> tree
tree may already be present
[cat mat hello]

> exit
```
