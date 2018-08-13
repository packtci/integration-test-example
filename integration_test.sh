#! /bin/bash

GO_TEMPLATE_EXAMPLE_REPO=$(curl -X GET \
    --header "Accept: application/json" \
    "https://circleci.com/api/v1.1/project/github/packtci/go-template-example-with-circle-ci?circle-token=$CIRCLECI_API_TOKEN_GITHUB" | jq '.[0].author_name' | tr -d "\n")

if [[  -n ${GO_TEMPLATE_EXAMPLE_REPO} ]]; then
    echo "The current owner was shown"
    exit 0
else 
    echo "No owner own"
    exit 1
fi
