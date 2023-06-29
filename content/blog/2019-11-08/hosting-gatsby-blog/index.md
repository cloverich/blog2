---
title: Setting up a Gatsby blog
date: "2019-11-08T00:00:00.000Z"
description: "An up and running guide to setting up, customizing, and hosting a blog using a static site generator"
---

In this post, I'm going to cover creating your own blog with Gatsby, a React-based Static Site Generator.

- [Motivation](#motivation)
- [Content and static site generators](#content-and-static-site-generators)
- [Up and running with Gatsby](#gatsby)
- [Customizing](#further-customizations)
- [Hosting](#hosting)
- [Syndication](#syndication)
- [Next steps](#wrap-up)


## motivation
Before getting to the background, lets (briefly) start with why. You could publish blog posts to [Medium](https://writingcooperative.com/a-quick-guide-to-medium-9f61864ab26), [Wordpress](https://wordpress.com/), or [Ghost](https://ghost.org/) and be up and running immediately. That's a perfectly fine route to go. Creating your own blog is something you do because you want to own your website. You want a custom domain name and control of the content. You want custom pages, or to use it as a platform for a side project. Maybe you just want something to hack on. Its more work than using one of the afformentioned blog platforms, but thanks to abundant open source tooling, only a little.

This post is about getting started on that route with a site generator named Gatsby. Its a tutorial and is targeted towards developers familiar with React. There are many other Gatsby blogging tutorials out there -- this one focuses on some specific customizations I personally used. It also includes practical hosting information. If you want to follow along, you'll need [node](https://nodejs.org/en/) and [yarn](https://yarnpkg.com/en/docs/install). 

## content and static site generators
Of the numerous options for creating your own blog so-called static site generators (hereafter "SSG"s) hit a particular sweet spot. You author content (i.e. blog posts) in a format like markdown, then generate html documents using a build script. Templates are used to generate layouts and non-content pages. The result is a complete website - a set of html files and associated assets (images, css, fonts). Because the pages are pre-generated, there's no need to serve content directly from a database. The files are typically uploaded to a [CDN](https://www.cloudflare.com/learning/cdn/what-is-a-cdn/) giving you cost efficient scalability. The entire process can be automated out of a Github repository, providing a nice "author - commit - publish" workflow. 

There are many established SSGs in widespread use. You can find one in your language or platform of choice [here]((https://www.staticgen.com/)). They all share at least two important features:

- the ability to generate content pages (blog posts) from text files (like markdown)
- the ability to generate non-content pages from templates (ex: an about page)

![high level overview of an ssg](./ssg-overview.png?resize=blogImages)

Templates are a core part of any particular SSGs identity. Most languages have at least one, and usually several templating languages suitable for generating html. Javascript has a dozen or more, but React was not traditionally considered one of them. That's because React's core abstraction is not static html generation, but components. Components have lifecycle's and state. They respond to user events and update the DOM. React's ability to [render to static markup](https://reactjs.org/docs/react-dom-server.html#rendertostaticmarkup) is just a feature.

But given React's popularity, the transition to using it as a template engine was natural and perhaps inevitable: If you already know JSX, why bother learning another templating language that does (seemingly) the same thing.  Multiple React based SSG's now exist, but many depart from traditional SSGs in an interesting way: They can re-claim a statically generated html page and turn it back into a React application, after its been loaded. There are many reasons you might do that, including pre-rendering, progressive enhancement, and [many others](https://github.com/react-static/react-static/issues/1180). There are good reasons to choose one approach over another, but I won't cover much of that here. For the purposes of this post, we'll focus on the fact that **React-based SSGs offer a familiar templating language and ecosystem to the developer already familiar with React**.

## gatsby
[Gatsby](https://www.gatsbyjs.org/) is one such React-based static site generator. I chose Gatsby as a starting point because it **provides boilerplate that includes  everything needed to run a blog out of the box**. After installing and making a few tweaks, you'll have a site that:

- Looks good
- Can be served entirely from a CDN
- Includes numerous optimizations ([example](https://image-processing.gatsbyjs.org/))

Its magical but [well documented](https://www.gatsbyjs.org/docs/). Most of its functionality is provided by plugins, and its site includes a [nice UI for exploring](https://www.gatsbyjs.org/plugins/) these. We'll cover some of these in this tutorial, but for now lets install the blog boilerplate and get the site running:



```sh
# Install the gatsby-cli tool and make it globally available (eww)
yarn global add gastby-cli

# generate a new project using "gatsby-starter-blog" and cd into it
gatsby new <blog-name> https://github.com/gatsbyjs/gatsby-starter-blog
cd <blog-name>

# start the development server
gastby develop
```

The tool scaffolds a node project and installs its dependencies. `gatsby develop` starts a development server and if you navigate to http://localhost:8000, you should see something like this:

![starter blog index with some pre-filled content](./gatsby-blog-starter.png?resize=blogImages)

The starter is pre-populated with a bio section and some content. Its merely an example of how to populate your blog, but note that it looks quite similar to [Dan Abramov's popular blog](https://overreacted.io/). This is what makes Gatsby a great starting point: **by editing this content and some metadata you are ready to publish your site**. I'll cover how to do this next

## Digging into Gatsby
Personally, I do not enjoy walking through a tutorial without some understanding of the system I"m using. This section discusses a bit about how  Gatsby works. 

Conceptually, Gastby's build phase can be broken into two steps: Data loading and data consuming. The interface between these two steps is... a GraphQL schema. Navigate to http://localhost:8000/___graphql to see your sites schema. You don't need to understand much of how GraphQL works to understand Gatsby however, only that its used to bridge the main build steps. For example, consider how Gatsby creates pages from blog posts: 

- **Load content**: Read markdown files from our `/content` directory and load that information into the schema
- **Consume content**: Read blog post data out of the schema and generate pages with it

Of course its [more complicated](https://gist.github.com/sw-yx/09306ec03df7b4cd8e7469bb74c078fb) than that, but that's how it works at the highest level. 

In addition, I also find it useful to understand the purpose of the various files in your site. There are three conceptual categories:

- **configuration files**: in the root of the project. These control Gatsby itself, including its plugins, where it looks for content, what pages it generates
- **React components**: representing layouts, pages, and parts of pages.
- **content files**: In `/content`, these are markdown files and images

See [this post](https://www.narative.co/articles/understanding-the-gatsby-lifecycle) for a more details.

## Make the site your own
With that out of the way, we can dive into editing your sites files to remove the placeholder content and add your own. We'll do the following:

- change the site metadata in `gatsby-config.js`
- change the bio component (unless you are, in fact, [Kyle Matthews](https://www.gatsbyjs.org/contributors/kyle-mathews/))
- delete the existing blog posts
- add your own blog post

First take a look at `gatsby-config.js`. This includes site metadata and configuration for various plugins used in the site. We'll modify some of the plugins in a moment but for now, take note of the `siteMetadata` section; it includes data used to populate the bio section:

```js:title=gatsby-config.js
module.exports = {
  siteMetadata: {
    title: `Gatsby Starter Blog`,
    author: `Kyle Mathews`,
    description: `A starter blog demonstrating what Gatsby can do.`,
    siteUrl: `https://gatsby-starter-blog-demo.netlify.com/`,
    social: {
      twitter: `kylemathews`,
    },
  },
```

Change title, author, description. After saving, the site will live update and should now show your name instead of Kyle's. The social section can be changed as well, but first delete it outright and save. This will produce a build error, as the `Bio` component expects this data to be present; it should display a trace in your browser:

![build errors are reflected in the browser](./build-error.png?resize=blogImages)

The error is not particularly helpful, but it does identify the component that threw the exception (`Bio.js`). Re-add the social section if you want to display your twitter handle, otherwise leave it out. 

Navigate to `src/components/Bio.js`. This file exports a single function - a stateless React component. It has two notable areas:

- A data query (`useStaticQuery`)
- A JSX expression that controls the displayed output

The data query defines the data available to the component, which includes an image used for the profile and the site metadata. You can replace the existing  profile pic (at `/content/assets/profile-pic.jpg`) with one of your own. If the filename is different, be sure to update the name in the query here as well. If you are wondering how Gatsby knows where to look for the file, its controlled by the `gatsby-source-filesystem` plugin, configured in `gatsby-config.js`:

```js:title=gatsby-config.js
plugins: [
  // ...
  {
    resolve: `gatsby-source-filesystem`,
    options: {
      path: `${__dirname}/content/assets`,
      name: `assets`,
    },
  },
```

We'll come back to this in a moment. Back in the component file (`Bio.js`), note how author and the twitter handle are pulled out of metadata; scan the components JSX and you should see where its used. Review the rest of the JSX and change it to suit your tastes. At this point the bio section should look the way you want. 

For the last step, we'll modify the blog content. You'll find the existing posts in `/content/blog/` and see that each post is in its own folder. Remove all but one, and modify that last one to be your first post. ...And that's it, this site is now yours. Almost.

## further customizations
You could publish your blog right now, but for my personal site I wanted to make the following changes

- customize typography
- add code highlighting
- adjust the url for posts

I'll explain how to make these changes next. I'll also explain a little more about how Gatsby works. If that's not interesting to you, [skip ahead](#hosting) to publishing.

### customize typography
Since this is a blog, typography will be the principle component of your sites design. [Typography.js](http://kyleamathews.github.io/typography.js/) is a library that helps simplify the process of configuring good typography settings by providing **pre-created typography themes and a way to easily tweak them**. This library comes pre-installed with the blog starter we've selected, so to customize your typogrphy you need only use [the website](http://kyleamathews.github.io/typography.js/) to pick a theme and settings you like, then apply them to your app. If the concept of typography feels overwhelming, check out [this guide](https://www.pierrickcalvez.com/journal/a-five-minutes-guide-to-better-typography) for a brief tutorial on how to think about this. 

This starter uses the wordpress 2016 theme (at the time of this writing). You can see the associated dependencies in `package.json`:

```json:title=/package.json
  "gatsby-plugin-typography": "^2.3.12",
  "typography": "^0.16.19",
  "typography-theme-wordpress-2016": "^0.16.19"
```

To configure the theme and customize it, the plugin (`gatsby-plugin-typography`) declares itself in `gatsby-config.js` and points to a file supplying the options:

```js:title=/gastby-config.js
{
  resolve: `gatsby-plugin-typography`,
  options: {
    pathToConfigModule: `src/utils/typography`,
  },
},
```

Navigate to `src/utils/typography` and you'll see how the plugin is customized to get the starter template's styles:

```js:title=/src/utils/typography.js
import Typography from "typography"
import Wordpress2016 from "typography-theme-wordpress-2016"

Wordpress2016.overrideThemeStyles = () => {
  return {
    "a.gatsby-resp-image-link": {
      boxShadow: `none`,
    },
  }
}

delete Wordpress2016.googleFonts

const typography = new Typography(Wordpress2016)

// Hot reload typography in development.
if (process.env.NODE_ENV !== `production`) {
  typography.injectStyles()
}

export default typography
```

The most important things to note:
- The wordpress-2016 theme is imported (as the package, `typography-theme-wordpress-2016`). 
- the `Typography` object is instantiated with the theme
- It is exported as the default export

The last point is particularly important, as Gatsby's configuration relies on the default export being the customized typography object. At this point you should select a typography theme of your choosing, then customize it. As an example, if you wanted to use the "US Web Design Standards" theme, you'd do this:

```sh
yarn remove typography-theme-wordpress-2016
yarn add typography-theme-us-web-design-standards
```

Then change the imports:

```ts
// replace this line:
import Wordpress2016 from "typography-theme-wordpress-2016"
// with this one: 
import USWebStd from "typography-theme-us-web-design-standards"
```


### adjust urls
 By default the project is configured generate a [slug](https://yoast.com/slug/) in a simple fashion: `https://yoursite.com/slug-here`. For my site I wanted to prefix all blog posts with `/post`, and include the date of the post in the url, ex: `/post/2019-10-21/my-first-post`. I'm doing this to maintain flexibility in my sites first path (`/post` above) and I also personally like seeing the date in the URL of posts I read. 

As is Gatsby is mostly copying the directory structure of your `/content` folder, so you could accomplish this by arranging your files however you want them to appear in the url. However, I'd prefer to not tie my url structure to my current folder hierarchy, and I'd like to change them independently of one another. The official Gatsby tutorial covers this:

- [Programmatically create pages from data](https://www.gatsbyjs.org/tutorial/part-seven/). 

What we'll want to do is hook into Gatsby's node creation step, modify the markdown content nodes to add a field called "slug", then use that field in page creation. All of this will take place in the file `gatsby-node.js`. To understand what we need to do, first understand the hooks we'll be using:

```js:title=/gatsby-node.js

// This hook is called when a GraphQL node is _created_
exports.onCreateNode = ({ node, getNode, actions }) => {

// This hook is called when a GraphQL node is _consumed_
exports.createPages = async ({ graphql, actions }) => {

```

With those two hooks we can do the following:
- in `onCreateNode`, check if the node is a markdown content node
- if so, generate a slug 
- in `createPages`, use the `slug` field to set the the url


```js:title=/gatsby-node.js
const { createFilePath } = require(`gatsby-source-filesystem`)

exports.onCreateNode = ({ node, getNode, actions }) => {
  const { createNodeField } = actions

  // Here, check if its a markdown Node:
  if (node.internal.type === `MarkdownRemark`) {

    // If so, create a slug based on the filepath
    const slug = createFilePath({ 
      node, 
      getNode, 
      basePath: `pages` 
    })

    // Add the slug's value to a new field in this node, and 
    // name it "slug" 
    createNodeField({
      node,
      name: `slug`,
      value: slug,
    })
  }
}

exports.createPages = async ({ graphql, actions }) => {
  const { createPage } = actions

  const blogPost = path.resolve(`./src/templates/blog-post.js`)

  // Load the markdown data
  const result = await graphql(
    `
      {
        allMarkdownRemark(
          sort: { fields: [frontmatter___date], order: DESC }
          limit: 1000
        ) {
          edges {
            node {
              fields {
                slug
              }
              frontmatter {
                title
              }
            }
          }
        }
      }
    `
  )

  // For clarity, use a more sensible name
  const posts = result.data.allMarkdownRemark.edges

  // Create a page for each post
  posts.forEach((post, index) => {
    const previous = index === posts.length - 1 ? null : posts[index + 1].node
    const next = index === 0 ? null : posts[index - 1].node

    createPage({
      // Create the URL we want
      path: post.node.fields.slug,
      component: blogPost,
      context: {
        slug: post.node.fields.slug,
        previous,
        next,
      },
    })
  })
}

```


### code highlighting
If you plan on including code snippets in your blog, you'll probably want to add code highlighting. This is not included out of the box with the default starter, but its a simple add. We'll use [Prisma](https://prismjs.com/index.html), again mostly because there is [documented support](https://www.gatsbyjs.org/packages/gatsby-remark-prismjs/) for it. Setup involves:

- Install dependencies
- Add the plugin to `gastby-config.js` 
- Select and add a CSS theme file to `gatsby-browser.js`

Understanding the basics of how the plugin works can give you intuition about the steps above. The plugin operates as part of the markdown phase and goes something like this:

- The (already installed) remark plugin parses your markdown and identifies code sgements
- Prism iterates each code segment and parses them based on the identified language
- The parsed keywords are wrapped with `span` elements and given CSS class names depending on the token (ex: `<span class="token string">foo</span>`)
- A CSS "theme" declares styles for those classes

First, install these dependencies:

```sh
yarn add prismjs gatsby-remark-prismjs prism-themes
```

Next, update `gatsby-config.js` to tell the remark plugin to use the prism plugin for code blocks:

```js:title=/gatsby-config.js
{
  resolve: `gatsby-transformer-remark`,
  options: {
    plugins: [
      // ... other plugins
      {
        resolve: `gatsby-remark-prismjs`,
        options: {
          classPrefix: "language-",
          inlineCodeMarker: null,
          aliases: {
            sh: "bash"
          },
          showLineNumbers: false,
          noInlineHighlight: false,
        },
      },
```

That sets up the parser, now we need a theme. Prism includes several themes in its `/themes` folder. You can view these by navigating to `./node_modules/prismjs/themes`, and apply them by importing the css file into `gatsby-browser.js`:

```js:title=/gatsby-browser.js
import "prismjs/themes/prism-dark.css"
```

In case you do not like any of the default themes, you can choose one of the extended themes from `prisma-themes`, which I had you install above. ex:

```js:title=/gatsby-browser.js
import "prism-themes/themes/prism-atom-dark.css"
```

Lastly if you want to customize it further, simply create your own css file and copy the contents of one of the themes into it. Import that file instead of the theme file and customize as you see fit. 

---

So that's it for customizations. There's myriad more plugins and customizations and as Gatsby is ultimately a React app, you could endlessly tinker and never write an actual blog post. Let's setup hosting now so that doesn't happen to you :)

## hosting
At this point you can serve your site from any webhost capable of serving static files. Some options:

- [Github pages](https://pages.github.com/)
- [AWS](https://aws.amazon.com/getting-started/projects/host-static-website/)
- [Netlify](https://www.netlify.com/)
- ...many more

Run `gatsby build` to generate the static website into the `/public` folder. You could upload these files directly to one of the above hosts, or use a CI system to do this step for you. I"m going to demonstrate Netlify, because it has a built-in CI system that will do this step for you, triggered directly off commits to a Github repostitory. After setting up a Github repository for your project and an account with Netlify, the workflow operates like this:

- Netlify listens for commits via a webhook
- After each commit, it runs `gatsby build`
- After files are built (into `/public`), it copies these to its CDN

At that point, your updated site is available online.

### getting started with Netlify
Create an account on [Netlify](https://www.netlify.com/), then navigate to Sites > New site from Git. Follow the prompts. I add Netlify to my Github account, then give it selective access to the repo I want to publish. Notice you can configure which branch Netlify deploys -- this is one option for setting up different deployments so you can demo work before releasing it. For now I'm publishing my site as is, so I tell it to deploy from Master. It auto-detects I'm using a Gatsby project and infers the build command and directory:

![I indicate which branch (master) to use. It auto-detects gatsby and pre-configures the build command to `gatsby build`](./netlify-build-setup.png?resize=blogImages)

To re-iterate the above configuration: Commits to your **master** branch trigger Netlify to update and run **`gatsby build`**, and then to take the output from **`/public`** and publish the results. After saving this configuration, Netlify has enough information to host your blog at a randomly generated url (mine was `upbeat-shockley-54ce48.netlify.com`). Once its initial run completes, you can see for yourself. 

### configuring your domain
Next comes the domain. Netlify offers to set one up for you which is a fine route to go. I happened to already have a domain name that I purchased previously. Netlify again provides prompts to help you out if you've never done something like this. The intuition is simple enough: Create DNS records so that when someone types `yourdomain.com` into their address bar, the browser resolves the domain name to the Netlify host that is serving your website. You have two options:

- Setup DNS records yourself, pointing to the Netlify host that has your website
- Defer DNS resolution to Netlify and let them manage the DNS records for you

The second case is a bit simpler, so I'll demo that next. Use the "Custom domains" section of your site in Netlify to add your custom domain, then follow its prompts. It will prompt you with something like this:

![Netlify DNS prompt](./netlify-dns-config.png?resize=blogImages)
![Netlify DNS update information](./netlify-nameservers copy.png?resize=blogImages)

In my case since I registered my domain name with [Namecheap](https://www.namecheap.com/), so I do the following: Login to Namecheap, navigate to my domain (pinecoder.dev), and find the "Nameservers" section. Change "Namecheap BasicDNS" to "Custom DNS" and add the Netlify domain servers provided above:

![Select "Custom DNS" from the dropdown](./namecheap-custom-dns.png?resize=blogImages)
![Add each Netlify DNS server from Netlify's list](./namecheap nameservers.png?resize=blogImages)

Once saved it will take a bit to update, but Netlify will soon detect that your domain is using its DNS. From there, it will autoprovision https for you using [Let's Encrypt](https://letsencrypt.org/donate/). Give it some time, but you'll soon have an https enabled custom domain. 

## wrap up
With that, you now you have a hosted blog with a custom domain. There's much more I could have covered, but the above is plenty for getting your blog started. If you want to dig in to Gatsby a bit further, start with [their in depth tutorial](https://www.gatsbyjs.org/tutorial/). Exploring [Gatsby plugins](https://www.gatsbyjs.org/plugins/) is a good next step. 

If you have feedback, <a href="mailto:chris@pinecoder.dev">reach out.</a>