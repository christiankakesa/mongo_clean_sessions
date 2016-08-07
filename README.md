# Mongo clean sessions
Utility program to cleanup MongoDB collection by a specific time field. 

## Installation
### Download your platform archive
Go to the **[releases]( https://github.com/fenicks/mongo_clean_sessions/releases)** section: https://github.com/fenicks/mongo_clean_sessions/releases, download and uncompress the archive corresponding to your platform:

    wget https://github.com/fenicks/mongo_clean_sessions/releases/download/v0.9/mongo_clean_sessions-v1.1-linux-amd64.tar.gz
    tar -xzf mongo_clean_sessions-v1.1-linux-amd64.tar.gz

Run the program as described in next sections.

### Build the binary
You need a working **Go** development environment ([see the installation guide](https://golang.org/doc/install))

    make

Run the program as described in next sections.

## Usage
    Usage of ./mongo_clean_sessions:
      -c string
         MongoDB collection to cleanup. (default "sessions")
      -f string
         MongoDB collection field with type 'time.Time'. (default "updated_at")
      -r int
         MongoDB retention delai in hour(s). Default is 7 days (168 hours). (default 168)
      -s Simulation mode, no deletion are send to the MongoDB database.
      -url string
         MongoDB connection URI. (default "mongodb://localhost:27017/test")

## Example
    ./mongo_clean_sessions -url=${MONGOLAB_URL} -c=sessions -f=updated_at -r=72

### Simulation mode
Just add a **-s** in command line:

    ./mongo_clean_sessions -s -url=${MONGOLAB_URL} -c=sessions -f=updated_at -r=72
