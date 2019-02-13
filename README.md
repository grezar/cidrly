# cidrly(1)

## Install
```
$ go get github.com/grezar/cidrly
```

## Usage

```
$ aws ec2 describe-subnets --filters "Name=vpc-id,Values=vpc-xxxxxxxx" --region ap-northeast-1 | cidrly
.
├── 10.0.0.0/24(subnet-a)
├── 10.0.2.0/28(subnet-b)
└── 10.0.2.16/28(subnet-c)
```
