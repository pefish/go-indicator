# go-indicator

[![view examples](https://img.shields.io/badge/learn%20by-examples-0C8EC5.svg?style=for-the-badge&logo=go)](https://github.com/pefish/go-indicator)

go-indicator

## Quick start

```go
func main() {
    result := NewIndicator().RSI([]float64{
        34220.4,
        34349.2,
        34352.5,
        34195.4,
        34142.6,
        34068.9,
    }, 3)
    fmt.Println(result)
}

```

## Document

[Doc](https://godoc.org/github.com/pefish/go-indicator)

## Security Vulnerabilities

If you discover a security vulnerability, please send an e-mail to [pefish@qq.com](mailto:pefish@qq.com). All security vulnerabilities will be promptly addressed.

## License

This project is licensed under the [Apache License](LICENSE).
