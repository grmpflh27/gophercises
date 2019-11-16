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
    csvFile, err := os.Open(path)
    if(err != nil){
        log.Fatal(err)
    } 
    r := csv.NewReader(csvFile)
    records, err := r.ReadAll()
    if(err != nil){
        log.Fatal(err)
    }
    return records
}


func gameLoop(question string, correctAnswer string) bool {
   
   fmt.Printf("What is %v:", question)
   reader := bufio.NewReader(os.Stdin)
   userAnswer, err := reader.ReadString('\n')
   if (err != nil){
       return false
   }
   
   userAnswer = strings.TrimSpace(userAnswer)
   return userAnswer == correctAnswer
}


func reportResults(correctCnt int, totalRecords int){
    // 3) report
    fmt.Printf("You answered %v out of %v questions correct\n", correctCnt, totalRecords)
}

func main() {
    shuffleFlagPtr := flag.Bool("shuffle", false, "Provide this flag to shuffle the questions")
    timerDurationPtr := flag.Int("duration", 30, "Set duration of quiz game")
    questionCsvPtr := flag.String("QA", "problems.csv", "Relative or absolute path to csv containing questions and answers in q1,a1\nq2,a2 format")
    flag.Parse()

    // 1) load csv and set cli args
    records := load(*questionCsvPtr)

    if (*shuffleFlagPtr){
        rand.Seed(time.Now().UnixNano())
        rand.Shuffle(len(records), func(i, j int) { records[i], records[j] = records[j], records[i] })
    }

    fmt.Printf("This is a funky quiz game\nDuration: %v seconds\nPress <Enter> to start", *timerDurationPtr)

    // start via confirm
    reader := bufio.NewReader(os.Stdin)
    reader.ReadString('\n')
    
    var correctCnt int = 0

    // start timer
    quizTimer := time.NewTimer(time.Duration(*timerDurationPtr) * time.Second)
    go func() {
        <-quizTimer.C
        fmt.Println("\n ... You are out of time ....")
        reportResults(correctCnt, len(records))
        os.Exit(0)
    }()


    // 2) game loop
    for _, rec := range records {
        question, answer := rec[0], rec[1]
        
        if (gameLoop(question, answer)){
            correctCnt++
        }

    }

    reportResults(correctCnt, len(records))
}
