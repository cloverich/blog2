---
title: Typed code and ADHD
date: "2024-05-13T00:00:00.000Z"
author: Christopher Loverich
tags: devlog
description: Considering the impact of ADHD and working memory, and whether a preference for typed programming languages is more than just subjective preference.
---


![Photo from a recent Washington hike about half way up Beacon Rock state park](assets/beacon_rock2.jpeg?resize=blogImages)

Navigating life with ADHD often involves developing strategies to manage daily challenges, particularly in professional environments. For me, a transformative moment came from understanding the [role of working memory in ADHD](https://guilfordjournals.com/doi/abs/10.1521/adhd.2008.16.6.8?journalCode=adhd). It reshaped how I approach tasks and manage my cognitive loads. There are many great [blogs](https://www.andrewaskins.com/how-i-run-a-company-with-adhd/) and [discussions](https://news.ycombinator.com/item?id=38274782) on this already, and I don't intend to pile on with everything that has worked for me, here.

Instead I wanted to reflect on a particular preference that, until recently, I hadn't even realized I was benefiting from: Typed code.


## Working with and without typed code
Since my career officially began in 2011, many of my opportunities required changing programming languages. After a few years of PHP and Python (which I loved), I briefly inherited an important project in Java. I was loathe to learn it and found the syntax and tooling clunky and slow. I was scared my career momentum would stall. Yet the opposite happened, and after the initial learning curve I was shocked at how productive I was with it. In particular I was able to write much more code at one time without compiling, and the code I did write had fewer bugs, typos, and other little issues that typically plauged my rough drafts.

When the project eventually finished, I took the next opportunity and with it another language change. This went on for some time, and while the language churn was not pleasant, the changes always came with genuine opportunities — typically pay bumps or more promising companies / products. I rarely felt technically limited and despite my preferences, did not dwell too much on which language I was forced to use at any given time. 

Eventually I landed at a Ruby on Rails shop, and things started to really slow down. Now in my early 40’s and an ever tired parent, I found picking up the syntax and the ecosystem a bit more tiring. Yet experience was on my side and the initial hump was eventually cleared. Despite this, my productivity never fully returned.  

It was around this time I began recognizing ADHD’s impact on my life, and in particular the ways I compensate for the working memory aspects of it. While many of my coping mechanisms helped (greatly), I began wondering if I was at my limits. Perhaps simply unable to work those extra hours given my parenting duties, or even being older and mentally slowing down in general was to blame. Had I simply lost my edge? Was it the multiple rounds of COVID? All of the above? I started to re-evaluate my approach to work, my systems for managing ADHD, and reluctantly accepting that I might simply be past my prime.

Then I picked up an old side project of mine in those 90 tired minutes I occasionally have to spare, using a typed language (Typescript) I was already familiar with. In an instant I was suddenly productive again. I was shocked at how easy everything suddenly felt, and how quickly I was moving. After burning through a large set of issues, feeling refreshed and efficient, I stepped back to evaluate whether the language itself was playing a role in this phenomena. After some reflection, it seemed I worked quite differently when using a typed language, and something more than preference or familiarity was going on. A typical session may look like:

* I start with a shell component, and add interfaces for the data
* I then write a few empty transformation functions
* I then write a couple interfaces for the objects these functions accept and return
* I begin refactoring the function signatures and interfaces, even entirely discarding early ideas.
* I realize (and am surprised) how quickly forget how things work even as I write them; I find myself hovering the interfaces I had just written to remind myself how they work and what they are for.
* I eventually begin to write implementations, and iterate on the interfaces; I mostly just follow the squiggly lines that remind me about things I forgot to do, to change, or typos I inevitably overlook.
* I do a final phase of refactoring, usually cleaning up the old ideas.

Stepping back to evaluate this process, I had clarity as to why this experience was so different. It was not the language familiarity, or the lack of constraints (given it is an unimportant side project). It was **the way my brain used the language and tooling itself**. 

## Is typed code actually better for managing ADHD?
I have two hypothesis about why I find typed code so much easier to work with.

1. **Offloading Working Memory**: Encoding thoughts into types reduces the need to hold complex details in memory, allowing me to focus on higher-level problem-solving.
2. **Immediate Feedback through Type Errors**: This feedback helps correct mistakes in real-time, reducing the cognitive load and enhancing task persistence.

These aspects of typed programming seem to create a structured framework that not only compensates for, but enhances my minds ability to think clearly and in a complex, focused, and motivated manner.

However, I remain somewhat skeptical about this conclusion. 
Everyone has their own set of preferences, strengths, weaknesses, styles, that lead them to one language or another. I tend not to judge any one preference or language as superior in any objective sense. I’ve seen plenty of efficient programmers using rails and python. And I like Typescript, and have maybe the most overall experience with it (thankfully, while front-end libraries change often, the underlying language does not).

Conversely, I've been thinking often of those Java days early in my career. It’s a language I don’t know very well, an ecosystem I am not super familar with, and a syntax I don’t particularly like. There are plenty of quirky odds and ends and when forced to use it, I am rarely immediately productive with it. Yet still I find my mind is much less distracted and much less fatigued when using it, and I think there is something to it. 

And I think when my next opportunity comes, I may weigh the language on offer a bit more highly.
