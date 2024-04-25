package googlecloud

import (
	"context"
	"fmt"
	"log"
	"time"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	gcal "google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"

	"backend/models"
	"backend/utils"
)

type Calendar interface {
	GetEventsForDay(date time.Time) ([]models.Event, error)
}

type calendar struct {
	service *gcal.Service
}

var _ Calendar = &calendar{}

func NewCalendar(ctx context.Context, json []byte) Calendar {
	config, err := google.ConfigFromJSON(json, gcal.CalendarReadonlyScope)
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}
	token := getTokenFromWeb(config)
	client := config.Client(context.Background(), token)

	svc, err := gcal.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		log.Fatalf("Unable to retrieve Calendar client: %v", err)
	}

	return &calendar{service: svc}
}

func (c *calendar) GetEventsForDay(date time.Time) ([]models.Event, error) {
	startTime, endTime := utils.StartOfDay(date), utils.EndOfDay(date)
	startTimeStr, endTimeStr := startTime.Format(time.RFC3339), endTime.Format(time.RFC3339)
	events, err := c.service.Events.List("primary").ShowDeleted(false).
		SingleEvents(true).TimeMin(startTimeStr).TimeMax(endTimeStr).MaxResults(50).OrderBy("startTime").Do()
	if err != nil {
		log.Fatalf("Unable to retrieve next ten of the user's events: %v", err)
	}

	fmt.Println("Upcoming events:")
	if len(events.Items) == 0 {
		return []models.Event{}, fmt.Errorf("no upcoming events found")
	}

	results := models.ToEvents(events.Items)
	return results, nil
}

func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser then type the "+
		"authorization code: \n%v\n", authURL)

	var authCode string
	if _, err := fmt.Scan(&authCode); err != nil {
		log.Fatalf("Unable to read authorization code: %v", err)
	}

	tok, err := config.Exchange(context.TODO(), authCode)
	if err != nil {
		log.Fatalf("Unable to retrieve token from web: %v", err)
	}
	return tok
}
