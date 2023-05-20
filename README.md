# comx
`comx` is a tool to compare the contents of two files and return the unique data


# Installation
```
go install github.com/SharokhAtaie/comx@latest
```

# Usage

```
sub1.txt:
sub.google.com
sub1.google.com
sub2.google.com
sub3.google.com
sub4.google.com


sub2.txt:
sub1.google.com
sub2.google.com
sub3.google.com
sub4.google.com
sub5.google.com
```
```
➜  comx -f1 sub1.txt -f2 sub2.txt               

 ██████╗ ██████╗ ███╗   ███╗██╗  ██╗
██╔════╝██╔═══██╗████╗ ████║╚██╗██╔╝
██║     ██║   ██║██╔████╔██║ ╚███╔╝ 
██║     ██║   ██║██║╚██╔╝██║ ██╔██╗ 
╚██████╗╚██████╔╝██║ ╚═╝ ██║██╔╝ ██╗
 ╚═════╝ ╚═════╝ ╚═╝     ╚═╝╚═╝  ╚═╝
                                Created by WatchDogs :)

[sub1.txt] Unique Content:
sub.google.com

[sub2.txt] Unique Content:
sub5.google.com
```

### Video
<a href="https://asciinema.org/a/6ZDqpKtGuHs8ZDtDVzcUqHIqv" target="_blank"><img src="https://asciinema.org/a/6ZDqpKtGuHs8ZDtDVzcUqHIqv.svg" /></a>


## Speacial Thanks to [ProjectDiscovery Team](https://github.com/projectdiscovery) for the great libraries.
