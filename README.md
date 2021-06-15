# sketch

`sketch` is a Go package that provides implementations for probabilistic data structures.

The implementations are based largely on literature and are a bit hacky, I don't recommend using
them for any reason other than _figuring out_ how PDS works.

Better suitable implementations can be found in Github.

# Trade Offs

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
The recommend _precision_ value (denoted m) should take small values for < 10e8 entries.

| n         | m (δ = 1%) | m (δ = 10%) |
| --------- | ---------- | ----------- |
| 1000      | 4          | 7           |
| 10000     | 4          | 8           |
| 100000    | 5          | 10          |
| 1000000   | 6          | 10          |
| 10000000  | 7          | 12          |
| 100000000 | 8          | 14          |

# References

- [A.Gakhov Probabilistic Data Structures](https://www.gakhov.com/books/pdsa.html)
