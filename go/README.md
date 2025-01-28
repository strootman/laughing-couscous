# Golang
This project uses the Cobra library to create a CLI.
This will be a constant WIP, and is mostly an excuse to practice golang features, but have a convenient reason to keep things segregated.

The APIs I'm currently trying out via the official Google API libraries are;

* [Youtube Data API] => [Youtube Data API v3 package]

`[ ]` Determine how to read env variables into go program

## Authentication
```
UPDATE: I'm punting on the (YT) Google APIs, because there is something about the auth system I'm not grokking. I get the same errors from curl as I do here, and the point is to get golang practice.
        So, switching gears.
```
This project is using [Application Default Credentials](https://cloud.google.com/docs/authentication/application-default-credentials) via `context.Background()`.
This project also requires a service account token, so you will need to replace the generated `application_default_credential.json` with your service account's client key.


[Youtube Data API v3 package]: https://pkg.go.dev/google.golang.org/api@v0.218.0/youtube/v3
[Youtube Data Api]: https://developers.google.com/youtube/v3/getting-started