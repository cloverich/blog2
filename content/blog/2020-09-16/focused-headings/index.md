---
title: Focused headings
tags: devlog
date: "2020-09-16T11:00:00.000Z"
description: "Devlog: Summary and implementation notes on \"Focused Headings\" in Chronicles."
---


I finished another feature on [Chronicles](https://github.com/cloverich/chronicles): "Focused headings" ([demo](https://youtu.be/qmnay59SvY4)). Focusing a heading finds all content under that heading, filters everything else out, then pins the heading to the top of the journal. 

A journal might typically have a range of topics across many days:

![assets/Screen_Shot_2020-09-16_at_6.22.33_AM.png](assets/Screen_Shot_2020-09-16_at_6.22.33_AM.png?resize=blogImages "Typical journal with multiple topics")

After "Focusing" a heading, the heading element is fixed to the top, and the view filters any content not under that heading:

![assets/Screen_Shot_2020-09-16_at_6.22.38_AM.png](assets/Screen_Shot_2020-09-16_at_6.22.38_AM.png?resize=blogImages "Focusing the blue heading")

The primary inspiration was wanting **a way to take running notes on similar topics, then have a way to surface them all in one stream** — as though the entire journal was about *just* that topic. The pattern works well for meeting notes, recurring themes, or any journal where you mix multiple topics. I most often use it to track work across multiple days (like this feature, for instance). Below i document a summary of the implementation choices and a few alternatives I may consider in the future.

### Background

First, there is some background information that is required, covering how documents are fetched and rendered today. 

When a journal is loaded, the app makes a (local) api search request for documents. It gets backed markdown parsed into a syntax tree ([MDAST](https://github.com/syntax-tree/mdast)). Individual `Document` components then render the markdown AST into html with [remark-html](https://github.com/remarkjs/remark-html).

![assets/container-mdast.png](assets/container-mdast.png?resize=blogImages "Container component passes parsed markdown trees to multiple Document components. They use remark-html to render the parse tree into html.")

So to implement the feature, I wanted to:

1. Search documents by heading when a heading was clicked
2. Filter the results to display only that content "under" a heading

## Implementation

There are three key components to the implementation:

- Generating a "focus" event when headings are clicked
- Processing the focus event
- Rendering the result

### Generating the "focus" event

Generating and catching a click event in html is straight forward, but there's a devil in the details here: I need the clicked *html* element to include information about its *source markdown*. When rendering markdown with remark-html, some of the source information is lost. Since my backend currently indexes raw markdown, I need the original markdown string to search on. After some research and hacking, I boil it down to a two step process:

- Annotate the source AST with the headings original string
- Render react components, instead of direct html, so I can emit a custom event on click

I use two libraries: [remark-rehype](https://github.com/remarkjs/remark-rehype) and [rehype-react](https://github.com/rehypejs/rehype-react) to get from my markdown AST to react components. Setting up the compiler is revealing:

```tsx
const compiler = remark()
  .use(remark2Rehype)
  .use(rehype2React, {
    createElement: React.createElement,
    passNode: true,
    components: {
      h1: Heading as any,
      h2: Heading as any,
      h3: Heading as any,
      h4: Heading as any,
      h5: Heading as any,
      h6: Heading as any,
    },
  });
```

The `passNode` argument indicates I want the compiler to pass down the source element to the component rendering the output. This provides the mechanism to pass information from the original tree to be used in the emitted event. 

```tsx
export function annotateHeadings(root: Root) {
  for (const child of root.children) {
    if (child.type === "heading") {
      child.data.hProperties = {
        remarkString: remark.stringify(child),
      };
    }
  }
}
```

My custom `Heading` element can then pick that property up, and emit it:

```tsx
function Heading(props: HeadingProps) {
  const handler = (evt: React.MouseEvent<HTMLHeadingElement>) => {
    evt.target.dispatchEvent(
      new CustomEvent("focus-heading", {
        bubbles: true,
        detail: {
          depth: props.node!.tagName,
          content: props.node!.properties.remarkString,
        },
      })
    );
  };
```

### Processing the event

Once an event is generated, there's a few steps to processing it:

- Execute a new search
- Pin the heading to the top of the journal
- Render the filtered result

For this step, I refactored my code a bit. I migrated some disparate React (hook) state into a consolidated view model, which for now I've simply called the `JournalUIStore`. 

### Rendering a filtered result

I want the rendered documents to display only the content that is "under" the focused heading. I interpret that as all elements after the focused heading, until another heading of equal or greater precedence is encountered. Because my `Document` components already have a handle on the AST, they can simply filter down to the desired output.

This turned out to be straight forward, as the tree structure is actually quite flat. See for example this parsed document from [astexplorer](https://astexplorer.net):

![assets/Screen_Shot_2020-09-15_at_9.02.12_AM.png](assets/Screen_Shot_2020-09-15_at_9.02.12_AM.png?resize=blogImages "Parsed AST, in JSON")

Filtering out all elements between a heading of a given `depth` and the next heading, or end of the document, is then mostly just iterating over it. 

With that in place, the basic feature works. You can [view a demo here](https://youtu.be/qmnay59SvY4). 

## Future architecture decisions

I took a few shortcuts, and practiced "make it work, then make it right" a bit because I did not start with much intuition on how to do this feature. In the end, the tradeoffs are probably good for this phase. I have a few alternative implementations and optimizations on my mind, and it's probably best to make note of them now. If the product lives long enough, I can re-visit. 

### Filter content server side

A key decision here was whether to filter the markdown on the backend, or in the UI. I initially wanted to filter on the backend, and keep the UI "dumb". This would have required moving too many pieces around, changing the search implementation, and spreading the feature across the UI and API (UI knows a heading is fixed, API is filtering content). Additionally the search result and document fetching have a simple cache that would need adjusted. Overall, although I still think this method has more potential, I opted to defer it for now.

### Pre-generate annotated headings

There's no strict reason I need to attach an event listener to each rendered heading element, except I have no other way to embed the original markdown string into the click event. Because I don't have a formalized concept of parsed node ids in my backend index, I couldn't think of mechanism for mapping the clicked html element to the original markdown node. 

But if I could, it opens the door to attaching the identifier to the html element directly in a `data-*` attribute. What that does is remove the need to render a *custom* element and event: I could simply catch click events from the container element, and then inspect the data-* attribute on the event target. The real performance gain there is I could **pre-render the html prior to the UI ever receiving it, and even cache it**. 

## Parting notes

- Restrained design
- Next steps

I am pretty happy to ship this feature. It felt more like a grind than the others but is one of the core features I've listed as the reason I am making this app in the first place. It was **hard to restrain myself from enhancing the design**.I can see its not optimal — you really have to know this features is there *and* why you'd want to use it to get any use out of it. However, I expect that the UX will change significantly as I land the next couple of features (collapsable headings and documents, search interface). Also, I don't have a strong intuition on what the final direction will be. It would be easy to sink a bunch of time optimizing design and UX here, only to throw it away or worse, be unsatisfied with the result anyways. Instead, I am opting to move forward towards the other defining features of the app, and re-visit the UX when the core features are in place.

As for next steps, I will either tackle **collapsable headings** or **the search interface**. I've put off the latter because the UI is a bit of departure from the simple dropdown I have today, and there are a few tricky UX cases. It is also going to leave the app in a less intuitive state, even while simplifying and enhancing the power of the interface. But ultimately together they form the set of features I think will make this app worth using.