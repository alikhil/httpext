# httpext

[![CircleCI](https://circleci.com/gh/alikhil/httpext.svg?style=svg)](https://circleci.com/gh/alikhil/httpext)

It's simple package for making development of http server a bit easier by simplifying http handlers implementations.

For example simple route implemented with a help of `httpext`:

```go

func handleRestore(s *session, r *http.Request) (*httpext.ResponseError, interface{}) {
  var vars = mux.Vars(r)
  var imageKey, passed = vars["imageKey"]

  if !passed {
    return httpext.ValidationError(errors.New("imageKey is not passed")), nil
  }

  var err = s.ctx.ImageService.Restore(imageKey)
  if err != nil {
    if err == ImageCanNotBeRestoredError {
      return httpext.OtherError(err, http.StatusPreconditionFailed), nil
    } else {
        return httpext.InternalServerError(errors.Wrap(err, "failed to restore image")), nil
    }
  }

  return httpext.NoError(), "ok"
}

```

Which later is wrapped with middleware function which knows how to respond according to errors:

```go
var appRouter = mux.NewRouter()
appRouter.HandleFunc("/restore/{imageKey}", httpext.Wrap(handleRestore)).Methods("POST")
```

As a wrapper you can use built-in `httpext.Wrap` or implement your own according to your needs.
