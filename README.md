
Time range library for Go.

Check if two ranges intersect.
```go
range1.Intersection(range2)
```


Get duration of ranges.
```go
range1.Duration()
```

Check if ranges are equivalent.
```go
range1.Equals(range2)
```


Lengthen ranges with duration.
```go
range1.Length(time.Hour * 3)
```