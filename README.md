# EC2 Filter

simple utility for listing EC2 instances with tag filter


### Usage

```sh
$ ./ec2-filter -h

NAME:
   ec2-filter - filter ec2 instances with tag

USAGE:
   ec2-filter [global options] command [command options] [arguments...]

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --region value, -r value  (default: "us-east-1")
   --key value, -k value
   --value value, -v value
   --help, -h                show help (default: false)
```

To list all instanaces

```sh
$ ./ec2-filter

Name   State    InstaceType  PrivateIP      InstanceId
demo1  Running  t3.micro     192.168.1.1    i-0e13dc5c7f811c72a
demo2  Running  t3.micro     192.168.1.2    i-03373efd80c8147c6
```

To list all instanaces with tag "Name", and its value is "demo1"

```sh
$ ./ec2-filter -k Name -v demo1

Name   State    InstaceType  PrivateIP      InstanceId
demo1  Running  t3.micro     192.168.1.1    i-0e13dc5c7f811c72a
```
