# sketch

![Go Report Card](https://goreportcard.com/badge/github.com/actuallyachraf/sketch)

---

`sketch` is a Go package that provides implementations for probabilistic data structures.

The implementations are based largely on literature and are a bit hacky, I don't recommend using
them for any reason other than _figuring out_ how PDS works or for demo's or mini-projects.

While the implementations are fast enough, they were not formally verified for correctness
_tests_ offer 1%~10% accuracy.

Better suitable implementations can be found in Github.

## Trade Offs

For **LinearCounter** the size-accuracy trade-offs for a few bitvector size values have been computed
and provided by [Ilya Katsov](https://highlyscalable.wordpress.com/2012/05/01/probabilistic-structures-web-analytics-data-mining/).

- Trade-off between accuracy (δ) and bit array size (m)

| n         | m (δ = 1%) | m (δ = 10%) |
| --------- | ---------- | ----------- |
| 1000      | 5329       | 268         |
| 10000     | 7960       | 1709        |
| 100000    | 26729      | 12744       |
| 1000000   | 154171     | 100880      |
| 10000000  | 1096582    | 831809      |
| 100000000 | 8571013    | 7061760     |

For **HyperLogLog** the accuracy-register count trade-offs are based on the estimated cardinality of input.
The recommended _precision_ value (denoted m) should take small values for < 10e8 entries.

| n         | m (δ = 1%) | m (δ = 10%) |
| --------- | ---------- | ----------- |
| 1000      | 4          | 7           |
| 10000     | 4          | 8           |
| 100000    | 5          | 10          |
| 1000000   | 6          | 10          |
| 10000000  | 7          | 12          |
| 100000000 | 8          | 14          |

## Benchmarks

### LinearCounter

```sh

BenchmarkLinearCounter
BenchmarkLinearCounter/BenchmarkLinearCounter-5329
BenchmarkLinearCounter/BenchmarkLinearCounter-5329-8         	 2298578	       629.2 ns/op	      64 B/op	       2 allocs/op
BenchmarkLinearCounter/BenchmarkLinearCounter-5329#01
BenchmarkLinearCounter/BenchmarkLinearCounter-5329#01-8      	 1894555	       599.8 ns/op	      64 B/op	       2 allocs/op

```

### HyperLogLog

```sh

BenchmarkHyperLogLog
BenchmarkHyperLogLog/BenchmarkHyperLogLog-5
BenchmarkHyperLogLog/BenchmarkHyperLogLog-5-8         	30878364	        38.76 ns/op	       0 B/op	       0 allocs/op
BenchmarkHyperLogLog/BenchmarkHyperLogLog-8
BenchmarkHyperLogLog/BenchmarkHyperLogLog-8-8         	32440388	        40.03 ns/op	       0 B/op	       0 allocs/op
BenchmarkHyperLogLog/BenchmarkHyperLogLog-5#01
BenchmarkHyperLogLog/BenchmarkHyperLogLog-5#01-8      	25808560	        38.98 ns/op	       0 B/op	       0 allocs/op

```

## References

- [A.Gakhov Probabilistic Data Structures](https://www.gakhov.com/books/pdsa.html)
