# Never trust your terminal: Don't use plain Git / whatever to review patches

tl;dr: Git / cat / less don't display control characters which makes it
easy to inject (evil) code.

```bash
$ git clone https://github.com/leonklingele/how-to-git-evil
$ cd how-to-git-evil
$ git log -U --reverse
..
diff --git a/evil b/evil
new file mode 100644
index 0000000..c320815
--- /dev/null
+++ b/evil
@@ -0,0 +1 @@
+console.log("Good");
..
```

However, executing `evil` doesn't produce the expected result:

```bash
$ node evil
EVIL
```

To see all control characters:

```bash
$ hexdump -c evil
0000000   c   o   n   s   o   l   e   .   l   o   g   (   "   E   V   I
0000010   L   "   )   ;   /   /  \b  \b  \b  \b  \b  \b  \b  \b  \b  \b
0000020  \b  \b  \b  \b  \b  \b  \b  \b  \b  \b  \b  \b   c   o   n   s
0000030   o   l   e   .   l   o   g   (   "   G   o   o   d   "   )   ;
0000040  \n
0000041
```

Even `cat` and `less` fall for it:
```bash
$ less evil
console.log("Good");
$ cat evil
console.log("Good");//
```

This is _crazy_ :o! :/ Never trust your terminal output.

The best and only workaround I've found so far: Pipe everything to cat -v
```bash
$ git log -U --reverse | cat -v | less
..
diff --git a/evil b/evil
new file mode 100644
index 0000000..c320815
--- /dev/null
+++ b/evil
@@ -0,0 +1 @@
+console.log("EVIL");//^H^H^H^H^H^H^H^H^H^H^H^H^H^H^H^H^H^H^H^H^H^Hconsole.log("Good");
..
```

## Build instructions

`$ go run makeevil.go`
