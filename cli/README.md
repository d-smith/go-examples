Some examples of the basic cli in action

<pre>
macla961071:cli > go run cli.go --help bar
bar bar barness
macla961071:cli > go run cli.go --help foo
foo blah blah blah flibby dibby do
macla961071:cli > go run cli.go --version
1.0.0
macla961071:cli > go run cli.go foo
2015/10/03 02:31:28 Doing some foo...
2015/10/03 02:31:28 Ok, done with the foo
macla961071:cli > go run cli.go bar
2015/10/03 02:31:32 Doing some bar...
2015/10/03 02:31:32 Ok, done with the bar
macla961071:cli > go run cli.go huh
usage: cli [--version] [--help] <command> [<args>]

Available commands are:
    bar    when you want bar
    foo    do some foo

macla961071:cli > 
</pre>