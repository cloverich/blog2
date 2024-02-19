---
title: Reignition
date: "2024-02-18T00:00:00.000Z"
author: Christopher Loverich
tags: devlog
description: "Recapping Chronicles release v0.4.0"
---

Today I [released v0.4.0 of Chronicles](https://github.com/cloverich/chronicles/releases/tag/v0.4.0). It is the culmination of **the last six weeks of feature work**, and the first real work in two years. 

![Github insights of the last months work](assets/github_insight_feb2024.png?resize=blogImages)

The [release notes](https://github.com/cloverich/chronicles/releases/tag/v0.4.0) summarize the associated changes; here I wanted to briefly reflect on the current state of the project and future plans.

# Iterating on base usability
I called the [last release](https://github.com/cloverich/chronicles/releases/tag/v0.3.0) “Base Usability — Well kind of”. When I first broke up the hypothetical design of the app, I envisioned a base usability state. An overly simplistic but usable writing app which I could incrementally layer features on. “Well kind of” was me appreciating that **while the app is technically usable, it lacks the various note taking features I consider table stakes. Its not good enough to really use, even for basic tasks**.

Yet, there is still benefit to making releases; to doing batches of work, marking them complete, summarizing the changes and documenting learnings. The last release closed out some long unfinished work; it was nearly two years since my last commit, as my new job in 2022 led me to put it on the shelf a while. The current release includes a fresh set of features, fixes, and improvements. The closing out of each is a mere stepping stone towards something I hope will one day be good. 

## Using sprints
Before really kicking things off again at the start of 2024, I took some time to organize some upcoming work. I was inspired by [vscode's iteration plans](https://github.com/microsoft/vscode/issues/204292)[^3], and how they seem to organize a grab bag of features, loosely couple them into "themes", and use that to iterate.
I setup two sprints ([1](https://github.com/cloverich/chronicles/issues/116), [2](https://github.com/cloverich/chronicles/issues/123)) for current and future work. I found limiting total size and using "themes" a natural way to organize. I tried to keep any individual ticket, and any set of features from being too large. This was extremely helpful to carry the work through, because with kids and a full time job, I **generally did so in 30-90 minute increments**.

I started using a traditional pull request flow, instead of pushing to master; it feels a bit silly for a project with one developer and no users. But it helped keep things atomic and organized. I kept the commits and PRs fairly clean, linking them back to issues, and adding documentation to the PR as I went. This also helps when creating Github releases; using a pull request flow means Github can auto-create links to the pull requests since the last tag, and summarizes them. I’ve used [commitlint and the conventional changelog](https://github.com/conventional-changelog/commitlint) approach before, but found it too cumbersome for a personal project. I am pretty happy with the current approach. 

## The devlog on top
 Over the years, I’ve become more organized and methodical in my regular note taking — the way I take notes as I develop. Those notes are more useful than they used to be. I enjoy reading them more than I used to, and do so more often. It has impacted how I steer the features of Chronicles; its the friction in my daily note taking I want it to eliminate in the end. I don’t know what impact this note taking application will have on my life; if anyone will use it; if it’s ultimately worth doing. But what I do know is the way jouraling has impacted my life in general, has been extremely positive. And at the moment, **of all the good ideas in the world that I am not doing, this one blip of a good idea has momentum** — its a rolling project I can incrementally build off of. Overall, I found the end result of this invigorating and easy; when I reflect on the last month, I feel the project is fully alive again.[^1]



# What comes next
> You don't burn out from going too fast. You burn out from going too slow and getting bored. [^2]

I am extremely excited about [the next sprint](https://github.com/cloverich/chronicles/issues/123). Unlike the last one, it will be **more focused on features**. Notably it will include several key features I rely on in other apps, such as:

* Internal note linking (Many)
* Tagging (Apple notes, Bear)
* Folding header support (Notion, HTML)
* Frontmatter support (Notion, kind of)
* Saved views

It includes the seed of some of the more unique features I’ll be able to support, but is mostly focused on feature parity with respect to my personal use, i.e. if I build it, will I (once again) use it? I hope with this sprint I’ll begin to find out.

Further out I am both very excited for and also kind of dreading this amorphous “design” feature I have penned. In general, **I’ve been entirely focused on basic feature integrations. I don’t like the design as is. I want to like it more**. I’m not anticipating doing anything radical, spending too much time on it. But there’s a mix of styles and design components in the app now; I use [evergreen components](https://evergreen.segment.com/) for efficiency; I have some custom styles to match my blog; plate [has its own design requirements](https://github.com/cloverich/chronicles/pull/142) — so I’ve got [shadcn](https://ui.shadcn.com/) and tailwind floating around now too. It’s just a little unappealing and not cohesive. It also impacts some of the dependency consolidation and [clean-up](https://github.com/cloverich/chronicles/issues/157) I’d like to do. If I get the features to the point of basic parity, I need to get the design at least cohesive and pleasing at a basic level. Not polished. Not fancy. Really, just make it so I don’t get annoyed looking at it. 

Like any amorphous work, **it has potential to drag on**. To be too vague; to involve more changes than anticipated, especially if I need to consolidate the component libraries. It has potential to derail my momentum. Yet beneath the dread is a strong layer of excitement. Creating things is invigorating. When those creations are aesthetically pleasing, it takes it to another level. Playing with something I've created, however small, is so deeply satisfying; it reminds me why I switched into this career in the first place. Its like an antidote to burnout.


# The end game
 It’s been over **7 years since I started on the first iteration of this application**; many of my original ideas were tried and discarded. That’s good actually, I’m finding what works best. Many of my other good ideas have indeed been implemented by others at this point. That’s great too; as much as I enjoy building this app, I’d long hoped I ultimately wouldn’t feel the need to. I’ve tried to ditch the project and move on several times. Yet 7 years on, here I am. 

So how am I deciding what to even work on? As ever I’m selecting what to work on based on the features I use in other applications; or the features I wish I could use. At this point, I’m less concerned with anyone else actually using the application, as for now and perhaps ever, nobody will. What is the new user experience like? I don’t even know. Of course, I’ll address what all of that means in an as of yet unplanned sprint — one that will be planned if I think the app has any true merit. Instead, what will be worked on next, and next after that, etc, is what I would personally need to drop the other apps entirely. If I can’t do that, there is no reason for this application to exist. 

I used to worry this might be [a forever project](https://siddhesh.substack.com/p/projects) , but I no longer feel that way about it. I have many pie in the sky ideas, and many less grandiose and practical ones. **But below that lies a fairly discrete set of features I believe make the application usable**. It will never be “done”, but I definitely see a point where it can be 1.0. That point is much closer than I’d thought before; it’s what made it so hard to pick back up; so hard to put small incremental time into. 

So the “end” game in this sense is relatively straight forward now. 
* **Core features**: features I need to realistically use it
* **Polish**: features that make me like using it (design mostly)
* **Beginner UX**: ensure the app is usable, understandable, and sufficiently secure, from the vantage point of a random, new user

Ideally I’d love to get to the point I can share this project as such, and be proud of it. “This represents the best of my work.” It is pretty far from that now. What is not far, however, is how I now approach the work. It has near-term goals that are relatively well defined. I try and treat it as though people are using it: I try not to break it. I use a pull request flow, I document the work I am doing as I do it. I do proper releases, and make release notes. Ultimately this is all in service of practicing for the real deal. I don’t know this application will be something. But I think at some point, one of my projects (or contributions to someone else’s) might be. I find with that mindset, I’m able to actually enjoy the work that goes into leading this project to its next state. 

At any rate, I’m excited to make another attempt at this project, and for now enjoying the gratifying effect that building something has on my mentality. 


[^1]: As I typed this note out in Apple notes, it kept auto-correcting "reignition" it to resignation. It might have been a good name two years ago, when I switched from Chronicles to Notion and Apple notes. But frustration with those platforms is what led me back. Undesirable autocorrections which I cannot seem to turn off, are one of my least favorite parts of Apple notes. But I enjoyed this one.

[^2]: Alongside getting back into Chronicles development, I also recently picked up electric guitar -- after 20 years away. I was revisiting some of my former musical interests which included the old Metallica albums; I spontaenously ran into this quote from their original bassist and thought it was apt to my current state of mind. A lot of people feel burnt out, including myself. What has always surprised me, when I feel overworked, is how doing _more_ work can relieve the feeling. Ultimately what's happening in many cases is not over work, but a lot of work about work, a lot of ceremony, a lot of feature negotation and design by committee. Being able to break free of all that and just hack away on a side project, to move _fast_ again, is indescribably refreshing. 

[^3]: What I love most is how its all organized using another issue, with simple links, headings, descriptions. There's no Roadmap feature or formal "sprint" model, no date field for each ticket or story points. It instead feels more like the diligent work of a motivated project manager or team lead. Going through and organizing, figuring out what's realistic in the short term, and grouping related items so the next version will be a cohesive set of changes, and ideally the development work related. It was a good reminder to me of how often those things get in the way in the office; instead having this monthly document that just lists out the shit they want to ship, is all that's really needed. 