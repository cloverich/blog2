---
title: "Staff Engineer: Leadership beyond the management track"
author: Christopher Loverich
date: "2021-11-22T00:00:00.000Z"
tags: books
description: "Thoughts while reading Will Larson's Staff Engineering book"
---


This post is a review of (and thoughts about) Will Larson's [Staff Engineering](https://staffeng.com/book) book.

**Introduction**

I've worked at a few successful startups, yet rarely one's that had proper engineering roles as such. If your tech career spawned out of IT like me,  you might lack a conceptual framework for thinking about engineering leadership. There are hero programmer's and prestige seeker's. If there is an Engineering Manager it might be someone from finance with an MBA.

I still remember reading [Camille Fournier's The Managers Path](https://www.amazon.com/Managers-Path-Leaders-Navigating-Growth/dp/1491973897) and it being quite revelatory to me. Ah -- this is what distinguishes those successful tech companies from the software start-ups that kind of fizzle out, or the larger ones that never transition successfully into real software companies. They simply lack leaders with a background in software engineering, the supposed backbone of their company. How could they succeed? I wouldn't appoint an engineer as CFO. But the converse is, or was, surprisingly common. 

It took me much longer to come around to another kind of leadership, the one that we apparently call Staff ("plus") Engineering. I became a bit fascinated with the concept, though my understanding of it was limited. Thus when I heard about [Will Larson's upcoming book on the subject](https://staffeng.com/book), I was sure I would read it. At this point in my career I understand the general purpose of staff engineering roles and the perils of their absence. Yet, I lack a structured way for thinking about them. I'd hoped his book would give me some of the vocabulary and structure I crave. To that end it was successful.

**Reflecting on my favorite parts.** 

The book is broken in to two parts. The first is good, and relatively short. The author surveys their own leadership career, and reviews the similarities of the people he interviewed, and then summarizes it into key concepts. That's the part I"ll focus on here. The second part is all interviews with established Staff Engineers. They aren't bad. Just low fidelity. A few nuggets and a sense of camaraderie. Nearly every book I have read would be better if it were half the length. Yet something about a 100 page book feels less legitimate. I digress.

**Focus and effectiveness**

A recurring theme throughout the book is ambiguity and problem selection. This definitely resonates with me. In my first Engineering Management (EM) role it became difficult for me to focus; there were so many problems to solve. I eventually realized this was quite normal for an EM, and figuring out what to work on becomes both more important and more stressful. Maybe essential stress is the right word because, unlike other stressors, this felt like part of the job. As people above you are reduced in number, so too are those who can help you figure out what to work on. I expected this was similar for a Staff Engineering role.

> "Sure there's work that you're faster at or better at than some other folks, but much more important is the sort of work that simply won't happen if you don't do it". 

Working on the right thing is easy when it is code. But it won't always be code. It might be technical strategy, or even recruiting. That's where the job seems the most difficult, to **always be searching and selecting** for what is high impact. You will quickly find more problems than you can possibly tackle yourself. You should be searching, selecting, focusing, completing. Re-evaluating on some cadence. You need a good system to stay organized, aligned, and focused. Any one of those is a tough skill to develop. You won't always deliver high impact work but you must occasionally do so. There's a level of stress that goes with that, but it also develops you into a better engineer. And that is quite rewarding. Even it it takes a little time. 

See also - [Being Glue](https://noidea.dog/glue)

**Building an Engineering Strategy**

> Having a well structured plan can help substantially reduce scope while getting to most of the value
>
> Damian Schenkelman - Principal Engineer at Auth0

The short section on developing an engineering strategy was one of my favorite parts. Writing technical specs is something that I began doing relatively late in my career. When it first sunk in, what it really meant, it hit me how amateur my work and mode of operation had been before. It has taken me a long time to transition out of hacking and into thoughtful and intentional speccing and building. And a bit like your coding becomes capable of more expressive and complex ideas as you improve, so do your specs. If you haven't written any before, you might be surprised how bad you are at it. I used to think it merely exposed how incomplete your thinking was. But I don't think that is it, at least not the whole story. After coding day in and out for years, you get quite good at using code for organizing and structuring complex thoughts. Writing has the same effect but is a different skill. A skill that takes its own time to develop. It is worth it though, as even extremely laborious specs take far less time to write than the equivalent code, and can reach a wider audience. In addition to controlling scope, they allow you to set a vision and recognize your own gaps. And most importantly, serve as a point of communication and alignment. You can point to it and say not only this is what we think it should be done, but historically as a way to say "this is how the system came to be". 

Getting good at writing and evangelizing spec's, and building them into your culture, is certainly a key aspect of being an effective Staff Engineer. This book does a good job of laying out a way to take it even further, and progress from a spec into a multi year engineering strategy. 

See also:
- [The Power of “Yes, if”: Iterating on our RFC Process](https://engineering.squarespace.com/blog/2019/the-power-of-yes-if)
- [Delivering on an architecture strategy](https://blog.thepete.net/blog/2019/12/09/delivering-on-an-architecture-strategy/)

**On Access to interesting work**

One often cited reason for why getting the Staff Engineer title is worth it is to obtain access to so called interesting work. This one is a bit funny to me. Sometimes the most important thing isn't technical at all. Sometimes it is pretty boring. Finding joy with this mindset I think has to come from being motivated by the success of the business, because in those moments where you could choose something technically interesting vs something with high impact, you basically have to choose the one with impact. The book and the interviews back this up to an extent. It is not an absolute rule, but for me it stands in relief because I don't think Staff roles and interesting work necessarily go hand in hand, as one (me, previously) might naively aspire towards. Instead I think finding interesting work should happen through a mix of job hunting and moving towards business domains that are interesting. That is to say, there is no rule that says "become staff engineer, find interesting work." I think instead a better rubric is: If you are waiting for interesting work to come to you, you might be waiting an awful long time. If interesting work is what you want most, you should go find it. Career progressing into a Staff role is perhaps best viewed as a side effect of that pursuit. Because at worse it will, just like a management position you don't want, get in your way. 

**Technical Quality team**

> Process rollout requires humans to change how they work, which you shouldn't undertake lightly.

On improving technical quality, I especially appreciated the roadmap for a technical quality team. In particular I liked how he emphasized fixing hot spots over implementing best practices. Hot spots, best practices, and leverage points, in that order. Thinking of examples from my career, it is easy to categorize and address.

For example, is it painful to write and run tests? Invest in speedier test cycle. Figure out and fix that slow startup time. Or, perhaps, delete the slow or flaky tests outright. Make it easier to validate and QA key workflows. In an ideal world all the things would work. In this messy practical one, you have to focus on the key problem: Shipping working software in a timely fashion. If tests are hard to write and take a long time to run, people won't write them. If you address your lack of tests by then adding test coverage metrics, forcing people to write more, they will slowly hate working at the company. The best ones will leave. The good ones will follow. 

In other words, you can't put a metrics based quality program ahead of fixing obvious problems. And you likewise don't need permission or a system of metrics to implement best practices. Use your experience and judgement to decide if you can put out a few fires, or if you need to invest in a higher order decision. Having a a structured way to think about these steps helps.

**A Network of peers**

The very idea of "networking" was, for a long time, both confusing and intimidating to me. Until I realized it didn't need to be anything more than keeping in touch with the people you respect and like. As easy as grabbing lunch or starting a chat with former colleagues. And putting in the effort to maintain contact. Working with great people is the best way to find and do great work. It is much simpler than I thought. 

As soon as you gain some experience, it can be easy to reach for the best paying or most prestigious job. But instead you should probably find the one that has the people doing the things you want to do. People that you like and respect. Because those people become your network. **And they likewise become your access to new and interesting opportunities**. Those directly follow from their experience as well. By merely directing yourself towards the people doing things you like, or find challenging and rewarding, you'll over time become exposed to more of it. That Salesforce position probably pays a lot \_right now\_. But if you don't want to be doing it in ten years, don't do it now. (This reminds me a bit of [So Good They Can't Ignore You](https://www.amazon.com/Good-They-Cant-Ignore-You/dp/1455509124). Recommended)

Thinking in terms of your network also helps with your personal growth and maturity. Because another way to end up with a good network is to ask yourself: **How can I be remembered positively by my peers**? What this means changes a bit as you progress in your career. Early on, it can be easy as trying in earnest and leaving thoughtful code reviews, or writing maintanable code. As you become more senior however, it becomes more ambiguous and more impactful. 

Working with and supporting others becomes especially important. I have a recent example where I failed to do so that I think of often. At one company, we were using a project management tool which was not well suited to how we operated. I watched multiple teams tripping over themselves trying to bend it to their will. We were not doing anything magical at this company. Typical sprints and agile methodology you'd find in any startup. One senior colleague began championing JIRA. And I did not support them. I personally dislike it, and thought there were better options. Yet how I proceeded next was not right. I discussed this openly. I agreed with their assessment but pushed for an alternative solution. Some agreed with me, some with him, but far more importantly, most didn't say anything at all. There was no heated debate nor hurt feelings. Simply a moment that I incorrectly judged as unimportant. 

How I look at it now is quite a bit different. This particular colleague was trying to affect change. In an org that needed it, who had a problem committing to technical decisions of that nature. I should have recognized I was letting perfect be the enemy of the good. Reframing it from how I can get what I think is best, to how I can empower my peers to be effective, it would have been obvious what I needed to do. Meet with them directly, then [disagree and commit](https://en.wikipedia.org/wiki/Disagree_and_commit). Provide a united front and not just agree with their decision, but support it. Help them champion it, and help them implement it. It could be as simple as working with the people that trusted me most at the company, perhaps people that weren't in their inner circle, and win a few more votes. This was an obvious win-win opportunity. And in a deep sense of irony, I myself helped push and implement JIRA at a later point, after said person had given up their quest and left for greener pastures. 

**On being "in the room"**

> If you've ever wanted to know how the sausage was made from a leadership angle, maybe consider if you actually want to know how the sausage is made. >
> 
> Bert Fan, Staff engineer at Slack

There are several sections on getting the promotion, and getting into and staying in the room. I haven't worked at a company where people were more interested in career progression than solving useful problems or working with interesting tech. While I am certain I want to work at a company that understands the value of Staff Engineering roles, the set of sections around promotion and being in the room reminded me that career growth is the primary motivation for some people. And that not everyone can get the title, and there are some hurt feelings to think of, ego's to look out for, and just more people problems that need navigated. I imagine some companies end up with Staff Engineers who do not even like to code. They might still be good at it. Yet, I struggle to think I would enjoy working at such a place. I am infatuated with the idea of staff engineering because I know what it looks like when you work with peers who, at some base level, enjoy the craft. From a prior profession, I know what it looks like to work with those who do not. I am still amazed that the intersection between interest and effectiveness is not nearly as tight as I once thought. It is good to be cognizant of what comes with the territory and recognize the trade offs you are making. I don't know if I have a point to make. I guess it just makes me think.

**Is it worth the read?** 

Did you read this far? You should probably read the book if so. It is better written, more complete, and honestly not much longer than this reiew anyways. As far as leadership books go, it is not as good a book as [The Manager's Path](https://www.amazon.com/Managers-Path-Leaders-Navigating-Growth/dp/1491973897) or the various other well known leadership books, a few of which I've read. But the second half not withstanding, I don't think it should have been written differently. I actually think it is well done. It covers fertile ground -- this is a niche not well covered by other leadership books.

If you weren't brought up in a structured environment, yet are seeking to advance not only your own career but the craft and experience of those around you, you eventually need mentorship and guides. This book can help. It can provide a structure for thinking about career growth -- for yourself, and for others. To help you understand what you can expect of others. What you can ask of leadership. What you can ask for yourself. And also that sometimes growth and success come from asking forgiveness instead of permission, so long as you follow a few guidelines. 

 