# email

The email package provides the `Sender` interface as an abstraction on a email
sending service. It requires a generic `Email` struct to represent an email. It
is up to the implementation to convert this to its specific email format.

Current implementation are a mock implementation and an AWS SES implementation.

## SESSender

`SESSender` is an `Sender` implementation that uses AWS SES to send the emails.
The AWS account has to have permission to do the `ses:SendEmail` action. The
SESSender requires an object which implements the SenderSESAPI, which is
implemented by the `ses.SES` object. Use `NewSESSender` to create a new
`SESSender`. See [ses_test.go](ses_test.go) for an example.

## Mock

`MockSender` is an implementation of the `Sender` interface using [testify
mocks](https://godoc.org/github.com/stretchr/testify/mock) created by
[mockery](https://github.com/vektra/mockery). Use `./generate_mocks.sh` to
update the mock structs when changing the `Sender` or `SenderSESAPI` interface.
