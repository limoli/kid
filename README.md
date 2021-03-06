# KID

[![GitHub Actions](https://github.com/limoli/kid/workflows/Go/badge.svg?branch=main)](https://github.com/limoli/kid/actions?query=workflow%3AGo+branch%3Amain)
[![Go Report Card](https://goreportcard.com/badge/limoli/kid)](https://goreportcard.com/report/limoli/kid)
[![GoDoc](https://godoc.org/github.com/limoli/kid?status.svg)](https://godoc.org/github.com/limoli/kid)
[![GitHub license](https://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/limoli/kid/master/LICENSE)

## About

KID number (actually KID, customer identification) is a number used in bill payment in Norway to identify the payment 
(regardless of who pays the bill). It can be from 2 and up to 25 digits long. The last digit of the KID number is a 
check digit. This check digit is calculated from the MOD10 or MOD11 algorithm.

### Algorithms

- [Luhn MOD10 and MOD11](https://en.wikipedia.org/wiki/Luhn_algorithm)

## Usage

#### Validate a KID number
##### Using MOD10
```go
    kid.Verify(code, new(Luhn10))
```

##### Using MOD11
```go
    kid.Verify(code, new(Luhn11))
```

#### Generate a KID number
##### Using MOD10
```go
    luhn10 := kid.Luhn10{}
    checkDigit, err := luhn10.Generate(code)
    // e.g. 7, nil
```

##### Using MOD11
```go
    luhn11 := kid.Luhn11{}
    checkDigit, err := luhn11.Generate(code)
    // e.g. '-', nil
```

#### Handle errors

- `ErrorInvalidCodeLength`: KID has an invalid length  
- `ErrorInvalidNumber`: KID has an invalid character (not a digit)

## License

Released under the [MIT License](https://github.com/limoli/kid/blob/main/LICENSE).