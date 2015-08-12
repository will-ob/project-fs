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

Then:

```
todo config
```

Uninstall
-------------

```
make uninstall
```

Env Vars
----------

The `todo` commands use the following env vars:

| Env Var  | Description  |
|---|---|
| HOME  | Used to store user-specific program settings and data. |
| EDITOR  | Used to determine which program should be used for editing todo lists. |


Commands
---------------

| Command  | Description  |
|---|---|
| `todo`  | Open todo list for current directory. |
| `todo config` | Configure the project-fs daemon. |
| `todo init`  | Create a new todo list in the current directory. |
| `todo init <name> `| Link to existing todo list in the current directory. |
| `todo ls`  | List available todos. |
| `todo name`  | Print the name of the todo list in the current directory. |
| `todo <name>` | Open named todo list. |
| `todo rm` | Remove todo list reference in current directory. |
| `todo rm <name>` | Delete named todo list. |


Development
--------------------

### Environment

| Var  | Required  | Description | Example |
|---|:-:|---|---|
| `PROJECT_API_URL`     | x | Url of project api | `http://localhost:3333/projects/`  |
| `PROJECT_API_KEY`     | x | Valid API token. | `1643tej8qdfqsgm70b7hb5554riptbuvvnukp8pha8fnf3lgbv1e`  |
| `UNSAFE_TLS`     |   | Ignore certificate errors (eg. from self-signed dev certs) | `true`  |

### Testing

Must have [bats](https://github.com/sstephenson/bats) available on the command line for testing.

To test the `todo` command, run `make test`. 



### Paths

| Path  | Contents |
|---|---|
| `/opt/project-fs`  | System-wide binaries and bash tools. |
| `/home/<user>/.project-fs` | User configuration and mount point. |
| `/usr/local/bin/todo` | `ln -s => /opt/project-fs/todo` |
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
