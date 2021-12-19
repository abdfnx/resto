## **Resto Run** command

> The `resto run` command is used to run requests from `Restofile`

example of `Restofile`

```restofile
request {
   method "POST"
   url "https://api.spacex.land/graphql"
   contentType "application/graphql"
}

body {
   openBodyEditor "true"
   # if `openBodyEditor` prop is false or not set
   readFrom "examples/spacex.gql"
}

auth {
   type "bearer"
   token: "MY_TOKEN"
   # or basic auth
   type "basic"
   username: "USERNAME"
   password: "P@$$w0rd"
   # to use from env variable
   password "env:MY_PASSWORD"
}
```

examples:

```bash
resto run

# from file
resto run --file examples/basic_request/Restofile
```
