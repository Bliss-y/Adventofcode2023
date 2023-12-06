package main

import (
	"fmt"
	"os"
    "unicode"
	"strconv"
	"strings"
)

var INPUT_DIR = "input"

func main() {
    day3();
}

func day2(){
    input,_ := os.ReadFile("inputs/2.txt");
    games := strings.Split(string(input), "\r\n");
    total := 0;
        limit := make(map[string]int); 
    for i,game:= range games {
        if len(game) == 0 {
            fmt.Println("end")
            break;
        }
        limit["red"] = 0;
        limit["blue"] = 0;
        limit["green"] = 0;
        f := strings.Split(game, ": ");
        gameIdstr := strings.Split(f[0], " ")[1];
        sets := strings.Split(f[1], "; ");
        for _,set := range sets {
            balls := strings.Split(set, ", ");
            for _, ball := range balls {
                ballh := strings.Split(ball, " ");
                amount,_ := strconv.Atoi(ballh[0]);
                if amount > limit[ballh[1]] {
                    limit[ballh[1]] = amount;
                }
            }
        }
        total += limit["red"]*limit["blue"]*limit["green"]
        gameIdnum,_ := strconv.Atoi(gameIdstr);
        if(gameIdnum != i+1){
            panic("gameIdnum is not what it should be !")
        }

    }
    fmt.Println("day2:", total);
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}

type position struct {
    x int;
    y int;
}

func day3() {
    input,_ := os.ReadFile("inputs/3.txt");
    lines := strings.Split(string(input), "\r\n");
    positions := make([]position,0)
    currentGears := make(map[position][]int);
    for y, line := range lines {
        if len(line) == 0 {
            break;
        }
        nums := make([]rune, 0, 3)
        reading := false
        isPart := false
        partPos := position{0,0};
        directions := [8]position{{-1,0}, {1,0}, {0,1}, {0,-1}, {1,1}, {1,-1}, {-1, 1}, {-1,-1}}
        for i := 0; i < len(line); i++ {
            c := rune(line[i]);
            if unicode.IsDigit(c) {
                reading = true;
                nums = append(nums, c);
                for _,d := range directions {
                    finX := i + d.x;
                    finY := y + d.y;
                    if finX >= 0 && finX < len(line) && finY>=0 && finY < len(lines)-1{
                        if char := rune(lines[finY][finX]); !unicode.IsDigit(char) && char != '.' {
                            isPart = true;
                            partPos = position{finX, finY};
                            break;
                        }
                    }
                }
            }else {
                if reading{
                    if isPart {
                        n,_ := strconv.Atoi(string(nums));
                        if len(currentGears[partPos]) > 0 {
                            positions = append(positions, partPos);
                            currentGears[partPos] = append(currentGears[partPos], n);
                        } else {
                            currentGears[partPos] = []int{n};
                        }
                    }
                    nums = make([]rune, 0, 3);
                    isPart = false;
                }
            }
        }
        if reading{
            if isPart {
                n,_ := strconv.Atoi(string(nums));
                if len(currentGears[partPos]) > 0 {
                    positions = append(positions, partPos);
                    currentGears[partPos] = append(currentGears[partPos], n);
                } else {
                    currentGears[partPos] = []int{n};
                }
            }
            nums = make([]rune, 0, 3);
            isPart = false;
        }
    }
    total := 0
    for _,i := range positions {
        if nums :=  currentGears[i]; len(nums) == 2 {
            fmt.Println(i,currentGears[i]);
            total += nums[0] * nums[1];
        }
    }
    fmt.Println(total);
}
