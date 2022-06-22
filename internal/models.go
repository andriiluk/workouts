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
	ID   int
	Name string
	Tags []string
}

type Muscle struct {
	ID          int      `json:"id,omitempty"`
	Name        string   `json:"name,omitempty"`
	Description string   `json:"description,omitempty"`
	Tags        []string `json:"tags,omitempty"`
}

type Exercise struct {
	ID          int       `json:"id,omitempty"`
	Name        string    `json:"name,omitempty"`
	Description string    `json:"description,omitempty"`
	Tags        []*Tag    `json:"tags,omitempty"`
	Muscles     []*Muscle `json:"muscles,omitempty"`
}

type Workout struct {
	ID          int         `json:"id,omitempty"`
	Name        string      `json:"name,omitempty"`
	Description string      `json:"description,omitempty"`
	Tags        []*Tag      `json:"tags,omitempty"`
	Exercise    []*Exercise `json:"exercise,omitempty"`
}

type Routine struct {
	ID          int        `json:"id,omitempty"`
	Name        string     `json:"name,omitempty"`
	Description string     `json:"description,omitempty"`
	Tags        []*Tag     `json:"tags,omitempty"`
	Workouts    []*Workout `json:"workouts,omitempty"`
}
