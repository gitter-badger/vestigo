# Vestigo - A Vestige Of Echo's URL Router

## Abstract

[Echo][echo-main] has a very fast URL router.  This repo is a vestige of just the URL Router,
broken out into a stand alone module.  There is such an abundance of parts and pieces that can be fit
together for go web services, it seems like a shame to have a very fast URL router require the use
of one framework, and one context model.  This library aims to give the world a fast, and featureful
URL router that can stand on it's own, without being forced into a particular web framework.

## Design

1. Radix Tree Based
2. Attach URL Parameters into Request (PAT style) instead of context

## Examples

```go

router := vestigo.NewRouter()

router.Get("/welcome", GetWelcomeHandler)
router.Post("/welcome", PostWelcomeHandler)
router.Post("/welcome/:action", PostActionWelcomeHandler)
router.Post("/welcome/:action$UUID$", PostUUIDActionWelcomeHandler)

func PostActionWelcomeHandler(w http.ResponseWriter, r *http.Request) {
    action := vestigo.Param(r, "action") // url params live in the request
}

```

## Licensing

Portions of the URL Router were taken from [Echo][echo-main] and are covered under their [License][echo-main-license].

The rest of the implementation is covered under The MIT License covered under this [License][vestigo-main-license].

# Contributing

If you wish to contribute, please fork this repository, submit an issue, or pull request with your suggestions.  
Please use gofmt and golint before trying to contribute.


[echo-main]: https://github.com/labstack/echo
[echo-main]: https://github.com/labstack/echo/blob/master/LICENSE
[vestigo-main]: https://github.com/husobee/vestigo/blob/master/LICENSE