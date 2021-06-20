package main

import "fmt"

func has(list []int, value int) bool {
  for _, list_value := range list {
    if list_value == value {
      return true
    }
  }
  return false
}

func unique(list []int) (out []int) {
  for _, val := range list {
    if !has(out, val) {
      out = append(out, val)
    }
  }
  return
}

func join_unique(lists ...[]int) (out []int) {
  return unique(join(lists...))
}

func join(lists ...[]int) (out []int) {
  for _, list := range lists {
    out = append(out, list...)
  }
  return
}

func list_filter(list []int, f func(int, int) bool) (out []int) {
  for idx, val := range list {
    if f(val, idx) {
      out = append(out, val)
    }
  }
  return
}

func list_some(list []int, f func(int, int) bool) bool {
  for idx, val := range list {
    if f(val, idx) {
      return true
    }
  }
  return false
}

func list_map(list []int, f func(int, int) int) (out []int) {
  for idx, val := range list {
    out = append(out, f(val, idx))
  }
  return
}

func list_every(list []int, f func(int, int) bool) bool {
  for idx, val  := range list {
    if !f(val, idx) {
      return false
    }
  }
  return true
}

// when we have generics
// func list_reduce[K any, T any](list []K, f func(T, K, int) T, init T) (out T) {
//   out = init
//   for i, v := range list {
//     out = f(out, v, i)
//   }
//   return
// }

// when we have generics
// func array_has[T any](haystack []T, needle T) bool {
// 	for _, val := range haystack {
//     // if val == needle errors out with:
//     // cannot compare val == needle (operator == not defined for T)
//     // https://stackoverflow.com/questions/68053957/go-with-generics-operator-not-defined-for-t
// 		if reflect.ValueOf(val).Interface() == reflect.ValueOf(needle).Interface() {
// 			return true
// 		}
// 	}
// 	return false
// }


func main() {
  list1 := []int{1,2,3}
  list2 := []int{2,3,4}
  fmt.Println(join(list1, list2))
  fmt.Println(join_unique(list1, list2))

  double := func(val int, key int) int { return val * 2 }
  fmt.Println(list_map(list1, double))

  is_even := func(val int, key int) bool { return val % 2 == 0 }
  fmt.Println(list_filter(list2, is_even))

  is_one := func(val int, key int) bool { return val == 1 }
  fmt.Println(list_some(list1, is_one))
  fmt.Println(list_some(list2, is_one))

  less_than_ten := func(val int, key int) bool { return val < 10 }
  fmt.Println(list_every(list1, less_than_ten))
}
