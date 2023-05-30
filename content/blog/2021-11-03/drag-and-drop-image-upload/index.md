---
title: Drag and drop image upload
slug: "drag-and-drop-image-upload"
author: Christopher Loverich
date: "2021-11-03T00:00:00.000Z"
tags: devlog
description: "Notes while researching drag and drop image upload for Chronicles"
---

While planning drag and drop image support in [Chronicles](https://github.com/cloverich/chronicles), I had some confusion about handling files once dropped. There are many tutorials out there but they often provide a "copy this code" approach and offer little insight. Many recommended using base64 encoding but that did not seem right. The more general problem of file uploading through a browser is one I occassionally return to, so I wanted a conceptual deep dive I could use as a foundation. After reviewing and taking notes, I organized some of them into this blog post here. 

# Reviewing the design

> This section briefly touches on the design goals that motivated this research. I also review two implementations (Notion, Github) as a basis for the feature. Feel free to skip ahead.
> 

I reviewed a few existing apps to get a feel for their implementation, and broke what I saw down to this:

1. Drag then drop an image. So long as its a supported format, begin uploading
2. Provide an indication to the user that the file is being uploaded
3. Generate a unique filename on the backend, and persist the file to disk
4. Return a URL to the file which the frontend embeds in the document
5. Render the image

Because I think they are good examples, I include here how Github and Notions drag and drop work.

Github's image upload is minimal. When you drop an image onto their plain `<textarea>`, it begins the upload while adding a placeholder indicator:

![Screen Shot 2021-11-03 at 11.24.20 AM.png](assets/Screen_Shot_2021-11-03_at_11.24.20_AM.png)

The file upload happens via a POST request using `multipart/form-data` encoding.

![Screen Shot 2021-11-03 at 11.25.43 AM.png](assets/Screen_Shot_2021-11-03_at_11.25.43_AM.png)

Once the upload is complete, it replaces the placeholder with an `img` tag pointing to the url. The file url itself is a randomly generated ID (looks like a uuid) in place of the original filename and points to Github's CDN. There's no inline rendering of the image during or after upload, unless you click the preview tab. 

![Screen Shot 2021-11-03 at 11.24.28 AM.png](assets/Screen_Shot_2021-11-03_at_11.24.28_AM.png)

Notion's is a bit fancier. It renders the image while it uploads, providing a status bar in the bottom right:

![https://unsplash.com/@matthiasmullie](assets/Screen_Shot_2021-11-03_at_11.13.53_AM.png "https://unsplash.com/@matthiasmullie")

It appears to send it as a PUT request and get back a URL pointing to s3, which then replaces the image above.

![Screen Shot 2021-11-03 at 11.16.50 AM.png](assets/Screen_Shot_2021-11-03_at_11.16.50_AM.png)

The full process is a bit more involved but beyond the scope of this document. There's some dynamic checking for which server to upload to, and it appears to do things differently if you drop multiple images. Digging in further could interesting for a future post. 

# Implementation Overview

After reflecting on the implementation, I think there are four components to implementing basic drag and drop images:

1. Getting and using a reference to the file
2. Sending the file request on the server
3. Processing the file on the server
4. Displaying the image to the user

Where the backend is concerned, the few examples will use NodeJS.

# Getting a handle on the file

To start its helpful to understand how an HTML form lets users add image files and how we can interact with them. Consider this form with a [file input](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/input/file):

```tsx
<form enctype="multipart/form-data" method="post">
  <input id="files-input" type="file" multiple accepts="image/*" />
  <button type="submit">Upload</button>
</form>
```

On its own, a user can interact with this form to send files to the server without any javascript. The `enctype` is important and will be discussed when we look at sending to the server. The file input offers an API accessible via javascript, and this can be used to inspect files added by the user as well as to dynamically add files to the form:

```tsx
const filesInput = document.querySelector('#files-input');

for (const file of [...filesInput.files]) {
	console.log(file.name, file.type, file.size)
}

```

The files object is a [FileList](https://developer.mozilla.org/en-US/docs/Web/API/FileList). It [can be iterated](https://stackoverflow.com/questions/40902437/cant-use-foreach-with-filelist) using spread syntax, `Array.from`, or manually (using `files.length`) but is not an array itself. Each [File](https://developer.mozilla.org/en-US/docs/Web/API/File) object has properties for checking its name and type (extension) **but not its path**. Contents may be read via various API calls but **are not loaded automatically**; more on this later. 

## Drag and Drop

The "Drag and drop" APIs at a simplistic level give you a way to access a `FileList` from a user dropping files onto a designated area. You have to register event handlers on a container element, stop the default behavior, then do something with the FileList. The simplest thing is attach them to the form defined in the previous section:

```tsx
// Any area of the page you'd want users to drop files
const dropArea = document.querySelector('#files-drop-area')
const filesInput = document.querySelector('#files-input');

// dragover and dragenter default behaviors must be cancelled for 
// drop to work
dropArea.ondragover = dropArea.ondragenter = (event) => {
  event.preventDefault();
};

// handle dropped file(s)
dropArea.onDrop = (event) => {
  event.preventDefault();

  // this attaches the dropped files onto the HTML form setup previously
  fileInput.files = event.dataTransfer.files;
  
  form.submit()
}

```

In a nutshell, that is how most image upload forms work and why the preceding section was relevant: Add a normal form, then enhance it with drag and drop.

## Reading the file and base64

To send the file to a server, you don't need to read its contents — there are a few ways to stream the contents directly. However many tutorials detour into reading file contents, often as base64. I found it helpful to understand why you might do this. 

First, having a File object we can use a [FileReader](https://developer.mozilla.org/en-US/docs/Web/API/FileReader) to read its contents. There are multiple ways depending on your needs — as binary data (ArrayBuffer), a data url, or text (w/ encoding specified). The [readAsDataUrl](https://developer.mozilla.org/en-US/docs/Web/API/FileReader/readAsDataURL) method encodes the image in base64, then embeds it as a [data url](https://developer.mozilla.org/en-US/docs/Web/HTTP/Basics_of_HTTP/Data_URIs). Many tutorials would utilize this to display an image preview: 

```tsx
const img = document.createElement("img");
someContainer.appendChild(img);

const reader = new FileReader();
reader.onload = (e) => {
	img.src = e.target.result;
}

// read a file obtained from a FileList
reader.readAsDataURL(file);
```

Base64 was [once popular](https://bunny.net/blog/why-optimizing-your-images-with-base64-is-almost-always-a-bad-idea/) for inlining images for performance reasons, although it seems that should be done in only specific cases. There are some storage and transmission use cases for base64 but for basic uploading and saving a file to disk, it is not required. 

But should "I need an image preview" be synonymous with base64 encoding? [Apparently not](https://levelup.gitconnected.com/managing-image-uploads-in-your-browser-base64-vs-objecturls-part-1-cec5ee9af7be). Instead of reading and encoding the image as a base64 string, you can use Object URLs: 

```tsx
// Rather than using a FileReader, the src can be set to the
// ObjectUrl's return value
img.src = URL.createObjectURL(this.files[i]);

// confusingly, you must clean up the reference after its been loaded
// (or after you are done using it in your SPA)
img.onload = function() {
	URL.revokeObjectURL(this.src);
}
```

Object URL's are a bit mysterious to me, and something I'd like to dive deeper on in the future. For now my take aways are that they provide a way of avoiding the overhead of base64 encoding when dealing with File contents. I.e. you could use the object url for display purposes then rely on the original `File` object to upload it, if relevant. If I dealt with a lot of images or an image focused app I might spend some time playing with these API's a bit further. 

# Submitting the file to the server

Once you have a handle on the file, you're ready to send it to the server; there are a few options:

- An HTML form
- `FormData` API
- fetch (XHR)
- json

## HTML Forms and `enctype`

The example above used an HTML form to send the file to the server, but there's an important bit about the "enctype". An HTML form without files is, by default, handled a bit differently than one with files. The content-type at submit is set to `application/x-www-form-urlencoded`. Sending images this way will not work unless they are encoded as text (i.e. base64) but, as mentioned previously, you don't need to do that. Instead you can send the form as "multipart", which is an alternative method (see [HTML 4 spec here](https://www.w3.org/TR/html401/interact/forms.html#h-17.13.4)) where the browser can send a request with multiple "types" of data. You can opt into this behavior by adding the `enctype="multiepart/form-data"` attribute to the form. With this in place, the browser will stream the files binary contents within the form submission request. As long as the server is setup to handle such a request (see below), the front-end component is complete. 

## `FormData` API

The [FormData API](https://developer.mozilla.org/en-US/docs/Web/API/FormData/Using_FormData_Objects) allows for dynamically constructing a set of form fields in javascript. This API shows up commonly in drag and drop tutorials because it lets you skip setting up or using an existing HTML form. 

```tsx
// handle dropped file(s)
dropArea.onDrop = (event) => {
  event.preventDefault();

  const formData = new FormData()
  const files = event.dataTransfer.files;
  
  // this attaches the dropped files onto the HTML form setup previously
  for (let i=0; i<files.length; i++)
    formData.append(`file_${i}`, files[i], files[i].name);
  }
  
  // sends as multipart/form-data
	fetch('https://mybackend.dev/fileupload', {
	  body: formData,
	  method: 'POST',
	}
}
```

For many kinds of SPA's (like a notes app) it may be awkward to stuff an invisible and otherwise unused HTML form into the page; this API lets you skip that step.

## Using Fetch

The above two examples are what you'll most commonly find in tutorials. But there's also an extremely simple third way to do this, which incidentally is what sent me down this rabbit hole in the first place. You can send `File` objects directly with fetch: 

```tsx

// after getting files from ondrop event
for (const file of [...event.dataTransfer.files]) {
   fetch('https://mybackend.dev/uploadfile', {
	    method: 'POST',
	    body: file,
	 });
}
```

The browser will [*infer* the mime type](https://stackoverflow.com/questions/1201945/how-is-mime-type-of-an-uploaded-file-determined-by-browser) and pass that as a header. If your use case is sufficiently constrained (ex: don't care much about file size; dynamically generating names on server) this might be suitable. Its definitely great for prototypes and in my use case, this is how I did it. This is also quite simple to "handle" on the server; more on that below.

## What about JSON?

When thinking about multi-part having disparate sections, it reminds me a lot of the JSON protocol. Moreover 99.9% of the time I'm using fetch, I'm sending JSON. My mind naturally wondered about it here — can I pass some additional data in a JSON object, and simply include the file as one of the fields? Well yes, but there's an important caveat: JSON doesn't support binary data. This would require base64 encoding the data with the associated shortcomings. It *may* be something you need to do, **but if all you want is some metadata with the file(s), you are better off sending as `multipart/form-data` or even with custom headers in a fetch request.** 

Driving this point home, here's a great stackoverflow question on the subject: [Binary Data in JSON String. Something better than Base64](https://stackoverflow.com/questions/1443158/binary-data-in-json-string-something-better-than-base64). Note how the multipart response (not the accepted answer) is somewhat revelatory to some readers. Its natural to have JSON blinders on when 99% of the data you pass around is in that format, that's the interchange default on the front-end. Yet for binary data there are at least two better options: Fetch and `multipart/form-data`. Both will let you stream data without first encoding in a different format.

# Processing the data on the backend server

How you handle the request on the backend depends on how you sent it. For this example I'll use Node as an example, since that's the server I had already setup. 

To keep it simple, all we'll ask the backend do is the following:

- Parse the request as necessary
- Generate a unique filename and write the sent file to disk
- Return a URL pointing to the file

The most important concept to understand what needs to happen, is to understand how the file was encoded into the request. Specifically, if the file is embedded within JSON or a form, the request needs parsed. 

## `fetch` and `File`

When sending the `File` as the body of fetch directly, the body does not need parsed at all, because it was not encoded into base64 nor embedded into a particular protocol — it was sent as is. Thus processing the file can directly write the contents to disk:

```tsx

// Stream a requests contents to a file
// 
async function save(req: IncomingMessage, filePath: string) {
	// Take in the request & filepath, stream the file to the filePath
	return new Promise((resolve, reject) => {
		const stream = fs.createWriteStream(filePath);
		
		// Not actually sure if I need to setup the open listener
		stream.on("open", () => {
		 req.pipe(stream);
		});
		
		stream.on("close", () => {
		 resolve(filePath);
		});
		
		// If something goes wrong, reject the primise
		stream.on("error", (err) => {
		 reject(err);
		});
	});
}

// upload a file
// note error and validation handling omitted for brevity
async function upload(ctx: RouterContext) {
	// use provided extension, but randomly generated name
  const extension = mime.getExtension(ctx.request.headers["content-type"]);
  const filename = `${cuid()}.${extension}`;
  const filepath = path.join(this.assetsPath, filename);

  await Files.save(ctx.req, filepath);
  ctx.response.status = 200;
  ctx.response.body = {
    filename,
    filepath,
  };
};
```

This code takes advantage of the [browser automatically setting the mime type](https://stackoverflow.com/questions/1201945/how-is-mime-type-of-an-uploaded-file-determined-by-browser) and uses a [mime-to-extension](https://github.com/broofa/mime) library — both good reads to get a handle on dealing with validation and errors (omitted for brevity; see also [here](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/input/file#unique_file_type_specifiers)). I chose [cuid](https://github.com/ericelliott/cuid) for generating filenames and other than the validation and file naming aspects, there's no special handling of the request required. Streaming to a file and returning a url to it is all that is required to have a functioning file upload. 

## Forms handling

If the file comes in from a multipart/form-data upload, you'll need to **parse the request before using the files**. In Node, typical request parsing libraries may or may not include multipart form parsing out of the box. As an example, the [koa-body](https://github.com/koajs/koa-body#hello-world---quickstart) library for koa does and might look like this:

```tsx
const Koa = require('koa');
const koaBody = require('koa-body');

const app = new Koa();

app.use(koaBody({ multipart: true }));
app.use(ctx => {
  for (const file of ctx.request.files) {
    console.log('file metadata', file);
  }
  // ...
}); 
```

Peeling back the layers, the koa-body library delegates multipart form parsing to [formidable](https://github.com/node-formidable/formidable); the [file objects](https://github.com/node-formidable/formidable#file) are metadata and the underlying file data are written to disk by default. This might be surprising behavior — in general I found that dealing with multipart forms is somewhat involved. You need to figure out what you want to do with the file(s) as you parse them, whether you want them stored as intermediate files or streamed to s3 for instance. Most express and koa middleware libraries are based on one of two underlying Node libraries:

- [Formidable](https://github.com/node-formidable/formidable)
- [Busboy](https://github.com/mscdex/busboy)

Were I to build a production worthy file upload service, I'd likely invest my time in researching those two options, and selecting the one that fits my use case best. Working up towards a higher level framework wrapper (like koa-body) would be well informed by said research. Using either directly would also be perfectly suitable.

# Epilogue

I hope the above review of drag and drop file uploading was helpful. After building an understanding of multipart forms and using `File` references, I found I was able to more easily evaluate various blog posts, form libraries, and questions related to them. I was a bit surprised at the amount of misinformation out there and I suppose the topic is a bit more complicated than people initially think. Some of the most helpful resources in researching this topic were:

- [How To Make A Drag-and-Drop File Uploader With Vanilla JavaScript](https://www.smashingmagazine.com/2018/01/drag-drop-file-uploader-vanilla-js/)
- MDN articles: [File](https://developer.mozilla.org/en-US/docs/Web/API/File), [FileReader](https://developer.mozilla.org/en-US/docs/Web/API/FileReader), [file inputs](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/input/file#unique_file_type_specifiers), and [Using files from web applications](https://developer.mozilla.org/en-US/docs/Web/API/File/Using_files_from_web_applications)
- [Why "optimizing" your images with base64 is almost always a bad idea](https://bunny.net/blog/why-optimizing-your-images-with-base64-is-almost-always-a-bad-idea/)
- [Managing image uploads in your Browser: Base64 vs ObjectURLs](https://levelup.gitconnected.com/managing-image-uploads-in-your-browser-base64-vs-objecturls-part-1-cec5ee9af7be)
- [HTML 4 spec: multipart/form-data](https://www.w3.org/TR/html401/interact/forms.html#h-17.13.4) (its short and easy to follow)

Lastly, I did not review validating, resizing, and in general more sophisticated processing of file uploads. These would all need thorough evaluation prior to real production work loads. In particular I would love to dive further into Notion and simliar productionized upload processes to see what can be gleaned from their APIs. 