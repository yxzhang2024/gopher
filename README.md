# Gopher Usage Guide

## Overview

This guide provides instructions on how to run the Gopher tool, a utility designed to analyze Go projects. It also covers the necessary steps to clean up any previous scan results. In the `basicRules`, we provide the CryDict Rules for golang basic cryptography library. If you want to try writing your own CryDict documentation, you can refer to `CryDict_tutorial/readme.md`. In the compressed file(script), we provide specific code for comparing with cryptogo testing. In the `test_code` folder, we provide the projects for testing.

## Installing Go Environment

Before you begin, make sure that the Go language environment is installed on your computer. If not, please follow these steps to install it:

1. **Visit the Go Official Download Page**:
   Visit https://golang.org/dl/ and download the latest stable release for your operating system.

2. **Install Go**:
   - For Linux users, extract the downloaded package and add the `go` directory to your system path (PATH) environment variable.
   - For Windows users, run the downloaded installer and make sure to check the option "Add Go to PATH" during installation.

3. **Verify Installation**:
   Open a terminal or command prompt window and type `go version` to check if Go has been correctly installed.

## Setting Up Workspace

Assuming you have installed Go and want to place the Gopher project within the standard Go workspace structure, usually located at `$HOME/go/src` (Linux) or `%USERPROFILE%\go\src` (Windows). Here are the setup steps:

1. **Create Go Workspace**:
   If you haven't created a Go workspace yet, manually create a directory named `go` and within it create a subdirectory named `src`. For example, on Linux you can use the following command:
   ```sh
   mkdir -p $HOME/go/src
   ```
   For Windows users, you can do this:
   ```cmd
   md %USERPROFILE%\go\src
   ```

2. **Clone or Copy the Gopher Project**:
   Place the Gopher project into the `src` directory so its path becomes `$HOME/go/src/gopher` (Linux) or `%USERPROFILE%\go\src\gopher` (Windows). 

## Setup

Before running Gopher, ensure that the system environment variable `GO111MODULE` is set to `on`.

Here’s how you can do it:

1. **On Unix systems (Linux):**
   - Open your terminal.
   - Run the following command:
     ```sh
     export GO111MODULE=on
     ```
   This will set the `GO111MODULE` environment variable for the current shell session only.

   If you want to make this setting permanent, add the above line to your shell's startup file (like `.bashrc`, `.bash_profile`, or `.zshrc` if you're using ZSH).

2. **On Windows:**
   - Open Command Prompt.
   - Use the `set` command:
     ```cmd
     set GO111MODULE=on
     ```
   This will be active only for the current Command Prompt window.

   To set it permanently, you might need to modify the system's environment variables through the Control Panel or System Properties.

Setting `GO111MODULE=on` tells Go to use modules for source management by default, which is useful when working on projects that use Go modules for dependency management.

## Running Gopher

To execute the Gopher tool, follow these steps:

1. **Navigate to the Gopher Directory**:
   ```sh
   cd $HOME/go/src/gopher
   ```

2. **Execute the Gopher Tool**:
   Run the Gopher executable with the target project directory as an argument. 
   Please note that some code repositories require downloading code from the network, and poor network conditions may cause delays. Use the appropriate command based on your operating system:
   ```sh
   # For Linux
   chmod 777 gopher
   ./gopher ./test_code/beego-2.2.0

   # For Windows
   ./gopher.exe .\test_code\beego-2.2.0
   ```

   Replace `./test_code/beego-2.2.0` with the path to the project directory you wish to analyze. The provided test code originates from the [Beego 2.2.0 repository](https://github.com/beego/beego/tree/v2.2.0). The detection process takes less than 1 minute. After running Gopher, you will receive an error alert about MD5 and an error alert indicating the skipping of TLS validation.

## Obtaining scanning results

You can obtain the scan results from the command line window. In addition, the scan results of gopher will be saved in the corresponding project's `scan_desults/results.txt`.





Note: If there are too many files in `$GOPATH/pkg`, the scan speed will significantly decrease on computers with mechanical hard drives. The scan speed might also drop on computers with solid-state drives.