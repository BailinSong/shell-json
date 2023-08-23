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
   `./shell-json <keyPath[ ...[keyPath]]>`
   ```sh
   echo '{"key": "value", "nested": {"subkey": "subvalue"},"key4":["value4","value5","value6","value7"]}' | ./shell-json key nested.subkey key4.[0]
   ```


6. Check the formatted output.
   ```sh
   value subvalue    value4
   ```
# Shell-JSON

这是一个用 Go 编写的简单命令行工具，用于从 JSON 数据中提取指定键值。

## 特点

- 从 JSON 数据中提取指定键值

- 通过标准输入（stdin）输入 JSON 数据

- 根据提供的键路径格式化输出

- 易于使用命令行参数

## 使用方法

1.安装 Go（如果尚未安装）并设置 Go 环境。

2.克隆版本库：

   ```sh
   
   git clone https://github.com/BailinSong/shell-json.git
   
   ```

3.导航至项目目录

   ```sh
   
   cd shell-json
   
   ```

4.构建项目：

   ```sh
   
   go build
   
   ```

5.通过 stdin 提供 JSON 数据并指定关键路径作为命令行参数，运行该工具。例如

`./shell-json <keyPath[ ...[keyPath]]>`"。

   ```sh
   
   echo '{"key"："value", "nested"：{"subkey"："subvalue"},"key4":["value4","value5","value6","value7"]}'| ./shell-json key nested.subkey key4.[0]
   
   ```

6.检查格式化输出。

   ```sh
   
   value subvalue value4
   
   ```


## GitHub Repository

[GitHub Repository](https://github.com/BailinSong/shell-json)



