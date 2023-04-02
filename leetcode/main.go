package main

/*import (
    "strings"
    "strconv"
)*/

func runningSum(nums []int) []int {
    var result = make([]int, len(nums))
    
    for i := range(nums) {
        result[i] = 0
    }

    for j := 0; j < len(nums); j++ {
        if j == 0 {
            result[j] = nums[j]
        }else{
            result[j] = result[j - 1] + nums[j]
        }
    }

    return result
}

func sum(arr []int) int {
    var sum int = 0
    for i := range(arr) {
        sum = sum + arr[i]
    }
    return sum
}

func pivotIndex(nums []int) int {
    var index int = -1
    
    if len(nums) == 1 {
        index = 0
    }else{
        for i := range(nums) {
            if i != len(nums) - 1 {
                if sum(nums[0: i]) == sum(nums[i + 1:]) {
                    index = i
                    break
                }
            }else {
                if sum(nums[0: i]) == 0 {
                    index = i 
                    break
                }
            }
        }
    }
    
    return index
}

func isIsomorphic(s string, t string) bool {
    var isIso bool = true
    data := make(map[uint8]uint8)
    data2 := make(map[uint8]uint8)

    // checking for the first string
    for i := 0; i < len(s); i++ {
        c := s[i]
        if _, ok:= data[c]; ok {
            if t[i] != data[c] {
                isIso = false
                break
            }
        } else {
            data[c] = t[i]
        }
    }

    // checking for the second string
    for i := 0; i < len(t); i++ {
        c := t[i]
        if _, ok:= data2[c]; ok {
            if s[i] != data2[c] {
                isIso = false
                break
            }
        } else {
            data2[c] = s[i]
        }
    }
    
    return isIso
}

func getIndex(c rune, str string) int {
    var index int = -1
    for i, char := range(str) {
        if char == c {
            index = i   
        }
    }
    return index
}

func isSubsequence(s string, t string) bool {
    isSeq := true
    index := -1
    
    for i := 0; i < len(s); i++ {
        c := s[i]
        i := getIndex(rune(c), t)
        if i != -1 {
            if i >= index {
                index = i
            }else {
                isSeq = false
                break
            }
        }else {
            isSeq = false
            break
        }
    }

    return isSeq
}

func main(){
    println(isSubsequence("abc", "ahbgdc"))
}
