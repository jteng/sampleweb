package engine

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"time"
)

type ThirdpartyService struct {
	random *rand.Rand
}

//GetData will either return data from the service call or will cancel on timeout
func (s *ThirdpartyService) GetData(id string, ctx context.Context) <-chan string {
	result := make(chan string)
	go func() {
		select {
		case <-ctx.Done():
			log.Printf("cancel.....%s", id)
			return
		case <-time.After(s.randomDuration(id)):
			result <- fmt.Sprintf("got data for id %s", id)
		}

	}()
	return result
}

func (s *ThirdpartyService) randomDuration(id string) (d time.Duration) {
	if id == "777" {
		return time.Second * 3
	}
	return time.Duration(s.random.Intn(100)) * time.Millisecond
}

func NewService() *ThirdpartyService {
	return &ThirdpartyService{
		random: rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}
