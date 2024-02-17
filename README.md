## About nixconfig-go and requirements

This little go script was developed to search for `.nix` files on GitHub, although it can be used to search
for any files.
It has one configuration, `config.json` and one output directory which can be configured.

```
{
    "token":"",
    "output_directory":"",
    "config_type": ""
}
```

**token:** GitHub token with repository permissions

**output_directory:** Directory to place the downloaded nix files

**config_type:** The nix configuration type, for example `shell.nix`, `configuration.nix`, etc...


## Running nixconfig-go
You can set the page number and the per-page amount for the download, for example:

`bin/nixconfig-go --page 1 --per-page 100`


## Checking the output
the downloaded files will be at `${OUTPUTDIR}/page-X`

A successful search + download should look something like this:

```
➜  nixconfig-go git:(main) ✗ ls -l nix-config/page-1 
total 2552
-rw-r--r--@ 1 patrik.koska  staff   2168 Feb 17 04:40 0-configuration.nix
-rw-r--r--@ 1 patrik.koska  staff   1093 Feb 17 04:41 0-hardware-configuration.nix
-rw-r--r--@ 1 patrik.koska  staff   1044 Feb 17 04:41 0-users.nix
-rw-r--r--@ 1 patrik.koska  staff   2167 Feb 17 04:40 1-configuration.nix
-rw-r--r--@ 1 patrik.koska  staff   1495 Feb 17 04:41 1-hardware-configuration.nix
-rw-r--r--@ 1 patrik.koska  staff   1194 Feb 17 04:41 1-users.nix
-rw-r--r--@ 1 patrik.koska  staff   4368 Feb 17 04:40 10-configuration.nix
-rw-r--r--@ 1 patrik.koska  staff   1021 Feb 17 04:41 10-hardware-configuration.nix
-rw-r--r--@ 1 patrik.koska  staff    903 Feb 17 04:41 10-users.nix
-rw-r--r--@ 1 patrik.koska  staff   1899 Feb 17 04:40 11-configuration.nix
-rw-r--r--@ 1 patrik.koska  staff   3177 Feb 17 04:41 11-hardware-configuration.nix
-rw-r--r--@ 1 patrik.koska  staff    453 Feb 17 04:41 11-users.nix
-rw-r--r--@ 1 patrik.koska  staff   2228 Feb 17 04:40 12-configuration.nix
```

## Setting up a development environment
If you have nix installed and flakes enabled, you can start a minimal environment with
`nix develop` from the repository root


```
➜  nixconfig-go git:(main) nix develop
[pkoska@nixos:~/nixconfig-go]$ 


[pkoska@nixos:~/nixconfig-go]$ which go
/nix/store/zg65r8ys8y5865lcwmmybrq5gn30n1az-go-1.21.6/bin/go

[pkoska@nixos:~/nixconfig-go]$ go version
go version go1.21.6 linux/amd64

```

**GL with hunting!**
