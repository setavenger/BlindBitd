## blindbit-cli listaddresses

Lists all addresses belonging to the user

### Synopsis

Daemon has to be unlocked. Lists all addresses belonging to the user. 
If a user has not created any labels this will always 
return the standard silent payment address. 
The daemon returns the addresses in the order the labels were created. 
1. the base address then the labels with increasing m 1,2,... 
The change label address will not be returned.

```
blindbit-cli listaddresses [flags]
```

### Options

```
  -h, --help   help for listaddresses
```

### Options inherited from parent commands

```
  -s, --socket string   Set the socket path. This is set to blindbitd default value (default "~/.blindbitd/run/blindbit.socket")
```

### SEE ALSO

* [blindbit-cli](blindbit-cli.md)	 - A cli application to interact with the blindbit daemon

###### Auto generated by spf13/cobra on 7-May-2024