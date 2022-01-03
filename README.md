<p align="center">
  <img src="./.github/assets/logo.svg" height="120px" />
</p>

> a CLI app can send pretty HTTP & API requests with TUI.

![demo](https://user-images.githubusercontent.com/64256993/145669325-d9f122d9-c562-417f-a223-a7f2b1c49adb.gif)

## Installation

### Using script

* Shell

```
curl -fsSL https://git.io/resto | bash
```

* PowerShell

```
iwr -useb https://git.io/resto-win | iex
```

**then restart your powershell**

### Go package manager

  ```bash
  go install github.com/abdfnx/resto@latest
  ```

### GitHub CLI
  
  ```bash
  gh extension install abdfnx/gh-resto
  ```

### Via Docker

  ```bash
  docker run -it restohq/resto <CMD>
  ```

  > full container:

  ```bash
  docker run -it restohq/resto-full
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
  
* Install binary app from script URL and run it.

  ```bash
  resto install https://yarnpkg.com/install.sh
  resto i https://get.docker.com
  ```

* Send a request from Restofile

  ```bash
  resto run

  # from path
  resto run --file ./examples/restofile/basic_request/Restofile
  ```
  
* Get the latest release/tag from repository
  ```bash
  resto get-latest abdfnx/resto
  
  # use another registry
  resto get-latest 23028539 --registry gitlab.com
  
  # with access token
  resto get-latest spittet/node-postgresql --registry bitbucket.org --token YOUR-ACCESS-TOKEN
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
  
3. `install` command flags

  ```
  -H, --hidden         hide the output
  -s, --shell string   shell to use default: bash
  ```
  
4. `run` command flags

  ```
  -a, --all           Show all response headers & status
  -f, --file string   Path to Restofile (Default: PATH/Restofile)
  ```
  
5. `get-latest` command flags

  ```
  -r, --registry string   The registry to use
  -t, --token string      The access token to use it the registry requires authentication
  ```

### Shortcuts

- <kbd>Ctrl+P</kbd>: **Open Resto Panel**
- <kbd>Ctrl+H</kbd>: **Open Help Guide**
- <kbd>Ctrl+E</kbd>: **Open Settings**
- <kbd>Ctrl+S</kbd>: **Save Request Body**
- <kbd>Ctrl+U</kbd>: **Update Your Resto**
- <kbd>Ctrl+Q</kbd>: **Quit**

## Documentation

Refer to our [**Wiki**](https://github.com/abdfnx/resto/wiki) for the documentation.
