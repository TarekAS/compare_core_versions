
# Usage

A script that periodically compares the core_version of two Horizon environments.

```
# Build Go program:
go build .

# Run:
./compare_core_versions --name1=prod --url1=https://horizon.stellar.org/ --name2=dev --url2=https://horizon-testnet.stellar.org/

# For all options:
./compare_core_versions --help
```

You can set the `--frequency` flag to change how often to check for the different versions (in seconds).
