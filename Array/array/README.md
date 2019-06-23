# Arrays

Arrays are the most famous data structures in different programming languages. Different data types can be handled as elements in arrays such as `int`, `float32`, `double`, and others. 

The following code snippet shows the initialization of an array (`arrays.go`)

```go
var arr = [5]int{1, 2, 4, 5, 6}
```

An array's size can be found with the `len()` function. A `for` loop is used for accessing all the elements in an array, as follows: 

```go
var i int
for i=0; i < len(arr); i++ {
	fmt.Println("printing elements, " arr[i])
}
```

In the following code snippet, the `range` keyword is explained in detail. The `range` keyword can be used to `access` the index and `value` for each element. 

```go
var value=int
for i, value = range, arr {
	fmt.Println(" range ", value)
}
```

The `_` blank identifier is used if the index is ignored. The following code shows how a `_` blank identifier can be used: 

```go
for _, value = range arr {
	fmt.Println("blank range", value)
}
``` 

Go arrays are not dynamic but have a fixed size. To add more elements than the size, a bigger array needs to be created and all the elements of the old one need to be copied. 

An array is passed as a value through functions by copying the array. Passing a big array to a function might be a performance issue. 

