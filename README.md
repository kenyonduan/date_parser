# Usage
Expand golang default `time.Parser(layout, value)`.

### chinese
```go
tval, _ := time.Parse("2006-01-02", "2018-09-13")
eval, _ = date_parser.ParserLangDate("zh", "on 2018年09月13日", "on 2006年January02日")
assert.Equal(t, tval, eval)
```

### Germany
```go
tde, _ := time.Parse("2006-01-02", "2018-12-26")
eval, _ = date_parser.ParserLangDate("de", "am 26. Dezember 2018", "am 02. January 2006")
assert.Equal(t, tde, eval)
```

### Italy
```go
tval, _ := time.Parse("2006-01-02", "2018-09-13")
eval, _ = date_parser.ParserLangDate("it", "il 13 settembre 2018", "il 02 January 2006")
assert.Equal(t, tval, eval)
```
