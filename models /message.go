// models/message.go
package models

type Message struct {
    ID      string `json:"id"`
    Text    string `json:"text"`
    UserID  string `json:"user_id"`
    Created string `json:"created"`
}