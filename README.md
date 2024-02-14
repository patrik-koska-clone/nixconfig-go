This little go script was developed to search for .nix files on GitHub, although it can be used to search
for any files.
It has one configuration, config.json and one output directory which can be configured.

```
{
    "token":"",
    "output_directory":"",
    "config_type": ""
}
```

**token:** GitHub token with repository permissions

**output_directory:** directory to place the downloaded nix files

**config_type:** The nix configuration type, for example shell.nix, configuration.nix, etc...

A successful search + download should look something like this:

```
➜  nixconfig-go ls -l nix-config
total 320
-rw-r--r--@ 1 patrik.koska  staff  1672 Feb 14 12:30 0-configuration.nix
-rw-r--r--@ 1 patrik.koska  staff   479 Feb 14 12:33 0-hardware_configuration.nix
-rw-r--r--@ 1 patrik.koska  staff  1337 Feb 14 12:30 1-configuration.nix
-rw-r--r--@ 1 patrik.koska  staff    96 Feb 14 12:33 1-hardware_configuration.nix
-rw-r--r--@ 1 patrik.koska  staff   244 Feb 14 12:30 10-configuration.nix
-rw-r--r--@ 1 patrik.koska  staff   877 Feb 14 12:30 11-configuration.nix
-rw-r--r--@ 1 patrik.koska  staff  3767 Feb 14 12:30 12-configuration.nix
-rw-r--r--@ 1 patrik.koska  staff  2855 Feb 14 12:30 13-configuration.nix
-rw-r--r--@ 1 patrik.koska  staff   952 Feb 14 12:30 14-configuration.nix
-rw-r--r--@ 1 patrik.koska  staff  1522 Feb 14 12:30 15-configuration.nix
...
```

GL with hunting!