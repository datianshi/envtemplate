# envtemplate

## Purpose
Execute a golang template by injecting the environment variables

## How to

```
go get -u github.com/datianshi/envtemplate
envtemplate -file=[a file path]
```

## Example
In [spec/sample](spec/sample):

```
host={{.SMTP_HOST}}
username={{.SMTP_USERNAME}}
```

```
export SMTP_HOST=smtp.gmail.com
export SMTP_USERNAME=user@gmail.com
envtemplate -file=spec/sample
```

produces:

```
host=smtp.gmail.com
username=user@gmail.com
```
