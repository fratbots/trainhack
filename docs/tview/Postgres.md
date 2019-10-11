The `tview` package makes it easy to write terminal based demo applications. The following application is a simple PostgreSQL database browser written in ~300 lines of Go code:

[[https://github.com/rivo/tview/blob/master/demos/postgres.png]]

## How to Run

Apart from [github.com/rivo/tview](https://github.com/rivo/tview), the script also uses [github.com/lib/pq](https://github.com/lib/pq).

To run it, you must provide the PostgreSQL database connection string (see [here](https://godoc.org/github.com/lib/pq) for details):

```sh
go run src/postgres.go "user='xyz' password='abc' sslmode=disable"
```

## Source Code

You can find the source code for this demo application in [this gist](https://gist.github.com/rivo/2893c6740a6c651f685b9766d1898084).