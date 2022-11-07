package main

import (
    "os"
    "fmt"
    "log"
    "bytes"
    "net/http"
    "io/ioutil"
    "github.com/joho/godotenv"
)

func sendMessage(message string) {
    client := http.Client{};
    var jsonStr = []byte("{\"text\":\""+message+"\"}");

    request, err := http.NewRequest(
        "POST",
        os.Getenv("SLACK_URL"),
        bytes.NewBuffer(jsonStr));

    if err != nil {
        //Handle Error
    }

    request.Header = http.Header{
        "Content-Type": {"application/json"},
    };

    response, err := client.Do(request);
    body, err := ioutil.ReadAll(response.Body);

    fmt.Printf(string(body));
}

func main() {
    err := godotenv.Load(".env");

    if err != nil {
        log.Fatalf("Error loading .env file");
    }

    sendMessage("hello from golang");
}