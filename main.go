package main

import (
    "bufio"
    "flag"
    "fmt"
    "io"
    "log"
    "net/http"
    "os"
    "strings"
    "sync"
    "sync/atomic"
    "time"
)

const banner = `
  ______   _       _              _ _  _____          
 |  ____| | |     | |       /\   | | |/ ____|         
 | |__ ___| |_ ___| |__    /  \  | | | (___  _ __ ___ 
 |  __/ _ \ __/ __| '_ \  / /\ \ | | |\___ \| '__/ __|
 | | |  __/ || (__| | | |/ ____ \| | |____) | | | (__ 
 |_|  \___|\__\___|_| |_/_/    \_\_|_|_____/|_|  \___|
                                               v1.3
             Abid Ahmad [ 0xAb1d ]
`

func downloadFile(url, outputDir string, wg *sync.WaitGroup, notFound *os.File, logger *log.Logger, counter, total *uint64) {
    defer wg.Done()

    safeName := strings.NewReplacer("http://", "", "https://", "", "/", "_", ":", "").Replace(url)
    outputPath := outputDir + "/" + safeName

    resp, err := http.Get(url)
    if err != nil || resp.StatusCode != http.StatusOK {
        logger.Printf("Failed to download %s: %v\n", url, err)
        fmt.Fprintln(notFound, url) // Record the failed URL
    } else {
        outFile, err := os.Create(outputPath)
        if err != nil {
            logger.Printf("Failed to create file for %s: %v\n", url, err)
        } else {
            defer outFile.Close()
            _, err = io.Copy(outFile, resp.Body)
            if err != nil {
                logger.Printf("Failed to write data for %s: %v\n", url, err)
            }
        }
    }
    processed := atomic.AddUint64(counter, 1)
    percentage := float64(processed) / float64(*total) * 100
    fmt.Printf("\r[PROGRESS] - %.2f%% complete (%d/%d)", percentage, processed, *total)
}

func main() {
    fmt.Println(banner)

    warnings := []string{
        "[WARN] - Use responsibly. Your actions are your own.",
        "[WARN] - The developer disclaims any liability for misuse or damage.",
    }

    for _, w := range warnings {
        fmt.Println(w)
    }

    inputFile := flag.String("i", "", "Input file containing URLs")
    outputDir := flag.String("o", ".", "Output directory")
    help := flag.Bool("h", false, "Display help")
    flag.Parse()

    if *help || len(os.Args) == 1 {
        fmt.Println("Usage: ./fas -i input.txt -o outputDir")
        fmt.Println("For more options, use -h or --help")
        return
    }

    if *inputFile == "" {
        log.Fatal("[ERROR] - Input file is required.")
    }

    startTime := time.Now()
    fmt.Println("[INFO] - Processing started")

    setupOutputDirectory(outputDir)
    logger, file := setupLogger(*outputDir)
    defer file.Close()

    notFound := setupFile(*outputDir, "NotFound.txt", logger)
    defer notFound.Close()

    fmt.Printf("[IN] - Input file: %s\n", *inputFile)
    fmt.Printf("[INFO] - Output directory: %s\n\n", *outputDir)

    total := countLines(*inputFile, logger)
    processURLs(inputFile, outputDir, notFound, logger, &total)

    fmt.Printf("\n[OUT] - All files saved in %s\n", *outputDir)
    printElapsedTime(startTime)
}

func setupLogger(outputDir string) (*log.Logger, *os.File) {
    logFilePath := outputDir + "/fetchallsrc.log"
    file, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
    if err != nil {
        log.Fatalf("Error opening log file: %v", err)
    }

    return log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile), file
}

func setupOutputDirectory(outputDir *string) {
    if _, err := os.Stat(*outputDir); os.IsNotExist(err) {
        err := os.MkdirAll(*outputDir, os.ModePerm)
        if err != nil {
            log.Fatalf("Error creating output directory: %v", err)
        }
    }
}

func setupFile(outputDir, fileName string, logger *log.Logger) *os.File {
    filePath := outputDir + "/" + fileName
    file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
    if err != nil {
        logger.Fatalf("Error opening or creating %s: %v", fileName, err)
    }
    return file
}

func countLines(inputFile string, logger *log.Logger) uint64 {
    file, err := os.Open(inputFile)
    if err != nil {
        logger.Fatalf("Error opening input file for counting: %v", err)
    }
    defer file.Close()

    var count uint64
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        count++
    }
    if err := scanner.Err(); err != nil {
        logger.Fatalf("Error reading input file for counting: %v", err)
    }
    return count
}

func processURLs(inputFile *string, outputDir *string, notFound *os.File, logger *log.Logger, total *uint64) {
    file, err := os.Open(*inputFile)
    if err != nil {
        logger.Fatalf("Error opening input file: %v", err)
    }
    defer file.Close()

    var wg sync.WaitGroup
    var counter uint64

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        wg.Add(1)
        go downloadFile(scanner.Text(), *outputDir, &wg, notFound, logger, &counter, total)
    }

    if err := scanner.Err(); err != nil {
        logger.Printf("Error reading input file: %v", err)
    }

    wg.Wait()
}

func printElapsedTime(startTime time.Time) {
    elapsed := time.Since(startTime)
    fmt.Printf("\n[INFO] - Completed in %s\n", elapsed.Round(time.Second))
}

