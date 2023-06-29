---
title: The Fifth Prototype
date: "2020-08-19T00:00:00.000Z"
description: "A brief retrospective on rebooting and shipping a basic version of my journaling application"
tags: devlog
---



On July 2nd I re-booted a stalled side project of mine called Chronicles — a markdown based journaling application. On August 14th [I shipped the first (very basic) phase](https://github.com/cloverich/chronicles/releases) of it. Nobody should be interested in this release, but it is an important milestone for me — something usable shipped. This is a brief retrospective on the first milestone.

## Getting in a rhythm with Go

The thought of restarting this project was daunting. I'd built a lot of stuff and if I wanted to use the old code I had some complicated changes to make. The path towards the features I really cared about was not obvious. I still had questions about how the backend should work -- a seprate process? Hosted or local? At some point I decided the best path forward was to get back into a rhythm, whatever it took. I decided to  treat it like a hackathon and I wrote the first version of a new backed in Go. It was a great experience. I hit a few snags but had working code quickly. I found a [great markdown library](https://github.com/russross/blackfriday) that produces an AST I could interact with — a core requirement for the search, transform, and filtering features I am after. I toyed with the [webview](https://github.com/webview/webview) project and within a few days had a new, basic UI working on top. I felt the rush of getting something off the ground and was pretty excited to be in the flow again. Whatever the merits of a re-write, I'd went from months of not even looking at it, to working on it in a healthy rhythm a few nights a week. 

## Go → Deno

Once I got to a stopping point, I planned out the UI a bit further. I wanted to use [remark](https://github.com/remarkjs/remark) and [its AST](https://github.com/syntax-tree/mdast) to support some particular features and realized I'd need the backend and front-end to share an AST. The Go library I was using was not compatible out of the box, and though I briefly tried extending it, I realized I was innovating in the wrong place. Did I need to switch to Node to take advantage of Remark? I was dreading that. I had [played with Deno](/blog/2020-05-23/deno-test-drive) recently, am relatively high on it, and guessed I could get remark working in it instead. After a little work and surprisingly few hiccups, I'd ported my Go backend to Deno *and* gotten remark with a few plugins working. 

## Building a Roadmap

At some point the project started to feel real again, and I decided to pause and spend more time on the roadmap. At the start, I ruminated on the backend quite a bit. What were the hosting options? How would I manage privacy? And the monetization choices? How could I deliver on the core product vision and also turn it into a sustainable project? After some unstructured writing and much reflection, I found myself asking: Do I even have a good idea here? And if so can I even execute on it? The reality is I've toyed with this idea for ages, but never shipped anything meaningful. That was a difficult thing to accept. But in doing so I was able to isolate the thing I really cared about: Prototyping my ideal journaling application. 

When I isolated the features that were core to that idea, I realized that I'd often deferred features I care about to promote other less critical ones. Hosting, licensing, a website -- these are important to creating a sustainable business, but not to exploring my vision for the project. Performance, a good editing experience -- very important in general but they are polish. I needed to focus on getting to the core value. Ultimately, I needed to realize enough of the raw vision to know whether its worth following through on the rest. That is where I fell short previously, and once I organized my roadmap a bit it became clear how to re-prioritize to get the truly important things towards the top. 

## Deno → Electron and Node

One casualty in this process was my hope for avoiding Electron. I'd hoped I could ship a hosted or webview based version of the app, but it was obvious that would take a lot more work. It was a tough choice as I was really enjoying Deno (Electron requires Node), but once I decided to move the refactor went surprisingly fast. In the end I had a Node based API and Electron UI working in a single day. A few years from now I may have had other, better choices. But for a cross platform desktop application today, for those with many constraints (say, a fulltime job and kids), Electron is difficult to beat. 

## HTTP server

"Why is there is an http server running for this local application?" Because compromise and pragmatism. I had an http-based API for reading and saving documents already, and using it was the simplest thing to do. It took a few tricks but like the move to Node and Electron, I had something working in a single sitting. One day I hope the project has enough steam that I can remove the http sever and use IPC or even something else[^1] . But that's a journey for another day.

## Restraint

Compared to other projects I've worked on, this project has much ambition. Scope creep and complexity are frequent challenges. I used the goals I'd laid out in my roadmap to stay focused, to keep features under control, and to avoid shipping too many half implemented ideas. As an example, searching across multiple journals is a key feature I wanted my API to support and I got a working version of it early on. However the UX was pretty confusing and janky. Rather than leave it or attempt to add more to it, I decided to back off and replace it with a single journal select menu. I do not love how it works now, but it feels a bit more "complete". I gave several other features similar treatment and ultimately, I feel the project is more prepared to begin layering on top of. Restraint is a skill that's never come easily to me, and its something I recognize I need to willfully improve. If I take anything else away from this project, I would be happy to apply more restraint to anything I take on in the future. 





[^1]: In Electron you can use Node directly from the UI. So if I need to load local markdown files for a journaling app, I can do it directly. It has a performance impact since it shares the UI thread but if you design it reasonably well, its workable in simple scenarios. I'd previously worked on a web-worker and indexeddb solution to speed it up, but that's among the many reasons I never really shipped the prior version. So, this is on the docket for a future experiment. But again, to the roadmap -- its a bit down the road.