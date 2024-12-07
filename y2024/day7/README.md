# Learning

Concating `int`s:

Naïve implementation:
(took .5 seconds long on part2)

```go
strconv.Atoi(fmt.Sprintf("%d%d",lhs,rhs)))
```

Using logarithm:

```go
func concatInts2(lhs, rhs int) int {
	return int(float64(lhs)*math.Pow(10, math.Ceil(math.Log10(float64(rhs+1))))) + rhs
}
```

Something I found on other's solutions:

```go
func concatInts(lhs, rhs int) int {
	tmp := rhs
	for {
		lhs *= 10
		tmp /= 10
		if tmp <= 0 {
			break
		}
	}
	return lhs + rhs
}
```
