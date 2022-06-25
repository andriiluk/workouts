package internal

const (
	ExerciseSvcTagName = "exercise_svc"
	MuscleSvcTagName   = "muscle_svc"
	WorkoutSvcTagName  = "workout_svc"
	RoutineSvcTagName  = "routine_svc"
)

type Tag struct {
	Name        string `json:"name,omitempty"`
	ServiceName string `json:"service_name,omitempty"`
}

type Params struct {
	Tags []string
	Name string
	ID   int
}

type Muscle struct {
	ID          int      `json:"id,omitempty"`
	Name        string   `json:"name,omitempty"`
	Description string   `json:"description,omitempty"`
	Tags        []string `json:"tags,omitempty"`
}

type Exercise struct {
	Name        string    `json:"name,omitempty"`
	Description string    `json:"description,omitempty"`
	Tags        []*Tag    `json:"tags,omitempty"`
	Muscles     []*Muscle `json:"muscles,omitempty"`
	ID          int       `json:"id,omitempty"`
}

type Workout struct {
	Name        string      `json:"name,omitempty"`
	Description string      `json:"description,omitempty"`
	Tags        []*Tag      `json:"tags,omitempty"`
	Exercise    []*Exercise `json:"exercise,omitempty"`
	ID          int         `json:"id,omitempty"`
}

type Routine struct {
	Name        string     `json:"name,omitempty"`
	Description string     `json:"description,omitempty"`
	Tags        []*Tag     `json:"tags,omitempty"`
	Workouts    []*Workout `json:"workouts,omitempty"`
	ID          int        `json:"id,omitempty"`
}
