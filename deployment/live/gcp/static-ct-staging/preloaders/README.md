# How to build
docker build -t us-central1-docker.pkg.dev/static-ct-staging/docker-staging/preloader:PUT_HASH_HERE -t us-central1-docker.pkg.dev/static-ct-staging/docker-staging/preloader:latest -f ./Dockerfile .
gcloud auth login --project=static-ct-staging
docker -v push --all-tags  us-central1-docker.pkg.dev/static-ct-staging/docker-staging/preloader

# How to use
Then, manually edit root.hcl with the new container
