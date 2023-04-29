### Cyclone's Truecrypt / Veracrypt Container Truncater
- Tool to process Truecrypt / Veracrypt container files to be ran with hashcat
- Tested on debian linux & Windows 11
- Run tool in directory where Truecrypt *.tc or Veracrypt *.vc container file is located
- Once file is selected from menu, tool will truncate container and save a new file to "truncate_filename"
- "truncate_filename" can now be ran with hashcat using the appropriate mode
- - example: ./veracrypt_container_truncator.bin
- - 1 ) truecrypt.tc
- - outputs new file "truncate_truecrypt.tc"
- - hashcat -m 6211 -a 0 truncate_truecrypt.tc cyclone_hk_v2.txt -r cyclone_250.rule
- Tool will not overwrite any existing files

Compile from source code info:
- https://github.com/cyclone-github/scripts/blob/main/intro_to_go.txt
