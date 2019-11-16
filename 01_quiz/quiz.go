package main

import (
    "fmt"
    "flag"
    "log"
    "math/rand"
    "time"
    "os"
    "bufio"
    "strings"
    "encoding/csv"
)


func load(path string) [][]string {
    csvFile, _ := os.Open(path) 
    r := csv.NewReader(bufio.NewReader(csvFile))
    records, err := r.ReadAll()
    if(err != nil){
        log.Fatal(err)
    }
    return records
}


func gameLoop(question string, correctAnswer string) bool {
   
   fmt.Printf("What is %v:", question)
   reader := bufio.NewReader(os.Stdin)
   text, err := reader.ReadString('\n')
   if (err != nil){
       return false
   }
   
   text = strings.TrimSpace(text)
   identical := text == correctAnswer
   return identical
}


func main() {
    shuffleFlagPtr := flag.Bool("shuffle", false, "Provide this flag to shuffle the questions")
    flag.Parse()


    // 1) load csv
    records := load("problems.csv")

    if (*shuffleFlagPtr){
        rand.Seed(time.Now().UnixNano())
        rand.Shuffle(len(records), func(i, j int) { records[i], records[j] = records[j], records[i] })
    }

    fmt.Printf("This is a funky quiz game\n")
    var correctCnt int = 0
    // 2) game loop
    for _, rec := range records {
        question := rec[0]
        answer := rec[1]
        
        if (gameLoop(question, answer)){
            correctCnt++
        }

    }
    // 3) report
    fmt.Printf("You got %v out of %v answers correct\n", correctCnt, len(records))
}
