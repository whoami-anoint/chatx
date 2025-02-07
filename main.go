package main

import (
    "fmt"
    "log"
    "net/http"

    "github.com/chatx/services"
    "github.com/chatx/utils"
)

func main() {
    // Create a new user service
    userService := services.NewUserService(repositories.NewUserRepository(db))

    // Create a new message service
    messageService := services.NewMessageService(repositories.NewMessageRepository(db))

    http.HandleFunc("/websocket", func(w http.ResponseWriter, r *http.Request) {
        conn, err := websocket.Upgrade(w, r, nil, 1024, 1024)
        if err != nil {
            log.Println(err)
            return
        }
        defer conn.Close()

        for {
            messageType, p, err := conn.ReadMessage()
            if err != nil {
                log.Println(err)
                return
            }

            // Handle incoming message
            fmt.Println(string(p))

            // Create a new message
            message := &models.Message{
                Text:    string(p),
                UserID:  "1",
                Created: time.Now().String(),
            }

            // Save the message to the database
            err = messageService.CreateMessage(message)
            if err != nil {
                log.Println(err)
                return
            }
        }
    })

    log.Fatal(http.ListenAndServe(":8080", nil))
}
