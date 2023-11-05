package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.40

import (
	"context"
	"example/internal/db"
	"example/internal/graph/model"
	"fmt"
)

// Name is the resolver for the name field.
func (r *chatResolver) Name(ctx context.Context, obj *db.Chat) (*string, error) {
	panic(fmt.Errorf("not implemented: Name - name"))
}

// Users is the resolver for the users field.
func (r *chatResolver) Users(ctx context.Context, obj *db.Chat) ([]*db.User, error) {
	panic(fmt.Errorf("not implemented: Users - users"))
}

// Messages is the resolver for the messages field.
func (r *chatResolver) Messages(ctx context.Context, obj *db.Chat) ([]*db.ChatMessage, error) {
	panic(fmt.Errorf("not implemented: Messages - messages"))
}

// Chat is the resolver for the chat field.
func (r *chatMessageResolver) Chat(ctx context.Context, obj *db.ChatMessage) (*db.Chat, error) {
	panic(fmt.Errorf("not implemented: Chat - chat"))
}

// User is the resolver for the user field.
func (r *chatMessageResolver) User(ctx context.Context, obj *db.ChatMessage) (*db.User, error) {
	panic(fmt.Errorf("not implemented: User - user"))
}

// Chat is the resolver for the chat field.
func (r *chatUserResolver) Chat(ctx context.Context, obj *db.ChatUser) (*db.Chat, error) {
	panic(fmt.Errorf("not implemented: Chat - chat"))
}

// User is the resolver for the user field.
func (r *chatUserResolver) User(ctx context.Context, obj *db.ChatUser) (*db.User, error) {
	panic(fmt.Errorf("not implemented: User - user"))
}

// Chat is the resolver for the chat field.
func (r *queryResolver) Chat(ctx context.Context, id string) (*db.Chat, error) {
	panic(fmt.Errorf("not implemented: Chat - chat"))
}

// Chats is the resolver for the chats field.
func (r *queryResolver) Chats(ctx context.Context, limit *int, offset *int) (*model.ChatConnection, error) {
	panic(fmt.Errorf("not implemented: Chats - chats"))
}

// Chat returns ChatResolver implementation.
func (r *Resolver) Chat() ChatResolver { return &chatResolver{r} }

// ChatMessage returns ChatMessageResolver implementation.
func (r *Resolver) ChatMessage() ChatMessageResolver { return &chatMessageResolver{r} }

// ChatUser returns ChatUserResolver implementation.
func (r *Resolver) ChatUser() ChatUserResolver { return &chatUserResolver{r} }

type chatResolver struct{ *Resolver }
type chatMessageResolver struct{ *Resolver }
type chatUserResolver struct{ *Resolver }
