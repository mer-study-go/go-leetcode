# Slices

**Go Slice** is an abstraction over **Go Array**. 

Multiple data elements of the **same** type are allowed by Go arrays. 

The definition of variables that can be hold several data elements of the same type are allowed by Go Array, but it does not have any provision of inbuilt methods to increase its size in Go. This shortcoming is taken care of by Slices. 

A Go slice can be **appended** to elements after the capacity has reached its size. Slices are dynamic and can double the current capacity in order to add more elements. 

## `len()` function

```go
var slice = []int{1, 3, 5, 6}
slice = append(slice, 23)
fmt.Println("Caapacity", cap(slice))
fmt.Println("Length", len(slice))
```

Output

```
Capacity 8
Length 5
```

## Slice function

Slices are passed by referring to functions. Big slices can be passed to functions without impacting performance. Passing a slice as a reference to a function is demonstrated in the code as follows (`slice.go`)

```go
// twiceValue method
func twiceValue(slice []int) {
	var i int
	var value int
	for i, value = range slice {
		slice[i] = 2 * value
	}
}

func main() {
	// slice
	twiceValue(slice)
	var i int
	for i = 0; i < len(slice); i++ {
		fmt.Println("new slice value", slice[i])
	}
}

```

For dynamic allocation, we use slice of slices. In the following code, slice of slices is explained as two-dimensional slices 

`twodslices.go`

