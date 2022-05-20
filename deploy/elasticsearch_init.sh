#!/bin/sh

ES_HOST=elasticsearch:9200

curl -s -X PUT $ES_HOST/foods \
    -H "Content-Type: application/json" \
    -d '
        {
            "mappings": {
                "properties": {
                    "name": {
                        "type": "text"
                    },
                    "description": {
                        "type": "text"
                    }
                },
                "dynamic": false
            }
        }    
    '