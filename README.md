# Files

This is a learning project, focused on distributed systems. The goal is to build a resilient system to share files via a url. The system will be comprised of multiple services.

## General Notes

- Services communicate using HTTP.
- Meta data can be stored in a NoSQL DB e.g. MongoDB.
- Each file node will periodically ping the controller service to make itself 'active' and available for work.

## Frontend Service

### Responsible for...

- Serving the upload page.
- Handling download requests.
    - Fetch file chunks, recombine and return to the user.
- Handling upload requests.
    - Chunk the files and distribute them between active file nodes.
    - Makes initial request to store file (with retry).
    - Once successfully uploaded, publish replication event to the message queue.

### Notes

- Split files into 4MB chunks.

### Endpoints

- GET / upload page
- GET /download/{id} download the given file ID
- POST /upload chunk and save a file

## Controller Service

### Responsible for...

- Storing file meta data
    - id (hash)
    - name
    - time uploaded
    - chunks
        - chunk id (hash)
        - location (array of file node ids)
        - order
- Storing file node information
    - id
    - last ping time
- Listing active file nodes 
    - node is active if it has pinged the controller service within a set time
- Listing file information

### Endpoints

- POST /file save a new file
- POST /file/{id} update an existing file
- GET /file/{id} return the meta data for the given file ID
- POST /ping create or update a file node (used for both initial registration and periodic check ins)
- GET /active-nodes returns a list of all active nodes 

## Message Queue

Redis/RabbitMQ

## Replicator Service

### Responsible for...

- Replicating file chunks across active file nodes.
    - Each chunk should exist on at least 2 nodes.
    - Retry replication on failure (same node and then different).

### Notes

- Subscribes to message queue, listening for chunk replicate event.

## Garbage Collector Service

### Responsible for...

- Removing expired chunks.
- Updating the controller service to remove file meta data.
- If file is only partially deleted (due to a file node being offline) try again at a later time.

## File Node

### Responsible for...

- Storing the file chunk
    - POST { hash: '', chunk: raw binary data }
    - Verify the hash matches the chunk
- Serving the file chunk

### Endpoints

- POST /chunk verify and store a chunk
- GET /chunk/{id} return a chunk
