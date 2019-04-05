# Serverless-OSINT
Implemented OSINT features with Serverless Framework.

## Usage
```
$ git clone
$ make
$ sls deploy
```

## API
### whois
`/whois/{domain}`: Get whois information related to the domain

```sh
$ curl -X GET https://<Api Gateway>/whois/{domain}
```

### lookupHost
`lookupHost/{domain}`: Get IP address related to the domain

```sh
$ curl -X GET https://<Api Gateway>/lookupHost/{domain}
```

### lookupMX
`lookupMX/{domain}`: Get MX records related to the domain

```sh
$ curl -X GET https://<Api Gateway>/lookupMX/{domain}
```

### lookupAddr
`lookupAddr/{addr}`: Get host names related to IP address

```sh
$ curl -X GET https://<Api Gateway>/lookupAddr/{IP address}
```

### lookupNS
`lookupNS/{domain}`: Get name servers related to the domain

```sh
$ curl -X GET https://<Api Gateway>/lookupNS/{domain}
```

