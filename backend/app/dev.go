package app

import (
	"fmt"

	"github.com/okteto/app/backend/k8s/client"
	"github.com/okteto/app/backend/k8s/deployments"
	"github.com/okteto/app/backend/k8s/secrets"
	"github.com/okteto/app/backend/k8s/volumes"
	"github.com/okteto/app/backend/model"
)

//DevModeOn activates a development environemnt
func DevModeOn(dev *model.Dev, s *model.Space) error {
	c, err := client.Get()
	if err != nil {
		return fmt.Errorf("error getting k8s client: ", err)
	}

	if err := secrets.Create(dev, s, c); err != nil {
		return err
	}

	if err := volumes.Create(dev, s, c); err != nil {
		return err
	}

	if err := deployments.Deploy(dev, s, c); err != nil {
		return err
	}

	return nil
}

//DevModeOff deactivates a development environment
func DevModeOff(dev *model.Dev, s *model.Space, removeVolumes bool) error {
	c, err := client.Get()
	if err != nil {
		return fmt.Errorf("error getting k8s client: ", err)
	}

	if err := secrets.Destroy(dev, s, c); err != nil {
		return err
	}

	if removeVolumes {
		if err := volumes.Destroy(dev, s, c); err != nil {
			return err
		}
	}

	if err := deployments.Destroy(dev, s, c); err != nil {
		return err
	}

	return nil
}