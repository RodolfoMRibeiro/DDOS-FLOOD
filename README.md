# DDoS Project

[![Go Report Card](https://goreportcard.com/badge/github.com/your-username/ddos-project)](https://goreportcard.com/report/github.com/your-username/ddos-project)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

DDoS Project is an educational tool implemented in Go, aimed at understanding the underlying concepts of concurrency and network programming through the simulation of a Distributed Denial of Service (DDoS) attack.

> :warning: **Disclaimer**: This project is for educational purposes only. Please do not use it for any illegal or malicious activities. See [Disclaimer](#disclaimer) section for more details.

## Table of Contents
- [Introduction](#introduction)
- [Features](#features)
- [Getting Started](#getting-started)
    - [Prerequisites](#prerequisites)
    - [Installation](#installation)
- [Usage](#usage)
- [Contributing](#contributing)
- [Disclaimer](#disclaimer)
- [License](#license)

## Introduction
This project serves as an educational tool for learning concurrent programming and network communication using Go. It simulates a DDoS attack, providing insights into how such attacks are orchestrated and how one can develop strategies to defend against them.

## Features
- Clean and intuitive Go codebase, ideal for educational purposes.
- Efficiently leverages goroutines and channels for concurrency.
- Ability to specify custom target URLs.
- Offers basic logging and real-time reporting functionalities.
- Enhanced control over threading.
- Powerful and scalable flooding mechanism.
### Default Header Settings:
 - :white_check_mark: Random user-agents
 - :white_check_mark: Random Accept-All
 - :white_check_mark: Random Referrer

## Getting Started
### Prerequisites
- You should have Go programming language installed (version 1.19 or higher).

### Installation
1. Clone the repository:
   ```shell
   git clone https://github.com/RodolfoMRibeiro/DDOS-FLOOD.git
   ```
1. Change into the project directory:
   ```shell
   cd DDOS-FLOOD
   ```
1. Build the project:
   ```shell
   go build .
   ```
1. Run the binary file:
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
