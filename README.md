# Go CLI Booking Application

This is a Go Command Line Interface Conference Application for booking tickets to the "Go Conference." It allows users to book tickets for the conference, sends them confirmation emails, and sends ticket notifications via SMS. The CLI application uses environment variables for configuration and handles concurrent requests using goroutines and a wait group.

## Table of Contents

- [Overview](#overview)
- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Usage](#usage)
- [Application Structure](#application-structure)
- [License](#license)

## Overview

The "Go Command Line Interface Conference Booking Application" is a simple command-line program that provides the following functionality:

1. Welcomes users to the Go Conference.
2. Collects user information, including first name, last name, phone number, email, and the number of tickets they want to book.
3. Validates user input for correctness.
4. Books tickets if the input is valid and updates the remaining ticket count.
5. Sends a confirmation email to the user.
6. Sends a ticket notification via SMS.
7. Keeps track of all bookings.

## Prerequisites

Before you can use this application, make sure you have the following prerequisites:

- Go installed on your system.
- Required environment variables properly configured (see the `config.LoadEnvVariables` function).
- SMTP email server and Termii API credentials (for sending emails and SMS).

## Installation

To install and run the Go Command Line Interface Conference Booking Application, follow these steps:

1. Clone the repository:

   ```bash
   git clone <repository_url>
   ```

2. Navigate to the project directory:

   ```bash
   cd <project_directory>
   ```

3. Build the application:

   ```bash
   go build
   ```

4. Run the application:

   ```bash
   ./<application_name>
   ```

## Usage

1. When you run the application, it will welcome you to the Go Command Line Interface Conference and display the number of available tickets.

2. Enter your information when prompted, including your first name, last name, phone number, email, and the number of tickets you want to book.

3. The application will validate your input. If it's valid, it will book the tickets, send you a confirmation email, and notify you via SMS.

4. If your input is not valid, the application will display error messages, and you can try again.

5. The application keeps track of all bookings, and you can view the list of first names of users who have booked tickets.

## Application Structure

The application consists of the following main components:

- `main.go`: The entry point of the application that handles user input, validation, booking, and communication.

- `config`: A package for loading environment variables.

- `helper`: A package containing helper functions for input validation.

- `mail.v2`: A third-party library for sending emails.

The application uses goroutines and a wait group (`sync.WaitGroup`) to handle concurrent tasks such as sending emails and SMS notifications.

## License

This Go Conference Booking Application is licensed under the MIT License.

Feel free to use and modify this application for your needs and add it to your portfolio.

---

_Note: Make sure to configure the necessary environment variables before running the application for proper functionality._
