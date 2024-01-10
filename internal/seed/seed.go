package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/interngowhere/web-backend/ent"
	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

func main() {
	// load .env file
	err := godotenv.Load("../../.env")

	if err != nil {
	  log.Fatalf("Failed to load .env file: %v", err)
	}
	
	// Setup database
	db, err := ent.Open(
		os.Getenv("DB_DIALECT"), 
		fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", 
			os.Getenv("DB_HOST"), 
			os.Getenv("DB_PORT"),
			os.Getenv("DB_USER"),
			os.Getenv("DB_NAME"),
			os.Getenv("DB_PASSWORD")))

	if err != nil {
        log.Fatalf("Failed to establish connection to database: %v", err)
    }

	log.Println("Established connection to database")

	defer db.Close()
	defer log.Println("Database connection closed")

	ctx := context.Background()

	// Run the auto migration tool.
    if err := db.Schema.Create(ctx); err != nil {
        log.Fatalf("Failed to create schema resources: %v", err)
    }

	log.Println("Created schema resources")

	// Clear all tables
	_, err = db.CommentKudo.Delete().Exec(ctx); handleError(err)
	_, err = db.Comment.Delete().Exec(ctx); handleError(err)
	_, err = db.Tag.Delete().Exec(ctx); handleError(err)
	_, err = db.ThreadKudo.Delete().Exec(ctx); handleError(err)
	_, err = db.Thread.Delete().Exec(ctx); handleError(err)
	_, err = db.Moderator.Delete().Exec(ctx); handleError(err)
	_, err = db.Topic.Delete().Exec(ctx); handleError(err)
	_, err = db.User.Delete().Exec(ctx); handleError(err)

	log.Println("Cleared all tables")

	u1, err 	:=	CreateUser(ctx, db, SeedUser1); handleError(err)
	u2, err 	:=	CreateUser(ctx, db, SeedUser2); handleError(err)
	topic, err  :=	CreateTopic(ctx, db, SeedTopic, u1); handleError(err)
	thread, err :=	CreateThread(ctx, db, SeedThread, topic, u1); handleError(err)
	c1, err 	:=	CreateComment(ctx, db, SeedComment1, thread, u2, 0); handleError(err)
	c2, err 	:=	CreateComment(ctx, db, SeedComment2, thread, u1, c1.ID); handleError(err)
	c3, err 	:=	CreateComment(ctx, db, SeedComment3, thread, u2, c1.ID); handleError(err)
	tag1, err 	:=	CreateTag(ctx, db, SeedTag1, thread); handleError(err)
	tag2, err 	:=	CreateTag(ctx, db, SeedTag2, thread); handleError(err)
	tag3, err 	:=	CreateTag(ctx, db, SeedTag3, thread); handleError(err)
	_,_,_,_,_ = c2, c3, tag1, tag2, tag3
}

func handleError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func CreateUser(ctx context.Context, client *ent.Client, user ent.User) (*ent.User, error) {
    u, err := client.User.
        Create().
		SetEmail(user.Email).
		SetFirstName(user.FirstName).
		SetLastName(user.LastName).
		SetUsername(user.Username).
        Save(ctx)
    if err != nil {
        return nil, fmt.Errorf("failed to create user: %w", err)
    }
    log.Printf("Created user: %v", u.Username)
    return u, nil
}

func CreateTopic(ctx context.Context, client *ent.Client, topic ent.Topic, user *ent.User) (*ent.Topic, error) {
	t, err := client.Topic.
        Create().
		SetTitle(topic.Title).
		SetDescription(topic.Description).
		SetShortDescription(topic.ShortDescription).
		AddTopicModerators(user).
        Save(ctx)
    if err != nil {
        return nil, fmt.Errorf("failed to create topic: %w", err)
    }
    log.Printf("Created topic: %v", t.Title)
    return t, nil
}

func CreateThread(
		ctx context.Context,
		client *ent.Client,
		thread ent.Thread,
		topic *ent.Topic,
		user *ent.User,
	) (*ent.Thread, error) {
	t, err := client.Thread.
        Create().
		SetTitle(thread.Title).
		SetDescription(thread.Description).
		SetTopics(topic).
		SetUsers(user).
		AddKudoedUsers(user).
        Save(ctx)
    if err != nil {
        return nil, fmt.Errorf("failed to create thread: %w", err)
    }
    log.Printf("Created thread: %v", t.Title)
    return t, nil
}

func CreateComment(
	ctx context.Context,
	client *ent.Client,
	comment ent.Comment,
	thread *ent.Thread,
	user *ent.User,
	parentID int,
) (*ent.Comment, error) {
	c, err := client.Comment.
		Create().
		SetParentID(parentID).
		SetContent(comment.Content).
		SetThreads(thread).
		SetCommentAuthors(user).
		AddKudoedUsers(user).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to create comment: %w", err)
	}
	log.Printf("Created comment: %v", c.Content)
	return c, nil
}

func CreateTag(
	ctx context.Context,
	client *ent.Client,
	tag ent.Tag,
	thread *ent.Thread,
) (*ent.Tag, error) {
	t, err := client.Tag.
		Create().
		SetTagName(tag.TagName).
		AddTaggedThreads(thread).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to create tag: %w", err)
	}
	log.Printf("Created tag: %v", t.TagName)
	return t, nil
}