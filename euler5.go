package main

import  (
    "fmt"
    "math/big"
)

func isPrime(value *big.Int, primes []big.Int ) bool {
    devisibleByPrime := false
    fmt.Print(value)
    fmt.Print(" ")    
    for _, prime := range primes {
        var mod big.Int
        mod = *mod.Mod(value, &prime)
        fmt.Print(mod.String())
        fmt.Print(" ")
        compare := mod.Cmp(big.NewInt(0))
        if compare == 0 {
            devisibleByPrime = true
        }
    }
    fmt.Println()
    return !devisibleByPrime
}

func isAnswer(value big.Int, bound big.Int) bool {
    one  := big.NewInt(1)
    zero := big.NewInt(0)
    mod  := big.NewInt(1)
    for i := one; i.Cmp(&bound) == -1; i.Add(i, one) {
        mod.Mod(&value, i)
        if mod.Cmp(zero) != 0 {
            return false
        }
    }
    return true
}

func main() {
    upperBound := big.NewInt(20)
    one  := big.NewInt(1)

    var primes []big.Int
    primes = append(primes, *big.NewInt(2))
    for i := big.NewInt(3); i.Cmp(upperBound) == -1; i.Add(i, one) {
        if isPrime(i, primes) {
            primes = append(primes, *i)
            fmt.Println(primes)
        }
    }

    fmt.Println(primes)

    increment := big.NewInt(1)
    for _, prime := range primes {
        increment = increment.Mul(increment, &prime)
    }

    for check := increment; true; check.Add(check,increment) {
        if(isAnswer(*check, *upperBound)) {
            fmt.Println(check)
            return
        }
    }
}