# DDoS Project

[![Go Report Card](https://goreportcard.com/badge/github.com/your-username/ddos-project)](https://goreportcard.com/report/github.com/your-username/ddos-project)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](https://opensource.org/licenses/MIT)

A DDoS project implemented in Go to learn about concurrency and network programming.

## Table of Contents
- [Introduction](#introduction)
- [Features](#features)
- [Getting Started](#getting-started)
- [Usage](#usage)
- [Contributing](#contributing)
- [Disclaimer](#disclaimer)
- [License](#license)

## Introduction
This project is developed to explore the concepts of concurrent programming and network communication using Go. It aims to simulate a DDoS (Distributed Denial of Service) attack, helping you understand how such attacks work and how to defend against them.

## Features
- Simple and intuitive Go codebase for learning purposes.
- Utilizes goroutines and channels to achieve concurrency.
- Supports custom target URLs or IP addresses.
- Basic logging and reporting functionalities.

## Getting Started
### Prerequisites
- Go programming language (version X.X.X)
- Additional dependencies (if any)

### Installation
1. Clone the repository:
   ```shell
   git clone https://github.com/your-username/ddos-project.git
   ```
1. Change into the project directory:
   ```shell
   cd DDOS-FLOOD
   ```
1. Build the project:
   ```shell
   go build .
   ```
1. Clone the repository:
   ```shell
   ./ddos-flood
   ```
## Usage
1. Run the project by executing the binary generated during the installation process.
2. Enter the `URL`, `Number of Worker` and `Attack Durantion` on the terminal
3. Monitor the console output and log files for attack statistics and progress.
4. Terminate the program to stop the DDoS attack or wait until the end.

## Contributing
Contributions are welcome! If you'd like to contribute to this project, please follow these steps:
1. Fork this repository.
2. Create a new branch: `git checkout -b feature/your-feature-name`.
3. Make your changes and commit them: `git commit -m 'Add some feature'`.
4. Push to the branch: `git push origin feature/your-feature-name`.
5. Open a pull request.

Please ensure that your code adheres to the existing code style and includes appropriate tests.

## Disclaimer
This code is provided for educational purposes only. It is essential to understand that launching a DDoS attack or engaging in any malicious activities is illegal and unethical. The purpose of this project is solely for learning about concurrent programming and network communication in a controlled and responsible environment.

The author of this project take no responsibility for any misuse or damage caused by the code in this repository. Please use this code responsibly and respect the laws and regulations of your jurisdiction.

## License
This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for more details.
