[![GitHub version](https://img.shields.io/github/release/jeffalyanak/check_namecheap.svg)](https://github.com/jeffalyanak/check_namecheap/releases/latest)
[![License](https://img.shields.io/github/license/jeffalyanak/check_namecheap.svg)](https://github.com/jeffalyanak/check_namecheap/blob/master/LICENSE.txt)
[![Donate](https://img.shields.io/badge/donate--green.svg)](https://jeff.alyanak.ca/donate)
[![Matrix](https://img.shields.io/matrix/check_namecheap:social.rights.ninja.svg)](https://matrix.to/#/#check_namecheap:social.rights.ninja)

# Namecheap Domain Expiry Checker

Icinga/Nagios plugin, checks the domain expiry status using the namecheap API.

User configurable `warning` and `critical` levels

## Installation and requirements

* Golang 1.13.8


## Usage

```bash
Usage of ./check_namecheap:
  Required:
  -apiuser string
        API Username
  -clientip string
        Client IP
  -domain string
        domain to check
  -key string
        API Key
  -username string
        Username

  Optional:
  -crit int
        days until critical (default 7)
  -warn int
        days until warning (default 15)
```

## License

Namecheap Domain Expiry Checker is licensed under the terms of the GNU General Public License Version 3.