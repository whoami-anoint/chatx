// repositories/user_repository.go
package repositories

import (
    "database/sql"
    "errors"

    "github.com/chatx/models"
)

type UserRepository struct {
    db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
    return &UserRepository{db: db}
}

func (ur *UserRepository) GetUser(id string) (*models.User, error) {
    row := ur.db.QueryRow("SELECT * FROM users WHERE id = $1", id)
    user := &models.User{}
    err := row.Scan(&user.ID, &user.Username, &user.Email)
    if err != nil {
        return nil, errors.New("user not found")
    }
    return user, nil
}

// repositories/message_repository.go
package repositories

import (
    "database/sql"
    "errors"

    "github.com/chatx/models"
)

type MessageRepository struct {
    db *sql.DB
}

func NewMessageRepository(db *sql.DB) *MessageRepository {
    return &MessageRepository{db: db}
}

func (mr *MessageRepository) GetMessage(id string) (*models.Message, error) {
    row := mr.db.QueryRow("SELECT * FROM messages WHERE id = $1", id)
    message := &models.Message{}
    err := row.Scan(&message.ID, &message.Text, &message.UserID, &message.Created)
    if err != nil {
        return nil, errors.New("message not found")
    }
    return message, nil
}
