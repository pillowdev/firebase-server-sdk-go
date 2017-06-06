package firebase

import (
	"errors"
	"fmt"
	"os"
)

// Options is storage for configurable Firebase options.
type Options struct {
	// ProjectID is the project ID associated with the application.
	ProjectID string
	// ServiceAccountPath is the path to load the Service Account.
	ServiceAccountPath string
	// ServiceAccountCredential is the credential for the Service Account.
	ServiceAccountCredential *GoogleServiceAccountCredential
}

// ensureServiceAccount sets the Service Account associated with the Firebase Options.
func (o *Options) ensureServiceAccount() error {
	if o.ServiceAccountCredential != nil {
		// credential already loaded
		return nil
	}
	if o.ServiceAccountPath == "" {
		return errors.New("ServiceAccountPath cannot be empty.")
	}

	f, err := os.Open(o.ServiceAccountPath)
	if err != nil {
		return fmt.Errorf("Service Account file cannot be opened: %s %v", o.ServiceAccountPath, err)
	}
	defer f.Close()

	c, err := loadCredential(f)
	if err != nil {
		return err
	}
	o.ServiceAccountCredential = c
	return nil
}

// ensureProjectID sets the Project ID associated with the Firebase Options.
func (o *Options) ensureProjectID() error {
	if o.ProjectID == "" {
		return errors.New("ProjectID cannot be empty.")
	}
	return nil
}
