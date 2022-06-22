package musclesvc

import (
	"context"

	"github.com/andriiluk/workouts/internal"
	log "github.com/sirupsen/logrus"
)

type errorHandle func(ctx context.Context, err error)

func (e errorHandle) Handle(ctx context.Context, err error) {
	e(ctx, err)
}

var ErrorLogHandler errorHandle = func(ctx context.Context, err error) {
	log.WithError(err).Warn("Error handler")
}

type LogginMiddleware struct {
	next Service
}

func WithLoggingMidleware(s Service) *LogginMiddleware {
	return &LogginMiddleware{s}
}

func (l *LogginMiddleware) AddMuscle(ctx context.Context, m *internal.Muscle) (id int, err error) {
	log.WithFields(log.Fields{"muscle": m}).Debug("add muscle request")
	defer func() {
		log.WithFields(log.Fields{"id": id, "err": err}).Debug("add muscle response")
	}()

	id, err = l.next.AddMuscle(ctx, m)
	return
}

func (l *LogginMiddleware) PutMuscle(ctx context.Context, m *internal.Muscle) (err error) {
	log.WithFields(log.Fields{"muscle": m}).Debug("put muscle request")
	defer func() {
		log.WithFields(log.Fields{"muscle": m, "err": err}).Debug("put muscle response")
	}()

	err = l.next.PutMuscle(ctx, m)
	return
}

func (l *LogginMiddleware) DeleteMuscleByName(ctx context.Context, name string) (err error) {
	log.Debug("DeleteMuscleByName", name)
	defer func() {
		log.WithField("error", err).Debug("delete muscle by name")
	}()

	err = l.next.DeleteMuscleByName(ctx, name)
	return
}

func (l *LogginMiddleware) DeleteMuscleByID(ctx context.Context, id int) (err error) {
	log.Debug("DeleteMuscleByID", id)
	defer func() {
		log.WithField("err", err).Debug("delete muscle by id")
	}()

	err = l.next.DeleteMuscleByID(ctx, id)
	return
}

func (l *LogginMiddleware) GetMuscleByID(ctx context.Context, id int) (m *internal.Muscle, err error) {
	log.WithFields(log.Fields{"id": id}).Debug("get muscle by id")
	defer func() {
		log.WithFields(log.Fields{"muscle": m, "err": err}).Debug("get muscle by muscle response")
	}()

	m, err = l.next.GetMuscleByID(ctx, id)
	return
}

func (l *LogginMiddleware) GetMuscleByName(ctx context.Context, name string) (m *internal.Muscle, err error) {
	log.Debug("GetMuscleByName", name)
	defer func() {
		log.WithFields(log.Fields{
			"muscle": m,
			"err":    err,
		}).Debug("get muscle by name response")
	}()

	m, err = l.next.GetMuscleByName(ctx, name)
	return
}

func (l *LogginMiddleware) GetMusclesByTags(ctx context.Context, tags ...string) (data []*internal.Muscle, err error) {
	log.Debug("GetMuscleByTags", tags)
	defer func() {
		log.WithFields(log.Fields{"data": data, "err": err}).Debug("get muscle by tags response")
	}()

	data, err = l.next.GetMusclesByTags(ctx, tags...)
	return
}
