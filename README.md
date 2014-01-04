SocialVibes (Compute Engine Project)
=====================

[![GoDoc](https://godoc.org/github.com/FraBle/SocialVibes-ComputeEngine?status.png)](https://godoc.org/github.com/FraBle/SocialVibes-ComputeEngine)
[![Google+ SocialVibes](http://b.repl.ca/v1/Google+-SocialVibes-brightgreen.png)] (https://plus.google.com/118015654972563976379/)

- [SocialVibes Website](http://gcdc2013-socialvibes.appspot.com "Social Vibes")
- [Google Cloud Developer Challenge 2013](http://www.google.com/events/gcdc2013/ "Google Cloud Developer Challenge 2013")
- [SocialVibes App Engine Project](https://github.com/FraBle/SocialVibes-AppEngine "SocialVibes App Engine Project")

![SocialVibes Cover](https://raw.github.com/FraBle/SocialVibes-AppEngine/master/static/images/ui/cover.png "Social Vibes Cover")

## What is Social Vibes?
Social Vibes visualizes any public Google+ event. During an event, you can start the Social Vibes slideshow to show the latest event pictures combined with music of your choice.

## How to use it?
You can select one of the public events you've created with your Google+ account or any public event by pasting the URL. For the music, you can select one of our suggestions or paste any YouTube playlist URL. After selecting an event and a playlist, you can start the slideshow.

[![How to use SocialVibes](https://raw.github.com/FraBle/SocialVibes-AppEngine/master/static/images/ui/youtube.png)](http://www.youtube.com/watch?v=KCPR8WcQYIY)

## How to set it up?

##### You need a running SocialVibes Compute Engine instance!

- You need Google Go (http://golang.org/doc/install) with all environment variables set
- You need Task Queues set up (see https://github.com/FraBle/SocialVibes-AppEngine#how-to-set-it-up)
- Checkout the repository
- You need a socialvibes.toml file under socialvibes/socialvibes.toml, e.g.:

```toml
[taskqueue]
clientID = "<<Your Client ID for web application>>"
clientSecret = "<<Your Client secret for web application>>"
```
_You can find the Client ID and secret in the Google Cloud Console (APIs & auth > Credentials):_
https://cloud.google.com/console
- You need to `go get` the external packages and install dependencies
- You need to compile with `go install`

## Used packages

- [github.com/gorilla/rpc/v2](http://www.gorillatoolkit.org/pkg/rpc/v2 "github.com/gorilla/rpc/v2")
- [github.com/gorilla/rpc/v2/json](http://www.gorillatoolkit.org/pkg/rpc/v2/json "github.com/gorilla/rpc/v2/json")
- [code.google.com/p/google-api-go-client/taskqueue/v1beta2](https://code.google.com/p/google-api-go-client/ "code.google.com/p/google-api-go-client/taskqueue/v1beta2")
- [code.google.com/p/goauth2/oauth](https://code.google.com/p/goauth2/ "code.google.com/p/goauth2/oauth")
- [github.com/stvp/go-toml-config](https://github.com/stvp/go-toml-config "github.com/stvp/go-toml-config")
- [github.com/PuerkitoBio/goquery](https://github.com/PuerkitoBio/goquery "github.com/PuerkitoBio/goquery")
- [github.com/sourcegraph/webloop](https://github.com/sourcegraph/webloop "github.com/sourcegraph/webloop")
