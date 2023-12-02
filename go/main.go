package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var INPUT_DIR = "input"

func main() {
    //day1("tests/1.txt");
    //day1("inputs/1.txt");
    day2()
}

func day2(){
    input,_ := os.ReadFile("inputs/2.txt");
    games := strings.Split(string(input), "\r\n");
    limit := make(map[string]int); 
    limit["red"] = 12;
    limit["blue"] = 13;
    limit["green"] = 14;
    total := 0;
    for i,game:= range games {
        if len(game) == 0 {
            fmt.Println("end")
            break;
        }
        large := false;
        f := strings.Split(game, ": ");
        gameIdstr := strings.Split(f[0], " ")[1];
        sets := strings.Split(f[1], "; ");
        for _,set := range sets {
            balls := strings.Split(set, ", ");
            for _, ball := range balls {
                ballh := strings.Split(ball, " ");
                amount,_ := strconv.Atoi(ballh[0]);
                if amount > limit[ballh[1]] {
                    fmt.Println(amount,ballh, i+1)
                    large = true;
                    break;
                }

            }
            if(large){
                break;
            }
        }
        if large {
            continue;
        }
        gameIdnum,_ := strconv.Atoi(gameIdstr);
        if(gameIdnum != i+1){
            panic("gameIdnum is not what it should be !")
        }
        total += gameIdnum;

    }
    fmt.Println("day2:", total);
}

