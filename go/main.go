package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var INPUT_DIR = "input"

func main() {
    day1("inputs/1.txt");
    day2();
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

