package test

import (
	"net/http/httptest"
	"testing"

	"gin-admin/internal/app/ginadmin/schema"
	"gin-admin/pkg/util"
	"github.com/stretchr/testify/assert"
)

func TestRole(t *testing.T) {
	const router = "v1/roles"
	var err error

	w := httptest.NewRecorder()

	// post /Permissions
	addPermissionItem := &schema.Permission{
		Name:     util.MustUUID(),
		Sequence: 9999999,
		Actions: []*schema.PermissionAction{
			{Code: "query", Name: "query"},
		},
		Resources: []*schema.PermissionResource{
			{Code: "query", Name: "query", Method: "GET", Path: "/test/v1/Permissions"},
		},
	}
	engine.ServeHTTP(w, newPostRequest("v1/Permissions", addPermissionItem))
	assert.Equal(t, 200, w.Code)
	var addNewPermissionItem schema.Permission
	err = parseReader(w.Body, &addNewPermissionItem)
	assert.Nil(t, err)

	// post /roles
	addItem := &schema.Role{
		Name:     util.MustUUID(),
		Sequence: 9999999,
		Permissions: []*schema.RolePermission{
			{
				PermissionID:    addNewPermissionItem.RecordID,
				Actions:   []string{"query"},
				Resources: []string{"query"},
			},
		},
	}
	engine.ServeHTTP(w, newPostRequest(router, addItem))
	assert.Equal(t, 200, w.Code)
	var addNewItem schema.Role
	err = parseReader(w.Body, &addNewItem)
	assert.Nil(t, err)
	assert.Equal(t, addItem.Name, addNewItem.Name)
	assert.Equal(t, addItem.Sequence, addNewItem.Sequence)
	assert.Equal(t, len(addItem.Permissions), len(addNewItem.Permissions))
	assert.NotEmpty(t, addNewItem.RecordID)

	// query /roles?q=page
	engine.ServeHTTP(w, newGetRequest(router,
		newPageParam(map[string]string{"q": "page"})))
	assert.Equal(t, 200, w.Code)
	var pageItems []*schema.Role
	err = parsePageReader(w.Body, &pageItems)
	assert.Nil(t, err)
	assert.Equal(t, len(pageItems), 1)
	if len(pageItems) > 0 {
		assert.Equal(t, addNewItem.RecordID, pageItems[0].RecordID)
		assert.Equal(t, addNewItem.Name, pageItems[0].Name)
	}

	// put /roles/:id
	engine.ServeHTTP(w, newGetRequest("%s/%s", nil, router, addNewItem.RecordID))
	assert.Equal(t, 200, w.Code)
	var putItem schema.Role
	err = parseReader(w.Body, &putItem)
	putItem.Name = util.MustUUID()
	engine.ServeHTTP(w, newPutRequest("%s/%s", putItem, router, addNewItem.RecordID))
	assert.Equal(t, 200, w.Code)
	var putNewItem schema.Role
	err = parseReader(w.Body, &putNewItem)
	assert.Nil(t, err)
	assert.Equal(t, putItem.Name, putNewItem.Name)
	assert.Equal(t, len(putItem.Permissions), len(putNewItem.Permissions))

	// delete /Permissions/:id
	engine.ServeHTTP(w, newDeleteRequest("%s/%s", "v1/Permissions", addNewPermissionItem.RecordID))
	assert.Equal(t, 200, w.Code)
	err = parseOK(w.Body)
	assert.Nil(t, err)

	// delete /roles/:id
	engine.ServeHTTP(w, newDeleteRequest("%s/%s", router, addNewItem.RecordID))
	assert.Equal(t, 200, w.Code)
	err = parseOK(w.Body)
	assert.Nil(t, err)
}
