# secrets-app-http-client
Sample secret sharing application cli client build as part of technical review of Manning liveProject "Create a Custom HTTP Client to Interact with the HTTP Server"

This is a CLI client for [secrets-app](https://github.com/borkod/secrets-app)

Usage:
```
secrets create -data <secret data> -url <secrets app server url>

secrets create -data mysecretvalue -url http://localhost:8080

secrets view -id <secret id> -url <secrets app server url>

secrets view -id 11eed548a8f140ba4781f3eb11554680 -url http://localhost:8080"
```