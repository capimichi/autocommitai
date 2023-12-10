# Auto Commit AI

This project is designed to automate the process of committing changes to a Git repository.

## Getting Started

Download the executable for your operating system from the [releases page]

## Usage

- Run the executable

```bash
./auto-commit-ai
```

### Parameters

The RootCommand function in the Command package defines the root command for the autocommitai CLI application. It has two parameters:

- ignore-untracked: This is a boolean flag. If set to true, the application will ignore untracked files when committing. The default value is false.

- default-choice: This is a string flag. It sets the default choice for the action to take when the application encounters a file. The default value is "".

You can use these parameters when running the autocommitai command. For example:

```bash
./autocommitai --ignore-untracked=true --default-choice="1"
```

## Parameters

## Structure

The directory structure of the project is as follows:

README.md: This is the markdown file that contains the documentation for the project.

go.mod and go.sum: These files are used by Go's dependency management system.

internal: This directory contains the internal packages of the project.

Command: This directory contains the RootCommand.go file which likely defines the root command for a CLI application.

Config: This directory contains the DefaultConfig.go file which likely contains default configuration settings.

Helper: This directory contains helper functions for the project. It includes AutoCommitHelper.go, BardHelper.go, GitHelper.go, and TextHelper.go.

Model: This directory contains the data models used in the project, which are CommitMessage.go and GitFile.go.

Service: This directory contains the AutoCommitAiService.go file which likely contains the main service logic for the Auto Commit AI.

main.go: This is the entry point of the Go application.
