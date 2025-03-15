/*
MQ Endpoint
*/

package logger_mq

import (
	logger_application "github.com/shortlink-org/shortlink/boundaries/platform/logger/internal/application"
	"github.com/shortlink-org/shortlink/pkg/logger"
	"github.com/shortlink-org/shortlink/pkg/mq"
	"github.com/shortlink-org/shortlink/pkg/mq/query"
)

type Event struct {
	mq  mq.MQ
	log logger.Logger

	service *logger_application.Service
}

func New(dataBus mq.MQ, log logger.Logger, service *logger_application.Service) (*Event, error) {
	if dataBus == nil {
		return nil, ErrMQIsNil
	}

	event := &Event{
		mq:  dataBus,
		log: log,

		service: service,
	}

	// Subscribe
	event.Subscribe()

	return event, nil
}

func (e *Event) Subscribe() {
	getEventNewLink := query.Response{
		Chan: make(chan query.ResponseMessage),
	}

	go func() {
		// if err := e.mq.Subscribe(context.Background(), v1.MQ_EVENT_LINK_CREATED, getEventNewLink); err != nil {
		// 	e.log.Error(err.Error())
		// }
	}()

	go func() {
		for {
			msg := <-getEventNewLink.Chan

			// Convert: []byte to link.Link
			// myLink := &v1.Link{}
			// if err := proto.Unmarshal(msg.Body, myLink); err != nil {
			// 	e.log.Error(fmt.Sprintf("Error unmarsharing event new link: %s", err.Error()))
			// 	msg.Context.Done()
			//
			// 	continue
			// }
			//
			// e.service.Log(msg.Context, myLink)
			msg.Context.Done()
		}
	}()
}
