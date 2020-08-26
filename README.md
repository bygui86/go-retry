
# go-retry

Sample project to explore retry mechanism in Golang

## run

```shell script
go run main.go
```

## links

- https://github.com/avast/retry-go

### alternatives

- giantswarm/retry-go - slightly complicated interface.
- sethgrid/pester - only http retry for http calls with retries and backoff
- cenkalti/backoff - Go port of the exponential backoff algorithm from Google's HTTP Client Library for Java. Really complicated interface.
- rafaeljesus/retry-go - looks good, slightly similar as this package, don't have 'simple' Retry method
- matryer/try - very popular package, nonintuitive interface (for me)
