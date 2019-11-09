

### Benchmarking of the functions to print the nodes block

// pkg: github.com/jbaxx/gocode/bst
// BenchmarkPrintDataBytes-8                3323289               363 ns/op
// BenchmarkPrintDataBytesMath-8            4296769               281 ns/op
// BenchmarkPrintDataStrings-8              2971926               402 ns/op

#### Output of the functions
[12]  
[10]  
[ 9]  
[ 7]  
[ 6]  
[ 5]  
[ 4]  
[ 3]  
[ 1]  


```go
func printDataStrings(writer io.Writer, data, max int) {

    maxRuneCount := utf8.RuneCountInString(strconv.Itoa(max))

    var sdataBlock strings.Builder
    sdataBlock.WriteString(strconv.Itoa(data))

    var s string

    for maxRuneCount > sdataBlock.Len() {
        s = sdataBlock.String()
        sdataBlock.Reset()
        sdataBlock.WriteString(" " + s)
    }

    fmt.Fprintf(writer, "[%s]", sdataBlock.String())

}
```


```go
const BLANK = " "

func printDataBytes(writer io.Writer, data, max int) {

    maxRuneCount := utf8.RuneCountInString(strconv.Itoa(max))

    s := []byte(strconv.Itoa(data))

    for maxRuneCount > len(s) {
        s = append([]byte(" "), s...)
    }

    fmt.Fprintf(writer, "[%s]", string(s))

}
```

```go
func printDataBytesMath(writer io.Writer, data, max int) {

    maxRuneCount := utf8.RuneCountInString(strconv.Itoa(max))

    s := []byte(strconv.Itoa(data))

    if l := maxRuneCount - len(s); l > 0 {
        s = append(bytes.Repeat([]byte(BLANK), l), s...)
    }

    fmt.Fprintf(writer, "[%s]", string(s))

}
```


