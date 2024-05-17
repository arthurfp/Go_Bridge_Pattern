# Device Management System Using Design Patterns in Go

## Overview
This repository demonstrates the application of various design patterns in Go to create a robust device management system. The project highlights the implementation of user profiles, advanced device controls, and dynamic user interaction, making it ideal for environments where multiple device settings need to be managed efficiently.

## Pattern Description
The project utilizes multiple design patterns, including Singleton for device initialization, Bridge for separating device commands from the device types, and Observer for notifying changes in device states. These patterns enhance the modularity, scalability, and maintainability of the application.

## Project Structure
- **cmd/**: Contains the application entry point (`main.go`), demonstrating interactions with the device management system.
- **pkg/**
  - **device/**: Houses the device interfaces and their implementations.
  - **remote/**: Contains the implementations of different remote controls.
  - **user/**: Manages user profile settings and interactions, enabling personalized control settings for each user.

## User Profile Management
This feature allows the system to maintain individual user preferences for device settings like volume and channel. Each user's preferences are applied automatically when they interact with a device, enhancing the personalized experience.

## Getting Started

### Prerequisites
Ensure you have Go installed on your system. You can download it from [Go's official site](https://golang.org/dl/).

### Installation
Clone this repository to your local machine:
```bash
git clone https://github.com/arthurfp/Go_Bridge_Pattern.git
cd Go_Bridge_Pattern
```

### Running the Application
To run the application, execute:
```bash
go run cmd/main.go
```

### Running the Tests
To execute the tests and verify the functionality:
```bash
go test ./...
```

### Contributing
Contributions are welcome! Please feel free to submit pull requests or open issues to discuss proposed changes or enhancements.

### Author
Arthur Ferreira - github.com/arthurfp