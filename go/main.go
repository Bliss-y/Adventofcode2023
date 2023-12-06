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
    day5();
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

func day4() {
    input, _ := os.ReadFile("inputs/4.txt");
    cards := strings.Split(string(input), "\r\n");
    total := 0;
    matchingns := make([]int, 0, len(cards));
    for _,e := range cards {
        if(len(e) <=0){break;}
        numbers := strings.Split(strings.Split(e, ": ")[1], " | ")
        winnings := strings.Split(numbers[0], " ");
        matchers := strings.Split(numbers[1], " ")
        matchingn := 0;
        for _,n:= range winnings {
            for _,m := range matchers {
                if n!="" && n == m {
                    matchingn ++;
                }
            }
        }
        matchingns = append(matchingns, matchingn);
    }
    amounts := make([]int, len(matchingns), len(matchingns));
    for i,e := range matchingns {
        amounts[i] += 1;
        fmt.Println(i, amounts[i]);
        for j := 0; j < e; j++ {
            index := i+j+1
            if index < len(amounts) {
                amounts[index]+=amounts[i];
            }
        }
        total += amounts[i];
    }
    fmt.Println(total);
}

type Seed struct {
    val int;
    next *Seed;
}

func getNumOrThrow(x string) int{
    i, err := strconv.Atoi(x);
    if err != nil {
        panic(err)
    }
    return i;
}

func day5() {
    input,_ := os.ReadFile("tests/5.txt");
    lines := strings.Split(string(input), "\r\n");
    seeds := make(map[string][]int);
    mapStart := 0;
    things := []string{"seed","soil", "fertilizer", "water", "light", "temperature", "humidity", "location"}
    for i,e := range lines {
        if i == 1 {
            // just seeds 
            s := strings.Split(strings.Split(e, "seeds: ")[1], " ");
            seeds["seed"] = make([]int, len(s), len(s));
            for _, n := range s {
                val,_ := strconv.Atoi(n);
                seeds["seed"] = append(seeds["seed"], val)
            }
            continue;
        }
        if e == ""{
            // end of a mapping 
            mapStart++;
            continue;
        }
        vals := strings.Split(e, " ");
        if len(vals) != 3 {
            seeds[things[mapStart]] = make([]int, len(seeds["seed"]));
            continue;
        }
        valn := make([]int, 3);
        for in,n := range vals {
            valn[in] = getNumOrThrow(n);
        }
        for i, x := range seeds[things[mapStart -1]]{
            if valn[0] < x && valn[3] > x {
                seeds[things[mapStart]][i] = valn[1] + x - valn[0]
            }
        }
    }
    fmt.Println(seeds);
}
