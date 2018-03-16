# ilcd
This is a Go package for reading
[ILCD data sets](http://eplca.jrc.ec.europa.eu/LCDN/developer.xhtml). It
currently supports only a subset of the ILCD format but allows very fast and
memory efficient extraction of information from the different ILCD data set
types.

## Usage

#### Iterating over each data set in an ILCD package
The `ilcd.ZipReader` provides methods for extracting each data set of a specific
type via the `Each*` functions:

```go
zip, err := ilcd.NewZipReader("an_ILCD_package.zip")
// check err ...
err = zip.EachProcess(func(p *ilcd.Process) bool {
  fmt.Println(p.FullName("en"))
  return true // returning false will stop the iteration
})
```

#### Transforming and ILCD package

```go
reader, err := ilcd.NewZipReader("input.zip")
check(err)
writer, err := ilcd.NewZipWriter("output.zip")
check(err)

reader.Map(writer, func(entry *ilcd.ZipFile) (string, []byte) {
  if entry.Type() == ilcd.FlowDataSet {
    data, err := entry.Read()
    check(err)
    return entry.Path(), data
  }
  return "", nil // ignore other data sets
})

check(reader.Close())
check(writer.Close())
```
