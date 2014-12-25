encore-cmd is meant to explore a script-based configuration management tool.

It uses a cmdfile akin to Dockerfile syntax. However, it only provides a limited functionality since we want to separate configuration management from system administration tasks.

encore-cmd follows the Unix philosophy of building single-purpose tools.

Guidelines:

- This program is meant to promote interface, not abstraction
- As such, my choice of shell is Bash
- If you do not like Bash, a separate community of cmdfile need to fill that void
- Separate data from functions
- Embrace script because script serves as an interface (not abstraction)
- Code as data (script serves as data in the form of cmdfile)
- Simple syntax (like Dockerfile)
- Uses a push model (test cmdfile on local computer, then push it to remote machines via SSH)
- Cmdfile tasks alter machine state (not ad hoc query commands)
- Use existing tools (SSH, Bash). After all, Bash is the JavaScript of system programming
- Keep encore-cmd as small as possible. Extraneous tasks should be relegated as modules
- Imperative programming (code as code) complements declarative programming (code as data)
- And lastly, remember my 12-Rule App Manifesto at ITJUMPSTART.NET

You must have Go installed.

To run, type encore-cmd /path/of/cmdfile

Your feedback is important and very much welcome.

License: MIT
