# document-search-service

Document searcher service to search for documents using search query terms. Uses inverted index mappings to optimize search queries. Developed with Go / Gin, S3, DynamoDB.

<br/>
<br/>

## Directory structure

The directory structure is as follows:

### `/bucket`
Contains logic for interacting with Amazon S3 to store and retrieve documents.

### `/conf`
Holds configuration files and settings for the service.

### `/data`
Stores sample data or datasets used for indexing and search.

### `/dynamodb`
Manages interactions with DynamoDB, including table operations.

### `/models`
Defines data models used within the service.

### `/object`
Contains utilities for managing document objects in storage.

### `/routes`
Defines API routes for the service.

### `/searcher`
Handles the core search logic using an inverted index.

### `/utils`
Utility functions used across the project.

### `/main.go`
Main entry point for the application.

<br/>
<br/>

## Overview

### Design

The high level workflow of the document searcher can be found below. Similar services can be found <a href="https://whimsical.com/web-microservices-6uqvwWZtcBFsNJB2hepGy1">here</a> and below:

#### Document searcher workflow

<img width="518" alt="image" src="https://github.com/user-attachments/assets/daa66cc1-a116-4097-8624-905bc4dc9590">

#### Similar services

<img width="834" alt="image" src="https://github.com/user-attachments/assets/b54088e7-870c-46dd-9cf6-2e5ec27d9d5c">

### Examples

#### Sample inverted index mappings

<img width="1334" alt="image" src="https://github.com/user-attachments/assets/a6c5e3f1-7296-4913-9dba-d8e1daee2d45">

#### Sample search result from API

```
// Input
{
    "tableName": "document-indexer-index-mapping",
    "bucketName": "document-indexer-service-test-documents"
}
```

```
// Output
{
    "ok": true,
    "response": {
        "InvertedIndexMappings": {
            "Lorem": {
                "documentIDs": [
                    "lorem_ipsum_1.json",
                    "lorem_ipsum_2.json",
                    "lorem_ipsum_3.json",
                    "lorem_ipsum_4.json"
                ],
                "documentTermFrequencies": [
                    5,
                    2,
                    5,
                    5
                ],
                "documentTermLocations": [
                    [
                        2,
                        0,
                        12,
                        73,
                        89
                    ],
                    [
                        28,
                        64
                    ],
                    [
                        4,
                        52,
                        70,
                        115,
                        133
                    ],
                    [
                        7,
                        40,
                        59,
                        100,
                        107
                    ]
                ]
            },
            "The": {
                "documentIDs": [
                    "lorem_ipsum_2.json",
                    "lorem_ipsum_3.json",
                    "lorem_ipsum_4.json"
                ],
                "documentTermFrequencies": [
                    1,
                    2,
                    1
                ],
                "documentTermLocations": [
                    [
                        24
                    ],
                    [
                        111,
                        129
                    ],
                    [
                        105
                    ]
                ]
            }
        },
        "Documents": {
            "What is Lorem Ipsum?": "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.",
            "Where can I get some?": "There are many variations of passages of Lorem Ipsum available, but the majority have suffered alteration in some form, by injected humour, or randomised words which don't look even slightly believable. If you are going to use a passage of Lorem Ipsum, you need to be sure there isn't anything embarrassing hidden in the middle of text. All the Lorem Ipsum generators on the Internet tend to repeat predefined chunks as necessary, making this the first true generator on the Internet. It uses a dictionary of over 200 Latin words, combined with a handful of model sentence structures, to generate Lorem Ipsum which looks reasonable. The generated Lorem Ipsum is therefore always free from repetition, injected humour, or non-characteristic words etc.",
            "Where does it come from?": "Contrary to popular belief, Lorem Ipsum is not simply random text. It has roots in a piece of classical Latin literature from 45 BC, making it over 2000 years old. Richard McClintock, a Latin professor at Hampden-Sydney College in Virginia, looked up one of the more obscure Latin words, consectetur, from a Lorem Ipsum passage, and going through the cites of the word in classical literature, discovered the undoubtable source. Lorem Ipsum comes from sections 1.10.32 and 1.10.33 of \"de Finibus Bonorum et Malorum\" (The Extremes of Good and Evil) by Cicero, written in 45 BC. This book is a treatise on the theory of ethics, very popular during the Renaissance. The first line of Lorem Ipsum, \"Lorem ipsum dolor sit amet..\", comes from a line in section 1.10.32. The standard chunk of Lorem Ipsum used since the 1500s is reproduced below for those interested. Sections 1.10.32 and 1.10.33 from \"de Finibus Bonorum et Malorum\" by Cicero are also reproduced in their exact original form, accompanied by English versions from the 1914 translation by H. Rackham.",
            "Why do we use it?": "It is a long established fact that a reader will be distracted by the readable content of a page when looking at its layout. The point of using Lorem Ipsum is that it has a more-or-less normal distribution of letters, as opposed to using 'Content here, content here', making it look like readable English. Many desktop publishing packages and web page editors now use Lorem Ipsum as their default model text, and a search for 'lorem ipsum' will uncover many web sites still in their infancy. Various versions have evolved over the years, sometimes by accident, sometimes on purpose (injected humour and the like)."
        }
    }
}
```
