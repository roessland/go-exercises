# fetchall

Fetches a list of urls concurrently. Writes responses and errors to file (fetchlog.txt in current directory). Writes errors to stdout. Writes summary to stdout.

```
go run fetchall.go https://golang.org http://gopl.io https://godoc.org https://nope.io http://slowly.com/fdfdfdsf
Get https://nope.io: dial tcp: lookup nope.io: no such host
0.26s     8158  https://golang.org
0.28s     6811  https://godoc.org
0.41s     4154  http://gopl.io
0.96s    12518  http://slowly.com/fdfdfdsf
0.96s elapsed
```