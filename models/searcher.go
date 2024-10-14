package models

import (
	"document-search-service/dynamodb"
	"document-search-service/searcher"
)

type SearchQueryInput struct {
	TableName  string `json:"tableName"`
	BucketName string `json:"bucketName"`
}

type Response struct {
	Ok       bool        `json:"ok"`
	Response interface{} `json:"response"`
}

type SearchQueryResponse struct {
	InvertedIndexMappings dynamodb.InvertedIndexMappings
	Documents             searcher.Documents
}

func SearchDocuments(searchQuery string, searchQueryInput SearchQueryInput) (*Response, error) {
	documentInvertedIndexMappings, documents, err := searcher.SearchDocuments(searchQuery, searchQueryInput.TableName, searchQueryInput.BucketName)
	if err != nil {
		return &Response{
			Ok:       false,
			Response: nil,
		}, err
	}

	return &Response{
		Ok: true,
		Response: SearchQueryResponse{
			InvertedIndexMappings: documentInvertedIndexMappings,
			Documents:             documents,
		},
	}, nil
}