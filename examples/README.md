# Go WhatsApp Library Examples

This directory contains various examples demonstrating how to use the go-whatsapp library.

## Examples Overview

- **simple/**: Basic connection and version checking example
- **login/**: Shows how to log in using QR code
- **restoreSession/**: Demonstrates session restoration
- **receiveMessages/**: Example of receiving and handling messages
- **sendTextMessages/**: Shows how to send text messages
- **sendImage/**: Example of sending image messages
- **echo/**: Bot that echoes received messages
- **chatHistory/**: Retrieve chat history
- **checkNewVersion/**: Check for WhatsApp server version updates

## Running Examples

Each example is a standalone Go program. To run any example:

1. Navigate to the example directory
2. Initialize a Go module: `go mod init example`
3. Add the library dependency: `go mod edit -replace github.com/Rhymen/go-whatsapp=../..`
4. Install dependencies: `go get github.com/Rhymen/go-whatsapp`
5. Run the example: `go run .`

## Note

Most examples require an active internet connection and the ability to connect to WhatsApp's servers. Some examples require scanning a QR code with your WhatsApp mobile app for authentication.