# filterfakenews
This program will set your /etc/hosts file to forward fake news sites in your browser.

What this does: When you try to load a url say 'addictinginfo.org' in your browser it will just send you to 'npr.org' (This will be configurable in a future release.)

Currently the ip in the code is hard coded to npr.org but you can update that yourself if you know go.  I will make this more user friendly in the near future.


Cred for the list:
http://www.latimes.com/nation/politics/trailguide/la-na-trailguide-updates-want-to-keep-fake-news-out-of-your-1479260297-htmlstory.html
## Build:

```sh
$ go build main.go
```

## Run after build:

```sh
$ sudo ./main links.yml
```

## for non programmers:
Click clone or download and download zip

unzip, open terminal - http://www.macworld.co.uk/feature/mac-software/get-more-out-of-os-x-terminal-3608274/

```sh
$ cd ~/Downloads/filterfakenews-master/
$ sudo ./main links.yml
```

