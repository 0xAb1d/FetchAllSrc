# FetchAllSrc

## FetchAllSrc, also known as `fas`, is a streamlined command-line tool designed for fetching web resources or source code directly from provided URLs and saves them with clear, web-structure-mirrored names in a specified directory.

![Screenshot (1307)](https://github.com/user-attachments/assets/c3d0e991-54cf-458a-a6e5-8b034cdfd652)


<br>

## Features

- **⚡️ Concurrent Downloads**: FetchAllSrc uses goroutines to handle multiple downloads simultaneously, speeding up the process and efficiently using network resources.
- **🛠️ Error Handling and Logging**: The tool logs all download activities, capturing Failures and Success details in log files to aid troubleshooting.
- **📊 Progress Tracking**: Displays real-time download progress, giving users visibility into the status of each download task.
- **🔒 Safe Filename Conversion**: Automatically sanitizes URLs into valid filenames by replacing unsupported characters, ensuring compatibility across different operating systems.

<br>

## Installation

To install FetchAllSrc as `fas`, use the following commands:

```bash
go install github.com/0xAb1d/FetchAllSrc@latest
```
```bash
mv ~/go/bin/FetchAllSrc ~/go/bin/fas
```
Or,
```bash
mv $GOPATH/bin/FetchAllSrc $GOPATH/bin/fas
```

> Ensure that `$GOPATH/bin` is in your system's PATH to run `fas` from any terminal window.
 `export GOPATH=$HOME/go`
 `export PATH=$PATH:$GOPATH/bin`
 `export PATH=$PATH:~/go/bin`
 `source ~/.zshrc`
```
echo 'export GOPATH=$HOME/go' >> ~/.zshrc
echo 'export PATH=$PATH:$GOPATH/bin' >> ~/.zshrc
source ~/.zshrc
```
<br>

#### Building Locally

If you prefer to build the tool locally, follow these steps:

```bash
git clone https://github.com/0xAb1d/FetchAllSrc.git
cd FetchAllSrc
go build -o fas
```

> This command compiles the program and creates an executable file named `fas` in your current directory.

To move the fas executable to the Go binary directory, you can use the following commands:
```
mv fas $(go env GOPATH)/bin
```
or,
```
mv fas ~/go/bin
```
<br>


#### Binary Download

You can download the latest compiled binary from the [Releases](https://github.com/0xAb1d/FetchAllSrc/releases) page and run it directly:

```bash
# After downloading the binary
chmod +x fas && sudo mv fas /usr/local/bin/
```

> This command will make the binary executable, move it to `/usr/local/bin/` for global use, and immediately allow users to run `fas`.

<br>

## Usage

Run the tool with the following syntax:

```bash
fas -i input.txt -o outputDir
```
```bash
fas -i input.txt -o downloads
```

- `-i input.txt`: Specifies the input file containing list of URLs to download.
- `-o outputDir`: Specifies the output directory where the downloaded files will be saved.
  
<br>

## Examples

### Input File

Example content for `input.txt`:

```
https://example.com/assets/package/script.js
https://example.com/src/data/contents/module.json
https://example.com/backup/installation/database/db.config
```

### Output Files

Example content for `downloads` directory:

```
downloads/example.com_assets_package_script.js
downloads/example.com_src_data_contents_module.json
downloads/example.com_backup_installation_database_db.config
```

### Log File

Example entries in `fetchallsrc.log`:

```
INFO: Successfully downloaded https://example.com/assets/package/script.js
ERROR: Failed to download https://example.com/src/data/contents/module.json: 404 Not Found
INFO: Successfully downloaded https://example.com/backup/installation/database/db.config
```

### NotFound.txt

If some files fail to download, `NotFound.txt` will list:

```
https://example.com/src/data/contents/module.json
```

## Sample screenshots

![Screenshot (1305)](https://github.com/user-attachments/assets/40814e85-5898-4735-b89e-84f535762fa7)

![Screenshot (1306)](https://github.com/user-attachments/assets/09ce4ede-c292-4d6b-8326-37b8b01460d6)
> Slash `/` converted to underscore `_`

![Screenshot (1303)](https://github.com/user-attachments/assets/7095d276-a35e-4c14-87cc-977d185313e2)

![Screenshot (1304)](https://github.com/user-attachments/assets/6be3a59f-bce0-486b-92f1-9654f1f5ac7d)
> `fetchallsrc.log` Log file

![Screenshot (1301)](https://github.com/user-attachments/assets/ed1826b9-ecdc-4e79-a07d-222285f2664d)
> `NotFound.txt` file

![Screenshot (1302)](https://github.com/user-attachments/assets/b89f96f1-3ee0-4db6-b59e-d50f0094bf44)
> Downloaded source code file

## License

FetchAllSrc is provided under the MIT License. See the [LICENSE](https://github.com/0xAb1d/FetchAllSrc/blob/main/LICENSE) file for more details.

