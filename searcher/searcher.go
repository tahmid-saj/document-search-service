package searcher

import "document-search-service/dynamodb"

func SearchDocuments(searchQuery, tableName, bucketName string) (*dynamodb.InvertedIndexMappings, error) {
	// parse through the searchQuery

	// retrieve searchQuery's term from dynamodb if it exists

	// retrieve the document names

	// retrieve the document contents

}

// helper functions

func retrieveDocumentNames() {

}

func retrieveDocumentContents() {

}