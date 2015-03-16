# go-dbf
A pure Go library for reading and writing [dBase/xBase](http://en.wikipedia.org/wiki/DBase#File_formats) database files

This project is a part of go-shp library. [go-shp](https://github.com/jonas-p/go-shp) is a pure Go implementation of Esri [Shapefile](http://en.wikipedia.org/wiki/Shapefile) format.

You can incorporate the library into your local workspace with the following 'go get' command:

```go
go get github.com/LindsayBradford/go-dbf/godbf
```

Code needing to call into the library needs to include the following import statement:
```go
import (
  "github.com/LindsayBradford/go-dbf/godbf"
)
```

There's no real documentation as yet. Here is a very simple snippet of example 'load' code to get you going:
```go
  dbfTable, err := godbf.NewFromFile("exampleFile.dbf", "UTF8")

  exampleList := make(ExampleList, dbfTable.NumberOfRecords())

  for i := 0; i < dbfTable.NumberOfRecords(); i++ {
    exampleList[i] = new(ExampleListEntry)

    exampleList[i].someColumnId, err = dbfTable.FieldValueByName(i, "SOME_COLUMN_ID")
  }
```

##Licence: 
  This library is made available as-is under the [Apache Licence 2.0](http://www.apache.org/licenses/LICENSE-2.0).
