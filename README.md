# Resto

> CLI app can send pretty HTTP & API requests from your terminal with GUI.

![demo](https://user-images.githubusercontent.com/64256993/145669325-d9f122d9-c562-417f-a223-a7f2b1c49adb.gif)

## Installation

1. Install it from script

  * Shell

  ```
  curl -fsSL https://git.io/resto | bash
  ```

  * PowerShell

  ```
  iwr https://git.io/resto-win | iex
  ```

2. Go package manager

  ```bash
  go install github.com/abdfnx/resto@latest
  ```

3. GitHub CLI
  
  ```bash
  gh extension install abdfnx/gh-resto
  ```

## Usage

* Open Resto UI

  ```bash
  resto
  ```

* Send a request to a URL

  ```bash
  resto get https://api.github.com
  ```

* Send a request to a URL and use resto editor

  ```bash
  resto post https://localhost:3000/v1/login --content-type json --editor
  ```

* Read Body from stdin

  ```bash
  cat schema.graphql | resto post https://api.spacex.land/graphql --content-type graphql --body-stdin
  ```

* Use Authentecation with Basic Auth or Bearer Token

  ```bash
  # Bearer Token
  resto delete https://api.secman.dev/api/logins/13 --content-type json --token TOKEN
  
  # Basic Auth
  resto delete https://api.secman.dev/api/logins/13 --content-type json --username USERNAME --password PASSWORD
  ```

* Save response to a file

  ```bash
  resto get http://localhost:3333/api/v1/hello --save response.json
  ```

### Flags

1. `GET` & `HEAD` flags

  ```
  -H, --headers           Just show the response headers
  -j, --just-body         Just show the response body
  -p, --password string   The password to use for basic authentication
  -s, --save string       Save the response body to a file
  -t, --token string      The bearer token to use for authentication
  -u, --username string   The username to use for basic authentication
  ```

2. `POST`, `PUT`, `PATCH`, `DELETE` flags

  ```
  -b, --body string           The body of the request
  -i, --body-stdin            Read the body from stdin
  -c, --content-type string   The content type of the body
  -e, --editor                Open the editor to edit the body
  -H, --headers               Just show the response headers
  -j, --just-body             Just show the response body
  -p, --password string       The password to use for basic authentication
  -s, --save string           Save the response to a file
  -t, --token string          The bearer token to use for authentication
  -u, --username string       The username to use for basic authentication
  ```

## Documentation

Refer to [**resto website**](https://resto.deno.dev) for the documentation. Or you can check out the [**Wiki**](https://github.com/abdfnx/resto/wiki).
