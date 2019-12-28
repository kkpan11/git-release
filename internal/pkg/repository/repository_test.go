package repository_test

import (
	"fmt"
	"github.com/anton-yurchenko/git-release/internal/pkg/repository"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestReadTag(t *testing.T) {
	assert := assert.New(t)

	// TEST: env.var correct
	tag := "refs/tags/v1.0.0"
	var version string

	err := os.Setenv("GITHUB_REF", tag)
	assert.Equal(nil, err, "preparation: error setting env.var 'GITHUB_REF'")

	m := new(repository.Repository)

	err = m.ReadTag(&version)

	assert.Equal(nil, err)
	assert.Equal("v1.0.0", m.Tag)
	assert.Equal("1.0.0", version)

	// TEST: env.var incorrect
	err = os.Setenv("GITHUB_REF", "malformed-var")
	assert.Equal(nil, err, "preparation: error setting env.var 'GITHUB_REF'")

	m = new(repository.Repository)

	err = m.ReadTag(&version)

	assert.EqualError(err, "malformed env.var 'GITHUB_REF': expected to match regex 'refs/tags/.*([0-9]+.[0-9]+.[0-9]+)', got 'malformed-var'")

	// TEST: env.var not set
	err = os.Setenv("GITHUB_REF", "")
	assert.Equal(nil, err, "preparation: error setting env.var 'GITHUB_REF'")

	m = new(repository.Repository)

	err = m.ReadTag(&version)

	assert.EqualError(err, "env.var 'GITHUB_REF' is empty or not defined")
}

func TestReadCommitHash(t *testing.T) {
	assert := assert.New(t)

	// TEST: env.var set
	expected := "123abc"

	err := os.Setenv("GITHUB_SHA", expected)
	assert.Equal(nil, err, "preparation: error setting env.var 'GITHUB_SHA'")

	m := new(repository.Repository)

	err = m.ReadCommitHash()

	assert.Equal(nil, err)
	assert.Equal(expected, m.CommitHash)

	// TEST: env.var not set
	err = os.Setenv("GITHUB_SHA", "")
	assert.Equal(nil, err, "preparation: error setting env.var 'GITHUB_SHA'")

	m = new(repository.Repository)

	err = m.ReadCommitHash()

	assert.EqualError(err, "env.var 'GITHUB_SHA' is empty or not defined")
}

func TestReadProjectName(t *testing.T) {
	assert := assert.New(t)

	// TEST: env.var correct
	user := "user"
	project := "project"

	err := os.Setenv("GITHUB_REPOSITORY", fmt.Sprintf("%v/%v", user, project))
	assert.Equal(nil, err, "preparation: error setting env.var 'GITHUB_REPOSITORY'")

	m := new(repository.Repository)

	err = m.ReadProjectName()

	assert.Equal(nil, err)
	assert.Equal(user, m.Owner)
	assert.Equal(project, m.Project)

	// TEST: env.var incorrect
	err = os.Setenv("GITHUB_REPOSITORY", "value")
	assert.Equal(nil, err, "preparation: error setting env.var 'GITHUB_REPOSITORY'")

	m = new(repository.Repository)

	err = m.ReadProjectName()

	assert.EqualError(err, "malformed env.var 'GITHUB_REPOSITORY': expected to match regex '.*/.*', got 'value'")

	// TEST: env.var not set
	err = os.Setenv("GITHUB_REPOSITORY", "")
	assert.Equal(nil, err, "preparation: error setting env.var 'GITHUB_REPOSITORY'")

	m = new(repository.Repository)

	err = m.ReadProjectName()

	assert.EqualError(err, "env.var 'GITHUB_REPOSITORY' is empty or not defined")
}

func TestGetOwner(t *testing.T) {
	assert := assert.New(t)

	expected := "user"

	m := repository.Repository{
		Owner: expected,
	}

	result := m.GetOwner()

	assert.Equal(expected, result)
}

func TestGetProject(t *testing.T) {
	assert := assert.New(t)

	expected := "project"

	m := repository.Repository{
		Project: expected,
	}

	result := m.GetProject()

	assert.Equal(expected, result)
}

func TestGetTag(t *testing.T) {
	assert := assert.New(t)

	expected := "1.0.0"

	m := repository.Repository{
		Tag: expected,
	}

	result := m.GetTag()

	assert.Equal(expected, *result)
}

func TestGetCommitHash(t *testing.T) {
	assert := assert.New(t)

	expected := "123"

	m := repository.Repository{
		CommitHash: expected,
	}

	result := m.GetCommitHash()

	assert.Equal(expected, *result)
}