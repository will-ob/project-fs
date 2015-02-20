
Probably want to make a debian package for installation

http://askubuntu.com/questions/4983/what-are-ppas-and-how-do-i-use-them
https://wiki.debian.org/IntroDebianPackaging
http://stackoverflow.com/questions/7542430/inotify-and-bash
https://wiki.debian.org/HowToPackageForDebian
http://superuser.com/questions/181517/how-to-execute-a-command-whenever-a-file-changes

There will be a cli tool for accessing / managing projects

- also used to sym-link to data directory
- data directory - should be a hidden folder in the user's home directory

Need to install daemon

- Inotifywait for changes to those files
  - but how to intercept reads..

Read could be on-the-fly from go, but write could be to a temporary file that is watched
