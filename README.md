# ilcd
This is a Go package for reading
[ILCD data sets](http://eplca.jrc.ec.europa.eu/LCDN/developer.xhtml). It
currently supports only a subset of the ILCD format but allows very fast and
memory efficient extraction of information from the different ILCD data set
types. Here is an example for reading all processes in an ILCD zip file:

```go
zip, err := ilcd.NewZipReader("an_ILCD_package.zip")
// check err ...
err = zip.EachProcess(func(p *ilcd.Process) bool {
    fmt.Println(p.FullName("en"))
	return true // returning false will stop the iteration
})
```
