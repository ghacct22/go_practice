package main

import (
    "fmt"
    "math/big"
)

func main() {
    factorial   := big.NewInt(100)
    product     := big.NewInt(1)
    one         := big.NewInt(1)
    ten         := big.NewInt(10)
    for ; factorial.Cmp(one) != 0; factorial.Sub(factorial, one) {
        product.Mul(product, factorial)
    } 

    sum := big.NewInt(0)
    mod := big.NewInt(0)
    for product.Cmp(one) == 1 {
        mod.Mod(product, ten)
        sum.Add(sum, mod)
        product.Div(product, ten)
    }

    fmt.Println(sum)
}