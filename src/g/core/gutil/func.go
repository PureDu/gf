package gutil

import (
    "time"
    "math/rand"
)

// 框架自定义函数库

// 获得一个 min, max 之间的随机数
func Rand (min, max int) int {
    //fmt.Printf("min: %d, max: %d\n", min, max)
    if min >= max {
        return min
    }
    rand.Seed(time.Now().UnixNano())
    n := rand.Intn(max)
    if n < min {
        return Rand(min, max)
    }
    return n
}