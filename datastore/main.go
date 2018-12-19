package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"cloud.google.com/go/datastore"
	"google.golang.org/api/iterator"
)

type Task struct {
	Id          string
	Description string
}

func main() {
	http.HandleFunc("/", indexHandler)
	// [START setting_port]
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}
	log.Printf("Listening on port %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	switch {
	case r.Method == http.MethodGet:
		tasks := getClient()
		fmt.Fprint(w, tasks)
	case r.Method == http.MethodPut:
		v := r.URL.Query().Get("value")
		putClient(v)
		fmt.Fprint(w, "Success")
	default:
		fmt.Fprint(w, "Success")
	}
}

func getClient() []Task {
	ctx := context.Background()
	projectID := ""
	client, err := datastore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	var tasks []Task
	query := datastore.NewQuery("Task")
	it := client.Run(ctx, query)
	for {
		var task Task
		_, err := it.Next(&task)
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Error %v", err)
		}
		tasks = append(tasks, task)
	}
	return tasks
}

func putClient(v string) {
	ctx := context.Background()
	projectID := ""
	client, err := datastore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	kind := "Task"
	taskKey := datastore.IncompleteKey(kind, nil)
	task := Task{
		Description: v,
	}
	if _, err := client.Put(ctx, taskKey, &task); err != nil {
		log.Fatalf("Failed to save task: %v", err)
	}
	fmt.Printf("Saved %v: %v\n", taskKey, task.Description)
}
