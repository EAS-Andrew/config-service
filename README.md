# Config-Service

Config-Service is a Golang application designed to traverse specified directories, each corresponding to a unique application team, and to load YAML configuration files within these directories. It supports configurations for environments, deployment repositories, components, and artifact repositories, allowing for unified configuration availability across different operation types. The application is designed with extensibility, modularity, concurrency, and efficiency in mind.

## Overview

The application utilizes Go's strong concurrency features to process multiple directories in parallel, enhancing efficiency and performance. It adopts a modular architecture, enabling easy addition of new operation types or configuration schemas. The project is structured around key functionalities implemented in separate Go files: main.go for the entry point, scanner.go for directory traversal, parser.go for parsing configuration files, and processor.go with an interface and implementations for processing different kinds of configurations.

## Features

- Directory-Based Configuration Loading: Scans directories named in the format AL<5digit number> and loads all YAML configuration files within.
- Configuration Types Support: Handles various types of configurations such as Environment, DeploymentRepository, Component, and ArtefactRepository.
- Unified Configuration Availability: Makes all configurations from a single directory available to each type of operation the application supports.
- Extensibility and Modularity: Facilitates easy addition of new operation types or configuration schemas.
- Concurrency and Efficiency: Processes multiple directories in parallel, optimizing performance.
- Robust Error Handling and Validation: Manages issues like malformed YAML files and unsupported configuration schemas.
- Comprehensive Logging and Monitoring: Tracks the application's processing stages and performance.

## Getting started

### Requirements

- Go version 1.21.5 or newer
- Access to directories structured according to the specified format

### Quickstart

1. Clone the repository to your local machine.
2. Navigate to the cloned directory.
3. Run `go build .` to compile the application.
4. Execute the application with `./config-service`.

To customize the root directory for configuration file scanning, modify the `traverseDirectory` function call in `main.go`.

### License

Copyright (c) 2024. All rights reserved.

The project is proprietary and not open source.