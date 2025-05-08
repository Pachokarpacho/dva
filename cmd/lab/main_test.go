package main

import "testing"

func TestIsPrime(t *testing.T) {
    // Тест-кейсы
    testCases := []struct {
        input    int
        expected bool
    }{
        {1, true},
        {2, true},
        {3, true},
        {4, false},
        {7, true},
        {10, false},
        {17, true},
        {0, false},
        {-4, false},
    }

    for _, tc := range testCases {
        t.Run(fmt.Sprintf("isPrime(%d)", tc.input), func(t *testing.T) {
            result := isPrime(tc.input)
            if result != tc.expected {
                t.Errorf("isPrime(%d) = %v; want %v", tc.input, result, tc.expected)
            }
        })
    }
}