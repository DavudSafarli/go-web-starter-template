package contracts

import (
	"context"
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/DavudSafarli/go-web-starter-template/domains/appname"
	"github.com/stretchr/testify/require"
)

type StorageContract struct {
	Subject appname.Storage
}

var random = rand.New(rand.NewSource(time.Now().UnixNano()))

func getUniqueResource(t *testing.T) appname.Resource {
	return appname.Resource{
		Body: fmt.Sprint("body-", t.Name(), random.Int()),
	}
}

func (c StorageContract) Test(t *testing.T) {
	c.TestCreateResource(t)
	c.TestFindResource(t)
	c.TestGetResources(t)
	c.TestDeleteResource(t)
	c.TestUpdateResource(t)
}

func (c StorageContract) TestCreateResource(t *testing.T) {
	t.Run(`Created resource has an ID and other fields are correct`, func(t *testing.T) {
		input := getUniqueResource(t)

		resource, _ := c.Subject.CreateResource(context.Background(), input)
		defer c.Subject.DeleteResource(context.Background(), resource.ID)

		require.NotEmpty(t, resource.ID)
		require.Equal(t, resource.Body, input.Body)
	})

	t.Run(`Created resource can be find`, func(t *testing.T) {
		input := getUniqueResource(t)

		createdResource, _ := c.Subject.CreateResource(context.Background(), input)
		defer c.Subject.DeleteResource(context.Background(), createdResource.ID)

		foundResource, err := c.Subject.FindResource(context.Background(), appname.Resource{ID: input.ID})
		require.Nil(t, err)
		require.Equal(t, createdResource, foundResource)
	})
}

func (c StorageContract) TestFindResource(t *testing.T) {
	t.Run(`Created resource can be find by ID`, func(t *testing.T) {
		record := c.createResourceGet(t)
		criteria := appname.Resource{ID: record.ID}

		foundResource, err := c.Subject.FindResource(context.Background(), criteria)

		require.Nil(t, err)
		require.Equal(t, record, foundResource)
	})
}

func (c StorageContract) TestGetResources(t *testing.T) {
	t.Run(`Created resources are included in response`, func(t *testing.T) {
		record1 := c.createResourceGet(t)
		record2 := c.createResourceGet(t)

		resources, err := c.Subject.GetResources(context.Background())

		require.Nil(t, err)
		require.Contains(t, resources, record1)
		require.Contains(t, resources, record2)
	})
}

func (c StorageContract) TestDeleteResource(t *testing.T) {
	t.Run(`Deleting non-existing record returns false`, func(t *testing.T) {
		nonExistingRecordID := -3151

		ok, err := c.Subject.DeleteResource(context.Background(), nonExistingRecordID)

		require.Nil(t, err)
		require.False(t, ok)
	})

	t.Run(`Deleting existing record returns true`, func(t *testing.T) {
		record := c.createResourceGet(t)

		ok, err := c.Subject.DeleteResource(context.Background(), record.ID)

		require.Nil(t, err)
		require.True(t, ok)
	})

	t.Run(`Record cannot be found after Deletion`, func(t *testing.T) {
		record := c.createResourceGet(t)

		c.Subject.DeleteResource(context.Background(), record.ID)

		_, err := c.Subject.FindResource(context.Background(), appname.Resource{ID: record.ID})
		require.Nil(t, err)
		require.Equal(t, err, appname.ErrResourceNotFound)
	})
}

func (c StorageContract) TestUpdateResource(t *testing.T) {
	t.Run(`Update returns the updated record`, func(t *testing.T) {
		record := c.createResourceGet(t)
		input := appname.Resource{ID: record.ID, Body: "updated-body"}

		updatedRecord, err := c.Subject.UpdateResource(context.Background(), input)

		require.Nil(t, err)
		require.Equal(t, input, updatedRecord)
	})

	t.Run(`Updated record is persisted`, func(t *testing.T) {
		record := c.createResourceGet(t)
		input := appname.Resource{ID: record.ID, Body: "updated-body"}

		c.Subject.UpdateResource(context.Background(), input)

		// fetch record of the ID
		recordAfterUpdate, err := c.Subject.FindResource(context.Background(), appname.Resource{ID: input.ID})
		require.Nil(t, err)
		require.Equal(t, input, recordAfterUpdate)
	})

	t.Run(`Updating a record with non-existing ID returns ErrResourceNotFound`, func(t *testing.T) {
		nonExistingRecordID := -3151
		input := appname.Resource{ID: nonExistingRecordID, Body: "updated-body"}

		_, err := c.Subject.UpdateResource(context.Background(), input)

		require.Equal(t, appname.ErrResourceNotFound, err)
	})
}

func (c StorageContract) createResourceGet(t *testing.T) appname.Resource {
	record := getUniqueResource(t)
	createdResource, err := c.Subject.CreateResource(context.Background(), record)
	require.Nil(t, err)
	t.Cleanup(func() {
		c.Subject.DeleteResource(context.Background(), createdResource.ID)
	})
	return createdResource
}
