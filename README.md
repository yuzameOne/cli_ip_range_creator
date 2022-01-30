# cli_ip_range_creator - ip address range generator.

### what is he doing????

- the program reads a file with a range of ip addresses.
- inside the file are lines like this:
			31.22.48.0-31.22.50.15
			31.22.50.16-31.22.51.1
			31.22.51.2-31.22.52.39
			31.22.52.40-31.22.52.255
			31.22.53.0-31.22.53.254
			31.22.53.255-31.22.55.3
- the program reads the input file line by line and increments each line iteratively and writes all received lines to a single file.
### Example:
- we have the string 94.28.1.0-94.28.2.255 
- the result of running the program on this line
			94.28.1.0
			94.28.1.1
			94.28.1.2
			..............
			..............
			94.28.1.255
			94.28.2.0
			94.28.2.1
			94.28.2.3
			..............
			..............
			94.28.2.255
- and so on for each line in the file.

### how to use it ?
- compile or run the program as a script and pass the following arguments:
	 -  -ptf  (path to file) 
	 - -ptsf (path to  save file)

#### Example
`go run main.go -ptf [path_to_file]  -ptsf [path_where_save_file]` 

- -ptsf   - optional argument, if you do not specify it then the finished file will be saved next to the file main.go

## License

MIT
