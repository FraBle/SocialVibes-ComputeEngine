// Package aggregation contains the GalleryAggregator for Google+ event pictures.
package aggregation

import (
    "encoding/base64"
    "encoding/json"
    "fmt"
    "net/http"
    "net/http/httptest"
    "strings"
    timepkg "time"

    "github.com/PuerkitoBio/goquery"
    "github.com/sourcegraph/webloop"
    "code.google.com/p/goauth2/oauth"
    "code.google.com/p/google-api-go-client/taskqueue/v1beta2"

    "socialvibes/model"
    "socialvibes/config"
)

// The GalleryAggregator takes a Google+ event ID, 
// renders the event page and aggregates all event picture URLs.
func GalleryAggregator(eventId string) {
    
    // Create a webloop renderer for the given event
    renderer := &webloop.StaticRenderer{
            TargetBaseURL:         "https://plus.google.com/events/gallery/" + eventId + "?sort=1",
            WaitTimeout:           timepkg.Second * 5,
            ReturnUnfinishedPages: true,
    }

    // Create a http.ResponseWriter, so that GoQuery can work with it
    w := httptest.NewRecorder()
    
    // Create an empty http.Request for the given event
    r, err := http.NewRequest("GET", "", nil)
    if err != nil {
            fmt.Printf("Error creating request: %v\n", err)
    }
    renderer.ServeHTTP(w, r)
    
    // Transform the response into a parsable document
    var document *goquery.Document
    document, err = goquery.NewDocumentFromReader(w.Body)
    if err != nil {
            fmt.Printf("Error creating request: %v\n", err)
    }

    // Parse all event picture URLs from the response document
    var pictures []model.Picture
    document.Find(".Bea.VLb").Each(func(i int, s *goquery.Selection) {
            picUrl, _ := s.Attr("src")
            pictures = append(pictures, model.Picture{picUrl})
    })

    // Request current API Access Token
    authResp, err := http.Get("http://metadata/computeMetadata/v1beta1/instance/service-accounts/default/token")
    if err != nil {
        fmt.Printf("Error getting authorization information: %v\n", err)
    }
    defer authResp.Body.Close()
    authDecoder := json.NewDecoder(authResp.Body)
    authData := new(model.AuthorizationResponse)
    authDecoder.Decode(&authData)
    
    // Create a new authorized API client
    transport := &oauth.Transport{
            Config: config.OAuthConfig,
            Token: &oauth.Token{
                    AccessToken: authData.Access_token,
            },
    }
    taskapi, err := taskqueue.New(transport.Client())
    if err != nil {
        fmt.Printf("Error generating task queue service: %v\n", err)
    }

    // Encode the event pictures object which will be passed to the task as payload
    pictures_json, err := json.Marshal(pictures)
    if err != nil {
            fmt.Printf("Error encoding event to json: %v\n", err)
    }
    pictures_base64 := base64.StdEncoding.EncodeToString(pictures_json)

    // Insert the task
    _, err = taskapi.Tasks.Insert("s~gcdc2013-socialvibes", "picture", &taskqueue.Task{
            PayloadBase64: pictures_base64,
            QueueName:     "picture",
    }).Do()
    if err != nil {
        fmt.Printf("Error inserting task: %v\n", err)
    }

    // Notify the task consumer in the App Engine via RPC
    timepkg.Sleep(3 * timepkg.Second)
    response := `{"method":"EventService.PullTask","params":[{"PullType":"picture", "EventId":"` + eventId + `"}], "id":"1"}`
    _, err = http.Post("http://gcdc2013-socialvibes.appspot.com/rpc", "application/json", strings.NewReader(response))
    if err != nil {
            fmt.Printf("Error notifying the task consumer in the app engine via rpc: %v\n", err)
    }
}