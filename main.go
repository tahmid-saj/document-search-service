package main

import (
	"document-search-service/routes"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// documentInvertedIndexMappings, documents, err := searcher.SearchDocuments("The", "document-indexer-index-mapping", "document-indexer-service-test-documents")
	// if err != nil {
	// 	return
	// }
	// fmt.Println(documentInvertedIndexMappings)
	// fmt.Println(documents)

	godotenv.Load()

	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(os.Getenv("PORT"))
}