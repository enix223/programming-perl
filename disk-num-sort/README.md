# Disk numbers sort

## Gen 10e7 unique random numbers

```
python gen.py ./rand-num-10e7.bin
```

## Sort numbers and output to file

```
python py/sort.py rand-num-10e7.bin rand-num-10e7-sort-py.bin
```


## Sort numbers and output to file (golang)

```
go run go/sort.go -infile rand-num-10e7.bin -outfile rand-num-10e7-sort-go.bin
```