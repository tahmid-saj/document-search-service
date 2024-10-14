package searcher

import (
	"document-search-service/dynamodb"
	"document-search-service/object"
	"log"
	"strings"
)

type DocumentTitle string
type DocumentContent string
type Documents map[DocumentTitle]DocumentContent

func SearchDocuments(searchQuery, tableName, bucketName string) (
	documentInvertedIndexMappings dynamodb.InvertedIndexMappings, documents Documents, err error) {
	searchQueryList := strings.Split(searchQuery, " ")

	documentInvertedIndexMappings = make(dynamodb.InvertedIndexMappings)
	documents = make(Documents)

	// parse through the searchQuery
	for _, searchQueryTerm := range searchQueryList {
		// retrieve searchQuery's term from dynamodb if it exists
		searchQueryTermInvertedIndex, err := dynamodb.ReadItem(searchQueryTerm, tableName)
		if err != nil {
			log.Print(err)
			return nil, nil, err
		}

		// retrieve the document title and content
		for _, documentID := range searchQueryTermInvertedIndex.DocumentTermMatrix.DocumentIDs {
			// check if the documentID is not already in the search result return list
			if _, exists := documents[DocumentTitle(documentID)]; !exists {
				document, err := retrieveDocument(bucketName, documentID)
				if err != nil {
					log.Print(err)
					return nil, nil, err
				}
				
				documents[DocumentTitle(document.Title)] = DocumentContent(document.Content)
			}

			// check if the term is not already in the search result return list
			if _, exists := documentInvertedIndexMappings[dynamodb.Term(searchQueryTerm)]; !exists {
				documentInvertedIndexMappings[dynamodb.Term(searchQueryTerm)] = searchQueryTermInvertedIndex.DocumentTermMatrix
			}
		}
	}

	return documentInvertedIndexMappings, documents, nil
}

// helper functions

func retrieveDocument(bucketName, documentID string) (document *object.Document, err error) {
	s3Client := InitS3()

	document, err = object.ReadObject(s3Client, bucketName, documentID)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	return document, nil
}