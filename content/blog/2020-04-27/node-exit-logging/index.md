---
title: Node exit logging
date: "2020-04-27T00:00:00.000Z"
description: "Node supports exit codes and error logging but has important caveats I've learned the hard way."
---

Node supports exit codes and error logging but has important caveats I've learned the hard way. What follows is a summary of what I learned while fixing a failed Github Actions workflow.

## Errors and exit codes

Consider this script which throws an error:

```js title="fail.js"
throw new Error('all i do is fail');
```

This script fails, and two important things happen:

1. It logs an error
2. It exits with a non-zero exit code

```sh
> node ./fail.js

# error is logged
/home/code/fail.js:1
throw new Error('all i do is fail');
^

Error: all i do is fail
    at Object.<anonymous> (/home/code/fail.js:1:7)
    at Module._compile (internal/modules/cjs/loader.js:955:30)
    at Object.Module._extensions..js (internal/modules/cjs/loader.js:991:10)
    at Module.load (internal/modules/cjs/loader.js:811:32)
    at Function.Module._load (internal/modules/cjs/loader.js:723:14)
    at Function.Module.runMain (internal/modules/cjs/loader.js:1043:10)
    at internal/main/run_main_module.js:17:11

# Print the exit code
> echo $?
1
```

These unsurprising outcomes can change unexpectedly when you add async code to the mix. 

## Errors in top-level async code

Its typical in node for the main script to make asynchronous calls. Node did not support [top-level await][2] prior to v13, so until recently code like this was quite common:

```js
(async () => {
  throw new Error('all i do is fail');
})();
```

If the code inside the async block throws an error, you see the following:


```sh
(node:35343) UnhandledPromiseRejectionWarning: Error: all i do is fail
    at /home/code/fail.js:2:9
    at Object.<anonymous> (/home/code/fail.js:3:3)
    at Module._compile (internal/modules/cjs/loader.js:955:30)
    at Object.Module._extensions..js (internal/modules/cjs/loader.js:991:10)
    at Module.load (internal/modules/cjs/loader.js:811:32)
    at Function.Module._load (internal/modules/cjs/loader.js:723:14)
    at Function.Module.runMain (internal/modules/cjs/loader.js:1043:10)
    at internal/main/run_main_module.js:17:11
(node:35343) UnhandledPromiseRejectionWarning: Unhandled promise rejection. This error originated either by throwing inside of an async function without a catch block, or by rejecting a promise which was not handled with .catch(). (rejection id: 1)
(node:35343) [DEP0018] DeprecationWarning: Unhandled promise rejections are deprecated. In the future, promise rejections that are not handled will terminate the Node.js process with a non-zero exit code.
```

I've often seen the following recommended in response:

```js
(async () => {
  throw new Error('all i do is fail');
})().catch(console.error)
```

I've written code like this, but it is wrong. It prevents the `unhandledRejection` and prints a normal looking error log, but the program exits cleanly.

```sh
> node fail.js 

# Looks like an error
Error: all i do is fail
    at /home/code/fail.js:2:9
    at Object.<anonymous> (/home/code/fail.js:3:3)
    at Module._compile (internal/modules/cjs/loader.js:955:30)
    at Object.Module._extensions..js (internal/modules/cjs/loader.js:991:10)
    at Module.load (internal/modules/cjs/loader.js:811:32)
    at Function.Module._load (internal/modules/cjs/loader.js:723:14)
    at Function.Module.runMain (internal/modules/cjs/loader.js:1043:10)
    at internal/main/run_main_module.js:17:11

# But clean exit reported!
> echo $?
0 
```

Instead, you can manually indicate the program failed with `process.exit`:

```sh
(async () => {
  throw new Error('all i do is fail');
})().catch((err) => {
  console.error('program failed:', err);
  process.exit(1);
});
```


This solves both problems: The error is logged and the program exits with a non-zero exit code. Until recently, I thought this was all there was to it. 

## Caveats

Working with larger programs, you'll invariably end up using a logging library. What better place to use it than to log some errors!? 


```ts
// ...
.catch(err => {
  logger.error('Something went wrong!', err);
  process.exit(1);
})
```

I was surprised to learn that the logger may or may not print before the program exits. Interestingly, the following may also be problematic:

```ts
// ...
.catch(err => {
  process.stderr.write('Something went wrong!, but you may not see this!');
  process.exit(1);
})
```

From the node [`process.exit`](https://nodejs.org/api/process.html#process_process_exit_code) docs:

> ...writes to process.stdout in Node.js are sometimes asynchronous and may occur over multiple ticks of the Node.js event loop. Calling process.exit(), however, forces the process to exit before those additional writes to stdout can be performed.

Elsewhere in that document:

```
process.stdout and process.stderr differ from other Node.js streams in important ways:

    They are used internally by console.log() and console.error(), respectively.

    Writes may be synchronous depending on what the stream is connected to and whether the system is Windows or POSIX:
        Files: synchronous on Windows and POSIX
        TTYs (Terminals): asynchronous on Windows, synchronous on POSIX
        Pipes (and sockets): synchronous on Windows, asynchronous on POSIX

These behaviors are partly for historical reasons, as changing them would create backwards incompatibility, but they are also expected by some users.
```

In the [pino logging library][pino-logger] I am using, they addressed this scenario directly and explained in the [v5 release notes][pino-v5-notes]:

> The new pino.final method supplies a special logger that always writes synchronously. This is important for log messages written during the final tick, since the process will have died before asynchronously scheduled operations would resolve.

There is a special section of the pino documentation that calls this out explicitly -- [exit-logging][exit-logging]:

```ts
process.on('uncaughtException', pino.final(logger, (err, finalLogger) => {
  finalLogger.error(err, 'uncaughtException')
  process.exit(1)
}))

process.on('unhandledRejection', pino.final(logger, (err, finalLogger) => {
  finalLogger.error(err, 'unhandledRejection')
  process.exit(1)
}))
```

Those two examples have an additional side effect in that an unhandled promise rejection in _any_ method will crash the program, but the example is illustrative: **If you use a logging library, take special care when exit logging.**

## Summary
- Non-zero exit codes let calling processes reason about your program exits
- Logging errors before crashing is important, but has caveats
- Uncaught errors may or may not crash Node
- Logging libraries have existing solutions for proper exit logging, but you must use them properly to benefit


## See also
- [Linux and Unix exit codes][exit-codes]
- [Let It Crash: Best Practices for Handling Node.js Errors on Shutdown][let-it-crash]
- [pino exit-logging][exit-logging]
- [top-level await][2]
- [Github issue: Terminate node on unhandled promise rejection](https://github.com/nodejs/node/issues/20392)


[1]: https://help.github.com/en/actions/reference/
[2]: https://v8.dev/features/top-level-await
<!-- workflow-syntax-for-github-actions#jobsjob_idsteps -->
[unhandledRejection]: https://nodejs.org/api/process.html#process_event_unhandledrejection
[exit-codes]: https://shapeshed.com/unix-exit-codes/
[node-console]: https://nodejs.org/api/console.html
[note-on-process-io]: https://nodejs.org/api/process.html#process_a_note_on_process_i_o
[exit-logging]: https://getpino.io/#/docs/help?id=exit-logging
[pino-v5-notes]: https://www.nearform.com/blog/announcing-pino-v5-0-0/
[pino-logger]: https://github.com/pinojs/pino
[let-it-crash]: https://blog.heroku.com/best-practices-nodejs-errors
[what-is-the-event-loop]: https://www.youtube.com/watch?v=8aGhZQkoFbQ
[event-loop-timers]: https://nodejs.org/uk/docs/guides/event-loop-timers-and-nexttick/
