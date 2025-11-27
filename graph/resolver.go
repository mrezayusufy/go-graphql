package graph

import (
	"sync"
	"time"

	"github.com/mrezayusufy/go-graphql/graph/model"
)
type Resolver struct{
    users []*model.User
    todos []*model.Todo
    userID int
    todoID int
    mu sync.Mutex
}

func NewResolver() *Resolver {
    r := &Resolver{
        users: make([]*model.User, 0),
        todos: make([]*model.Todo, 0),
        userID: 1,
        todoID: 1,
    }
    
    user1 := &model.User{
        ID: "1",
        Name: "John Doe",
        Email: "john@gmail.com",
        CreatedAt: time.Now(),
    }
    user2 := &model.User{
        ID: "2",
        Name: "Jane Smith",
        Email: "jane@gmail.com",
        CreatedAt: time.Now(),
    }
    
    r.users = append(r.users, user1, user2)
    
    todo1 := &model.Todo{
        ID: "1",
        Title: "Learn Graphql",
        Completed: false,
        UserID: "1",
        User: user1,
        CreatedAt: time.Now(),
    }
    
    todo2 := &model.Todo{
        ID: "2",
        Title: "Build a todo App",
        Completed: true,
        UserID: "1",
        User: user1,
        CreatedAt: time.Now(),
    }
    
    todo3 := &model.Todo{
        ID: "3",
        Title: "Go Shopping",
        Completed: false,
        UserID: "2",
        User: user2,
        CreatedAt: time.Now(),
    }
    
    r.todos = append(r.todos, todo1, todo2, todo3)
    return r
}
