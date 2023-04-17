package service

import (
	"time"

	"github.com/anaxaim/tui/server/pkg/model"
	"github.com/anaxaim/tui/server/pkg/repository"
)

type credentialService struct {
	credentialRepository repository.CredentialRepository
}

func NewCredentialService(credentialRepository repository.CredentialRepository) CredentialService {
	return &credentialService{
		credentialRepository: credentialRepository,
	}
}

func (c *credentialService) List() (model.Credentials, error) {
	return c.credentialRepository.List()
}

func (c *credentialService) Create(credential *model.Credential) (*model.Credential, error) {
	currentTime := time.Now()
	credential.CreatedAt = currentTime
	credential.CreatedAtString = currentTime.Format("15:04:05 02/01/2006")

	return c.credentialRepository.Create(credential)
}

func (c *credentialService) Get(id string) (*model.Credential, error) {
	return c.getCredentialByID(id)
}

func (c *credentialService) Delete(id string) error {
	credential, err := c.getCredentialByID(id)
	if err != nil {
		return err
	}

	return c.credentialRepository.Delete(credential)
}

func (c *credentialService) getCredentialByID(id string) (*model.Credential, error) {
	return c.credentialRepository.GetCredentialByID(id)
}
