package diffsquares

import "math"

func Difference(n int) int{
    return SquareOfSum(n) - SumOfSquares(n)
}

func SquareOfSum(n int) int{
    var summation float64
    for i := 1; i <= n; i++ {
        summation += float64(i)
    }
    return int(math.Pow(summation, 2))
}

func SumOfSquares(n int) int{
    var summation float64
    for i := 1; i <= n; i++ {
        summation += math.Pow(float64(i),2)
    }
    return int(summation)
}
