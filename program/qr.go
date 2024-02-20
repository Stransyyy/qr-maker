package main

import (
	qr "github.com/skip2/go-qrcode"
)

func Generate(url string) ([]byte, error) {
	return qr.Encode(url, qr.Medium, 256)
}
