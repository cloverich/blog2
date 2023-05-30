---
title: "Adding a search tags interface"
author: Christopher Loverich
date: "2020-10-09T00:00:00.000Z"
tags: chronicles, devlog
description: "Notes on adding a search interface to chronicles"
---

## Introduction

In my last attempt to build [this app](https://pinecoder.dev/blog/2020-08-19/the-fifth-prototype), search organically grew in prominence as I added new features. Originally the app was just date ordered markdown documents, and search was free text with a few added features. In this version, I wanted search to be more central and flexible. Rather than an adjunct on the selected journal, I thought it could be the main driver of what was on screen. From the current versions inception I knew it could play a central role. However, I was surprised how much I struggled to come up with a design for it.

### Inspiration

The goal is flexibility. Instead of streamlining particular views (like selected journal) via tailored UI, I wondered if a generic search input could work instead. Search "tags" could fill a variety of use cases without tying the application to any particular UX. The application still (very) experimental, and I feel it is important to maintain flexible, generic components to drive as much of the experience as possible. 

To that end, I'd always found Github's search syntax inspiring. Plain text supports many kinds of use cases using a `prefix:value` structure:

![./github-search.png](./github-search.png)

Their [documentation](https://docs.github.com/en/free-pro-team@latest/github/searching-for-information-on-github/searching-issues-and-pull-requests) refers to the prefixes (`is:`, `label:`) above as *qualifiers,* but within this post and the code I refer to them as "search tags". At any rate, they support a wide variety of possible searches and modifiers (like `>` and `<` for dates). 

In addition, the Github interface has many supporting features which auto add, remove, or modify text on your behalf:

![./github-dropdown-example.png](./github-dropdown-example.png "The filters menu saves users from needing to remember and type various kinds of searchs")

 In looking for a generic search interface, I think this is a good fit. While it won't win any design awards, its a nice way to experiment with a variety of features without needing to commit to a particular design. Great for the experimental nature of this application.

## Specifying tags

There's a variety of search tags I could add to the app. Before beginning work on the interface, I'd thought through my commonest needs and decided on these three tags to start:

### `in:`

This tag specifies which journals to search. Example: `in:work` or `in:chornicles`. Its the least interesting in that it merely replicates existing behavior (selecting a journal to view) into the search input. 

### `focus:`

The focus tag is a search and a command. It takes my previously implemented [focused headings](https://pinecoder.dev/blog/2020-09-16/focused-headings) feature and re-implements it as a search tag. This one is tricky because it needed to not only update the UI if you enter this tag, but also add this tag to the input based on a user clicking on a heading (how I focusing a heading works now). 

### `filter:`

This tag is also a search and command. It searches for documents containing the markdown node (ex: `link`, `inlineCode`) and then *filters* rendered documents to only contain the matched content. The intent of this tag is to generalize support for the following two examples:

- Show me all links in this journal: `filter:link`
- Show me all SQL snippets: `filter:code[lang="sql"]`.

There's several others I had in mind (including free text), but I thought this group required implementing several key behaviors in the app. Thus I thought if I could get these working well, adding others would be a simple exercise.

## Implementing the interface

The existing interface exposes a single dropdown menu to select the active journal

![./dropdown.png](./dropdown.png "The existing, single select dropdown menu")

From the UI's perspective, there is no search, just the 10 most recently created documents in the selected journal. However the backend actually interprets a selection as a search request. Selecting "chronicles" above would produce a query that looks like:

```tsx
query: {
  journals: ["chronicles"],
}
```

So converting the dropdown to free text would be mostly an exercise in UX.

Despite wanting free text, I'd like to make it easy to quickly search by journal, in effect maintaining the existing behavior. After thinking on it, and exploring the [evergreen-ui components](https://evergreen.segment.com) (currently using in the app), I decided the [side sheet](https://evergreen.segment.com/components/side-sheet) menu could offer a solution

- Click a button to slide out a menu listing the available journals
- Clicking a journal sets the search input to: `in:<journal name>`

![./sidesheet-mockup.png](./sidesheet-mockup.png "An excalidraw mockup of the sidesheet select interface")

This is in effect still a select menu. Select menu's are of course powerful, which I believe explains their ubiquity. Less obvious until you've built interfaces for a while is how many unique interfaces can be boiled down to select menus. If I were stopping here, it would be a silly interface. But of course, the idea was to maintain the existing behavior, and provide a canvas for bigger ideas. 

Once i'd decided on a starting point, the work flowed quickly. With the help of evergreen's tag input and sidebar components. Initially the tag input was read only, click sidebar to enter it. After fiddling with the UI design just a bit, I settle on the search input and two buttons (toggle sidebar, create entry).

![./intag-example.png](./intag-example.png)

It worked well — but at this point all I had was a fancy but slightly *worse* select menu.

### Multi-select, No select

Next I wanted to implement searching multiple journals, where each journal can be independently added or removed from the search. Searching for journals involves a query that looks like: 

```sql
query: {
  journals: ["blog-ideas", "work"]
}
```

The prior, single select menu would operate on it by simply replacing the 0th element (`query.journals[0] = selected`). After migrating to multi-select, the UI reflected the possible query operations more closely: `in:` tags are a direct representation the `journals` array. 

I enabled clicking the `x` on each tag, and it immediately opened a new possibility: **Clearing all search tags**. This could mean displaying the empty view, but I opted for the opposite: No `in:` tags means "display content from *all* journals". Interestingly, this behavior actually aligns closely with a typical user experience: Having only one journal. I had not intended this initially, and on reflection realized how important this behavior would be. `in:` tags are only necessary when you have multiple journals. Otherwise, you can ignore the feature all together, and the application will work as expected. 

### The `focus:` tag

Because I'd already implemented the focused heading behavior, implementing the tag involved two simple steps:

- When the app is focused, make sure it has a `focus` tag
- When clearing a focus tag, clear the focused state

Since these states were already merely derived from the query state, I only needed logic in the `add` and `remove` handlers from my search input that would update the query state; everything else worked out of the box. This became my [first stopping point.](https://github.com/cloverich/chronicles/commit/5ac96cf8fdc002425feffbff8641783431393fc0) 

### The `filter:` tag

This tag has two parts:

- Parsing is slightly more complex: it must support the simpler `filter:code` and attribute based `filter:code[lang="sql"]`
- Filtering markdown: Documents must understand how to filter markdown content based on the filter's value.

I added a constraint which made parsing relatively straight forward — I support only a single attribute (`lang="sql"` above). In this way I could leverage a regex to do all the work for me. 

For filtering actual content, I took the same path as focused headings and filter content within the rendering component. Given the content available in each document is merely a sub-set of the full content, the backend would be perfectly capable of handling this filtering behavior. However, the initial design didn't support this, and I don't have a compelling enough reason to change it yet. There are numerous performance enhancements I could implement — caching in a service worker, pre-parsing and rendering in in the backend, and more. But at this point, the app is still fairly basic the the performance is suitable. Keeping notes on where the low hanging fruit lives is sufficient for now. ([commit](https://github.com/cloverich/chronicles/commit/a61ea0d4d4dde8fa4e40cdc6da68486da99307fb))

## Epilogue

The app now has a generic search interface ([release here](https://github.com/cloverich/chronicles/releases/tag/v0.2.1)), and a few tags demonstrating key behaviors:

- Finding content in multiple journals
- Filtering content by type (i.e. code, links, headings)
- Search tags that support custom layouts (focused headings)

There are plenty more features I'd like here (wildcard support, saved searches, free text), and this groundwork gives me a canvas to add those and more ambitious features with relative ease.  However, while this application is beginning to take shape, I still consider it an experiment. Thus its important to not linger too long on rounding out everything. Instead, I need to keep pushing ahead on the features I think will define whether or not the project is worth continuing. 

Lastly, there were numerous challenges in getting the new UX functioning properly. It is difficult to capture in a single blog post just how many little changes were necessary, and I think this is part of the reason UI development can be so challenging. In a follow up, I'd like to dig into one part of that — representing and reacting to search state.