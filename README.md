# mini-curl
Minimalistic client to request urls

The program can make GET or POST requests to a given url. 
On http code 200 the program exit code is 0. On any error the exit code is -1.

# Parameters

```
Usage: mini-curl [--username USERNAME] [--password PASSWORD] [--body BODY] URL

Positional arguments:
  URL                    url to request

Options:
  --username USERNAME, -u USERNAME
                         optional: when set http basic authorization is added to the request header
  --password PASSWORD, -p PASSWORD
                         optional: needed for http basic authorization
  --body BODY, -b BODY   optional: is send as post body
  --help, -h             display this help and exit
```

# Examples
Simple GET request
```bash
mini-curl https://google.com
```

Simple POST request
```bash
mini-curl --body "postbody" https://google.com
```