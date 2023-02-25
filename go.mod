module github.com/ngamux/ngamux-example

go 1.20

require (
	github.com/ngamux/middleware v0.0.0-20230224184957-9d8a61cbd79b
	github.com/ngamux/ngamux v1.2.0
)

replace github.com/ngamux/ngamux => ../ngamux
replace github.com/ngamux/middleware => ../middleware
