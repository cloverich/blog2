---
title: Taking Deno for a spin
date: "2020-05-23T00:00:00.000Z"
description: "Deno hit 1.0, and I was curious how usable it was. Here's a tiny script and my up and running thoughts"
---

When I first heard about Deno [^1], I thought it would fizzle. A better Node sounded nice but I couldn't see why it would take off. When it hit 1.0 and [blew up](https://news.ycombinator.com/item?id=23172483) on hackernews, I was surprised and a little excited. Maybe it had more steam than I thought. Node has warts and Typescript is great after all. Reading the announcement [blog post](https://deno.land/v1), two points stood out to me:

- Core behaviors use async / await, not `EventEmitter`
- It has a standard library, and its based on Go's

Despite being familiar with the Node ecosystem, I'm no less annoyed at having to spell out my dependencies for all the utterly basic needs a standard library provides. Its not very important for large, established projects. But when you write a lot of one-off, unrelated scripts, the pain adds up. I found working with streams particularly difficult and was surprised at how many incompatible ways there were to do the same thing. When streams work they are great. When they don't — its awfully confusing. 

So for day-to-day scripts, maybe `ts-node` or Go would be a better option. But I have a small problem a one-off script can solve — so figured I'd take Deno for a test drive.

## Installation, looking around

Before laying out my simple problem, I setup a hello world script. I use the installation instructions from the [install page](https://deno.land/#installation):

```bash
brew install deno
mkdir rename-files && cd rename-files
echo 'console.log("hello, world!")' > hello-world.ts

deno run index.ts
Compile file:///hello-world.ts
hello, world
```

Poking around, many methods I'd expected to pull from a library are exposed on the `Deno` global (ex: `Deno.cwd()`) — these are listed under the [runtime API](https://doc.deno.land/https/github.com/denoland/deno/releases/latest/download/lib.deno.d.ts). The [standard library](https://deno.land/std) is a separate thing, and it has... interesting documentation. It appears to be mostly commented source code, with the occassional README.md rendered as HTML. I like the simplicity but don't linger long. The installation page provides example usage:

```tsx
import { serve } from "https://deno.land/std@0.50.0/http/server.ts";
const s = serve({ port: 8000 });
console.log("http://localhost:8000/");

// async iteration of a never-ending stream of requests?
for await (const req of s) {
  req.respond({ body: "Hello World\n" });
}
```

If you are familiar with Node, the import statement might stand out the most. This is the import from url behavior touted as an alternative to npm and `package.json`. Specifying the version of the standard library seems odd to me — I'd expected it to be auto(magically) tied to whatever version of Deno I was running with the script (and, downloaded along side it). Feels odd, but I may not have enough context to understand. More on this later.

After some trial and error, I get the script to run with `deno run --allow-net example.ts`. Unlike Node, Deno requires you to grant permissions for various io facilities; they can be listed with `deno run -h`. This was straight forward to figure out. 

With the basic setup out of the way, I think I am ready to tackle the problem at hand.

## A simple script

I've been journaling for years, and 4-5 years ago settled on a certain format for storing files:

```tsx
2019/
  03/
    2019-03-02.md
    2019-03-15.md
2020/
  05/
    2020-05-01.md
    2020-05-02.md
```

I have a couple journals prior to that using various other conventions but they are always markdown files grouped into a named folder:

```tsx
an-old-journal/
  foo_was_interesting.md
  2015_tuesday_again.md
  // etc
```

I'd like to convert them in bulk to the other format. I can do this by:

- walking the old journal directories
- renaming them to `YYYY/MM/YYYY-MM-DD.md` based on the files [mtime](https://unix.stackexchange.com/questions/2464/timestamp-modification-time-and-created-time-of-a-file)
- merging them into an output directory

## Walk the file tree

Reading a directory of files, filtering down to markdown — we can do this with `readDir` in Node, but I usually use [klaw](https://www.npmjs.com/package/klaw). When I dabbled in Go I remembered this was part of its standard library ([walk](https://golang.org/pkg/path/filepath/#Walk)), so I go searching the equivalent in Deno's. I was pleasantly surprised [to find it](https://deno.land/std/fs/walk.ts). The page itself is commented source code... some pages render markdown (if there is a README.md). I don't linger on the longevity of this style but appreciate the simplicity today. Sure enough, [walk](https://deno.land/std/fs/walk.ts) is in the Deno standard library. From the standard library page:

> These modules are tagged in accordance with Deno releases. So, for example, the v0.3.0 tag is guaranteed to work with deno v0.3.0. You can link to v0.3.0 using the URL [https://deno.land/std@v0.3.0/](https://deno.land/std@v0.3.0/). Not specifying a tag will link to the master branch.

I import walk per the instructions, then write a simple loop using [async iterator](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Statements/for-await...of) syntnax:

```tsx
import { walk } from "https://deno.land/std@v1.0.0/fs/mod.ts";

async function readIn(srcDir: string) {
  const processed = new Map<string, string>();

  for await (const entry of walk(srcDir, {
    includeDirs: false,
    exts: ["md", "mdown"],
  })) {
    console.log(entry);
  }
}
```

Unfortunately, this results in:

```bash
Compile file:///Users/cloverich/code/scripts/convert-ls-chron-deno/index.ts
Download https://deno.land/std@v1.0.0/fs/mod.ts
error: Uncaught Error: Import 'https://deno.land/std@v1.0.0/fs/mod.ts' failed: 404 Not Found
    at unwrapResponse ($deno$/ops/dispatch_json.ts:43:11)
    at Object.sendAsync ($deno$/ops/dispatch_json.ts:98:10)
    at async processImports ($deno$/compiler.ts:736:23)
    at async processImports ($deno$/compiler.ts:753:7)
    at async compile ($deno$/compiler.ts:1316:31)
    at async tsCompilerOnMessage ($deno$/compiler.ts:1548:22)
    at async workerMessageRecvCallback ($deno$/runtime_worker.ts:74:9)
```

So much for that. I must have written the import incorrectly. After some digging I reluctantly remove the version and link the master branch. But this results in another error, and it is quite confusing:

```bash
Compile file:///Users/cloverich/code/scripts/hello-world.ts
error: TS2339 [ERROR]: Property 'utime' does not exist on type 'typeof Deno'.
    await Deno.utime(dest, statInfo.atime, statInfo.mtime);
               ~~~~~
    at https://deno.land/std/fs/copy.ts:90:16

# ... many more errors
```

I google the errors but later realize the answer is in the [fs README](https://deno.land/std/fs)

> All the following modules are exposed in mod.ts This feature is currently unstable. To enable it use deno run --unstable

Unexpected for 1.0 but, lets see:

```tsx
deno run --unstable hello-world.ts 
Compile file:///Users/cloverich/code/scripts/convert-ls-chron-deno/hello-world.ts

{
 path: "/Users/cloverich/wiki/notes/2013-07-11.md",
 name: "2013-07-11.md",
 isFile: true,
 isDirectory: false,
 isSymlink: false
}

# ... etc
```

## Read and write files

Now that I can walk the file tree, I need to:

- Read file contents
- Check the `mtime`

In Node I could read a file with `fs.readFile`. I would usually wrap this w/ util.promisify to make it more palatable. Its a bit different in Deno:

```tsx
const decoder = new TextDecoder("utf-8");
const contents = decoder.decode(await Deno.readFile(entry.path));
```

Interestingly, there's no imports required for this. Deno exposes some Web APIs in its global namespace (I need to learn more about this); `TextEncoder` seems part of the [Encoding API](https://developer.mozilla.org/en-US/docs/Web/API/Encoding_API). At any rate, in Node the encoding could be specified in the call to readFile. After some searching, Deno provides a convenience helper in std ([`read_file_str`](https://deno.land/std/fs/read_file_str.ts).)

Moving on to writing, I naively write out the file in an if / else, assuming I'll find an `appendFile` fs equivalent in Deno.

```tsx
if (await exists(outfile)) {
  // appendFile?
} else {
  await writeFile(outfile, encoder.encode(contents))
}
```

But `writeFile`  actually accepts an `append` option, so the check becomes one intuitive line: 

```tsx
await writeFile(outfile, encoder.encode(contents), { append: true })
```

Deno provides a further enhancement with [`write_file_str`](https://deno.land/std/fs/write_file_str.ts) but it does not accept an options object to toggle append file mode. Shucks. 

## The final program

Well, I could explain all of my thinking but its not interesting. So here's the final short script:

```tsx
import { walk, ensureDir } from "https://deno.land/std/fs/mod.ts";
const { readFile, writeFile } = Deno;

// See readFile / writeFile calls
const decoder = new TextDecoder("utf-8");
const encoder = new TextEncoder();

/**
 * Merge a directories content into another, using my folder / filename conventions
 * 
 * @param dateToContents - Dictionary of truncated date strings (YYYY-MM-DD) to file contents
 * @param dir - The output directory
 */
async function mergeInto(dateToContents: Map<string, string>, dir: string) {
  for (const [dateStr, contents] of dateToContents.entries()) {

    // Generate folders + filename: YYYY/MM/YYYY-MM-DD.md
    // ex: 2020/03/2020-03-15.md
    const ymdir = `${dir}/${dateStr.slice(0, 4)}/${dateStr.slice(5, 7)}`;
    const outfile = ymdir + "/" + dateStr + ".md";
    
    await ensureDir(ymdir);
    await writeFile(outfile, encoder.encode(contents), { append: true });
  }
}

/**
 * Walk a source directory and collect all file contents.
 * @param srcDir 
 */
async function readIn(srcDir: string) {
  const processed = new Map<string, string>();

  for await (const entry of walk(srcDir, {
    includeDirs: false,
    exts: ["md", "mdown"],
  })) {
    const fileInfo = await stat(entry.path);

    // I'll rename file to yyyy-mm-dd based on last modified time
    const dateStr = fileInfo.mtime!.toISOString().slice(0, 10);
    const contents = decoder.decode(await readFile(entry.path));

    if (processed.has(dateStr)) {
      // append
      const existing = processed.get(dateStr);
      processed.set(dateStr, existing + "\n" + contents);
    } else {
      processed.set(dateStr, contents);
    }
  }

  return processed;
}

const wikiLogs = await readIn("/Users/cloverich/wiki");
await mergeInto(wikiLogs, "/Users/cloverich/notes/some-journal");
```

The program is straight foward, and could be further simplified by simply appending as I read; by the time I realized that this one off script was nearly done. So I'll save that work for my next visit. 

## Take aways

I enjoyed writing this small script in Deno. Other than some confusion around importing and using the standard library, it was relatively smooth sailing. Not having to pull in 3rd party dependencies for trivial tasks (like `mkdirp` and `walk`) was refreshing, and I could see myself using Deno for more one off scripts like this. 

While Deno spiked in popularity now, is it a flash in the pan? If the runtime and standard library continue to mature, I think there's a valid argument that Deno may be a more practical choice than something like `ts-node` for people who are familiar with javascript but want less boilerplate. Will that evolve into a mature ecosystem? I'm not sure, but it doesn't seem as far fetched as it did a year ago. I think someone familiar with Node programming wouldn't be a fish out of water in Deno land, so I don't think there's a need to "keep up" with its change; I'd argue you should use whichever one feels more practical or pleasant.

A more difficult question to answer (for me anyways), is why choose Deno over Go? I am more conflicted here. If an advantage of Deno is the standard library borrowed from Go — why bother? I've dipped my toes into it a few times, and found it mostly intuitive. I remember string parsing and JSON manipulation to be a bit more involved, even frustrating. When you are used to manipulating JSON as first class dynamic objects, its difficult to imagine anyting else being more practical. Particularly for one-off scripting needs.


[^1]: Ryan Dahl's [10 Things I regret about Node.js][1] in June 2018, where he reflected on mistakes he made in Node's early development. At the end he introduced Deno as a toy project he was working on.

[1]: https://www.youtube.com/watch?v=M3BM9TB-8yA
