package firebase

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnsureServiceAccount(t *testing.T) {
	o := &Options{}
	err := o.ensureServiceAccount()
	assert.EqualError(t, err, "ServiceAccountPath cannot be empty.")

	o.ServiceAccountPath = "testdata/service-account-appengine.json"
	assert.NoError(t, o.ensureServiceAccount())
	c := o.ServiceAccountCredential
	assert.NotNil(t, c)
	assert.Equal(t, "myapp-dev", c.ProjectID)
	assert.Equal(t, "myapp-dev@appspot.gserviceaccount.com", c.ClientEmail)
	assert.NotNil(t, c.PrivateKey)
}

func TestEnsureProjectID(t *testing.T) {
	o := &Options{}
	err := o.ensureProjectID()
	assert.EqualError(t, err, "ProjectID cannot be empty.")
}
