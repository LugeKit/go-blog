module github.com/k1/go-blog

go 1.17

replace (
	github.com/k1/go-blog/conf => /home/k1/Desktop/go-blog/conf
	github.com/k1/go-blog/middleware => /home/k1/Desktop/go-blog/middleware
	github.com/k1/go-blog/model => /home/k1/Desktop/go-blog/model
	github.com/k1/go-blog/pkg => /home/k1/Desktop/go-blog/pkg
	github.com/k1/go-blog/router => /home/k1/Desktop/go-blog/router
	github.com/k1/go-blog/runtime => /home/k1/Desktop/go-blog/runtime
)

require github.com/BurntSushi/toml v0.4.1
