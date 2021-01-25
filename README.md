# cli_ip_range_creator - ip address range generator

### what is he doing????

- takes a range of IP addresses like this..... 94.28.104.0-94.28.111.255
- take the number 104 compare with the number 111, 0 and 255 compare and rotate in a loop
- and we get a file like this :
- 5.100.67.0   
- 5.100.67.255
- 5.100.70.0
- 5.100.70.255
- 5.100.73.0
- 5.100.73.255

### how to use it ?

- compile the file :  go build main.go
- run the program  :  ./main  [path_to_file] [path_where_save_file]
- path_to_file required argument, without it the program will exit with an error
- path_where_save_file optional argument, if you do not specify it then the finished file will be saved next to the file ./main

## example of starting the program :

- clone github  repository https://github.com/yuzameOne/cli_ip_range_creator.git
- compile the file :  go build main.go
- run the program : ./main example.txt 