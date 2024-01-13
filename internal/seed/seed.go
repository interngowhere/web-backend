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

	u1, err 	 		:=	CreateUser(ctx, db, SeedUser1); handleError(err)
	u2, err 	 		:=	CreateUser(ctx, db, SeedUser2); handleError(err)
	topic1, err  		:=	CreateTopic(ctx, db, SeedTopic1, u1); handleError(err)
	topic2, err  		:=	CreateTopic(ctx, db, SeedTopic2, u1); handleError(err)
	topic3, err  		:=	CreateTopic(ctx, db, SeedTopic3, u1); handleError(err)
	thread1, err 		:=	CreateThread(ctx, db, SeedThread1, topic1, u1); handleError(err)
	thread2, err 		:=	CreateThread(ctx, db, SeedThread2, topic2, u1); handleError(err)
	thread3, err 		:=	CreateThread(ctx, db, SeedThread3, topic2, u1); handleError(err)
	thread4, err 		:=	CreateThread(ctx, db, SeedThread4, topic3, u1); handleError(err)
	c1, err 	 		:=	CreateComment(ctx, db, SeedComment1, thread1, u2, 0); handleError(err)
	c2, err 	 		:=	CreateComment(ctx, db, SeedComment2, thread1, u1, c1.ID); handleError(err)
	c3, err 	 		:=	CreateComment(ctx, db, SeedComment3, thread1, u2, c1.ID); handleError(err)
	tag1, err 	 		:=	CreateTag(ctx, db, SeedTag1, thread1); handleError(err)
	tag2, err 	 		:=	CreateTag(ctx, db, SeedTag2, thread1); handleError(err)
	tag3, err 	 		:=	CreateTag(ctx, db, SeedTag3, thread1); handleError(err)
	threadKudo1, err 	:= 	CreateThreadKudo(ctx, db, u2, thread1); handleError(err)
	commentKudo1, err 	:= 	CreateCommentKudo(ctx, db, u1, c1); handleError(err)

	_,_,_,_,_,_,_,_,_,_ = thread2, thread3, thread4, c2, c3, tag1, tag2, tag3, threadKudo1, commentKudo1
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
		SetSlug(topic.Slug).
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
		SetSlug(thread.Slug).
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

func CreateThreadKudo(
	ctx context.Context,
	client *ent.Client,
	user *ent.User,
	thread *ent.Thread,
) (*ent.ThreadKudo, error) {
	t, err := client.ThreadKudo.
		Create().
		SetThread(thread).
		SetUser(user).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to add kudo by %v to thread: %w", t.UserID, err)
	}
	log.Printf("Added kudo by %v to thread: %v", t.UserID, t.ThreadID)
	return t, nil
}

func CreateCommentKudo(
	ctx context.Context,
	client *ent.Client,
	user *ent.User,
	comment *ent.Comment,
) (*ent.CommentKudo, error) {
	c, err := client.CommentKudo.
		Create().
		SetComment(comment).
		SetUser(user).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to add kudo by %v to comment: %w", c.UserID, err)
	}
	log.Printf("Added kudo by %v to comment: %v", c.UserID, c.CommentID)
	return c, nil
}