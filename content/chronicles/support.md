---
title: "Chronicles Support"
date: 2026-11-01T00:00:00-08:00
# Unlisted: App Store Connect links here. Not in nav, the post list, RSS, or the sitemap.
_build:
  list: never
---

For help with Chronicles on iPhone or Mac, email [support@pinecoder.dev](mailto:support@pinecoder.dev). Say which device you're on and the app version (Settings › About). For a sync problem, Settings › iCloud › Log has a Copy button; pasting the log into your email helps. It names records by ID, not by their contents.

## Where are my notes stored?

On your device, in the app's own storage. With sync on, an encrypted copy is also kept in your private iCloud. Your notes are never sent to me.

## How does sync work?

Turn it on in Settings › iCloud. The first time, this device's library is merged with the one in iCloud: notes on only one side are added to the other, notes that differ keep both versions, and nothing is deleted. After that, edits move between your devices on their own. If the same note's text changed on two devices, both versions are kept, one as a "conflicted copy".

Everything is encrypted on your device before it goes to iCloud: notes, titles, journal names, tags, and attachments. The key is kept in your iCloud Keychain, so your other devices signed in to the same Apple Account open the library without any extra step.

## What is the recovery code for?

When you set up sync, Chronicles shows a recovery code (32 letters and digits) and waits until you confirm you've saved it before anything goes to iCloud. Keep it somewhere safe, like a password manager.

Chronicles asks for it when:

- you turn on sync on a device that can't reach the key in your iCloud Keychain (Settings › iCloud shows "Needs the recovery code"), or
- you import an encrypted export on a device that doesn't have that library's key.

Without the keychain and the code, the copy in iCloud can't be read, by anyone. The libraries on your devices are unaffected. I can't recover the code for you.

## I lost my recovery code

If one of your devices still has the library, use Settings › iCloud › Remove Library from iCloud on it, then turn sync on again. That sets up a new encrypted library in iCloud from this device, with a new recovery code. Your other devices stop syncing until you turn sync on there again.

## How do I get all my notes out?

Settings › Import & export › Export. Choose one `.chronicles.zip` file, a folder with one Markdown file per note, or an encrypted file (which needs sync's encryption set up). Attachments come along. Any of these can be imported again on another device.

## Privacy

Chronicles collects no data. See the [privacy policy](/chronicles/privacy/).
