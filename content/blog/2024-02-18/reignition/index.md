---
title: Reignition
date: "2024-02-18T00:00:00.000Z"
author: Christopher Loverich
tags: devlog
description: "Recapping Chronicles release v0.4.0"
---

As of today, 2024-02-18, I [released v0.4.0 of Chronicles](https://github.com/cloverich/chronicles/releases/tag/v0.4.0) and intend to **wrap up the last six weeks of feature work on Chronicles** and take a short break to catch up on life. I hope the momentum from this work, a desire to see things through, and handling a few life issues that need more attention, will not turn into me putting this project aside for another 2 years. I hope it has the opposite effect. But to that end, and to the end of approaching things differently in general, I want to spend a little time wrapping up. Doing a release; adding notes. Writing an epilogue here. Writing several other epilogues on the features I did, specifically so I can understand the design choices and how things worked if I need to back reference. I also want to brain dump my state of mind in regards to the project. 

![Github insights of the last months work](assets/github_insight_feb2024.png?resize=blogImages)

# Release v0.4.0
I called the [last release](https://github.com/cloverich/chronicles/releases/tag/v0.3.0) “Base Usability — Well kind of”. “Base Usability” was the original name for the set of features that went into it; “Well kind of” was me appreciating that the actual base usability I seek will be much bigger in scope than a single release, and better suited to be done across several sprints and releases. Moreover the last release was merely closing out some long unfinished work; the current release includes a fresh set of features, fixes, and improvements[1] (well detailed in the release notes). Additionally I [used a sprint](https://github.com/cloverich/chronicles/issues/116) to organize the work, and was able to follow through with it. I found my planning is more realistic and cohesive than before. This was extremely helpful to carry the work through, because with kids and a full time job, I **generally did so in 30-90 minute increments**.

Using a traditiojnal pull request flow, instead of pushing to master, helped keep things atomic and organized. I kept the commits and PRs fairly clean and complete. Linking them back to issues, and adding documentation to the PR as I went. This was in part motivated by my prior experience dropping off and picking up of the project; **the easier it is to understand where you left off, the easier it is to get moving again**. This also helps when creating Github releases; using a pull request flow means Github can auto-create links to the pull requests since the last tag, and summarizes them. I’ve used [commitlint and the conventional changelog](https://github.com/conventional-changelog/commitlint) approach before, but found it too cumbersome for a personal project. I am pretty happy with the current approach. 

Taken together, the PRs, release notes, and my personal progress notes have become a sort of rolling epilogue. Over the years, I’ve become more organized and methodical in my regular note taking — taking notes as I develop. Those notes are more useful than they used to be. I enjoy reading them more than I used to, and do so more often. I don’t know what impact this note taking application will have on my life; if anyone will use it; if it’s ultimately worth doing. But what I do know is the way jouraling has impacted my life, has been extremely positive. And at the moment, **of all the good ideas in the world that I am not doing, this one blip of a good idea has momentum** — its a rolling project I can incrementally build off of. Overall, I found the end result of this invigorating and easy; when I reflect on the last month, I feel the project is fully alive again. Hence, reignition[^1]


# What comes next
I am extremely excited about [the next sprint](https://github.com/cloverich/chronicles/issues/123). Unlike the last one, it will be **much more focused on features**. Notably it will include several key features I rely on in other apps, such as:

* Internal note linking (Many)
* Tagging (Apple notes, Bear)
* Folding header support (Notion, HTML)
* Frontmatter support (Notion, kind of)
* Saved views

It includes the seed of some of the more unique features I’ll be able to support, but is mostly focused on feature parity with respect to my personal use, i.e. if I build it, will I (once again) use it? I hope with this sprint I’ll begin to find out.

One I am both very excited for and also kind of dreading, is an amorphous “design” feature. Up to the point of those features, as now, **I’m entirely focused on basic feature integrations. I don’t like the design as is. I want to like it more**. I’m not anticipating doing anything radical, spending too much time on it. But there’s a mix of styles and design components in the app now; I use [evergreen components](https://evergreen.segment.com/) for efficiency; I have some custom styles to match my blog; plate [has its own design requirements](https://github.com/cloverich/chronicles/pull/142) — so I’ve got [shadcn](https://ui.shadcn.com/) and tailwind floating around now too. It’s just a little unappealing and not cohesive. It also impacts some of the dependency consolidation and [clean-up](https://github.com/cloverich/chronicles/issues/157) I’d like to do. If I get the features to the point of basic parity, I need to get the design at least cohesive and pleasing at a basic level. Not polished. Not fancy. Really, just make it so I don’t get annoyed looking at it. 

Ultimately I am excited because if I can do that, I know it will instantly be more satisfying to use on a day to day basis. But like the refactoring work in the last sprint, **it has potential to drag on**. To be too amorphous; to involve more changes than anticipated, especially if I need to consolidate the component libraries. It has potential to derail my momentum. I guess it’s an opportunity to see if I actually am better at tackling that kind of work, in a restrained fashion. 

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


[^1]: As I type this, Apple notes auto-corrects it to Resignation. This gave me a good chuckle. It might have been a good name for the last release, when I switched from Chronicles to Notion and Apple notes. Frustration with those platforms is what led me back. Undesirable autocorrections which I cannot seem to turn off, are one of my least favorite parts of Apple notes.
