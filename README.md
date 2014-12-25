encore-cmd is meant to explore a script-based configuration management tool.

It uses a cmdfile akin to Dockerfile syntax. However, it only provides a limited functionality since we want to separate configuration management from system administration tasks.

encore-cmd follows the Unix philosophy of building single-purpose tools.

Guidelines:

- This program is meant to promote interface, not abstraction
- As such, my choice of shell is Bash
- Bash is the JavaScript of system programming (https://github.com/progrium/bashstyle)
- If you do not like Bash, a separate community of cmdfile needs to fill that void
- Embrace script because script serves as an interface (not abstraction)
- Code as data (script serves as data in the form of cmdfile)
- Simple syntax (like Dockerfile)
- Uses a push model (test cmdfile on local computer, then push it to remote machines via SSH)
- Cmdfile tasks alter machine state (not ad hoc query commands)
- Keep encore-cmd as small as possible. Extraneous tasks should be relegated as modules
- Imperative programming (code as code) complements declarative programming (code as data)
- And lastly, remember my 12-Rule App Manifesto at ITJUMPSTART.NET
- This program has been tested under Turnkey Linux 13.0 Core / Debian 7.2 Wheezy

Why script?

Short answer: user experience.
Long answer: user experience plus standardization of shell scripting in the form of cmdfile (just like Dockerfile).

If you go the route of abstraction (like popular CM tools today), you are forever chasing a moving target (AWS API, differences between DEB and RPM, APIs among RDBMS, NoSQL and a myriad host of software).

encore-cmd does not aim to replace your favorite shell. Rather, it adds error handling line by line from your cmdfile so it would abort at the first occurrence of an error.

Cmdfile
-------

Cmdfile is the only argument required for now
Cmdfile takes inspiration from Dockerfile
Cmdfile tasks must be sequential (no loops or conditionals)
Cmdfile tasks are Bash commands and external programs
Cmdfile is for humans, not machines
Cmdfile is simple and intuitive (no YAML)

Limitations of cmdfile
----------------------

It is not a shell (so no variable declaration and substitution)
No pipe commands
No backslash (commands must be put on each line)
No &&
No cd (Instead, use GO chdir directoryname)

encore-cmd uses the excellent Go runtime and even borrows some syntax from Golang os/exec package.

To run, type encore-cmd /path/of/cmdfile

ToDo
----

- Variable substitution in cmdfile (using Mustache template for example)
- Modules (a la Webmin)
- Comprehensive test suite
- SSH key management on master node
- Seamless file transfer from master to children nodes (using scp)
- Tight integration between master node and encore-cmd on children nodes
- Documentation

Your feedback is important and very much welcome.

License
-------

MIT
