# go-pattern-sum

go cli for basic logs extractions summary intended to replace grep/awk/sed/wc/echo using golang regexp.

* Read stdin
* search for given pattern
* if pattern found,
* if value is a number, you can
  * output value
  * print sum of value
  * output some stats
  * append with a tag

## Example

```Shell

For basic postfix logs, containing lines like
Feb  5 22:06:57 myhost postfix/qmgr[1451]: 7C7E620FC7A: from=<incognitouser1>, size=1326, nrcpt=45 (queue active)

if you want to sum nrcpt 

$ ./go-pattern-sum.exe   -h 
Usage of go-pattern-sum.exe:
  -P string
        pattern with numeric value to sum (default ", nrcpt=(?P<value>[0-9]+) ")
  -p    only print values
  -s    Show sum count/min/max/avg instead of only sum
  -t string
        tag to add after printing sum value

$ ./go-pattern-sum.exe  -s -t "file test.txt" < test.txt
1823 53/1/50/34 file test.txt

1823 is the sum for the 53 values founds , min is 1 max 50, avg 34 for default pattern ", nrcpt=(?P<value>[0-9]+) "

```
