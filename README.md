# Tail4Windows

##A simple implementation of the TAIL Linux command to be used by Windows users

####SYNTAX tail [OPTION] [FILE]

Implemented options are:

-f --follow - Shows the last part of a file while it grows 

-n --lines NUM - Shows the last NUM number of lines of the file

-v --version - Displays the version of tail.exe

-h --help - Displays this HELP information

#####Ex. tail -n 10 system.log

#####Displays the last 10 lines of the system.log file 

###Build

Make sure that you have a valid GOLANG installation and execute the ***build.cmd*** file. This will result in compilation and generation of tail.exe file.
Once the ***tail.exe*** has been generated you can add it to the PATH and use it exactly like you do in Linux environment.

Currently only a part of the options have been implemented since some of the options do not make sens to exist in the Windows enviornment.

####Enjoy!