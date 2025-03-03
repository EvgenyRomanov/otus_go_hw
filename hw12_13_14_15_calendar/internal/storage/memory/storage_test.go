package memorystorage

import (
	"context"
	"math/rand"
	"sync"
	"testing"
	"time"

	"github.com/EvgenyRomanov/otus_go_hw/hw12_13_14_15_calendar/internal/storage"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func randomTimeGenerator() time.Time {
	// https://stackoverflow.com/questions/40944233/generating-random-timestamps-in-go
	return time.Unix(rand.Int63n(time.Now().Unix()-94608000)+94608000, 0)
}

func TestCreateAndGetAndUpdateEvent(t *testing.T) {
	st := New()

	event := &storage.Event{
		ID:    uuid.New(),
		Title: "Event title",
	}

	err := st.CreateEvent(context.Background(), event)
	assert.NoError(t, err)

	// error if already exists
	err = st.CreateEvent(context.Background(), event)
	assert.Equal(t, storage.ErrEventAlreadyExists, err)

	newTitle := "Event after update"
	eventForUpdate := &storage.Event{
		Title:    newTitle,
		DateTime: randomTimeGenerator(),
	}
	err = st.UpdateEvent(context.Background(), event.ID, eventForUpdate)
	assert.NoError(t, err)

	// check after update
	updatedEvent, err := st.GetEvent(context.Background(), event.ID)
	assert.NoError(t, err)
	assert.Equal(t, newTitle, updatedEvent.Title)

	// update event that doesn't exist
	err = st.UpdateEvent(context.Background(), uuid.New(), &storage.Event{Title: "1"})
	assert.Equal(t, storage.ErrEventNotFound, err)

	// get event that doesn't exist
	_, err = st.GetEvent(context.Background(), uuid.New())
	assert.Equal(t, storage.ErrEventNotFound, err)

	// get events for notification
	st = New()
	eventUUID := uuid.New()

	_ = st.CreateEvent(context.Background(), &storage.Event{
		ID:               eventUUID,
		Title:            "Event title",
		DateTime:         time.Now(),
		TimeNotification: time.Now().Add(time.Millisecond * 5),
	})

	// event that already has been notify
	_ = st.CreateEvent(context.Background(), &storage.Event{
		ID:               uuid.New(),
		Title:            "Event title",
		DateTime:         time.Now(),
		TimeNotification: time.Now().Add(time.Millisecond * 4),
		NotifyAt:         time.Now(),
	})

	_ = st.CreateEvent(context.Background(), &storage.Event{
		ID:               uuid.New(),
		Title:            "Event title",
		DateTime:         time.Now(),
		TimeNotification: time.Now().Add(time.Millisecond * 1000),
	})

	event, _ = st.GetEvent(context.Background(), eventUUID)

	time.Sleep(time.Millisecond * 10)
	events, err := st.GetEventsForNotifications(context.Background())
	assert.NoError(t, err)
	assert.Len(t, events, 1)
	assert.Equal(t, event, events[0])
}

func TestUpdateWithBusyTimeEvent(t *testing.T) {
	time := randomTimeGenerator()
	id := uuid.New()

	st := New()

	event := &storage.Event{
		ID:       id,
		Title:    "Event title",
		DateTime: time,
	}

	err := st.CreateEvent(context.Background(), event)
	assert.NoError(t, err)

	err = st.UpdateEvent(context.Background(), id, &storage.Event{Title: "1", DateTime: time})
	assert.Equal(t, storage.ErrEventDateTimeIsBusy, err)
}

func TestDeleteEvent(t *testing.T) {
	st := New()
	event := &storage.Event{
		ID:       uuid.New(),
		Title:    "Event Title",
		DateTime: time.Now(),
	}

	// Create an event
	err := st.CreateEvent(context.Background(), event)
	assert.NoError(t, err)

	// Delete the event
	err = st.DeleteEvent(context.Background(), event.ID)
	assert.NoError(t, err)

	// Try deleting a non-existing event
	err = st.DeleteEvent(context.Background(), uuid.New())
	assert.Equal(t, storage.ErrEventNotFound, err)
}

func TestConcurrent(t *testing.T) {
	st := New()
	UUID := uuid.New()

	event := &storage.Event{
		ID:    UUID,
		Title: "Event Title",
	}
	err := st.CreateEvent(context.Background(), event)
	assert.NoError(t, err)

	var wg sync.WaitGroup

	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			event := &storage.Event{
				ID:       UUID,
				Title:    uuid.New().String(),
				DateTime: randomTimeGenerator(),
			}
			err := st.UpdateEvent(context.Background(), UUID, event)
			assert.NoError(t, err)
		}()
	}

	wg.Wait()

	updatedEvent, err := st.GetEvent(context.Background(), UUID)
	assert.NoError(t, err)
	assert.NotNil(t, updatedEvent)
	assert.NotContains(t, updatedEvent.Title, "Event Title")

	errCh := make(chan error, 50)
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			err := st.DeleteEvent(context.Background(), UUID)
			if err != nil {
				errCh <- err
			}
		}()
	}

	wg.Wait()
	close(errCh)

	// count how many errors do we have. 49 -- because we delete exactly one
	assert.Equal(t, 49, len(errCh))
}
