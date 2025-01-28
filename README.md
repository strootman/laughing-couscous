# laughing-couscous
"laughing-couscous" is a codename for a Youtube CLI utility I'm writing for fun.

`[ ]` Determine how to read env variables into go program

## Youtube Data API
This project uses the Youtube Data API ([youtube v3 package]). If you haven't already, use the following link to [get started](https://developers.google.com/youtube/v3/getting-started).
Follow the instructions as best you can.

This project is using [Application Default Credentials](https://cloud.google.com/docs/authentication/application-default-credentials) via `context.Background()`.
This project also requires a service account token, so you will need to replace the generated `application_default_credential.json` with your service account's client key.


[youtube v3 package]: https://pkg.go.dev/google.golang.org/api@v0.218.0/youtube/v3