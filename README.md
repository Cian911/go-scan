### go-scan

A dump, yet concurrent port scanner written in go.

#### Build

To build, run the following.

```bash
go build -o goscan main.go
```

#### Usage

To use this, you can run compiled binary, passing in any optional flags you desire.

```bash
-address      The address you wish to scan
-range        The range of ports you wish to scan. Default is 100.
-buffers      The number of channel buffers you wish to use, the more can speed up the scan. Default is 100.
```

```bash
./goscan -address ciangallagher.net -range 80 -buffers 150
```
