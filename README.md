encore-cmd is meant to explore a script-like and SSH-based configuration management (CM) tool using a subset of Bash-like syntax. It aims to meet the common 80% of use cases that are relevant in CM using your favorite shell (in this case, it is Bash). However, if you want to use other shell, encore-cmd lets you do that by calling your script in a cmdfile.

It uses a cmdfile akin to Dockerfile syntax. However, it only provides a limited functionality since we want to separate configuration management from system administration tasks.

encore-cmd follows the Unix philosophy of building single-purpose tools.

Guidelines:

- Code as data (script serves as data in the form of cmdfile)
- Simple syntax (like Dockerfile)
- Uses a push model (test cmdfile on local computer, then push it to remote machines via SSH)
- Cmdfile tasks alter machine state (not ad hoc query commands)
- Keep encore-cmd as small as possible. Extraneous tasks should be relegated as modules
- And lastly, remember my 12-Rule App Manifesto at ITJUMPSTART.NET
- This program has been tested under Turnkey Linux 13.0 Core / Debian 7.2 Wheezy

Why script-like?

Answer: Syntax. Because developers are familiar with the syntax of script (command followed by optional arguments), a script-based CM would be more intuitive rather than YAML-based or any other syntax for that matter.

encore-cmd does not aim to replace your favorite shell. Rather, it adds error handling line by line from your cmdfile so it would abort at the first occurrence of an error. 

Moreover, it lets you separate which tasks can be automated on your local computer or on the remote machines you want to configure. For example, you may download a hefty file to your local computer and then simply push it to your remote machines. That way, you do not have to use curl or wget on each of those remote computers.

Why Bash?

Because it is the JavaScript of system programming (https://github.com/progrium/bashstyle). Bash is like SQL. It lets you interface with the machine, rather than build a leaky abstraction like ORM.

Why SSH?

Because it is secure and easy to setup.

Why yet another tool?

- you want a simple and developer-friendly configuration management tool
- you love the simplicity and deployment of Go binaries
- you love the Gopher way of building single-purpose tools

Cmdfile
-------

- Cmdfile is the only argument required for now
- Cmdfile takes inspiration from Dockerfile
- Cmdfile tasks must be sequential (no loops or conditionals)
- Cmdfile is simple and intuitive (no YAML)

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

Run encore-cmd at multiple command prompt. The above example runs SSH commands in series. To do this in parallel, you have to put the following in a separate cmdfile, then run it.

- #cmdlocal2.txt
- RUN scp cmdremote.txt root@192.168.1.103:/home
- RUN ssh root@192.168.1.103 ./encore-cmd /home/cmdremote.txt

To run it in parallel, open two different prompts, then type

- encore-cmd cmdlocal.txt

On another command prompt, type

- encore-cmd cmdlocal2.txt


Tested
------

encore-cmd can be used to configure Debian and Ubuntu servers from Windows (tested using Git Bash from GitHub). However, the output is not color-coded unlike with Linux-based distros. Needs testing from the community for RPM-based servers. My setup is just Virtual Box and Turnkey Linux (no need for Vagrant which is just another abstraction). Remember: use interface, not abstraction.


ToDo
----

- Variable substitution in cmdfile (using Mustache template for example)
- Modules (a la Webmin but CLI-based)
- Wiki documentation

Your feedback is important and very much welcome.

License
-------

MIT
