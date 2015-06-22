Project FS
=========

Markdown todo-lists, sync'd online.

Make command-line todo lists available in multiple locations.

Install
------------

```
git clone git@github.com:will-ob/project-fs.git && cd project-fs
make install
```

Uninstall
-------------

```
make uninstall
```

Commands
---------------

| Command  | Description  |
|---|---|
| `todo`  | Find or create todo list for current directory.|
| `todo <name>` | Open named todo list. |
| `todo config` | Configure the project-fs daemon. |
| `todo rm` | Delete todo list in current directory. |
| `todo rm <name>` | Delete named todo list. |


Development
--------------------

### Environment

| Var  | Required  | Description | Example |
|---|:-:|---|---|
| `PROJECT_API_URL`     | x | Url of project api | `http://localhost:3333/projects/`  |
| `UNSAFE_TLS`     |   | Ignore certificate errors (eg. from self-signed dev certs) | `true`  |


### Paths

| Path  | Contents |
|---|---|
| `/opt/project-fs`  | Binary executable |
| `/home/<user>/.project-fs` | User configuration & file cache |
| `/usr/local/bin/todo` | `ln -s => /opt/project-fs/todo`|
| `/var/log/project-fs.log` | Log file of daemon (filesystem backend). (Note: `todo` cmd logs to stdout and stderr, not this file.)|

License
-------------

Copyright 2015 Will O'Brien

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
