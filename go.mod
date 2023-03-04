module github.com/ngamux/ngamux-example

go 1.20

require (
	github.com/ngamux/ctx v0.0.0-20230304032138-79dba02181ff
	github.com/ngamux/middleware v0.0.0-20230224184957-9d8a61cbd79b
	github.com/ngamux/ngamux v1.2.0
)

replace (
	github.com/ngamux/ctx => ../ctx
	github.com/ngamux/middleware => ../middleware
	github.com/ngamux/ngamux => ../ngamux
)
