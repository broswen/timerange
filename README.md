
Time range library for Go.


Create ranges with RFC3339
```go
range1, err := New("2021-01-01T05:00:00Z", "2021-01-02T08:00:00Z")
fmt.Println(range1)
// 2021-01-01T05:00:00Z -> 2021-01-02T08:00:00Z
```

Check if two ranges intersect.
```go
range1.Intersection(range2)
// bool
```

Get duration of ranges.
```go
range1.Duration()
// time.Duration
```

Check if ranges are equivalent.
```go
range1.Equals(range2)
// bool
```


Lengthen ranges with duration.
```go
range1.Length(time.Hour * 3)
```