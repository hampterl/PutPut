# PutPut Multitool

A command-line multitool written in **Go**.

> **Project status:** This tool is still under development.

---

## Table of Contents

* [Overview](#overview)
* [Features](#features)
* [Requirements](#requirements)
* [Installation](#installation)
* [Usage](#usage)

    * [Encoding / Decoding](#encoding--decoding)
    * [Encryption / Decryption (AES-256)](#encryption--decryption-aes-256)
    * 
* [Warnings & Disclaimer](#warnings--disclaimer)
* [Author](#author)
* [License](#license)

---

## Overview

**PutPut Multitool** is a terminal based utility that provides simple encoding, decoding, encryption,
and decryption features for text and files.

The tool runs entirely in the command line (CMD, PowerShell, Bash, etc.) and is designed mainly for **learning and educational purposes**.

---

## Features

* Encode / decode text

    * Base64
    * Hex
* Encrypt / decrypt text using AES-256
+ Hashing with SHA-256

---

## Requirements

* **Go (Golang)** installed (only required if building from source)
* A terminal (CMD, PowerShell, Bash, etc.)

---

## Installation

### Option 1: Download Release (Recommended)

Download the latest release from GitHub:

[https://github.com/hampterl/PutPut/releases](https://github.com/hampterl/PutPut/releases)

> Note: If no release is available yet, build the tool from source.

---

### Option 2: Build from Source

1. Download the repository as a `.zip` file and extract it
2. Open a terminal (CMD or PowerShell on Windows)
3. Navigate to the project directory:

```bash
cd "path/to/PutPut/cmd"
```

4. Run the tool:

```bash
go run main.go
```

or

```bash
go run .
```

---

## Usage

After starting the tool, a menu is displayed with numbered options.

**Only enter the number**, not the brackets.

If you enter a wrong option, type `0` to return to the main menu.

---

## Encoding / Decoding

### Menu Options

```
|1| Encode
|2| Encrypt/Decrypt
|3| Hashing
|0| Return to main menu
```

### Example (Base64 Encoding)

### Menu Options

```
|1| Encode text with Base64
|2| Encode file with Base64

|3| Decode Base64 text
|4| Decode file with Base64

|0| Return to main menu
```

### Text Encryption Example

```
> 1
Text to encode: Hello World!
Encoded text: SGVsbG8gV29ybGQh
Save txt into file? (y/n): n
NO file saved!
Press enter to return to main menu...
```

You may optionally save the result to a file.

---

## Encryption / Decryption (AES-256)

### Menu Options

```
|1| Encrypt text with AES-256
|2| Encrypt file with AES-256

|3| Decrypt AES-256 text
|4| Decrypt file with AES-256

|0| Return to main menu
```

### Text Encryption Example

```
> 1
AES Key (16, 24, 32 bytes): 1234567890123456
Text to encrypt: Hello World!
Encrypted text (hex): 2173213iuaisd01872e0d712312edboasd78
Save txt into file? (y/n): n
NO file saved!
Press enter to return to main menu...
```

> The AES key **must be exactly 16, 24, or 32 bytes long**.

---

## File Encryption / Decryption

### How It Works

* You must provide:

    * With base64:
    * Input file path (absolute or relative) **including file name and extension**
    * Output path **including file name and extension**

    * With Aes:
    * Input file path (absolute or relative) **including file name and extension**
    * Output path **including file name and extension**
    * AES key (16 / 24 / 32 bytes)

* Encrypted files are saved with `.enc` appended

* The original file is **not overwritten automatically**

---

### File Encryption Example

```
> 3
Filepath: encryptThisFile.txt
Save location: encryptThisFile.txt
AES Key (16, 24, 32 bytes): 1234567890123456
File encrypted! SAVE YOUR KEY!
Press enter to return to main menu...
```

> **IMPORTANT:**
>
> * Always include the file name **and extension**
> * Otherwise, you may encrypt a directory or create an invalid output file

---

## Warnings & Disclaimer

**READ CAREFULLY**

* This tool **will encrypt files you choose to encrypt**
* If you lose your key, **decryption is impossible**
* Never share your encryption key
* Always create backups before encrypting files

**USE AT YOUR OWN RISK**

---

## Legal Disclaimer

* This project is for **educational purposes only**
* **DO NOT** use this tool to encrypt sensitive or critical data
* **The developer is NOT responsible** for data loss, damage, or misuse

See the license for more information.

---

## Author

**hampterl**
GitHub: [https://github.com/hampterl](https://github.com/hampterl)

---

## License

This project is licensed under the **MIT License**.

See the [LICENSE](LICENSE) file for details.
