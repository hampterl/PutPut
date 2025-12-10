# PutPut Multitool
Multitool in golang by hampterl

## Getting Started

This is a command line interface tool. You have to run it in a terminal. E.g. cmd or powershell.

- You can download the latest release from the releases page.  
- Since there isn't a release out yet, you have to build it yourself.  

- On Windows download the .zip file and extract it.  
- Then you go to the directory where the main is and copy the filepath.  
- Open a command prompt and type cd "filpath", press enter.  
- Now you can run the tool with "go run main.go" or "go run ."  
- You have to have golang installed. If you choose to run it this way.  

## Download
Download latest release -> https://github.com/hampterl/PutPut/releases

## Features
- En/decoding text (base64, hex)  
- En/decrypting text (AES-256)  
- En/decrypting files (AES-256)  

After running the tool, you can see your options next to the logo.
Since the tool is still in development, there are only 2 options to choose from.
You have to type in the number that is in the brackets next to the option. These: |"Here the number"|, but only the number!
Not the brackets. If you mistake the number, you can always type "0" to return to the main menu.

## Instructions 
### En/Decoding
For En/Decoding text:  
- Type in "1" to encode text.  
- Type in "2" to decode text.  

After that you have to type in the text you want to encode/decode.
It will go through the process and tell you when to type something in and return the result.
You can save the result to a file.  
But you can look at the process here, this is an example:

|1|: Encode text with Base64  
|2|: Decode Base64 to text  
|0|: Return to main menu\n>  
\> 1  
\> Text to encode: Hello World!  
\> Encoded text: SGVsbG8gV29ybGQh  
\> Save txt into file? (y/n): n  
\> NO file saved!  
Press enter to return to main menu...  

### En/Decrypting
For En/Decrypting, text and files:  
- Type in "1" to encrypt.  
- Type in "2" to decrypt.  
- Type in "3" to encrypt a file.  
- Type in "4" to decrypt a file.  

For text:  
After that you have to type in the text you want to encode/decode.
It will go through the process and tell you when to type something in and return the result.
You can save the result to a file.  
But you can look at the process here, this is an example:

|1|: Encrypt text with aes256  
|2|: Decrypt aes256 to text  
|3|: Encrypt File with aes256  
|4|: Decrypt File with aes256  
|0|: Return to main menu\n>  
\> 1  
Aes Key(16, 24, 32 bytes): 1234567890123456  
Text to encode: Hello World!  
Encrypted text (hex): 2173213iuaisd01872e0d712312edboasd78   <- This is just an example of how it would look.  
Save txt into file? (y/n): n  
NO file saved!  
Press enter to return to main menu...  


For file encrypting:  
You have to first type in the filepath WITH the ending e.g. .txt, .exe... Either you enter the absolute path or the relative path.
Then you have to enter the outputpath (where it should save) and a key with either 16, 24 or 32 bytes.  

DISCLAIMER AND WARNING:  
This WILL encrypt the file you chose! IF you lose your key and overwrite your original, you will NOT be able to decrypt the file!  
Never share your key with anyone! Make a backup before encrypting! Use at own Risk!  

Your encrypted file will be saved where you told it to, with .enc at the end. So it won't automatically overwrite the original.  
Here an example of how it would look:  
|1|: Encrypt text with aes256  
|2|: Decrypt aes256 to text  
|3|: Encrypt File with aes256  
|4|: Decrypt File with aes256  
|0|: Return to main menu\n>
\> 3  

Instructions in README.md!  
DISCLAIMER AND WARNING:  
This WILL encrypt the file you chose! IF you lose your key and overwrite your original, you will NOT be able to decrypt the file!  
Never share your key with anyone! Make a backup before encrypting! Use at own Risk!  
  
Filepath: encryptThisFile.txt  <- if it's in the same directory as the main, this is enough  
Save location: encryptThisFile.txt  <- if you want to save it in the same directory, write in the same as in Filepath  
Aes Key(16, 24, 32 bytes): 1234567890123456  
File encrypted! SAVE YOUR KEY!  
Press enter to return to main menu...  

The encrypted file will appear in the save location directory.  
*IMPORTANT!* You have to add the name and ending of the file in both fields, 
or it will encrypt the directory, or save it without a name.

## Legal Disclaimer
USE FOR EDUCATIONAL PURPOSES ONLY!

NEVER USE THIS TO ENCRYPT SENSITIVE DATA!

THE DEVELOPER OF THIS TOOL IS NOT RESPONSIBLE FOR ANY DAMAGE BY MISSUSE! SEE LICENSE FOR MORE INFORMATION!

## Author
hampterl on github

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
