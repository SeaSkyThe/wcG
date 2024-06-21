# wcG - a simple wc clone

In this project I'm just doing a clone of wc Unix tool.

If you are not familiar with wc, wc is a tool that serves to basically print newline, word, and byte counts for files.

If you want to know more about it, just read the manual by running the following command in your terminal:
```bash
man wc
```
```
# How to use

Our version can be compiled with

```bash
go build -o wcg
```

Then we will have an executable to use wherever we want to, use it directly, add to your PATH, anything.

To run it we will have some flags:
- -c: will output the number of bytes in a file.
- -l: will output the number of lines in a file.
- -w: will output the number of words in a file.
- -m: will output the number of characters in a file.
- no flags: will output the equivalent of using -c, -l and -w flags.

Usage example:
```bash
wcg -c test.txt
```
