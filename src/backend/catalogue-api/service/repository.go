package service

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"catalogue-api/models"

	"github.com/Azure/azure-sdk-for-go/sdk/data/azcosmos"
	"github.com/pkg/errors"
)

type CosmosRepository struct {
	client    *azcosmos.Client
	container *azcosmos.ContainerClient
}

// NewCosmosRepository initializes a new CosmosRepository with a Cosmos client and container.
func NewCosmosRepository() (*CosmosRepository, error) {
	// Load environment variables
	cosmosDbEndpoint := os.Getenv("COSMOS_URI")
	cosmosDbKey := os.Getenv("COSMOS_KEY")
	databaseName := os.Getenv("COSMOS_DATABASE_NAME")
	containerName := os.Getenv("COSMOS_CONTAINER_NAME")

	// Create a key credential using the CosmosDB key
	cred, err := azcosmos.NewKeyCredential(cosmosDbKey)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create CosmosDB key credential")
	}

	// Create a CosmosDB client with the endpoint and credentials
	client, err := azcosmos.NewClientWithKey(cosmosDbEndpoint, cred, nil)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create Cosmos client")
	}

	// Get the Cosmos container reference (handle the two return values)
	container, err := client.NewContainer(databaseName, containerName)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get Cosmos container reference")
	}

	return &CosmosRepository{
		client:    client,
		container: container, // Store the container client
	}, nil
}

// GetCatalogueItems retrieves all items from the catalogue.
func (repo *CosmosRepository) GetCatalogueItems(ctx context.Context) ([]models.ListResponse, error) {
	query := "SELECT * FROM c"
	result := []models.ListResponse{}

	// Create an empty partition key if needed (or use a valid partition key value)
	partitionKey := azcosmos.PartitionKey{}

	// Execute the query using the new QueryItemsPager
	pager := repo.container.NewQueryItemsPager(query, partitionKey, nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, errors.Wrap(err, "failed to retrieve catalogue items")
		}

		// Unmarshal each item in the page into ListResponse
		for _, item := range page.Items {
			var listResponse models.ListResponse
			if err := json.Unmarshal(item, &listResponse); err != nil {
				return nil, errors.Wrap(err, "failed to unmarshal item into ListResponse")
			}
			result = append(result, listResponse)
		}
	}

	return result, nil
}

// GetCatalogueItemByID retrieves a specific catalogue item by its ID.
func (repo *CosmosRepository) GetCatalogueItemByID(ctx context.Context, id string) (*models.GetAnItemResponse, error) {
	query := fmt.Sprintf("SELECT * FROM c WHERE c.id = '%s'", id)
	result := []models.GetAnItemResponse{}

	// Create an empty partition key if needed
	partitionKey := azcosmos.PartitionKey{}

	// Execute the query
	pager := repo.container.NewQueryItemsPager(query, partitionKey, nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, errors.Wrap(err, "failed to retrieve catalogue item by ID")
		}

		// Unmarshal each item in the page into GetAnItemResponse
		for _, item := range page.Items {
			var getItemResponse models.GetAnItemResponse
			if err := json.Unmarshal(item, &getItemResponse); err != nil {
				return nil, errors.Wrap(err, "failed to unmarshal item into GetAnItemResponse")
			}
			result = append(result, getItemResponse)
		}
	}

	if len(result) == 0 {
		return nil, nil // Item not found
	}
	return &result[0], nil
}

// GetCatalogueSize retrieves the total number of items in the catalogue.
func (repo *CosmosRepository) GetCatalogueSize(ctx context.Context) (int32, error) {
	query := "SELECT VALUE COUNT(1) FROM c"
	result := []models.GetSizeResponse{}

	// Create an empty partition key if needed
	partitionKey := azcosmos.PartitionKey{}

	// Execute the query
	pager := repo.container.NewQueryItemsPager(query, partitionKey, nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return 0, errors.Wrap(err, "failed to query catalogue size")
		}

		// Unmarshal each item in the page into GetSizeResponse
		for _, item := range page.Items {
			var sizeResponse models.GetSizeResponse
			if err := json.Unmarshal(item, &sizeResponse); err != nil {
				return 0, errors.Wrap(err, "failed to unmarshal size response")
			}
			result = append(result, sizeResponse)
		}
	}

	if len(result) == 0 {
		return 0, nil
	}
	return result[0].Size, nil
}
