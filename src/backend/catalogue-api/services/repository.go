package services

import (
	"context"
	"encoding/json"
	"os"

	"catalogue-api/dto"

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
func (repo *CosmosRepository) GetCatalogueItems(ctx context.Context) ([]dto.HodResponse, error) {
	query := "SELECT * FROM c"
	result := []dto.HodResponse{}

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
			var listResponse dto.HodResponse
			if err := json.Unmarshal(item, &listResponse); err != nil {
				return nil, errors.Wrap(err, "failed to unmarshal item into ListResponse")
			}
			result = append(result, listResponse)
		}
	}

	return result, nil
}
