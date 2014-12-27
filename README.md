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

If you go the route of abstraction, you are forever chasing a moving target (APIs of public cloud providers, differences between DEB and RPM, APIs among RDBMS, NoSQL and a myriad host of software).

encore-cmd does not aim to replace your favorite shell. Rather, it adds error handling line by line from your cmdfile so it would abort at the first occurrence of an error.

Cmdfile
-------

- Cmdfile is the only argument required for now
- Cmdfile takes inspiration from Dockerfile
- Cmdfile tasks must be sequential (no loops or conditionals)
- Cmdfile tasks are Bash commands and external programs
- Cmdfile is for humans, not machines
- Cmdfile is simple and intuitive (no YAML)

Limitations of cmdfile
----------------------

- It is not a shell (so no variable declaration and substitution)
- No pipe commands
- No backslash (commands must be put on each line)
- No &&
- No cd (Instead, use GO chdir directoryname)
- No eval

encore-cmd uses the excellent Go runtime and even borrows some syntax from Golang os/exec package.

Usage
-----

You must have Go installed.

go get github.com/ibmendoza/encore-cmd

To run, type encore-cmd /path/of/cmdfile

encore-cmd can work on your local computer or remote computer via SSH. For seamless remote operation, running an SSH agent is recommended. Of course, this assumes that the remote computer has encore-cmd already installed at the folder of your own choosing.

Example
-------

Assuming you have two cmdfiles on your local computer (cmdlocal.txt)

- RUN scp cmdremote.txt root@192.168.1.102:/home
- RUN ssh root@192.168.1.102 ./encore-cmd /home/cmdremote.txt

- RUN scp cmdremote.txt root@192.168.1.103:/home
- RUN ssh root@192.168.1.103 ./encore-cmd /home/cmdremote.txt


Here is other cmdfile (cmdremote.txt).

- RUN uptime


Like Fabric (but without Python)
--------------------------------

From your local computer, type

encore-cmd cmdlocal.txt


Parallel SSH execution
----------------------

Run encore-cmd at multiple command prompt

Tested
------

encore-cmd can also be used to manage Debian-based servers from Windows (tested using Git Bash from GitHub). However, the output is not color-coded unlike with Linux-based distros. Needs testing from the community for RPM-based servers. My setup is just Virtual Box and Turnkey Linux (no need for Vagrant which is just another abstraction). Remember: use interface, not abstraction.


ToDo
----

- Variable substitution in cmdfile (using Mustache template for example)
- Modules (a la Webmin but CLI-based)
- Wiki documentation

Your feedback is important and very much welcome.

License
-------

MIT
