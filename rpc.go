package main

import (
	"net/http"
	"encoding/json"
	"fmt"

	"code.google.com/p/google-api-go-client/taskqueue/v1beta2"
	"code.google.com/p/goauth2/oauth"
)

// EventArgs represents the parameter of the PullTask RPC service method.
type EventArgs struct {
	EventId  string
	PullType string
}

// EventReply represents the reply message of the PullTask RPC service method.
type EventReply struct {
	Message string
}

// EventService represents the RPC service.
type EventService struct{}

// PullTask is the only service method of EventService.
// It pre-checks the RPC parameter, calls PullTasks() and sets the reply message.
// It returns any error encountered. 
func (eventService *EventService) PullTask(r *http.Request, args *EventArgs, reply *EventReply) error {
	// Pre-check the given pull type
	if args.PullType == "picturerequest" {
		PullTasks(r, args.PullType, args.EventId)
	}
	reply.Message = "Ok"
	return nil
}

// PullTasks leases all tasks which affect the same event ID from the Google App Engine Task Queue
// and starts the GalleryAggregator for that event.
func PullTasks(r *http.Request, pullType, eventId string) {
	fmt.Printf("PullTask; PullType: %v, EventId: %v", pullType, eventId)
    
    // Request current API Access Token
    authResp, err := http.Get("http://metadata/computeMetadata/v1beta1/instance/service-accounts/default/token")
    if err != nil {
        fmt.Printf("Error getting authorization information: %v\n", err)
    }
    defer authResp.Body.Close()
    authDecoder := json.NewDecoder(authResp.Body)
    authData := new(authorizationResponse)
    authDecoder.Decode(&authData)

    // Create a new authorized API client
    transport := &oauth.Transport{
        Config: OAuthConfig,
        Token: &oauth.Token{
            AccessToken: authData.Access_token,
        },
    }
    taskapi, err := taskqueue.New(transport.Client())
    if err != nil {
        fmt.Printf("Error generating task queue service: %v\n", err)
    }
       
    // Lease all tasks with the given event ID as tag
	var tasks []*taskqueue.Task
    taskList, err := taskapi.Tasks.Lease("s~gcdc2013-socialvibes", pullType, 100, 3600).GroupByTag(true).Tag(eventId).Do()
	if err != nil {
		fmt.Printf("Error leasing tasks: %v\n", err)
	}
	if len(taskList.Items) > 0 {
		tasks = taskList.Items
	} else {
		// Return if no tasks found
		fmt.Printf("Error: no tasks found\n")
		return
	}

	// We now found all tasks associated to the given event id.
	// Therefore we can now aggregate the pictures, send them to the app engine and distribute to the clients
	go GalleryAggregator(eventId)

	// Delete all found tasks
	for _, tmptask := range tasks {
		err = taskapi.Tasks.Delete("s~gcdc2013-socialvibes", pullType, tmptask.Id).Do()
	}
}