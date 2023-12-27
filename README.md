# URL Shortener

A simple URL shortener service built with Go and the Gin web framework.

## Overview

This URL shortener provides a convenient way to shorten long URLs, making them easier to share. It includes features such as URL shortening and redirection.

## Features

- Shorten long URLs to compact, easy-to-share links.
- Redirect users to the original long URL when they access the shortened link.
- Automatic deletion of expired URLs.
- Simple and lightweight implementation using the Go programming language and Gin web framework.

## Getting Started

### Prerequisites

- Go installed on your machine.
- MongoDB or another database for storing URL mappings.

### Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/emeraldls/url-shortener.git

   cd url-shortener

   go get -u ./...

   go run main.go


