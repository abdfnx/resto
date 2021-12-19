## **Resto Install** command

> The `resto install` command is used to install binary app from script URL and run it.

this command is alternative to `curl -sL URL | bash`

```bash
resto i https://deno.land/x/install/install.sh

deno -V
```

### docs

#### flags

```bash
-s, --shell string   shell to use default: bash
```

examples:

```bash
resto i https://deno.land/x/install/install.sh --shell sh
```
