# :book: Newton [![GoDoc][doc-img]][doc] [![Build Status][ci-img]][ci] [![Coverage Status][cov-img]][cov] [![Report Status][rep-img]][rep] [![Go Version][go-img]][go]
`Newton` is a cloud-native serverless application made for those who love books.

## Installation

`git clone https://github.com/maestre3d/newton.git`

Note that _Newton_ only supports the two most recent minor versions of Go.

### Used Infrastructure and 3rd-party software
- Hashicorp Terraform

**Amazon Web Services (AWS)**
- CLi v2
- Lambda
- DynamoDB (+ DAX)
- S3
- SNS
- SQS
- SES
- Cognito
- CloudFront
- ACM
- Route53
- CloudWatch
- CloudTrail
- X-Ray

See the [documentation][docs] and [FAQ](FAQ.md) for more details.

## Maintenance
This library is currently maintained by
- [maestre3d][maintainer]

## Development Status: Alpha

All APIs are under development, breaking changes will be made in the 0.x series
of releases. Users of semver-aware dependency management systems should pin
_newton_ to `^1`.

Released under the [MIT License](LICENSE).

[doc-img]: https://pkg.go.dev/badge/github.com/maestre3d/newton
[doc]: https://pkg.go.dev/github.com/maestre3d/newton
[docs]: https://github.com/maestre3d/newton/tree/master/docs
[ci-img]: https://github.com/maestre3d/newton/workflows/Go/badge.svg?branch=master
[ci]: https://github.com/maestre3d/newton/actions
[go-img]: https://img.shields.io/github/go-mod/go-version/maestre3d/newton?style=square
[go]: https://github.com/maestre3d/newton/blob/master/go.mod
[rep-img]: https://goreportcard.com/badge/github.com/maestre3d/newton
[rep]: https://goreportcard.com/report/github.com/maestre3d/newton
[cov-img]: https://codecov.io/gh/maestre3d/newton/branch/master/graph/badge.svg
[cov]: https://codecov.io/gh/maestre3d/newton
[maintainer]: https://github.com/maestre3d
