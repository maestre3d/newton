#!/bin/bash

# Requires dynamo local JVM binaries
function start_dynamo() {
    java -Djava.library.path="$1/"DynamoDBLocal_lib -jar "$1/"DynamoDBLocal.jar -sharedDb -inMemory -port 8001
}

start_dynamo "$1"