Importing a new Code List
================

A script used to convert CSV code lists to json files for mongo import

### Getting started

* Install mongodb
* Install go


### Prepare your data

Code lists should be broken down into a file per list - don't try and import time, and geography and SIC
codes all in the same input file.

Code lists must be in a CSV file, where the codes are the first column and the labels are the second.
Headers should be the code list ID and list name respectively.

This is most easily achieved by taking the 2 dimension columns from an established V4 formatted file, and replacing the headers.

The header `geography_abcd1234, geography` from a V4 would become `abcd1234, geography` before importing here.


### Run the script
The script requires one argument: -input

`-input` is the path to the CSV file you're wanting to import. If the file is in this scripts directory
this will just be the filename

```sh

$ go run add.go -input example.csv
# ... the example CSV file is in this directory ...

$ go run add.go -input $HOME/files/example.csv
# ... the CSV is in a different directory ...
```

### Outcomes

The script will change 3 files:

* `codelists.json` will have a line added containing the json document for the list you've just added
* A new file will be created of the format `<listID>.json` - this file will contain json documents for
each individual code in the code list, and the filename will contain the code list ID provided in the
header of your input file.
* `setup.sh` will have a line added to ensure the new file is included in the mongo import when the
setup script is run

### Setting up a remote environment

If you want to run the `setup.sh` script against a remote/tunnelled environment, set the environment
variable similar to `export MONGODB_ADDR=mongodb://user:pass@localhost:tunnelled_port` before
running the script, e.g.:

```
    $ export MONGODB_ADDR=mongodb://user:pass@localhost:tunnelled_port
    $ ./scripts/setup.sh
```

### TODO
* Setup script is currently destructive. This is ok for a developer environment but shouldn't be
the case in production. A different script will be needed to constructively add new values to
mongo in live environments.
