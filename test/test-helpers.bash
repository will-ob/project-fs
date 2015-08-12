
DIR=$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )

function createProject {
  source $DIR/../.env.test
  curl 'https://api.f7ops.dev/v0.1/projects' \
    -X POST \
    -H 'content-length: 0' \
    -H 'content-type: application/json' \
    -H 'accept: */*' \
    -H "x-api-key: $PROJECT_API_KEY" \
    --insecure \
    --silent \
    --show-error
}

export PROJECT_DIR=$( realpath $DIR/.. )
