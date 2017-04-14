**Messy**

Messy is a tool used to send XMPP messages from the command line. This can be used in a script or cron for notifications or to just shoot off a quick message.

**Usage**

Example : ```messy -server foo.bar:5222 -username testytesterton@foo.bar -password supersecrettest -remote receiver@foo.bar -message "Super useful notification message!"```

**Flags**

```

-server: Server and port to connect to. Ex. foo.bar:5222

-username: The JID to connect as. Ex. test@testy.net

-password: The password for the account.

-remote: The JID of the user to be sent the message.

-message The message to send. Omitting this flag will cause messy to read from stdin.
```
