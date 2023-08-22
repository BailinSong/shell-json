# Shell-JSON

This is a simple command-line tool written in Go for extracting specified key values from JSON data.

## Features

- Extract specific key values from JSON data
- Input JSON data via standard input (stdin)
- Format output based on provided key path
- Easy to use with command line arguments

## Usage

1. Install Go (if not already installed) and set up your Go environment.

2. Clone the repository:

   ```sh
   git clone https://github.com/BailinSong/shell-json.git
   ```

3. Navigate to the project directory:

   ```sh
   cd shell-json
   ```

4. Build the project:

   ```sh
   go build
   ```

5. Run the tool by providing JSON data via stdin and specifying the key path as command line arguments. For example:

   ```sh
   echo '{"key": "value", "nested": {"subkey": "subvalue"},"key4":["value4","value5","value6","value7"]}' | ./shell-json <keyPath[ ...[keyPath]]>
   ./shell-json key nested.subkey key4.[0]
   ```


6. Check the formatted output.
  ```sh
   value subvalue    value4
   ```

## GitHub Repository

[GitHub Repository](https://github.com/BailinSong/shell-json)



