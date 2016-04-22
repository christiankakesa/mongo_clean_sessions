# congopro_clean_sessions
Congopro utility program to clean MongoDB "sessions" collection. 

## Usage

    Usage of ./main:
      -c string
            MongoDB collection to cleanup. (default "sessions")
      -r int
            MongoDB retention delai in hour(s). Default is 7 days (168 hours). (default 168)
      -url string
            MongoDB connection URI. (default "mongodb://localhost:27017/test")

## Example

    ./main -url=${CONGOPRO_MONGOLAB_URL} -c=sessions -r=72
