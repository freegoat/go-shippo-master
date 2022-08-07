package client

import (
	"encoding/json"
	"errors"
	"net/http"

	//"github.com/coldbrewcloud/go-shippo/models"
	"github.com/freegoat/go-shippo-master/models"
)

// PurchaseShippingLabel creates a new transaction object and purchases the shipping label for the provided rate.
func (c *Client) PurchaseShippingLabel(input *models.TransactionInput) (*models.Transaction, error) {
	if input == nil {
		return nil, errors.New("nil input")
	}

	output := &models.Transaction{}
	err := c.do(http.MethodPost, "/transactions/", input, output)
	return output, err
}

// PurchaseShippingLabel instant call return us a rate object instead of rate string
func (c *Client) PurchaseShippingLabelInstant(input *models.TransactionInput) (*models.TransactionFromInstantLabel, error) {
	if input == nil {
		return nil, errors.New("nil input")
	}

	output := &models.TransactionFromInstantLabel{}
	err := c.do(http.MethodPost, "/transactions/", input, output)
	return output, err
}

// RetrieveTransaction retrieves an existing transaction by object id.
func (c *Client) RetrieveTransaction(objectID string) (*models.Transaction, error) {
	if objectID == "" {
		return nil, errors.New("empty object ID")
	}

	output := &models.Transaction{}
	err := c.do(http.MethodGet, "/transactions/"+objectID, nil, output)
	return output, err
}

// ListAllTransactions lists all transaction objects.
func (c *Client) ListAllTransactions() ([]*models.Transaction, error) {
	list := []*models.Transaction{}
	err := c.doList(http.MethodGet, "/transactions/", nil, func(v json.RawMessage) error {
		item := &models.Transaction{}
		if err := json.Unmarshal(v, item); err != nil {
			return err
		}

		list = append(list, item)
		return nil
	})
	return list, err
}
