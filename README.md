# `SetMap`

A "set map" struct implemented by Golang. Just a practice for Golang generic.

## What is "set map"?

```c
map<K, set<V>>
```

## How am I implement it?

```go
type SetMap[K, V comparable] struct {
	m map[K]map[V]struct{}
}
```

## How to use?

run `setmap_test.go` and you would see.

# `SetMapaMteS`

A "set map" struct implemented by Golang, but can also get a list of key by the value.


## How am I implement it?

```go
type SetMapaMteS[K, V comparable] struct {
    SetMap[K, V]
    reverse SetMap[V, K]
}
```

## How to use?

run `setmap_test.go` and you would see.


