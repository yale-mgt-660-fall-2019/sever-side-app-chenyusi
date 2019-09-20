## Sever side applications

This is a minimal example of an HTTP server (equivalently, a web server, or a server side app.)
Go is a great language in which to write HTTP servers. E.g. see
fans [here](https://blog.iron.io/how-we-went-from-30-servers-to-2-go/)
and [here](https://www.youtube.com/watch?v=bAQ9ShmXYLY)
and [here](https://talks.golang.org/2016/applicative.slide#1)
and [here](https://medium.com/@kevalpatel2106/why-should-you-learn-go-f607681fad65).
There are many more...that's just what jumped out at my searching
around in a few minutes.

One of the nicest aspects of writing servers in Go is that everything
you need is included in [Go's standard library](https://golang.org/pkg/).
That is, it's "batteries included"&mdash;you don't have to learn another
framework in order to be productive.

### How to complete this assignment

Completing this assignment is easy. You need to edit `main.go` to make
it respond to HTTP GET requests at `/nickname` with your class nickname.
E.g. my nickname is "bald-chicken" and so my webserver should respond
at `/nickname` with "bald-chicken". It also needs to respond at the `/`
URL, but the content isn't important.

Once your app is working, you should deploy it to
[Heroku](https://heroku.com) (or some other
place if you're a ninja and prefer to host it elsewhere). To do that,
you'll need to sign up for a free Heroku account. Then you'll need
to create a Heroku app. From my computer I would do that with a command
like this, assuming I am in the directory where this file resides:

```
go mod init
heroku login
heroku create
git push heroku master
```

That will push your git master branch up to Heroku. The `go mod init` is
used to create a `go.mod` file, which is how Heroku knows that we're using
Go. Otherwise it's not needed. When you have a more complicated app that
file will list your "dependencies"&mdash;other code that needs to be imported.

### Testing (grading) your app

You can test your app at
[https://grading.656.mba/new/server-side-apps-1](https://grading.656.mba/new/server-side-apps-1).

You can see my completed homework running
at [http://server-side-app.solutions.656.mba/](http://server-side-app.solutions.656.mba/)

## Tips

- Make sure you commit your work using git and push it to GitHub
- Deploy your app to Heroku
- Submit both your GitHub repo URL and your app (Heroku) URL via
  the class website. If you don't do so, or you do so inaccurately,
  I won't be able to give you a grade that reflects your true
  awesomeness.
