# sketch

`sketch` is a Go package that provides efficient implementations for probabilistic data structures.

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

# References

- [A.Gakhov Probabilistic Data Structures](https://www.gakhov.com/books/pdsa.html)
