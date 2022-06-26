package exercisesvc

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

func (l *LogginMiddleware) AddExercise(ctx context.Context, m *internal.Exercise) (id int, err error) {
	log.WithFields(log.Fields{"exercise": m}).Debug("add exercise request")

	defer func() {
		log.WithFields(log.Fields{"id": id, "err": err}).Debug("add Exercise response")
	}()

	id, err = l.next.AddExercise(ctx, m)

	return
}

func (l *LogginMiddleware) PutExercise(ctx context.Context, m *internal.Exercise) (err error) {
	log.WithFields(log.Fields{"exercise": m}).Debug("put exercise request")

	defer func() {
		log.WithFields(log.Fields{"exercise": m, "err": err}).Debug("put exercise response")
	}()

	err = l.next.PutExercise(ctx, m)

	return
}

func (l *LogginMiddleware) DeleteExerciseByName(ctx context.Context, name string) (err error) {
	log.Debug("DeleteExerciseByName", name)

	defer func() {
		log.WithField("error", err).Debug("delete exercise by name")
	}()

	err = l.next.DeleteExerciseByName(ctx, name)

	return
}

func (l *LogginMiddleware) DeleteExerciseByID(ctx context.Context, id int) (err error) {
	log.Debug("DeleteExerciseByID", id)

	defer func() {
		log.WithField("err", err).Debug("delete exercise by id")
	}()

	err = l.next.DeleteExerciseByID(ctx, id)

	return
}

func (l *LogginMiddleware) GetExerciseByID(ctx context.Context, id int) (ex *internal.Exercise, err error) {
	log.WithFields(log.Fields{"id": id}).Debug("get exercise by id")

	defer func() {
		log.WithFields(log.Fields{
			"exercise": ex,
			"err":      err,
		}).Debug("get exercise by response")
	}()

	ex, err = l.next.GetExerciseByID(ctx, id)

	return
}

func (l *LogginMiddleware) GetExerciseByName(ctx context.Context, name string) (m *internal.Exercise, err error) {
	log.Debug("GetExerciseByName", name)

	defer func() {
		log.WithFields(log.Fields{
			"exercise": m,
			"err":      err,
		}).Debug("get Exercise by name response")
	}()

	m, err = l.next.GetExerciseByName(ctx, name)

	return
}

func (l *LogginMiddleware) GetExercisesByTags(ctx context.Context, tags ...string) (data []*internal.Exercise, err error) {
	log.WithField("tags", tags).Debug("GetExerciseByTags")

	defer func() {
		log.WithFields(log.Fields{"data": data, "err": err}).Debug("get exercise by tags response")
	}()

	data, err = l.next.GetExercisesByTags(ctx, tags...)

	return
}

func (l *LogginMiddleware) GetExercisesByMuscles(ctx context.Context, muscleNames ...string) (data []*internal.Exercise, err error) {
	log.WithField("muscles", muscleNames).Debug("GetExercisesByMuscles")

	defer func() {
		log.WithFields(log.Fields{"data": data, "err": err}).Debug("get exercise by muscle response")
	}()

	data, err = l.next.GetExercisesByMuscles(ctx, muscleNames...)

	return
}
