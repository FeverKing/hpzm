# hpzm - a high precision zone mapper





## What is this tool for?

This repo is based on `github.com/Ullaakut/nmap/v3` , it's mainly a console application for generating nmap fingerprints on specific devices and managing them. By using this tool, you can also merge your customize fingerprints into your local `nmap-os-db` and accessing it with your nmap.



## Why Should I Use This Tool?

As the nmap official site said in [Dealing with Misidentified and Unidentified Hosts]("https://nmap.org/book/osdetect-unidentified.html#osdetect-wrong") :

> ### Modifying the `nmap-os-db` Database Yourself
>
>
>
> People often ask about integrating a fingerprint themselves rather than (or in addition to) submitting it to Nmap.Org. While we don't offer detailed instructions or scripts for this, it is certainly possible after you become intimately familiar with [the section called “Understanding an Nmap Fingerprint”](https://nmap.org/book/osdetect-fingerprint-format.html). I hope this is useful for your purposes, but there is no need to send your own reference fingerprint creations to us. We can only integrate raw subject fingerprint submissions from the web form.

However, for some reasons people need to add specific fingerprint themself and submitting and "sit back & relax"  isn't a quit delightful process. For those people, you should consider using this tool to manage your own fingerprints.



## How Do I use?



### Install



#### Build From Source

If you are reading this section, I believe you can fully understand how to build a go project, so it won't be necessary to waste both of our time.



#### Click & Run

See the big `Releases` button on the right? Click it, choose files matching your OS and platform on the top, download it and there you go.



### Usage



1. run the application you built or downloaded.

2. type `help` for built-in hints.

   ```
   hpzm v0.0.1 - a High Precision Zone Mapper
   
   type 'help' for a list of commands
   
   
   a High Precision Zone Mapper
   
   Commands:
   =========
     addfp     scan a known target and add its fingerprint to the database
     clear     clear the screen
     deletefp  delete a fingerprint from the database
     detect    detect the alive hosts
     exit      exit the shell
     help      use 'help [command]' for command help
     listfp    list all fingerprints in the database
     merge     Merge custom fingerprints to the database
     version   print the version of the application
   
   Sub Commands:
   =============
   
   detect:
     os  detect the os of the target
   
   listfp:
     n  list all fingerprints in nmap database
   ```

3. Scan fingerprint example: `addfp -T 192.168.3.20 -t 3 -v SONY -m "55 TV"`

4. Merge customize fingerprint into nmap:

   ```bash
   merge
   # if failed, find your nmap-os-db file location and retry
   merge -N "YOUR_NMAP_OS_DB_LOCATION"
   ```



## Improve

Yes, current version may contains many bugs and has many place can be improved, if you has any suggestion, feel free to open an issue or make pull request.



------

## Contacts

iamfeverking@gmail.com