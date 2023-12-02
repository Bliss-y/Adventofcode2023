package main

import (    
    "fmt"
    "strings"
    "unicode"
    "strconv"
    "os"
)

func day1(filepath string){
    input, _ := os.ReadFile(filepath);
    lines := strings.Split(string( input), "\r\n");
    total := 0;
    for _,line := range lines {
        nums := []rune{0,0}
        for i := 0; i < len(line); i++ {
            c := rune(line[i]);
            if unicode.IsDigit(c) {
                if nums[0] == 0 {
                    nums[0] = c;
                    nums[1] = c;
                }else {
                    nums[1] = c;
                }
            } else {
                currentNum := rune(0);
                if len(line) >= i+3{
                    p := string(line[i:i+3])
                switch p{
                    case "one":
                        currentNum = '1';
                    case "two":
                        currentNum = '2';
                    case "six":
                        currentNum = '6';
                }
                }
                if len(line) >= i+4 {
                switch string(line[i:i+4]){
                    case "four":
                        currentNum = '4';
                    case "five":
                        currentNum = '5';
                    case "nine":
                        currentNum = '9'
                    case "zero":
                        currentNum = '0';
                }
                }
                if len(line) >= i + 5{
                switch string(line[i:i+5]){
                    case "three":
                        currentNum = '3';
                    case "eight":
                        currentNum = '8';
                    case "seven":
                        currentNum = '7'
                }
                }
                if(currentNum > 0){
                    if nums[0] == 0 {
                        nums[0] = currentNum;
                        nums[1] = currentNum;
                    } else {
                        nums[1] = currentNum;
                    }
                }
            }
        }
        n,_ := strconv.Atoi(string(nums))
        fmt.Println(n);
        total += n;
    }
    fmt.Println("day",filepath,"part1: ", total);
}
