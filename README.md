# gosslmate

_The gosslmate module queries the certificate transparency logs from sslmate.com._

## Table of contents

1. [Introduction](#introduction)
1. [Getting started](#getting-started)
    1. [Prerequisites](#prerequisites)
    1. [Installation](#installation)
1. [How to use](#how-to-use)
1. [Missing features](#missing-features)
1. [Known bugs](#known-bugs)
1. [License](/LICENSE)

## Introduction

This module queries the certificate transparency logs from sslmate.com. You can read more about CT logs here https://certificate.transparency.dev/.

Currently only the free api is supported, which has a rate limit of 100 queries per hour.

## Getting started

### Prerequisites

- Golang installed
- Operatingsystem: Linux or Windows, not tested on mac
- IDE like VS Code, if you want to contribute or change the code

### Installation

The recommended way to use this module is using the go cli.

    go get github.com/IT-Administrators/gosslmate

## How to use

After importing the module you can use it the following way.
First you create a query object. This object is used to manage what you want to query. 

```Go
// Create a query object.
sslq = NewSslMateQuery("example.com")
```

The following parameters are true by default as show in the documentation on sslmate.com.

```Go
SearchSubDomains:         true,
ShowDnsNames:             true,
ShowIssuer:               true,
ShowRevocationInfo:       true,
ShowProblemReportingInfo: true,
ShowCertData:             true,
```

Every other parameter is false by default and must be activated before running the query.

You can enable or disable a query parameter by running:

```Go
// Disable query parameter.
// Objname.Prorpertyname = false
sslq.ShowDnsNames = false
```

To run the configured query you have to call the ```GetCtLogs```function, with the ```sslq``` parameter. Save the result to a variable to use it later on. 

```Go
// Run query and save response to json object.
qres = GetCtLogs(sslq)
```

The response of this api is always an array of the ```sslMate```struct, even if the initial response is a single json object. To access the different properties you have to use the reference of the object just like accessing any other value in an array.

```Go
// Access properties.
qres[0].Issuer.Webiste

// Result:
// https://sectigo.com/
```

For testing purposes and development you can find examples files in the examples folder.

## Missing features

There is one parameter missing from the offitial documentation, which is the ```after``` paramater. This parameter is used to: ``Return issuances that were discovered by SSLMate after the issuance with the specified ID.``.

## Known bugs

Currently the returned ```sslMate``` struct is fully exported and can be changed. Dont do it. 

In the future this will be changed to a read only struct with different getter funtions for each struct property.. 

## License

[MIT](./LICENSE)