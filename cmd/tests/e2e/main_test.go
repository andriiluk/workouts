package main

import (
	"testing"

	"github.com/andriiluk/workouts/internal"
	"github.com/andriiluk/workouts/internal/musclesvc/client"

	"github.com/stretchr/testify/require"
)

func TestMuscleSvc(t *testing.T) {
	const addr = "localhost:8080"

	testMusclesEndpoints(t, addr)
}

func testMusclesEndpoints(t *testing.T, addr string) {
	t.Log("create new client")
	cli, err := client.NewHTTPClient(addr)
	require.NoError(t, err)

	var (
		muscleName = "test"
		muscleDesc = "test description"
		tags       = []string{"test", "tag"}
	)

	// Add muscle request
	t.Log("do post request")
	var id int
	id, err = cli.PostMuscle(muscleName, muscleDesc, tags...)
	require.NoError(t, err, "post muscle")
	require.Equal(t, 1, id, "muscle id")

	// Get muscle request with invalid id
	var m *internal.Muscle
	m, err = cli.GetMuscle(2)
	require.NoError(t, err, "get muscle")
	require.Empty(t, m, "muscle")

	// Get muscle request with valid id
	m, err = cli.GetMuscle(id)
	require.NoError(t, err, "get muscle")
	require.NotNil(t, m, "muscle")
	require.Equal(t, m.ID, id, "id")
	require.Equal(t, m.Name, muscleName, "name")
	require.Equal(t, m.Description, muscleDesc, "description")

	// Put muscle
	var (
		newMuscleName = "New name"
		newMuscleDesc = "New description"
		newMuscleTags = []string{"new", "tag"}
	)
	m.Name = "New name"
	m.Description = "New description"
	m.Tags = []string{"new", "tag"}
	err = cli.PutMuscle(m)
	require.NoError(t, err, "put muscle")

	var newMuscle *internal.Muscle
	newMuscle, err = cli.GetMuscle(m.ID)
	require.NoError(t, err, "get muscle")
	require.Equal(t, newMuscleName, newMuscle.Name, newMuscle)
	require.Equal(t, newMuscleDesc, newMuscle.Description, newMuscle)
	require.Equal(t, newMuscleTags, newMuscle.Tags, newMuscle)

	// Search muscles by tags
	muscles, err := cli.SearchMusclesByTags("new")
	require.NoError(t, err, "search muscles")
	require.Equal(t, 1, len(muscles), "len muscles")

	// Delete muscle request
	err = cli.DeleteMuscle(m.ID)
	require.NoError(t, err, "delete muscle request")
	m, err = cli.GetMuscle(id)
	require.NoError(t, err, "get muscle")
	require.Empty(t, m, "after delete request")
}
